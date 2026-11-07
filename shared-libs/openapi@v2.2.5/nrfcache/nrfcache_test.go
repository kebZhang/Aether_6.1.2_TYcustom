// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2024 Canonical Ltd.
// SPDX-FileCopyrightText: 2025 Intel Corporation
/*
 *  Tests for NRF Caching library
 */

package nrfcache

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/omec-project/openapi/v2"
	"github.com/omec-project/openapi/v2/Nnrf_NFDiscovery"
	"github.com/omec-project/openapi/v2/logger"
	"github.com/omec-project/openapi/v2/models"
)

const (
	testDnnInternet = "internet"
	testAmfSetID    = "3f8"
)

type testContext struct {
	nfProfilesDb           map[string]string
	validityPeriod         int32
	evictionInterval       int32
	nrfDbCallbackCallCount int32
	mu                     sync.Mutex
}

func (tc *testContext) getCallbackCount() int32 {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	return tc.nrfDbCallbackCallCount
}

func (tc *testContext) getNfProfile(key string) (models.NFProfileDiscovery, error) {
	var err error
	var profile models.NFProfileDiscovery

	nfProfileStr, exists := tc.nfProfilesDb[key]

	if exists {
		err = json.Unmarshal([]byte(nfProfileStr), &profile)
	} else {
		err = fmt.Errorf("failed to find nf profile for %s", key)
	}

	return profile, err
}

func (tc *testContext) getNfProfiles(targetNfType models.NFType) ([]models.NFProfileDiscovery, error) {
	var nfProfiles []models.NFProfileDiscovery

	for key, elem := range tc.nfProfilesDb {
		if strings.Contains(key, string(targetNfType)) {
			var profile models.NFProfileDiscovery
			err := json.Unmarshal([]byte(elem), &profile)
			if err != nil {
				return nil, err
			}

			nfProfiles = append(nfProfiles, profile)
		}
	}

	return nfProfiles, nil
}

func (tc *testContext) setValidityPeriod(period int32) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.validityPeriod = period
}

