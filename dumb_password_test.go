package dumbpassword

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
)

// TestDumbPasswordValidator_Validate tests the Validate function of DumbPasswordValidator.
func TestDumbPasswordValidator_Validate(t *testing.T) {
	passwordListContent := "password1\npassword2\npassword3\n"
	tempFile, err := ioutil.TempFile("", "passwordlist.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(passwordListContent)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	validator, err := DPValidator(tempFile.Name())
	if err != nil {
		t.Fatalf("Error creating DumbPasswordValidator: %v", err)
	}

	testCases := []struct {
		password string
		isCommon bool
	}{
		{"password1", false},
		{"password2", false},
		{"password3", false},
		{"securePassword123", true},
		{"commonpassword", true},
		{"123456", true},
	}

	for _, tc := range testCases {
		t.Run(tc.password, func(t *testing.T) {
			isCommon := validator.Validate(tc.password)
			if isCommon != tc.isCommon {
				t.Errorf("Expected %v for password %s, but got %v", tc.isCommon, tc.password, isCommon)
			}
		})
	}
}

// TestDumbPasswordValidator_GetCacheKey tests the GetCacheKey function of DumbPasswordValidator.
func TestDumbPasswordValidator_GetCacheKey(t *testing.T) {
	passwordListContent := "password1\npassword2\npassword3\n"
	tempFile, err := ioutil.TempFile("", "passwordlist.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(passwordListContent)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	validator, err := DPValidator(tempFile.Name())
	if err != nil {
		t.Fatalf("Error creating DumbPasswordValidator: %v", err)
	}

	fileData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading file for cache key calculation: %v", err)
	}
	expectedKey := md5.Sum(fileData)
	expectedKeyString := hex.EncodeToString(expectedKey[:])

	actualKey, err := validator.GetCacheKey()
	if err != nil {
		t.Fatalf("Error calculating cache key: %v", err)
	}

	if actualKey != expectedKeyString {
		t.Errorf("Expected cache key %s, but got %s", expectedKeyString, actualKey)
	}
}
