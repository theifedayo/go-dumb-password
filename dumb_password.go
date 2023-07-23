package dumbpassword

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"sync"
)

// DumbPasswordValidator is a struct representing the Dumb Password Validator.
type DumbPasswordValidator struct {
	passwordListPath string
	commonPasswords  []string
	mutex            sync.Mutex
}

// DPValidator creates a new instance of DumbPasswordValidator.
func DPValidator(passwordListPath string) (*DumbPasswordValidator, error) {
	commonPasswords, err := loadPasswordList(passwordListPath)
	if err != nil {
		return nil, err
	}

	return &DumbPasswordValidator{
		passwordListPath: passwordListPath,
		commonPasswords:  commonPasswords,
	}, nil
}

// Validate checks if the given password is common (dumb) and returns true if it's not common.
func (v *DumbPasswordValidator) Validate(password string) bool {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	lowercasePassword := strings.ToLower(password)
	for _, commonPassword := range v.commonPasswords {
		if lowercasePassword == commonPassword {
			return false
		}
	}
	return true
}

// loadPasswordList loads the list of common passwords from the file and returns it as a slice of strings.
func loadPasswordList(passwordListPath string) ([]string, error) {
	data, err := ioutil.ReadFile(passwordListPath)
	if err != nil {
		return nil, err
	}
	passwords := strings.Split(string(data), "\n")
	var commonPasswords []string
	for _, password := range passwords {
		commonPasswords = append(commonPasswords, strings.ToLower(password))
	}
	return commonPasswords, nil
}

// GetCacheKey calculates the cache key based on the MD5 hash of the password list file content.
func (v *DumbPasswordValidator) GetCacheKey() (string, error) {
	data, err := ioutil.ReadFile(v.passwordListPath)
	if err != nil {
		return "", err
	}
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:]), nil
}