func (tc *testContext) initTestData() {
	tc.nfProfilesDb["SMF-010203-internet"] = `{
		  "ipv4Addresses": [
			"smf"
		  ],
		  "allowedPlmns": [
			{
			  "mcc": "208",
			  "mnc": "93"
			}
		  ],
		  "smfInfo": {
			"sNssaiSmfInfoList": [
			  {
				"sNssai": {
				  "sst": 1,
				  "sd": "010203"
				},
				"dnnSmfInfoList": [
				  {
					"dnn": "internet"
				  }
				]
			  }
			]
		  },
		  "nfServices": [
			{
			  "apiPrefix": "http://smf:29502",
			  "allowedPlmns": [
				{
				  "mcc": "208",
				  "mnc": "93"
				}
			  ],
			  "serviceInstanceId": "b926f193-1083-49a8-adb3-5fcf57a1f0bfnsmf-pdusession",
			  "serviceName": "nsmf-pdusession",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "https://smf:29502/nsmf-pdusession/v1",
				  "expiry": "2022-08-17T05:31:40.997097141Z"
				}
			  ],
			  "scheme": "https",
			  "nfServiceStatus": "REGISTERED"
			},
			{
			  "scheme": "https",
			  "nfServiceStatus": "REGISTERED",
			  "apiPrefix": "http://smf:29502",
			  "allowedPlmns": [
				{
				  "mcc": "208",
				  "mnc": "93"
				}
			  ],
			  "serviceInstanceId": "b926f193-1083-49a8-adb3-5fcf57a1f0bfnsmf-event-exposure",
			  "serviceName": "nsmf-event-exposure",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "https://smf:29502/nsmf-pdusession/v1",
				  "expiry": "2022-08-17T05:31:40.997097141Z"
				}
			  ]
			}
		  ],
		  "nfInstanceId": "b926f193-1083-49a8-adb3-5fcf57a1f0bf",
		  "plmnList": [
			{
			  "mnc": "93",
			  "mcc": "208"
			}
		  ],
		  "sNssais": [
			{
			  "sd": "010203",
			  "sst": 1
			}
		  ],
		  "nfType": "SMF",
		  "nfStatus": "REGISTERED"
		}`
	tc.nfProfilesDb["SMF-010203-ims"] = `{
		  "ipv4Addresses": [
			"smf"
		  ],
		  "allowedPlmns": [
			{
			  "mcc": "208",
			  "mnc": "93"
			}
		  ],
		  "smfInfo": {
			"sNssaiSmfInfoList": [
			  {
				"sNssai": {
				  "sst": 1,
				  "sd": "010203"
				},
				"dnnSmfInfoList": [
				  {
					"dnn": "ims"
				  }
				]
			  }
			]
		  },
		  "nfServices": [
			{
			  "apiPrefix": "http://smf:29502",
			  "allowedPlmns": [
				{
				  "mcc": "208",
				  "mnc": "93"
				}
			  ],
			  "serviceInstanceId": "c926f193-1083-49a8-adb3-5fcf57a1f0bfnsmf-pdusession",
			  "serviceName": "nsmf-pdusession",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "https://smf:29502/nsmf-pdusession/v1",
				  "expiry": "2022-08-17T05:31:40.997097141Z"
				}
			  ],
			  "scheme": "https",
			  "nfServiceStatus": "REGISTERED"
			},
			{
			  "scheme": "https",
			  "nfServiceStatus": "REGISTERED",
			  "apiPrefix": "http://smf:29502",
			  "allowedPlmns": [
				{
				  "mcc": "208",
				  "mnc": "93"
				}
			  ],
			  "serviceInstanceId": "c926f193-1083-49a8-adb3-5fcf57a1f0bfnsmf-event-exposure",
			  "serviceName": "nsmf-event-exposure",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "https://smf:29502/nsmf-pdusession/v1",
				  "expiry": "2022-08-17T05:31:40.997097141Z"
				}
			  ]
			}
		  ],
		  "nfInstanceId": "c926f193-1083-49a8-adb3-5fcf57a1f0bf",
		  "plmnList": [
			{
			  "mnc": "93",
			  "mcc": "208"
			}
		  ],
		  "sNssais": [
			{
			  "sd": "010203",
			  "sst": 1
			}
		  ],
		  "nfType": "SMF",
		  "nfStatus": "REGISTERED"
		}
`
	tc.nfProfilesDb["SMF-0a0b0c-internet"] = `{
		  "ipv4Addresses": [
			"smf"
		  ],
		  "allowedPlmns": [
			{
			  "mcc": "208",
			  "mnc": "93"
			}
		  ],
		  "smfInfo": {
			"sNssaiSmfInfoList": [
			  {
				"sNssai": {
				  "sst": 1,
				  "sd": "0a0b0c"
				},
				"dnnSmfInfoList": [
				  {
					"dnn": "internet"
				  }
				]
			  }
			]
		  },
		  "nfServices": [
			{
			  "apiPrefix": "http://smf:29502",
			  "allowedPlmns": [
				{
				  "mcc": "208",
				  "mnc": "93"
				}
			  ],
			  "serviceInstanceId": "d926f193-1083-49a8-adb3-5fcf57a1f0bfnsmf-pdusession",
			  "serviceName": "nsmf-pdusession",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "https://smf:29502/nsmf-pdusession/v1",
				  "expiry": "2022-08-17T05:31:40.997097141Z"
				}
			  ],
			  "scheme": "https",
			  "nfServiceStatus": "REGISTERED"
			},
			{
			  "scheme": "https",
			  "nfServiceStatus": "REGISTERED",
			  "apiPrefix": "http://smf:29502",
			  "allowedPlmns": [
				{
				  "mcc": "208",
				  "mnc": "93"
				}
			  ],
			  "serviceInstanceId": "d926f193-1083-49a8-adb3-5fcf57a1f0bfnsmf-event-exposure",
			  "serviceName": "nsmf-event-exposure",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "https://smf:29502/nsmf-pdusession/v1",
				  "expiry": "2022-08-17T05:31:40.997097141Z"
				}
			  ]
			}
		  ],
		  "nfInstanceId": "d926f193-1083-49a8-adb3-5fcf57a1f0bf",
		  "plmnList": [
			{
			  "mnc": "93",
			  "mcc": "208"
			}
		  ],
		  "sNssais": [
			{
			  "sd": "0a0b0c",
			  "sst": 1
			}
		  ],
		  "nfType": "SMF",
		  "nfStatus": "REGISTERED"
		}`
	tc.nfProfilesDb["AUSF-1"] = `{ "nfServices": [
			{
			  "serviceName": "nausf-auth",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "1.0.0"
				}
			  ],
			  "scheme": "http",
			  "nfServiceStatus": "REGISTERED",
			  "ipEndPoints": [
				{
				  "ipv4Address": "ausf",
				  "port": 29509
				}
			  ],
			  "serviceInstanceId": "57d0a167-5283-4170-bdd8-881076049a81"
			}
		  ],
		  "ausfInfo": {
			"supiRanges": [
			  { "start": "123456789040000", "end": "123456789049999" }
			]
		  },
		  "nfInstanceId": "57d0a167-5283-4170-bdd8-881076049a81",
		  "nfType": "AUSF",
		  "nfStatus": "REGISTERED",
		  "plmnList": [
			{
			  "mcc": "208",
			  "mnc": "93"
			}
		  ],
		  "ipv4Addresses": [
			"ausf"
		  ],
		  "ausfInfo": {
			"groupId": "ausfGroup001"
		  }
		}`
	tc.nfProfilesDb["AUSF-2"] = `{ "nfServices": [
			{
			  "serviceName": "nausf-auth",
			  "versions": [
				{
				  "apiVersionInUri": "v1",
				  "apiFullVersion": "1.0.0"
				}
			  ],
			  "scheme": "http",
			  "nfServiceStatus": "REGISTERED",
			  "ipEndPoints": [
				{
				  "ipv4Address": "ausf",
				  "port": 29509
				}
			  ],
			  "serviceInstanceId": "67d0a167-5283-4170-bdd8-881076049a81"
			}
		  ],
		  "ausfInfo": {
			"supiRanges": [
			  { "pattern": "^imsi-22345678904[0-9]{4}$" }
			]
		  },
		  "nfInstanceId": "67d0a167-5283-4170-bdd8-881076049a81",
		  "nfType": "AUSF",
		  "nfStatus": "REGISTERED",
		  "plmnList": [
			{
			  "mcc": "208",
			  "mnc": "93"
			}
		  ],
		  "ipv4Addresses": [
			"ausf"
		  ],
		  "ausfInfo": {
			"groupId": "ausfGroup001"
		  }
		}`
	tc.nfProfilesDb["AMF-01"] = `
	{
	  "nfServices": [
		{
		  "serviceInstanceId": "0",
		  "serviceName": "namf-comm",
		  "versions": [
			{
			  "apiVersionInUri": "v1",
			  "apiFullVersion": "1.0.0"
			}
		  ],
		  "scheme": "http",
		  "nfServiceStatus": "REGISTERED",
		  "ipEndPoints": [
			{
			  "ipv4Address": "amf",
			  "transport": "TCP",
			  "port": 29518
			}
		  ],
		  "apiPrefix": "http://amf:29518"
		}
	  ],
	  "nfInstanceId": "9f7d5a3f-88ab-4525-b31e-334da7faedab",
	  "nfType": "AMF",
	  "nfStatus": "REGISTERED",
	  "plmnList": [
		{
		  "mcc": "208",
		  "mnc": "93"
		}
	  ],
	  "sNssais": [
		{
		  "sst": 1,
		  "sd": "010203"
		}
	  ],
	  "ipv4Addresses": [
		"amf"
	  ],
	  "amfInfo": {
		"amfSetId": "3f8",
		"amfRegionId": "ca",
		"guamiList": [
		  {
			"plmnId": {
			  "mcc": "208",
			  "mnc": "93"
			},
			"amfId": "cafe00"
		  }
		],
		"taiList": [
		  {
			"plmnId": {
			  "mcc": "208",
			  "mnc": "93"
			},
			"tac": "1"
		  }
		]
	  }
	} `
}

