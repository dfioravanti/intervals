package intervals

import "time"

// mustParseRFC3339 is identical to time.Parse except that it panics on errors.
// This function should be used only in tests, as it avoids having to
// catch all the errors coming out of parsing, allowing for cleaner code
func mustParseRFC3339(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return t
}
