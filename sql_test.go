// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"reflect"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	stringTest := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	badTypeTest := 6
	invalidTest := "f47ac10b-58cc-0372-8567-0e02b2c3d4"

	byteTest := make([]byte, 16)
	byteTestUUID := Must(Parse(stringTest))
	copy(byteTest, byteTestUUID[:])

	// sunny day tests

	var uuid UUID
	err := (&uuid).Scan(stringTest)
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan([]byte(stringTest))
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan(byteTest)
	if err != nil {
		t.Fatal(err)
	}

	// bad type tests

	err = (&uuid).Scan(badTypeTest)
	if err == nil {
		t.Error("int correctly parsed and shouldn't have")
	}
	if !strings.Contains(err.Error(), "unable to scan type") {
		t.Error("attempting to parse an int returned an incorrect error message")
	}

	// invalid/incomplete uuids

	err = (&uuid).Scan(invalidTest)
	if err == nil {
		t.Error("invalid uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid UUID returned an incorrect error message")
	}

	err = (&uuid).Scan(byteTest[:len(byteTest)-2])
	if err == nil {
		t.Error("invalid byte uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid byte UUID returned an incorrect error message")
	}

	// empty tests

	uuid = UUID{}
	var emptySlice []byte
	err = (&uuid).Scan(emptySlice)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty byte slice")
		}
	}

	uuid = UUID{}
	var emptyString string
	err = (&uuid).Scan(emptyString)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty byte slice")
		}
	}

	uuid = UUID{}
	err = (&uuid).Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning nil")
		}
	}
}

func TestValue(t *testing.T) {
	binTest := []byte{0xf4, 0x7a, 0xc1, 0x0b, 0x58, 0xcc, 0x03, 0x72, 0x85, 0x67, 0x0e, 0x02, 0xb2, 0xc3, 0xd4, 0x79}

	uuid, _ := FromBytes(binTest)

	val, _ := uuid.Value()
	if !reflect.DeepEqual(val, binTest) {
		t.Error("Value() did not return expected string")
	}
}