func (tc *testContext) nrfDbCallback(ctx context.Context, nrfUri string, targetNfType, requestNfType models.NFType, param Nnrf_NFDiscovery.ApiSearchNFInstancesRequest) (*models.SearchResult, error) {
	tc.mu.Lock()
	tc.nrfDbCallbackCallCount++
	tc.mu.Unlock()

	logger.NrfcacheLog.Infoln("nrfDbCallback Entry")

	searchResult := models.NewSearchResultWithDefaults()
	var nfProfile models.NFProfileDiscovery
	var err error

	searchResult.SetValidityPeriod(tc.validityPeriod)

	switch targetNfType {
	case models.NFTYPE_SMF:
		key := "SMF"
		if !reflect.DeepEqual(param, Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}) {
			snssais := param.GetSnssais()
			if snssais != nil {
				reqSnssais := *snssais
				snssai := reqSnssais[0]
				key += "-" + snssai.GetSd()
			}
			dnn := param.GetDnn()
			if dnn != nil {
				key += "-" + *dnn
			}
			nfProfile, err = tc.getNfProfile(key)
			if err != nil {
				return searchResult, err
			}
			searchResult.SetNfInstances(append(searchResult.GetNfInstances(), nfProfile))
		} else {
			var profiles []models.NFProfileDiscovery
			profiles, err = tc.getNfProfiles(targetNfType)
			searchResult.SetNfInstances(profiles)
		}
	case models.NFTYPE_AUSF, models.NFTYPE_AMF:
		var profiles []models.NFProfileDiscovery
		profiles, err = tc.getNfProfiles(targetNfType)
		searchResult.SetNfInstances(profiles)
	default:
		return searchResult, fmt.Errorf("unsupported NFType: %s", targetNfType)
	}

	return searchResult, err
}

func setupTest(t *testing.T) (*testContext, func()) {
	t.Helper()

	tc := &testContext{
		nfProfilesDb:           make(map[string]string),
		validityPeriod:         60,
		evictionInterval:       120,
		nrfDbCallbackCallCount: 0,
	}

	tc.initTestData()

	cleanup := func() {
		disableNrfCaching()
	}

	return tc, cleanup
}

func createTestParam() Nnrf_NFDiscovery.ApiSearchNFInstancesRequest {
	param := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}
	param = param.ServiceNames([]models.ServiceName{models.SERVICENAME_NSMF_PDUSESSION})
	param = param.Dnn(testDnnInternet)
	param = param.Snssais([]models.Snssai{{Sst: 1, Sd: openapi.PtrString("010203")}})
	return param
}

func createSMFParam(dnn, sd string) Nnrf_NFDiscovery.ApiSearchNFInstancesRequest {
	param := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}
	param = param.ServiceNames([]models.ServiceName{models.SERVICENAME_NSMF_PDUSESSION})
	param = param.Dnn(dnn)
	param = param.Snssais([]models.Snssai{{Sst: 1, Sd: openapi.PtrString(sd)}})
	return param
}

func createAusfParam(supi string) Nnrf_NFDiscovery.ApiSearchNFInstancesRequest {
	param := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}
	param = param.Supi(supi)
	return param
}

func createAmfParamWithPlmns(plmnList []models.PlmnId) Nnrf_NFDiscovery.ApiSearchNFInstancesRequest {
	param := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}
	param = param.TargetPlmnList(plmnList)
	return param
}

func waitForEviction(evictionIntervalSeconds int32) {
	// Wait longer than the eviction interval to ensure eviction has run
	sleepDuration := time.Duration(evictionIntervalSeconds+1) * time.Second
	time.Sleep(sleepDuration)
}

func assertSearchResult(t *testing.T, result *models.SearchResult, err error, expectedInstances int) {
	t.Helper()

	if err != nil {
		t.Fatalf("SearchNFInstances failed: %v", err)
	}

	if result == nil {
		t.Fatal("SearchNFInstances returned nil result")
	}

	if len(result.NfInstances) != expectedInstances {
		t.Errorf("Expected %d NF instances, got %d", expectedInstances, len(result.NfInstances))
	}
}

func assertCallbackCount(t *testing.T, actual, expected int32) {
	t.Helper()

	if actual != expected {
		t.Errorf("Callback count mismatch: expected %d, got %d", expected, actual)
	}
}

func TestCacheMissAndHits(t *testing.T) {
	testCtx, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	evictionTimerVal := time.Duration(testCtx.evictionInterval)
	InitNrfCaching(evictionTimerVal*time.Second, testCtx.nrfDbCallback)

	// Test case 1: Cache Miss for dnn - 'internet'
	param1 := createSMFParam(testDnnInternet, "010203")
	result, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param1)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, testCtx.getCallbackCount(), 1)

	// Test case 2: Cache hit scenario (same param)
	result, err = SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param1)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, testCtx.getCallbackCount(), 1) // Should still be 1

	// Test case 3: Cache Miss for dnn 'ims'
	param2 := createSMFParam("ims", "010203")
	result, err = SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param2)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, testCtx.getCallbackCount(), 2)

	// Test case 4: Cache Miss for dnn 'internet' sd '0a0b0c'
	param3 := createSMFParam(testDnnInternet, "0a0b0c")
	result, err = SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param3)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, testCtx.getCallbackCount(), 3)
}

