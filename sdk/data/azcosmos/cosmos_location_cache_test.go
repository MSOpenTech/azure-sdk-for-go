// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azcosmos

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"
)

var defaultEndpt *url.URL
var loc1Endpt *url.URL
var loc2Endpt *url.URL
var loc3Endpt *url.URL
var loc4Endpt *url.URL
var writeEndpts []url.URL
var readEndpts []url.URL
var endptsByLoc map[string]url.URL
var loc1 AcctRegion
var loc2 AcctRegion
var loc3 AcctRegion
var loc4 AcctRegion
var prefLocs []string

func TestMain(m *testing.M) {
	var err error
	defaultEndpt, err = url.Parse("https://default.documents.azure.com")
	if err != nil {
		fmt.Println("Unable to parse default endpoint URI")
		os.Exit(1)
	}
	loc1Endpt, err = url.Parse("https://location1.documents.azure.com")
	if err != nil {
		fmt.Println("Unable to parse location1 endpoint URI")
		os.Exit(1)
	}
	loc2Endpt, err = url.Parse("https://location2.documents.azure.com")
	if err != nil {
		fmt.Println("Unable to parse location2 endpoint URI")
		os.Exit(1)
	}
	loc3Endpt, err = url.Parse("https://location3.documents.azure.com")
	if err != nil {
		fmt.Println("Unable to parse location3 endpoint URI")
		os.Exit(1)
	}
	loc4Endpt, err = url.Parse("https://location4.documents.azure.com")
	if err != nil {
		fmt.Println("Unable to parse location4 endpoint URI")
		os.Exit(1)
	}

	loc1 = AcctRegion{name: "location1", endpoint: loc1Endpt.String()}
	loc2 = AcctRegion{name: "location2", endpoint: loc2Endpt.String()}
	loc3 = AcctRegion{name: "location3", endpoint: loc3Endpt.String()}
	loc4 = AcctRegion{name: "location4", endpoint: loc4Endpt.String()}

	writeEndpts = []url.URL{*loc1Endpt, *loc2Endpt, *loc3Endpt}
	readEndpts = []url.URL{*loc1Endpt, *loc2Endpt, *loc4Endpt}
	endptsByLoc = map[string]url.URL{"location1": *loc1Endpt, "location2": *loc2Endpt, "location3": *loc3Endpt, "location4": *loc4Endpt}

	prefLocs = make([]string, 0)

	status := m.Run()
	os.Exit(status)
}

func CreateDbAcct(useMultipleWriteLocations bool, enforceSingleMasterWriteLoc bool) AcctProperties {
	writeRegions := []AcctRegion{loc1, loc2, loc3}
	if !useMultipleWriteLocations && enforceSingleMasterWriteLoc {
		writeRegions = []AcctRegion{loc1}
	}
	readRegions := []AcctRegion{loc1, loc2, loc4}
	return AcctProperties{writeRegions: writeRegions, readRegions: readRegions, enableMultipleWriteLocations: useMultipleWriteLocations}
}

func ResetLocationCache() *LocationCache {
	lc := NewLocationCache(prefLocs, *defaultEndpt)
	lc.enableEndptDiscovery = true
	lc.connLimit = 10
	return lc
}

func TestMarkEndptUnavailable(t *testing.T) {
	lc := ResetLocationCache()
	var firstCheckTime time.Time
	// mark endpoint unavailable for first time
	err := lc.MarkEndptUnavailableForRead(*loc1Endpt)
	if err != nil {
		t.Fatalf("Received error marking endpoint unavailable: %s", err.Error())
	}
	if info, ok := lc.locationUnavailabilityInfoMap[*loc1Endpt]; ok {
		var zeroTime time.Time
		if firstCheckTime = info.lastCheckTime; firstCheckTime == zeroTime {
			t.Errorf("Expected lastCheckTime to be set, but was not")
		}
		if info.unavailableOps != read {
			t.Errorf("Expected unavailableOps to be 1 (read-only), but was %d", info.unavailableOps)
		}
	} else {
		t.Errorf("Expected locationUnavailabilityInfoMap to contain %s, but it did not", loc1Endpt.String())
	}
	// mark endpoint unavailable for second time
	time.Sleep(100 * time.Millisecond)
	err = lc.MarkEndptUnavailableForWrite(*loc1Endpt)
	if err != nil {
		t.Fatalf("Received error marking endpoint unavailable: %s", err.Error())
	}
	if info, ok := lc.locationUnavailabilityInfoMap[*loc1Endpt]; ok {
		var zeroTime time.Time
		if info.lastCheckTime == zeroTime || info.lastCheckTime == firstCheckTime {
			t.Errorf("Expected lastCheckTime to be updated, but was not. First check time: %s, last check time: %s", firstCheckTime, info.lastCheckTime)
		}
		if info.unavailableOps != all {
			t.Errorf("Expected unavailableOps to be 3 (read+write), but was %d", info.unavailableOps)
		}
	} else {
		t.Errorf("Expected locationUnavailabilityInfoMap to contain %s, but it did not", loc1Endpt.String())
	}
}

