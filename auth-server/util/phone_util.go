package util

import "regexp"

func IsValidPhoneNumber(phoneNumber string) bool {
	// Define a regular expression pattern for an Indonesian phone number
	pattern := `^(?:\+62|0)[2-9]\d{7,12}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Use the MatchString method to check if the phone number matches the pattern
	return regex.MatchString(phoneNumber)
}