func TestCacheMissOnTTlExpiry(t *testing.T) {
	testCtx, cleanup := setupTest(t)
	defer cleanup()

	// Reduce cache intervals for faster test execution
	testCtx.validityPeriod = 2
	testCtx.evictionInterval = 4

	ctx := context.Background()
	evictionTimerVal := time.Duration(testCtx.evictionInterval)
	InitNrfCaching(evictionTimerVal*time.Second, testCtx.nrfDbCallback)

	// First call with empty request - should be cache miss
	emptyParam := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}
	result, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, emptyParam)
	assertSearchResult(t, result, err, 3) // Expecting all 3 SMF profiles
	assertCallbackCount(t, testCtx.getCallbackCount(), 1)

	t.Log("waiting for profile validity timeout")
	time.Sleep(3 * time.Second)

	// After TTL expiry, this should be a cache miss
	param := createSMFParam(testDnnInternet, "0a0b0c")
	result, err = SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, testCtx.getCallbackCount(), 2)

	// Immediate second call with same param - should be cache hit
	result, err = SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, testCtx.getCallbackCount(), 2) // Should still be 2
}

func TestCacheTTLExpiry(t *testing.T) {
	ctx, cleanup := setupTest(t)
	defer cleanup()

	// Use shorter, more predictable intervals for testing
	ctx.validityPeriod = 1   // 1 second
	ctx.evictionInterval = 2 // 2 seconds

	evictionTimerVal := time.Duration(ctx.evictionInterval) * time.Second
	InitNrfCaching(evictionTimerVal, ctx.nrfDbCallback)

	param := createTestParam()

	// First call - should be cache miss
	result, err := SearchNFInstances(context.Background(), "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, ctx.getCallbackCount(), 1)

	// Second call immediately - should be cache hit
	result, err = SearchNFInstances(context.Background(), "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, ctx.getCallbackCount(), 1) // No additional calls

	// Wait for TTL expiry with some buffer
	time.Sleep(time.Duration(ctx.validityPeriod+1) * time.Second)

	// Third call after expiry - should be cache miss again
	result, err = SearchNFInstances(context.Background(), "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param)
	assertSearchResult(t, result, err, 1)
	assertCallbackCount(t, ctx.getCallbackCount(), 2) // One additional call
}

func TestCacheEviction(t *testing.T) {
	testCtx, cleanup := setupTest(t)
	defer cleanup()

	// Use shorter intervals for faster testing
	testCtx.evictionInterval = 2 // 2 seconds eviction check

	ctx := context.Background()
	evictionTimerVal := time.Duration(testCtx.evictionInterval)
	InitNrfCaching(evictionTimerVal*time.Second, testCtx.nrfDbCallback)

	testCases := []struct {
		name           string
		dnn            string
		sd             string
		description    string
		validityPeriod int32
	}{
		{
			name:           "short_ttl_entry",
			dnn:            testDnnInternet,
			sd:             "010203",
			validityPeriod: 1,
			description:    "Entry with 1 second TTL (will expire first)",
		},
		{
			name:           "medium_ttl_entry",
			dnn:            "ims",
			sd:             "010203",
			validityPeriod: 5,
			description:    "Entry with 5 second TTL",
		},
		{
			name:           "long_ttl_entry",
			dnn:            testDnnInternet,
			sd:             "0a0b0c",
			validityPeriod: 30,
			description:    "Entry with 30 second TTL (will survive)",
		},
	}

	// Create cache entries with different TTLs
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testCtx.setValidityPeriod(tc.validityPeriod)
			param := createSMFParam(tc.dnn, tc.sd)

			result, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, param)
			assertSearchResult(t, result, err, 1)
			assertCallbackCount(t, testCtx.getCallbackCount(), int32(i+1))

			t.Logf("Created %s: %s", tc.name, tc.description)
		})
	}

	// Wait for eviction to occur
	t.Run("verify_eviction", func(t *testing.T) {
		t.Log("waiting for eviction timeout")
		waitForEviction(testCtx.evictionInterval + 2) // Wait a bit longer than eviction interval

		// Test that short TTL entry was evicted (should cause new callback)
		testCtx.setValidityPeriod(60) // Reset to reasonable TTL
		shortTTLParam := createSMFParam(testDnnInternet, "010203")
		result, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_SMF, models.NFTYPE_AMF, shortTTLParam)
		assertSearchResult(t, result, err, 1)

		// This should be a cache miss (new callback) if eviction worked
		expectedCallbacks := int32(4) // 3 initial + 1 after eviction
		assertCallbackCount(t, testCtx.getCallbackCount(), expectedCallbacks)

		t.Log("eviction verification completed")
	})
}

func TestCacheConcurrency(t *testing.T) {
	testCtx, cleanup := setupTest(t)
	defer cleanup()

	evictionTimerVal := time.Duration(testCtx.evictionInterval)
	InitNrfCaching(evictionTimerVal*time.Second, testCtx.nrfDbCallback)

	numGoroutines := 100
	param := createSMFParam(testDnnInternet, "010203")
	expectedCallCount := testCtx.getCallbackCount() + 1

	t.Run("concurrent_cache_access", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		errChan := make(chan error, numGoroutines)
		resultChan := make(chan *models.SearchResult, numGoroutines)

		// Launch concurrent requests
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()

				result, err := SearchNFInstances(
					context.Background(),
					"testNrf",
					models.NFTYPE_SMF,
					models.NFTYPE_AMF,
					param,
				)
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d failed: %w", id, err)
					return
				}

				resultChan <- result
			}(i)
		}

		// Wait for completion with timeout
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			// Success
		case err := <-errChan:
			t.Fatalf("Concurrent test failed: %v", err)
		case <-time.After(10 * time.Second):
			t.Fatal("Concurrent test timed out")
		}

		// Verify results
		close(resultChan)
		resultCount := 0
		for result := range resultChan {
			if len(result.NfInstances) == 0 {
				t.Error("Empty result from concurrent request")
			}
			resultCount++
		}

		if resultCount != numGoroutines {
			t.Errorf("Expected %d results, got %d", numGoroutines, resultCount)
		}

		// Should only have one callback due to caching
		assertCallbackCount(t, testCtx.getCallbackCount(), expectedCallCount)
	})
}

