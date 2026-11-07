// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Infosys Limited
// SPDX-FileCopyrightText: 2024 Canonical Ltd.
// SPDX-FileCopyrightText: 2025 Intel Corporation
/*
 *  NRF Caching library
 */

package nrfcache

import (
	"container/heap"
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/omec-project/openapi/v2/Nnrf_NFDiscovery"
	"github.com/omec-project/openapi/v2/logger"
	"github.com/omec-project/openapi/v2/models"
)

const (
	defaultNfProfileTTl = time.Minute
)

type nfProfileItem struct {
	nfProfile  *models.NFProfileDiscovery
	expiryTime time.Time
	ttl        time.Duration
	index      int // index of the entry in the priority queue
}

// isExpired - returns true if the expiry time has passed.
func (item *nfProfileItem) isExpired() bool {
	return item.expiryTime.Before(time.Now())
}

// updateExpiryTime - sets new expiry time based on the current time
func (item *nfProfileItem) updateExpiryTime() {
	item.expiryTime = time.Now().Add(item.ttl)
}

func newNfProfileItem(profile *models.NFProfileDiscovery, ttl time.Duration) *nfProfileItem {
	item := &nfProfileItem{
		nfProfile: profile,
		ttl:       ttl,
	}
	item.updateExpiryTime()
	return item
}

// nfProfilePriorityQ : Priority Queue to store the profile. Queue is ordered by expiry time
type nfProfilePriorityQ []*nfProfileItem

// Len - Number of entries in the priority queue
func (npq nfProfilePriorityQ) Len() int {
	return len(npq)
}

// Less - Comparator for the sort interface used by the heap.
// entries will be sorted by increasing order of expiry time
func (npq nfProfilePriorityQ) Less(i, j int) bool {
	return npq[i].expiryTime.Before(npq[j].expiryTime)
}

// Swap - implemented for the sort interface used by the heap pkg.
// swaps the element at i and j.
func (npq nfProfilePriorityQ) Swap(i, j int) {
	npq[i], npq[j] = npq[j], npq[i]
	npq[i].index = i
	npq[j].index = j
}

// at - returns the element at index i
func (npq nfProfilePriorityQ) at(index int) *nfProfileItem {
	return npq[index]
}

// push - adds an entry to the priority queue. Invokes heap api to
// push the entry to the correct location in the queue
func (npq *nfProfilePriorityQ) push(item *nfProfileItem) {
	heap.Push(npq, item)
}

// update - update fields of existing entry. Invokes heap.Fix to re-establish the ordering.
func (npq *nfProfilePriorityQ) update(item *nfProfileItem, value *models.NFProfileDiscovery, ttl time.Duration) {
	item.nfProfile = value
	item.ttl = ttl
	item.updateExpiryTime()
	heap.Fix(npq, item.index)
}

// remove -removes an entry at given index.
func (npq *nfProfilePriorityQ) remove(item *nfProfileItem) {
	heap.Remove(npq, item.index)
}

// Push - implemented for heap interface. appends an element to the priority queue
func (npq *nfProfilePriorityQ) Push(item any) {
	n := len(*npq)
	entry := item.(*nfProfileItem)
	entry.index = n
	*npq = append(*npq, entry)
}

// Pop - implemented for heap interface. Removes the entry with expiry time
func (npq *nfProfilePriorityQ) Pop() any {
	old := *npq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*npq = old[0 : n-1]
	return item
}

// newNfProfilePriorityQ - New priority queue for storing NF Profiles.
func newNfProfilePriorityQ() *nfProfilePriorityQ {
	q := &nfProfilePriorityQ{}
	heap.Init(q)
	return q
}

// nrfCache : cache of nf profiles
type nrfCache struct {
	cache               map[string]*nfProfileItem // map[nf-instance-id] =*NfProfile
	priorityQ           *nfProfilePriorityQ       // sorted by expiry time
	evictionTicker      *time.Ticker
	done                chan struct{}
	nrfDiscoveryQueryCb nrfDiscoveryQueryFn // nrf query callback
	evictionInterval    time.Duration       // timer interval in which the cache is checked for eviction of expired entries
	mutex               sync.RWMutex
	discoveryMutex      sync.Mutex
}