func TestRefreshStaleEndpts(t *testing.T) {
	lc := ResetLocationCache()
	// mark endpoint unavailable for first time
	err := lc.MarkEndptUnavailableForRead(*loc1Endpt)
	if err != nil {
		t.Fatalf("Received error marking endpoint unavailable: %s", err.Error())
	}
	if info, ok := lc.locationUnavailabilityInfoMap[*loc1Endpt]; ok {
		info.lastCheckTime = time.Now().Add(-1*DefaultExpirationTime - 1*time.Second)
		lc.locationUnavailabilityInfoMap[*loc1Endpt] = info
	} else {
		t.Errorf("Expected locationUnavailabilityInfoMap to contain %s, but it did not", loc1Endpt.String())
	}
	// refresh stale endpoints, since time since last check is greater default expiration time
	lc.RefreshStaleEndpts()
	if len(lc.locationUnavailabilityInfoMap) != 0 {
		t.Errorf("Expected locationUnavailabilityInfoMap to be empty, but it was not")
	}
}

func TestIsEndptUnavailable(t *testing.T) {
	lc := ResetLocationCache()
	err := lc.MarkEndptUnavailableForRead(*loc1Endpt)
	if err != nil {
		t.Fatalf("Received error marking endpoint unavailable: %s", err.Error())
	}
	err = lc.MarkEndptUnavailableForWrite(*loc2Endpt)
	if err != nil {
		t.Fatalf("Received error marking endpoint unavailable: %s", err.Error())
	}

	if lc.IsEndptUnavailable(*loc1Endpt, none) {
		t.Errorf("Expected IsEndptUnavailable to return false, but it returned true for ops = none")
	}
	if lc.IsEndptUnavailable(*loc1Endpt, write) {
		t.Errorf("Expected IsEndptUnavailable to return false, but it returned true for ops = write when region is unavailable for read")
	}
	if lc.IsEndptUnavailable(*loc3Endpt, all) {
		t.Errorf("Expected IsEndptUnavailable to return false, but it returned true for an available region")
	}
	if !lc.IsEndptUnavailable(*loc1Endpt, read) {
		t.Errorf("Expected IsEndptUnavailable to return true, but it returned false for ops = read when region is unavailable for read")
	}

	if info, ok := lc.locationUnavailabilityInfoMap[*loc1Endpt]; ok {
		info.lastCheckTime = time.Now().Add(-1*DefaultExpirationTime - 1*time.Second)
		lc.locationUnavailabilityInfoMap[*loc1Endpt] = info
	} else {
		t.Errorf("Expected locationUnavailabilityInfoMap to contain %s, but it did not", loc1Endpt.String())
	}

	if lc.IsEndptUnavailable(*loc1Endpt, read) {
		t.Errorf("Expected IsEndptUnavailable to return false, but it returned true stale unavailability")
	}
}

