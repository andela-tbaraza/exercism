// Package leap implements a function to check if a year is a leap year
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 3

// IsLeapYear returns true if a year is a leap year and false otherwise
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 != 0 || year%400 == 0 {
			return true
		}
		return false
	}
	return false
}
