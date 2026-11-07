// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Infosys Limited
// SPDX-FileCopyrightText: 2024 Canonical Ltd.
/*
 *  Match the NF profiles based on the parameters
 */

// This file contains apis to match the nf profiles based on the parameters provided in the
// Nnrf_NFDiscovery.ApiSearchNFInstancesRequest. There is a match function provided for each NF type
// which must be updated with logic to compare profiles based on the applicable params in
// Nnrf_NFDiscovery.ApiSearchNFInstancesRequest

package nrfcache

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/omec-project/openapi/v2/Nnrf_NFDiscovery"
	"github.com/omec-project/openapi/v2/logger"
	"github.com/omec-project/openapi/v2/models"
)

// maxRegexpCacheEntries caps the SUPI-pattern cache to prevent unbounded growth.
const maxRegexpCacheEntries = 1024

var (
	regexpCache      sync.Map
	regexpCacheCount atomic.Int64
)

// allNFServices returns all NF services from a profile, merging the
// deprecated nfServices array with its TS 29.510 Rel-16 replacement,
// nfServiceList, so callers do not need to check both fields.
func allNFServices(profile *models.NFProfileDiscovery) []models.NFService {
	services := append([]models.NFService{}, profile.GetNfServices()...)
	for _, service := range profile.GetNfServiceList() {
		services = append(services, service)
	}
	return services
}

type matchFilterFunc func(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error)

type matchFilterMap map[models.NFType]matchFilterFunc

var matchFilters = matchFilterMap{
	models.NFTYPE_SMF:  matchSmfProfile,
	models.NFTYPE_AUSF: matchAusfProfile,
	models.NFTYPE_PCF:  matchPcfProfile,
	models.NFTYPE_NSSF: matchNssfProfile,
	models.NFTYPE_UDM:  matchUdmProfile,
	models.NFTYPE_UDR:  matchUdrProfile,
	models.NFTYPE_AMF:  matchAmfProfile,
}