func TestGetLocation(t *testing.T) {
	lc := ResetLocationCache()
	dbAcct := CreateDbAcct(lc.enableMultipleWriteLocations, false)
	err := lc.DbAcctRead(dbAcct) // requires unit test of update
	if err != nil {
		t.Fatalf("Received error Reading DB account: %s", err.Error())
	}
	if dbAcct.writeRegions == nil || len(dbAcct.writeRegions) == 0 {
		t.Fatal("Write Regions are empty")
	}
	actual := lc.GetLocation(*defaultEndpt)
	if actual == "" {
		t.Errorf("Expected GetLocation to return a valid location when provided the default endpoint, but it did not")
	}
	for _, region := range dbAcct.writeRegions {
		url, err := url.Parse(region.endpoint)
		if err != nil {
			t.Errorf("Failed to parse endpoint %s, %s", region.endpoint, err)
			continue
		}
		expected, actual := region.name, lc.GetLocation(*url)
		if expected != actual {
			t.Errorf("Expected GetLocation to return Write Region %s, but was %s", expected, actual)
		}
	}

	for _, region := range dbAcct.readRegions {
		url, err := url.Parse(region.endpoint)
		if err != nil {
			t.Errorf("Failed to parse endpoint %s, %s", region.endpoint, err)
			continue
		}
		expected, actual := region.name, lc.GetLocation(*url)
		if expected != actual {
			t.Errorf("Expected GetLocation to return Read Region %s, but was %s", expected, actual)
		}
	}
}

func TestGetEndptsByLocation(t *testing.T) {
	locs := []AcctRegion{loc1, loc2, loc3, loc4}
	newEndptsByLoc, parsedLocs, err := GetEndptsByLocation(locs)
	if err != nil {
		t.Fatalf("Received error getting endpoints by location: %s", err.Error())
	}
	if len(newEndptsByLoc) != len(endptsByLoc) {
		t.Errorf("Expected %d endpoints, but got %d", len(endptsByLoc), len(newEndptsByLoc))
	}
	for loc, endpt := range endptsByLoc {
		if newEndpt, ok := newEndptsByLoc[loc]; ok {
			if newEndpt != endpt {
				t.Errorf("Expected endpoint %s for location %s, but was %s", endpt.String(), loc, newEndpt.String())
			}
		} else {
			t.Errorf("Expected newEndptsByLoc to contain location %s, but it did not", loc)
		}
	}

	if len(parsedLocs) != len(locs) {
		t.Errorf("Expected parsedLocs to contain %d locations, but it contained %d", len(locs), len(parsedLocs))
	}
	// may need to fix this, believe that maps are unordered which could cause this to fail due to ordering
	for i, loc := range locs {
		if parsedLocs[i] != loc.name {
			t.Errorf("Expected parsedLocs to contain location %s, but it did not", loc.name)
		}
	}
}

func TestGetPrefAvailableEndpts(t *testing.T) {
	lc := ResetLocationCache()
	lc.enableMultipleWriteLocations = true
	lc.useMultipleWriteLocations = true
	dbAcct := CreateDbAcct(lc.enableMultipleWriteLocations, false)
	// will set write locations to loc1, loc2, loc3
	err := lc.DbAcctRead(dbAcct)
	if err != nil {
		t.Fatalf("Received error Reading DB account: %s", err.Error())
	}
	// marks loc1 unavailable, which will put it last in the preferred available endpoint list
	lc.MarkEndptUnavailableForWrite(*loc1Endpt)
	// loc1: unavailable, loc2: available, loc5: non-existent
	lc.locationInfo.prefLocations = []string{loc1.name, loc2.name, "location5"}
	prefWriteEndpts := lc.GetPrefAvailableEndpts(lc.locationInfo.availWriteEndptsByLocation, lc.locationInfo.availWriteLocations, write, lc.defaultEndpt)
	// loc2: preferred + available, default: fallback endpoint, loc1: unavailable + preferred
	expectedWriteEndpts := []*url.URL{loc2Endpt, defaultEndpt, loc1Endpt}

	for i, endpt := range expectedWriteEndpts {
		if endpt.String() != prefWriteEndpts[i].String() {
			t.Errorf("Expected endpoint %s, but was %s", endpt.String(), prefWriteEndpts[i].String())
		}
	}
}