// handleLookup - Checks if the cache has nf cache entry corresponding to the parameters specified.
// If entry does not exist, perform nrf discovery query. To avoid concurrency issues,
// nrf discovery query is mutex protected.
func (c *nrfCache) handleLookup(ctx context.Context, nrfUri string, targetNfType, requestNfType models.NFType, param Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (models.SearchResult, error) {
	// Check context cancellation
	select {
	case <-ctx.Done():
		return models.SearchResult{}, ctx.Err()
	default:
	}

	// Check first with read lock
	c.mutex.RLock()
	nfInstances := c.get(param)
	c.mutex.RUnlock()

	if len(nfInstances) > 0 {
		var result models.SearchResult
		result.SetNfInstances(nfInstances)
		return result, nil
	}

	// discoveryMutex serializes NRF round-trips without blocking concurrent cache hits.
	c.discoveryMutex.Lock()
	defer c.discoveryMutex.Unlock()

	// Re-check in case another goroutine already populated the entry.
	c.mutex.RLock()
	nfInstances = c.get(param)
	c.mutex.RUnlock()

	if len(nfInstances) > 0 {
		var result models.SearchResult
		result.SetNfInstances(nfInstances)
		return result, nil
	}

	logger.NrfcacheLog.Warnf("cache miss for nftype %s", targetNfType)
	searchResult, err := c.nrfDiscoveryQueryCb(ctx, nrfUri, targetNfType, requestNfType, param)
	if err != nil {
		return models.SearchResult{}, fmt.Errorf("NRF discovery failed: %w", err)
	}

	if searchResult == nil {
		return models.SearchResult{}, fmt.Errorf("NRF discovery returned nil result")
	}

	ttl := time.Duration(searchResult.GetValidityPeriod()) * time.Second
	c.mutex.Lock()
	nfInstances = searchResult.GetNfInstances()
	for i := range nfInstances {
		c.set(&nfInstances[i], ttl)
	}
	c.mutex.Unlock()

	return *searchResult, nil
}

// set - Adds nf profile entry to the map and the priority queue
func (c *nrfCache) set(nfProfile *models.NFProfileDiscovery, ttl time.Duration) {
	if ttl == 0 {
		ttl = defaultNfProfileTTl
	}

	// Convert seconds to duration if needed
	if ttl < time.Second {
		ttl = ttl * time.Second
	}

	item, exists := c.cache[nfProfile.GetNfInstanceId()]
	if exists {
		// if item.isExpired()
		c.priorityQ.update(item, nfProfile, ttl)
	} else {
		newItem := newNfProfileItem(nfProfile, ttl)
		c.cache[nfProfile.GetNfInstanceId()] = newItem
		c.priorityQ.push(newItem)
	}
}

// get - checks if nf profile corresponding to the search opts exist in the cache.
func (c *nrfCache) get(opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) []models.NFProfileDiscovery {
	var nfProfiles []models.NFProfileDiscovery

	if len(c.cache) == 0 {
		return nfProfiles
	}

	// Check if we have specific NF instance ID filter
	if targetNfInstanceId := opts.GetTargetNfInstanceId(); targetNfInstanceId != nil {
		if item, exists := c.cache[*targetNfInstanceId]; exists && !item.isExpired() {
			nfProfiles = append(nfProfiles, *item.nfProfile)
		}
		return nfProfiles
	}

	// General filtering
	isEmptyFilter := reflect.DeepEqual(opts, Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{})

	for _, element := range c.cache {
		if element.isExpired() {
			continue
		}

		if isEmptyFilter {
			nfProfiles = append(nfProfiles, *element.nfProfile)
			continue
		}

		if cb, ok := matchFilters[element.nfProfile.GetNfType()]; ok {
			if matchFound, err := cb(element.nfProfile, opts); err != nil {
				logger.NrfcacheLog.Errorf("match filter error for %s: %v", element.nfProfile.GetNfInstanceId(), err)
			} else if matchFound {
				nfProfiles = append(nfProfiles, *element.nfProfile)
			}
		}
	}

	return nfProfiles
}

// removeByNfInstanceId - removes nf profile with nfInstanceId from the cache and queue
func (c *nrfCache) removeByNfInstanceId(nfInstanceId string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, rc := c.cache[nfInstanceId]
	if rc {
		c.remove(item)
	}
	return rc
}

// remove -
func (c *nrfCache) remove(item *nfProfileItem) {
	c.priorityQ.remove(item)
	delete(c.cache, item.nfProfile.GetNfInstanceId())
}

// cleanupExpiredItems - removes the profiles with expired TTLs
func (c *nrfCache) cleanupExpiredItems() {
	logger.NrfcacheLog.Infoln("nrf cache: cleanup expired items")
	for c.priorityQ.Len() > 0 {
		item := c.priorityQ.at(0)
		if !item.isExpired() {
			break
		}

		logger.NrfcacheLog.Debugf("evicted nf instance %s", item.nfProfile.GetNfInstanceId())
		c.remove(item)
	}
}

// purge - release the cache and its resources.
func (c *nrfCache) purge() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	close(c.done)
	c.priorityQ = newNfProfilePriorityQ()
	c.cache = make(map[string]*nfProfileItem)
	c.evictionTicker.Stop()
}