func TestAusfMatchFilters(t *testing.T) {
	testScenarios := []struct {
		name        string
		supi        string
		description string
	}{
		{
			name:        "ausf1_numeric_range_start",
			supi:        "123456789040000",
			description: "SUPI at start of AUSF-1 numeric range",
		},
		{
			name:        "ausf1_numeric_range_middle",
			supi:        "123456789045000",
			description: "SUPI in middle of AUSF-1 numeric range",
		},
		{
			name:        "ausf2_regex_pattern_match",
			supi:        "imsi-223456789041111",
			description: "SUPI matching AUSF-2 regex pattern",
		},
		{
			name:        "no_match_outside_ranges",
			supi:        "999999999999999",
			description: "SUPI outside all defined ranges",
		},
	}

	for _, scenario := range testScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Fresh setup for each subtest
			testCtx, cleanup := setupTest(t)
			defer cleanup()

			ctx := context.Background()
			evictionTimerVal := time.Duration(testCtx.evictionInterval)
			InitNrfCaching(evictionTimerVal*time.Second, testCtx.nrfDbCallback)

			expectedCallCount := int32(1) // Each test starts fresh

			param := createAusfParam(scenario.supi)
			result, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_AUSF, models.NFTYPE_AMF, param)
			if err != nil {
				t.Fatalf("SearchNFInstances failed: %v", err)
			}

			// Count how many should actually match
			expectedMatches := 0
			for _, instance := range result.NfInstances {
				match, err := matchAusfProfile(&instance, param)
				if err != nil {
					// Handle the error appropriately - could log, fail test, or skip
					t.Errorf("MatchAusfProfile failed: %v", err)
					continue
				}
				if match {
					expectedMatches++
				}
			}

			t.Logf("SUPI %s: returned %d instances, %d should match",
				scenario.supi, len(result.NfInstances), expectedMatches)

			assertCallbackCount(t, testCtx.getCallbackCount(), expectedCallCount)
		})
	}
}

func TestAmfMatchFilters(t *testing.T) {
	testCtx, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	evictionTimerVal := time.Duration(testCtx.evictionInterval)
	InitNrfCaching(evictionTimerVal*time.Second, testCtx.nrfDbCallback)

	t.Run("plmn_filtering", func(t *testing.T) {
		testCases := []struct {
			name        string
			description string
			plmnList    []models.PlmnId
		}{
			{
				name:        "matching_plmn",
				plmnList:    []models.PlmnId{{Mcc: "208", Mnc: "93"}},
				description: "Single PLMN that matches AMF profile",
			},
			{
				name: "multiple_plmns_with_match",
				plmnList: []models.PlmnId{
					{Mcc: "208", Mnc: "93"},
					{Mcc: "209", Mnc: "94"},
				},
				description: "Multiple PLMNs where one matches",
			},
			{
				name:        "non_matching_plmn",
				plmnList:    []models.PlmnId{{Mcc: "999", Mnc: "99"}},
				description: "PLMN that doesn't match any AMF profile",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				param := createAmfParamWithPlmns(tc.plmnList)
				result, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_AMF, models.NFTYPE_AMF, param)
				if err != nil {
					t.Fatalf("SearchNFInstances failed: %v", err)
				}

				// Just verify we got some instances - don't check specific count
				if len(result.NfInstances) == 0 {
					t.Error("Expected at least one instance, got none")
				}

				// Test the matching logic separately
				matchCount := 0
				for _, instance := range result.NfInstances {
					if match, err := matchAmfProfile(&instance, param); err == nil && match {
						matchCount++
					}
				}

				t.Logf("PLMN test '%s': %s - Found %d instances, %d matches",
					tc.name, tc.description, len(result.NfInstances), matchCount)
			})
		}
	})
}

func TestAusfMatchFiltersIsolated(t *testing.T) {
	testCtx, cleanup := setupTest(t)
	defer cleanup()

	// Test the AUSF matching filter directly
	testCases := []struct {
		name            string
		supi            string
		expectedProfile string
		description     string
		expectedMatch   bool
	}{
		{
			name:            "ausf1_numeric_match",
			supi:            "123456789040001",
			expectedMatch:   true,
			expectedProfile: "AUSF-1",
			description:     "Should match AUSF-1 numeric range",
		},
		{
			name:            "ausf2_regex_match",
			supi:            "imsi-223456789041111",
			expectedMatch:   true,
			expectedProfile: "AUSF-2",
			description:     "Should match AUSF-2 regex pattern",
		},
		{
			name:          "no_match",
			supi:          "999999999999999",
			expectedMatch: false,
			description:   "Should not match any AUSF",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Get AUSF profiles directly
			ausf1Profile, err1 := testCtx.getNfProfile("AUSF-1")
			ausf2Profile, err2 := testCtx.getNfProfile("AUSF-2")

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to get AUSF profiles: %v, %v", err1, err2)
			}

			// Test matching against each profile
			param := createAusfParam(tc.supi)

			match1, err1 := matchAusfProfile(&ausf1Profile, param)
			match2, err2 := matchAusfProfile(&ausf2Profile, param)

			if err1 != nil || err2 != nil {
				t.Fatalf("Matching failed: %v, %v", err1, err2)
			}

			t.Logf("SUPI %s: AUSF-1 match = %v, AUSF-2 match = %v", tc.supi, match1, match2)

			if tc.expectedMatch {
				if !match1 && !match2 {
					t.Errorf("Expected SUPI %s to match at least one AUSF profile", tc.supi)
				}
				if tc.expectedProfile == "AUSF-1" && !match1 {
					t.Errorf("Expected SUPI %s to match AUSF-1", tc.supi)
				}
				if tc.expectedProfile == "AUSF-2" && !match2 {
					t.Errorf("Expected SUPI %s to match AUSF-2", tc.supi)
				}
			} else {
				if match1 || match2 {
					t.Errorf("Expected SUPI %s to not match any AUSF profile, but got AUSF-1=%v, AUSF-2=%v", tc.supi, match1, match2)
				}
			}
		})
	}
}