func TestReadEndpts(t *testing.T) {
	lc := ResetLocationCache()
	lc.locationInfo.prefLocations = []string{loc1.name, loc2.name, loc3.name, loc4.name}
	dbAcct := CreateDbAcct(lc.enableMultipleWriteLocations, false)
	err := lc.DbAcctRead(dbAcct)
	if err != nil {
		t.Fatalf("Received error Reading DB account: %s", err.Error())
	}

	lc.lastUpdateTime = time.Now().Add(-1*DefaultExpirationTime - 1*time.Second)
	expectedReadEndpts := []*url.URL{loc2Endpt, loc4Endpt, loc1Endpt}
	actualReadEndpts, err := lc.ReadEndpts()
	if err != nil {
		t.Fatalf("Received error getting read endpoints: %s", err.Error())
	}
	if len(expectedReadEndpts) != len(actualReadEndpts) {
		t.Errorf("Expected %d read endpoints, but got %d", len(expectedReadEndpts), len(actualReadEndpts))
	} else {
		for i, endpt := range expectedReadEndpts {
			if endpt.String() != actualReadEndpts[i].String() {
				t.Errorf("Expected endpoint %s, but was %s", endpt.String(), actualReadEndpts[i].String())
			}
		}
	}

	lc.lastUpdateTime = time.Now().Add(-1*DefaultExpirationTime - 1*time.Second)
	lc.MarkEndptUnavailableForRead(*loc2Endpt)
	expectedReadEndpts = []*url.URL{loc4Endpt, loc1Endpt, loc2Endpt}
	actualReadEndpts, err = lc.ReadEndpts()
	if err != nil {
		t.Fatalf("Received error getting read endpoints: %s", err.Error())
	}
	if len(expectedReadEndpts) != len(actualReadEndpts) {
		t.Errorf("Expected %d read endpoints, but got %d", len(expectedReadEndpts), len(actualReadEndpts))
	} else {
		for i, endpt := range expectedReadEndpts {
			if endpt.String() != actualReadEndpts[i].String() {
				t.Errorf("Expected endpoint %s, but was %s", endpt.String(), actualReadEndpts[i].String())
			}
		}
	}

}

func TestWriteEndpts(t *testing.T) {
	lc := ResetLocationCache()
	lc.enableMultipleWriteLocations = true
	lc.useMultipleWriteLocations = true
	lc.locationInfo.prefLocations = []string{loc1.name, loc2.name, loc3.name, loc4.name}
	dbAcct := CreateDbAcct(lc.enableMultipleWriteLocations, false)
	err := lc.DbAcctRead(dbAcct)
	if err != nil {
		t.Fatalf("Received error Reading DB account: %s", err.Error())
	}

	lc.lastUpdateTime = time.Now().Add(-1*DefaultExpirationTime - 1*time.Second)
	expectedWriteEndpts := []*url.URL{loc1Endpt, loc2Endpt, loc3Endpt, defaultEndpt}
	actualWriteEndpts, err := lc.WriteEndpts()
	if err != nil {
		t.Fatalf("Received error getting write endpoints: %s", err.Error())
	}
	if len(expectedWriteEndpts) != len(actualWriteEndpts) {
		t.Errorf("Expected %d write endpoints, but got %d", len(expectedWriteEndpts), len(actualWriteEndpts))
	} else {
		for i, endpt := range expectedWriteEndpts {
			if endpt.String() != actualWriteEndpts[i].String() {
				t.Errorf("Expected endpoint %s, but was %s", endpt.String(), actualWriteEndpts[i].String())
			}
		}
	}

	lc.lastUpdateTime = time.Now().Add(-1*DefaultExpirationTime - 1*time.Second)
	lc.MarkEndptUnavailableForWrite(*loc1Endpt)
	expectedWriteEndpts = []*url.URL{loc2Endpt, loc3Endpt, defaultEndpt, loc1Endpt}
	actualWriteEndpts, err = lc.WriteEndpts()
	if err != nil {
		t.Fatalf("Received error getting write endpoints: %s", err.Error())
	}
	if len(expectedWriteEndpts) != len(actualWriteEndpts) {
		t.Errorf("Expected %d write endpoints, but got %d", len(expectedWriteEndpts), len(actualWriteEndpts))
	} else {
		for i, endpt := range expectedWriteEndpts {
			if endpt.String() != actualWriteEndpts[i].String() {
				t.Errorf("Expected endpoint %s, but was %s", endpt.String(), actualWriteEndpts[i].String())
			}
		}
	}
}