func (c *nrfCache) startExpiryProcessing() {
	for {
		select {
		case <-c.evictionTicker.C:
			c.mutex.Lock()
			if c.priorityQ.Len() == 0 {
				c.mutex.Unlock()
				continue
			}

			c.cleanupExpiredItems()
			c.mutex.Unlock()

		case <-c.done:
			return
		}
	}
}

func newNrfCache(duration time.Duration, dbqueryCb nrfDiscoveryQueryFn) *nrfCache {
	if dbqueryCb == nil {
		panic("NrfDiscoveryQueryCb cannot be nil")
	}

	cache := &nrfCache{
		cache:               make(map[string]*nfProfileItem),
		priorityQ:           newNfProfilePriorityQ(),
		evictionInterval:    duration,
		nrfDiscoveryQueryCb: dbqueryCb,
		done:                make(chan struct{}),
	}

	cache.evictionTicker = time.NewTicker(duration)
	go cache.startExpiryProcessing()

	return cache
}

type nrfMasterCache struct {
	nrfDiscoveryQueryCb nrfDiscoveryQueryFn
	nfTypeToCacheMap    map[models.NFType]*nrfCache
	evictionInterval    time.Duration
	mutex               sync.Mutex
}

func (c *nrfMasterCache) getNrfCacheInstance(targetNfType models.NFType) *nrfCache {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	cache, exists := c.nfTypeToCacheMap[targetNfType]
	if !exists {
		logger.NrfcacheLog.Infof("creating cache for nftype %v", targetNfType)
		cache = newNrfCache(c.evictionInterval, c.nrfDiscoveryQueryCb)
		c.nfTypeToCacheMap[targetNfType] = cache
	}
	return cache
}

func (c *nrfMasterCache) clearNrfMasterCache() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for k, cache := range c.nfTypeToCacheMap {
		cache.purge()
		delete(c.nfTypeToCacheMap, k)
	}
}

func (c *nrfMasterCache) removeNfProfile(nfInstanceId string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var ok bool
	for _, cache := range c.nfTypeToCacheMap {
		if ok = cache.removeByNfInstanceId(nfInstanceId); ok {
			break
		}
	}
	return ok
}

var masterCache *nrfMasterCache

type nrfDiscoveryQueryFn func(ctx context.Context, nrfUri string, targetNfType, requestNfType models.NFType, param Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (*models.SearchResult, error)

func InitNrfCaching(interval time.Duration, cb func(ctx context.Context, nrfUri string, targetNfType, requestNfType models.NFType, param Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (*models.SearchResult, error)) {
	m := &nrfMasterCache{
		nfTypeToCacheMap:    make(map[models.NFType]*nrfCache),
		evictionInterval:    interval,
		nrfDiscoveryQueryCb: cb,
	}
	masterCache = m
}

func disableNrfCaching() {
	if masterCache != nil {
		masterCache.clearNrfMasterCache()
		masterCache = nil
	}
}

func SearchNFInstances(ctx context.Context, nrfUri string, targetNfType, requestNfType models.NFType, param Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (*models.SearchResult, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context cannot be nil")
	}

	if nrfUri == "" {
		return nil, fmt.Errorf("nrfUri cannot be empty")
	}

	if !targetNfType.IsValid() {
		return nil, fmt.Errorf("invalid target NFType: %v", targetNfType)
	}

	if !requestNfType.IsValid() {
		return nil, fmt.Errorf("invalid requester NFType: %v", requestNfType)
	}

	if masterCache == nil {
		return nil, fmt.Errorf("NRF cache is not initialized")
	}

	if !targetNfType.IsValid() {
		return nil, fmt.Errorf("invalid NFType: %v", targetNfType)
	}
	var searchResult models.SearchResult

	if masterCache == nil {
		return nil, fmt.Errorf("masterCache is not initialized")
	}

	c := masterCache.getNrfCacheInstance(targetNfType)
	if c == nil {
		logger.NrfcacheLog.Errorf("failed to find/create cache for nfType: %v", targetNfType)
		return nil, fmt.Errorf("unable to find/create cache for NF type: %v", targetNfType)
	}

	searchResult, err := c.handleLookup(ctx, nrfUri, targetNfType, requestNfType, param)
	if err != nil {
		logger.NrfcacheLog.With("nfType", targetNfType, "param", param).Errorln("handleLookup failed:", err)
		return nil, fmt.Errorf("handleLookup for nfType %v failed: %w", targetNfType, err)
	}
	for _, np := range searchResult.GetNfInstances() {
		logger.NrfcacheLog.Infof("%+v", np)
	}
	return &searchResult, err
}

func RemoveNfProfileFromNrfCache(nfInstanceId string) bool {
	return masterCache.removeNfProfile(nfInstanceId)
}