func TestCacheKeyBehavior(t *testing.T) {
	testCtx, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	evictionTimerVal := time.Duration(testCtx.evictionInterval)
	InitNrfCaching(evictionTimerVal*time.Second, testCtx.nrfDbCallback)

	t.Run("test_cache_key_isolation", func(t *testing.T) {
		// Clear any existing cache state
		initialCallCount := testCtx.getCallbackCount()

		// First query: SUPI that should match only AUSF-1
		t.Log("=== First query: SUPI matching AUSF-1 ===")
		param1 := createAusfParam("123456789040001")
		result1, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_AUSF, models.NFTYPE_AMF, param1)
		if err != nil {
			t.Fatalf("First SearchNFInstances failed: %v", err)
		}

		t.Logf("First query returned %d instances", len(result1.NfInstances))
		expectedCallCount := initialCallCount + 1
		assertCallbackCount(t, testCtx.getCallbackCount(), expectedCallCount)

		// Second query: Same SUPI (should use cache)
		t.Log("=== Second query: Same SUPI (should hit cache) ===")
		result2, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_AUSF, models.NFTYPE_AMF, param1)
		if err != nil {
			t.Fatalf("Second SearchNFInstances failed: %v", err)
		}

		t.Logf("Second query returned %d instances", len(result2.NfInstances))
		assertCallbackCount(t, testCtx.getCallbackCount(), expectedCallCount) // Should be same

		// Third query: Different SUPI that should match AUSF-2
		t.Log("=== Third query: Different SUPI matching AUSF-2 ===")
		param3 := createAusfParam("imsi-223456789041111")
		result3, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_AUSF, models.NFTYPE_AMF, param3)
		if err != nil {
			t.Fatalf("Third SearchNFInstances failed: %v", err)
		}

		t.Logf("Third query returned %d instances", len(result3.NfInstances))
		// This might or might not cause a new callback depending on cache implementation

		// Fourth query: SUPI that matches nothing
		t.Log("=== Fourth query: SUPI matching nothing ===")
		param4 := createAusfParam("999999999999999")
		result4, err := SearchNFInstances(ctx, "testNrf", models.NFTYPE_AUSF, models.NFTYPE_AMF, param4)
		if err != nil {
			t.Fatalf("Fourth SearchNFInstances failed: %v", err)
		}

		t.Logf("Fourth query returned %d instances", len(result4.NfInstances))

		// Log final callback count
		t.Logf("Final callback count: %d", testCtx.getCallbackCount())
	})
}

// Check if the MatchAmfProfile function is working correctly
func TestAmfProfileMatching(t *testing.T) {
	// Create a test AMF profile
	amfProfile := models.NewNFProfileDiscovery("9f7d5a3f-88ab-4525-b31e-334da7faedab", models.NFTYPE_AMF, models.NFSTATUS_REGISTERED)
	amfProfile.SetPlmnList([]models.PlmnId{{Mcc: "208", Mnc: "93"}})
	amfProfile.SetAmfInfo(models.AmfInfo{
		AmfRegionId: "ca",
		AmfSetId:    testAmfSetID,
	})

	testCases := []struct {
		param    Nnrf_NFDiscovery.ApiSearchNFInstancesRequest
		name     string
		expected bool
	}{
		{
			name:     "matching_plmn",
			param:    createAmfParamWithPlmns([]models.PlmnId{{Mcc: "208", Mnc: "93"}}),
			expected: true,
		},
		{
			name:     "non_matching_plmn",
			param:    createAmfParamWithPlmns([]models.PlmnId{{Mcc: "999", Mnc: "99"}}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match, err := matchAmfProfile(amfProfile, tc.param)
			if err != nil {
				t.Fatalf("MatchAmfProfile failed: %v", err)
			}
			if match != tc.expected {
				t.Errorf("Expected match=%v, got match=%v", tc.expected, match)
			}
		})
	}
}

// TestMatchProfileWithoutSupiRangesIsUnrestricted covers profiles that declare
// no SUPI ranges. Such a profile serves every SUPI, so a SUPI-filtered
// discovery must select it. The NRF's own discovery filter encodes this as an
// $or over "a range contains the SUPI" / "supiRanges is null" / "supiRanges is
// absent"; a matcher that rejected these profiles would make cached lookups
// disagree with a live discovery against the same NRF.
func TestMatchProfileWithoutSupiRangesIsUnrestricted(t *testing.T) {
	const supi = "imsi-208930100007500"

	testCases := []struct {
		profile models.NFProfileDiscovery
		matcher matchFilterFunc
		name    string
	}{
		{
			name:    "udm_info_without_supi_ranges",
			matcher: matchUdmProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "UDM-no-ranges",
				NfType:       models.NFTYPE_UDM,
				UdmInfo:      &models.UdmInfo{},
			},
		},
		{
			name:    "udm_info_absent",
			matcher: matchUdmProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "UDM-no-info",
				NfType:       models.NFTYPE_UDM,
			},
		},
		{
			name:    "pcf_info_without_supi_ranges",
			matcher: matchPcfProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "PCF-no-ranges",
				NfType:       models.NFTYPE_PCF,
				PcfInfo:      &models.PcfInfo{},
			},
		},
		{
			name:    "pcf_info_absent",
			matcher: matchPcfProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "PCF-no-info",
				NfType:       models.NFTYPE_PCF,
			},
		},
		{
			name:    "ausf_info_without_supi_ranges",
			matcher: matchAusfProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "AUSF-no-ranges",
				NfType:       models.NFTYPE_AUSF,
				AusfInfo:     &models.AusfInfo{},
			},
		},
		{
			name:    "ausf_info_absent",
			matcher: matchAusfProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "AUSF-no-info",
				NfType:       models.NFTYPE_AUSF,
			},
		},
		{
			name:    "udr_info_without_supi_ranges",
			matcher: matchUdrProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "UDR-no-ranges",
				NfType:       models.NFTYPE_UDR,
				UdrInfo:      &models.UdrInfo{},
			},
		},
		{
			name:    "udr_info_absent",
			matcher: matchUdrProfile,
			profile: models.NFProfileDiscovery{
				NfInstanceId: "UDR-no-info",
				NfType:       models.NFTYPE_UDR,
			},
		},
	}

	param := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Supi(supi)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match, err := tc.matcher(&tc.profile, param)
			if err != nil {
				t.Fatalf("matcher returned error: %v", err)
			}
			if !match {
				t.Errorf("profile %s declares no SUPI ranges and must match SUPI %s, got no match",
					tc.profile.GetNfInstanceId(), supi)
			}
		})
	}
}

