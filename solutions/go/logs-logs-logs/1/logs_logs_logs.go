package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var c string
	app := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
	for _, v := range log {
		if v == '❗' {
			c = app['❗']
			break
		}
		if v == '🔍' {
			c = app['🔍']
			break
		}
		if v == '☀' {
			c = app['☀']
			break
		} else {
			c = "default"
		}
	}

	return c
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var c string
	for _, v := range log {
		if v != oldRune {
			c += string(v)
		} else {
			c += string(newRune)
		}
	}
	return c
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) > limit {
		return false
	} else {
		return true
	}
}
