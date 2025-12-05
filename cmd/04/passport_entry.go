package main

import "strings"

var REQUIRED_ATTRIBUTES = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid", // Optional
}

// expects an array of "key:val" strings
func parsePassportEntry(entryParts []string) map[string]string {
	passport := make(map[string]string)

	for _, part := range entryParts {
		part := strings.TrimSpace(part)
		keyVal := strings.SplitN(part, ":", 2)
		passport[keyVal[0]] = keyVal[1]
	}

	return passport
}

func validatePassportEntry(passport map[string]string) bool {
	for _, attr := range REQUIRED_ATTRIBUTES {
		if _, ok := passport[attr]; !ok {
			return false
		}
	}

	return true
}