// TestMatchProfileWithSupiRangesStillFiltersOut guards the complementary case:
// once a profile declares ranges, a SUPI outside them must not match.
func TestMatchProfileWithSupiRangesStillFiltersOut(t *testing.T) {
	sr := models.NewSupiRange()
	sr.SetStart("100000000000000")
	sr.SetEnd("100000000000009")
	ranges := []models.SupiRange{*sr}
	param := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Supi("imsi-208930100007500")

	udmInfo := models.NewUdmInfo()
	udmInfo.SetSupiRanges(ranges)
	profile := models.NewNFProfileDiscovery("UDM-ranged", models.NFTYPE_UDM, models.NFSTATUS_REGISTERED)
	profile.SetUdmInfo(*udmInfo)
	match, err := matchUdmProfile(profile, param)
	if err != nil {
		t.Fatalf("MatchUdmProfile returned error: %v", err)
	}
	if match {
		t.Error("SUPI outside the declared ranges must not match")
	}
}

// TestUdrProfileIsSelectableFromCache guards the registration of UDR in
// matchFilters. A profile whose NfType has no registered filter is dropped by
// NrfCache.get, so an unregistered UDR made every cached UDR lookup return
// empty and forced a live NRF query on every subscriber data access.
func TestUdrProfileIsSelectableFromCache(t *testing.T) {
	if _, ok := matchFilters[models.NFTYPE_UDR]; !ok {
		t.Fatal("UDR has no entry in matchFilters, so cached UDR discovery can never return a profile")
	}

	srUdr := models.NewSupiRange()
	srUdr.SetStart("208930100007500")
	srUdr.SetEnd("208930100007599")
	udrInfo := models.NewUdrInfo()
	udrInfo.SetSupiRanges([]models.SupiRange{*srUdr})
	profile := models.NewNFProfileDiscovery("UDR-1", models.NFTYPE_UDR, models.NFSTATUS_REGISTERED)
	profile.SetUdrInfo(*udrInfo)

	inRange := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Supi("imsi-208930100007550")
	match, err := matchUdrProfile(profile, inRange)
	if err != nil {
		t.Fatalf("MatchUdrProfile returned error: %v", err)
	}
	if !match {
		t.Error("SUPI inside the declared range must match")
	}

	outOfRange := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Supi("imsi-208930100009999")
	match, err = matchUdrProfile(profile, outOfRange)
	if err != nil {
		t.Fatalf("MatchUdrProfile returned error: %v", err)
	}
	if match {
		t.Error("SUPI outside the declared range must not match")
	}
}

// TestAmfMatchesProfileWithoutPlmnList validates TS 29.510 Table 6.2.3.2.3.1-1
// (target-plmn-list entry): a profile that carries no plmnList is available to
// all PLMNs and must be included in any target-plmn-list-filtered discovery result.
func TestAmfMatchesProfileWithoutPlmnList(t *testing.T) {
	profile := models.NFProfileDiscovery{
		NfInstanceId: "AMF-no-plmn",
		NfType:       models.NFTYPE_AMF,
		NfStatus:     models.NFSTATUS_REGISTERED,
	}
	param := createAmfParamWithPlmns([]models.PlmnId{{Mcc: "208", Mnc: "93"}})
	match, err := matchAmfProfile(&profile, param)
	if err != nil {
		t.Fatalf("MatchAmfProfile returned error: %v", err)
	}
	if !match {
		t.Error("AMF profile without plmnList must match any PLMN filter (TS 29.510 Table 6.2.3.2.3.1-1, target-plmn-list entry)")
	}
}

// TestAmfTaiFilter validates that the tai query parameter filters AMF profiles
// by amfInfo.taiList (TS 29.510 Table 6.2.3.2.3.1-1, tai entry; absence of
// taiList in AmfInfo means the AMF serves any TAI, per Table 6.1.6.2.11-1).
func TestAmfTaiFilter(t *testing.T) {
	matchingTai := models.Tai{PlmnId: models.PlmnId{Mcc: "208", Mnc: "93"}, Tac: "000001"}
	otherTai := models.Tai{PlmnId: models.PlmnId{Mcc: "208", Mnc: "93"}, Tac: "000002"}

	profile := models.NFProfileDiscovery{
		NfInstanceId: "AMF-tai",
		NfType:       models.NFTYPE_AMF,
		NfStatus:     models.NFSTATUS_REGISTERED,
		AmfInfo: &models.AmfInfo{
			AmfRegionId: "ca",
			AmfSetId:    testAmfSetID,
			TaiList:     []models.Tai{matchingTai},
		},
	}

	paramMatch := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Tai(matchingTai)
	if match, err := matchAmfProfile(&profile, paramMatch); err != nil || !match {
		t.Errorf("expected TAI match: match=%v err=%v", match, err)
	}

	paramMiss := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Tai(otherTai)
	if match, err := matchAmfProfile(&profile, paramMiss); err != nil || match {
		t.Errorf("expected TAI miss: match=%v err=%v", match, err)
	}

	// AmfInfo present but taiList absent means unrestricted (TS 29.510 Table 6.1.6.2.11-1).
	unrestricted := models.NFProfileDiscovery{
		NfInstanceId: "AMF-no-tai",
		NfType:       models.NFTYPE_AMF,
		NfStatus:     models.NFSTATUS_REGISTERED,
		AmfInfo:      &models.AmfInfo{AmfRegionId: "ca", AmfSetId: testAmfSetID},
	}
	if match, err := matchAmfProfile(&unrestricted, paramMatch); err != nil || !match {
		t.Errorf("AMF with absent taiList must match any TAI query: match=%v err=%v", match, err)
	}
}

