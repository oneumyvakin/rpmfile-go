package rpmfile

import (
	"testing"
	"reflect"
)

func TestExtract_bin(t *testing.T) {
	var rpm Rpm_file
	store := []byte{0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x70}

	result, err := rpm.extract_bin(store, int32(0), int32(1))
	if err != nil {
		t.Fatalf("Error not nil %s", err)
	}
	if result[0] != 0x60 {
		t.Fatalf("Result %s not equal to 0x60", result[0])
	}

	result, err = rpm.extract_bin(store, int32(10), int32(1))
	if err != nil {
		t.Fatalf("Error not nil %s", err)
	}
	if result[0] != 0x70 {
		t.Fatalf("Result %s not equal to 0x70", result[0])
	}

	result, err = rpm.extract_bin(store, int32(0), int32(len(store)))
	if err != nil {
		t.Fatalf("Error not nil %s", err)
	}
	if !reflect.DeepEqual(result, store) {
		t.Fatalf("Result %s not equal to %s", result, store)
	}
}