package util

import "regexp"

func IsValidEmail(email string) bool {
	// Define a regular expression pattern for a valid email address
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Use the MatchString method to check if the email matches the pattern
	return regex.MatchString(email)
}