func matchSmfProfile(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error) {
	if profile == nil {
		return false, fmt.Errorf("profile cannot be nil")
	}
	serviceNames := opts.GetServiceNames()
	if serviceNames != nil && len(*serviceNames) > 0 {
		found := false
		for _, requiredService := range *serviceNames {
			for _, nfService := range allNFServices(profile) {
				if nfService.GetServiceName() == requiredService {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			logger.NrfcacheLog.Debugf("smf match failed: no service match for %s", profile.GetNfInstanceId())
			return false, nil
		}
	}

	snssais := opts.GetSnssais()
	if snssais != nil {
		matchCount := 0
		smfInfo, hasSmfInfo := profile.GetSmfInfoOk()
		for _, reqSnssai := range *snssais {
			// Snssai in the smfInfo has priority
			if hasSmfInfo && len(smfInfo.GetSNssaiSmfInfoList()) > 0 {
				for _, s := range smfInfo.GetSNssaiSmfInfoList() {
					snssai := s.GetSNssai()
					if snssai.GetSst() == reqSnssai.GetSst() && snssai.GetSd() == reqSnssai.GetSd() {
						matchCount++
					}
				}
			} else {
				for _, s := range profile.GetAllowedNssais() {
					if s.GetSst() == reqSnssai.GetSst() && s.GetSd() == reqSnssai.GetSd() {
						matchCount++
					}
				}
			}
		}

		// if at least one matching snssai has been found
		if matchCount == 0 {
			return false, nil
		}
	}

	// validate dnn within the S-NSSAI context when snssais filter is also active
	dnn := opts.GetDnn()
	if dnn != nil {
		// if a dnn is provided by the upper layer, check for the exact match
		// or wild card match
		dnnMatched := false

		smfInfo, hasSmfInfo := profile.GetSmfInfoOk()
		if hasSmfInfo {
		matchDnnLoop:
			for _, s := range smfInfo.GetSNssaiSmfInfoList() {
				if snssais != nil {
					entrySnssaiMatch := false
					for _, reqSnssai := range *snssais {
						if s.SNssai.GetSst() == reqSnssai.GetSst() && s.SNssai.GetSd() == reqSnssai.GetSd() {
							entrySnssaiMatch = true
							break
						}
					}
					if !entrySnssaiMatch {
						continue
					}
				}
				for _, d := range s.GetDnnSmfInfoList() {
					if d.GetDnn() == *dnn || d.GetDnn() == "*" {
						dnnMatched = true
						break matchDnnLoop
					}
				}
			}
		}

		if !dnnMatched {
			return false, nil
		}
	}
	logger.NrfcacheLog.Infof("smf match found, nfInstance Id %v", profile.GetNfInstanceId())
	return true, nil
}

func matchSupiRange(supi string, supiRange []models.SupiRange) bool {
	for _, s := range supiRange {
		if matchSingleSupiRange(supi, s) {
			return true
		}
	}
	return false
}

func matchSingleSupiRange(supi string, supiRange models.SupiRange) bool {
	// Handle regex pattern matching (preferred method)
	if pattern := supiRange.GetPattern(); pattern != "" {
		var r *regexp.Regexp
		if cached, ok := regexpCache.Load(pattern); ok {
			r = cached.(*regexp.Regexp)
		} else {
			var err error
			r, err = regexp.Compile(pattern)
			if err != nil {
				logger.NrfcacheLog.Errorf("invalid SUPI pattern '%s': %v", pattern, err)
				return false
			}
			if regexpCacheCount.Load() < maxRegexpCacheEntries {
				if _, loaded := regexpCache.LoadOrStore(pattern, r); !loaded {
					regexpCacheCount.Add(1)
				}
			}
		}
		return r.MatchString(supi)
	}

	// Handle numeric range
	start := supiRange.GetStart()
	end := supiRange.GetEnd()

	if start == "" || end == "" {
		return false
	}

	return isSupiInNumericRange(supi, start, end)
}

func isSupiInNumericRange(supi, start, end string) bool {
	supiNum := extractSupiNumber(supi)
	startNum := extractSupiNumber(start)
	endNum := extractSupiNumber(end)

	// If extraction failed for any, fall back to string comparison
	if supiNum == "" || startNum == "" || endNum == "" {
		return start <= supi && supi <= end
	}

	// For same-length numbers, string comparison works
	if len(supiNum) == len(startNum) && len(startNum) == len(endNum) {
		return startNum <= supiNum && supiNum <= endNum
	}

	// Different lengths require numeric comparison
	supiInt, err1 := strconv.ParseInt(supiNum, 10, 64)
	startInt, err2 := strconv.ParseInt(startNum, 10, 64)
	endInt, err3 := strconv.ParseInt(endNum, 10, 64)

	if err1 != nil || err2 != nil || err3 != nil {
		// Fallback to string comparison
		return start <= supi && supi <= end
	}

	return startInt <= supiInt && supiInt <= endInt
}

func extractSupiNumber(supi string) string {
	if strings.HasPrefix(supi, "imsi-") {
		return supi[5:]
	}
	if strings.HasPrefix(supi, "nai-") {
		return supi[4:]
	}
	// Return as-is if no known prefix
	return supi
}

func matchAusfProfile(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error) {
	if profile == nil {
		return false, fmt.Errorf("profile cannot be nil")
	}
	supi := opts.GetSupi()
	if supi != nil {
		// Unrestricted when no SUPI ranges are declared; see MatchPcfProfile.
		ausfInfo, ok := profile.GetAusfInfoOk()
		if !ok || len(ausfInfo.GetSupiRanges()) == 0 {
			logger.NrfcacheLog.Debugf("ausf match successful (unrestricted: no SUPI ranges) for %s", profile.GetNfInstanceId())
			return true, nil
		}

		matchFound := matchSupiRange(*supi, ausfInfo.GetSupiRanges())
		if matchFound {
			logger.NrfcacheLog.Debugf("ausf match successful for %s", profile.GetNfInstanceId())
		} else {
			logger.NrfcacheLog.Debugf("ausf match failed: SUPI range mismatch for %s", profile.GetNfInstanceId())
		}
		return matchFound, nil
	}

	logger.NrfcacheLog.Debugf("ausf match successful (no SUPI filter) for %s", profile.GetNfInstanceId())
	return true, nil
}

func matchNssfProfile(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error) {
	if profile == nil {
		return false, fmt.Errorf("profile cannot be nil")
	}
	logger.NrfcacheLog.Infoln("nssf match found")
	return true, nil
}

func matchAmfProfile(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error) {
	if profile == nil {
		return false, fmt.Errorf("profile cannot be nil")
	}

	if profile.GetNfType() != models.NFTYPE_AMF {
		return false, fmt.Errorf("profile is not AMF type: %v", profile.GetNfType())
	}
	targetPlmnList := opts.GetTargetPlmnList()
	if targetPlmnList != nil && len(*targetPlmnList) > 0 {
		profilePlmnList := profile.GetPlmnList()
		// A profile without plmnList is available to all PLMNs (TS 29.510 Table 6.2.3.2.3.1-1 [Query-11])
		if len(profilePlmnList) > 0 {
			found := false
			for _, targetPlmn := range *targetPlmnList {
				if slices.Contains(profilePlmnList, targetPlmn) {
					found = true
					break
				}
			}
			if !found {
				logger.NrfcacheLog.Debugf("amf match failed: no PLMN match for %s", profile.GetNfInstanceId())
				return false, nil
			}
		}
	}

	targetNfInstanceId := opts.GetTargetNfInstanceId()
	if targetNfInstanceId != nil && profile.GetNfInstanceId() != *targetNfInstanceId {
		logger.NrfcacheLog.Debugf("amf match failed: NF instance ID mismatch for %s", profile.GetNfInstanceId())
		return false, nil
	}

	amfInfo, hasAmfInfo := profile.GetAmfInfoOk()
	if hasAmfInfo {
		guamiOpt := opts.GetGuami()
		if guamiOpt != nil && !slices.Contains(amfInfo.GetGuamiList(), *guamiOpt) {
			logger.NrfcacheLog.Debugf("amf match failed: GUAMI mismatch for %s", profile.GetNfInstanceId())
			return false, nil
		}

		amfRegionId := opts.GetAmfRegionId()
		if amfRegionId != nil && amfInfo.GetAmfRegionId() != *amfRegionId {
			logger.NrfcacheLog.Debugf("amf match failed: AMF region ID mismatch for %s", profile.GetNfInstanceId())
			return false, nil
		}

		amfSetId := opts.GetAmfSetId()
		if amfSetId != nil && amfInfo.GetAmfSetId() != *amfSetId {
			logger.NrfcacheLog.Debugf("amf match failed: AMF set ID mismatch for %s", profile.GetNfInstanceId())
			return false, nil
		}
		tai := opts.GetTai()
		// Absent taiList and taiRangeList means unrestricted (TS 29.510 Table 6.1.6.2.11-1).
		// taiRangeList matching is not yet implemented; profiles with only taiRangeList pass through.
		if tai != nil && len(amfInfo.GetTaiList()) > 0 && !taiInList(*tai, amfInfo.GetTaiList()) {
			logger.NrfcacheLog.Debugf("amf match failed: TAI mismatch for %s", profile.GetNfInstanceId())
			return false, nil
		}
	} else {
		if opts.GetGuami() != nil || opts.GetAmfRegionId() != nil || opts.GetAmfSetId() != nil || opts.GetTai() != nil {
			logger.NrfcacheLog.Debugf("amf match failed: AMF filters provided but no AmfInfo for %s", profile.GetNfInstanceId())
			return false, nil
		}
	}

	logger.NrfcacheLog.Infof("amf match found = %v", profile.GetNfInstanceId())
	return true, nil
}

// taiInList matches PlmnId, Tac, and Nid; an absent Nid (PLMN context) matches only another absent Nid.
func taiInList(tai models.Tai, list []models.Tai) bool {
	for _, t := range list {
		if t.GetPlmnId() == tai.GetPlmnId() && t.GetTac() == tai.GetTac() && t.GetNid() == tai.GetNid() {
			return true
		}
	}
	return false
}

func matchPcfProfile(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error) {
	if profile == nil {
		return false, fmt.Errorf("profile cannot be nil")
	}
	supi := opts.GetSupi()
	if supi != nil {
		// A profile declaring no SUPI ranges is unrestricted and serves every
		// SUPI. The NRF's own discovery filter encodes this as an $or over
		// "a range contains the SUPI" / "supiRanges is null" / "supiRanges is
		// absent" (nrf producer/nf_discovery.go, [Query-18] supi). Rationale:
		// the cache must select the same profiles as the NRF it caches,
		// otherwise a cached lookup and a live discovery disagree.
		pcfInfo, ok := profile.GetPcfInfoOk()
		if !ok || len(pcfInfo.GetSupiRanges()) == 0 {
			logger.NrfcacheLog.Debugf("pcf match found = true (unrestricted: no SUPI ranges)")
			return true, nil
		}

		matchFound := matchSupiRange(*supi, pcfInfo.GetSupiRanges())
		logger.NrfcacheLog.Infof("pcf match found = %v", matchFound)
		return matchFound, nil
	}

	// No SUPI filter - match any profile
	logger.NrfcacheLog.Infof("pcf match found = true (no SUPI filter)")
	return true, nil
}

func matchUdmProfile(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error) {
	if profile == nil {
		return false, fmt.Errorf("profile cannot be nil")
	}
	supi := opts.GetSupi()
	if supi != nil {
		// Unrestricted when no SUPI ranges are declared; see MatchPcfProfile.
		udmInfo, ok := profile.GetUdmInfoOk()
		if !ok || len(udmInfo.GetSupiRanges()) == 0 {
			logger.NrfcacheLog.Debugf("udm match found = true (unrestricted: no SUPI ranges)")
			return true, nil
		}

		matchFound := matchSupiRange(*supi, udmInfo.GetSupiRanges())
		logger.NrfcacheLog.Infof("udm match found = %v", matchFound)
		return matchFound, nil
	}

	// No SUPI filter - match any profile
	logger.NrfcacheLog.Infof("udm match found = true (no SUPI filter)")
	return true, nil
}

// matchUdrProfile selects UDR profiles for a SUPI-filtered discovery.
//
// Without an entry in matchFilters a UDR profile is dropped unconditionally,
// so a cached UDR lookup returns nothing and every UDM/PCF data access falls
// through to a live NRF query. Rationale: UDR is resolved on every subscriber
// data access, which makes it the most frequently discovered NF in the core.
func matchUdrProfile(profile *models.NFProfileDiscovery, opts Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (bool, error) {
	if profile == nil {
		return false, fmt.Errorf("profile cannot be nil")
	}
	supi := opts.GetSupi()
	if supi == nil {
		logger.NrfcacheLog.Debugf("udr match successful (no SUPI filter) for %s", profile.GetNfInstanceId())
		return true, nil
	}

	// Unrestricted when no SUPI ranges are declared; see MatchPcfProfile.
	udrInfo, ok := profile.GetUdrInfoOk()
	if !ok || len(udrInfo.GetSupiRanges()) == 0 {
		logger.NrfcacheLog.Debugf("udr match successful (unrestricted: no SUPI ranges) for %s", profile.GetNfInstanceId())
		return true, nil
	}

	matchFound := matchSupiRange(*supi, udrInfo.GetSupiRanges())
	logger.NrfcacheLog.Debugf("udr match found = %v for %s", matchFound, profile.GetNfInstanceId())
	return matchFound, nil
}