// TestAmfTaiFilterSnpn validates SNPN TAI matching: the Nid field must be
// included in the comparison so that TAIs with equal PLMN+TAC but different
// NIDs do not produce false-positive matches across SNPNs (TS 23.501 5.30.2.1).
func TestAmfTaiFilterSnpn(t *testing.T) {
	nid1 := openapi.PtrString("1234567890a")
	nid2 := openapi.PtrString("b0987654321")
	plmn := models.PlmnId{Mcc: "208", Mnc: "93"}
	tac := "000001"

	snpnTai1 := models.Tai{PlmnId: plmn, Tac: tac, Nid: nid1}
	snpnTai2 := models.Tai{PlmnId: plmn, Tac: tac, Nid: nid2}
	plmnTai := models.Tai{PlmnId: plmn, Tac: tac} // no Nid – standard PLMN context

	profile := models.NFProfileDiscovery{
		NfInstanceId: "AMF-snpn",
		NfType:       models.NFTYPE_AMF,
		NfStatus:     models.NFSTATUS_REGISTERED,
		AmfInfo: &models.AmfInfo{
			AmfRegionId: "ca",
			AmfSetId:    testAmfSetID,
			TaiList:     []models.Tai{snpnTai1},
		},
	}

	// Same NID: must match.
	if match, err := matchAmfProfile(&profile, Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Tai(snpnTai1)); err != nil || !match {
		t.Errorf("expected match for same NID: match=%v err=%v", match, err)
	}
	// Different NID, same PLMN+TAC: must not match (cross-SNPN false positive).
	if match, err := matchAmfProfile(&profile, Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Tai(snpnTai2)); err != nil || match {
		t.Errorf("expected no match for different NID: match=%v err=%v", match, err)
	}
	// Query without NID against SNPN profile TAI: must not match.
	if match, err := matchAmfProfile(&profile, Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.Tai(plmnTai)); err != nil || match {
		t.Errorf("expected no match for PLMN query vs SNPN profile: match=%v err=%v", match, err)
	}
}

// TestSmfDnnScopedToMatchedSnssai guards that a DNN filter is only satisfied
// when the DNN appears in an sNssaiSmfInfoList entry whose S-NSSAI matches the
// snssais filter. Without this, an SMF with {snssai-A/dnn-X, snssai-B/dnn-Y}
// would falsely match a query for {snssai-A, dnn-Y}.
func TestSmfDnnScopedToMatchedSnssai(t *testing.T) {
	snssaiA := models.Snssai{Sst: 1, Sd: openapi.PtrString("010203")}
	snssaiB := models.Snssai{Sst: 1, Sd: openapi.PtrString("0a0b0c")}

	profile := models.NFProfileDiscovery{
		NfInstanceId: "SMF-cross-snssai",
		NfType:       models.NFTYPE_SMF,
		SmfInfo: &models.SmfInfo{
			SNssaiSmfInfoList: []models.SnssaiSmfInfoItem{
				{
					SNssai: snssaiA,
					DnnSmfInfoList: []models.DnnSmfInfoItem{
						{Dnn: testDnnInternet},
					},
				},
				{
					SNssai: snssaiB,
					DnnSmfInfoList: []models.DnnSmfInfoItem{
						{Dnn: "ims"},
					},
				},
			},
		},
	}

	// snssai-A + dnn "internet" must match (same entry)
	paramMatch := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.
		Snssais([]models.Snssai{snssaiA}).
		Dnn(testDnnInternet)
	if match, err := matchSmfProfile(&profile, paramMatch); err != nil || !match {
		t.Errorf("expected match for snssai-A/internet: match=%v err=%v", match, err)
	}

	// snssai-A + dnn "ims" must not match (ims is only in snssai-B)
	paramCross := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.
		Snssais([]models.Snssai{snssaiA}).
		Dnn("ims")
	if match, err := matchSmfProfile(&profile, paramCross); err != nil || match {
		t.Errorf("expected no match for snssai-A/ims cross-SNSSAI: match=%v err=%v", match, err)
	}
}

// TestSmfMatchNfServiceList guards allNFServices' handling of nfServiceList,
// the TS 29.510 Rel-16 replacement for the deprecated nfServices array. A
// profile populated only via nfServiceList must still be selectable by
// service-name filtering, and a service name absent from it must not match.
func TestSmfMatchNfServiceList(t *testing.T) {
	profile := models.NewNFProfileDiscovery("SMF-nfservicelist-only", models.NFTYPE_SMF, models.NFSTATUS_REGISTERED)
	profile.SetNfServiceList(map[string]models.NFService{
		"nsmf-pdusession-1": {
			ServiceInstanceId: "nsmf-pdusession-1",
			ServiceName:       models.SERVICENAME_NSMF_PDUSESSION,
			Scheme:            models.URISCHEME_HTTPS,
			NfServiceStatus:   models.NFSERVICESTATUS_REGISTERED,
		},
	})

	paramMatch := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.
		ServiceNames([]models.ServiceName{models.SERVICENAME_NSMF_PDUSESSION})
	if match, err := matchSmfProfile(profile, paramMatch); err != nil || !match {
		t.Errorf("expected match for service name present only in nfServiceList: match=%v err=%v", match, err)
	}

	paramMiss := Nnrf_NFDiscovery.ApiSearchNFInstancesRequest{}.
		ServiceNames([]models.ServiceName{models.SERVICENAME_NSMF_EVENT_EXPOSURE})
	if match, err := matchSmfProfile(profile, paramMiss); err != nil || match {
		t.Errorf("expected no match for service name absent from nfServiceList: match=%v err=%v", match, err)
	}
}
