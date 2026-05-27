package jsondiff

// findKey finds and return the object value that match key.
// It assumes to be on an opening curly bracket.
// The function expects a compact JSON input.
func findKey(json string, key string) string { _ = "STUB: not implemented"; return "" }

// skip semicolon

// skip comma

// findIndex finds and return the array value at the given index.
// It assumes to be on opening square bracket.
// The function expects a compact JSON input.
func findIndex(json string, idx int) string { _ = "STUB: not implemented"; return "" }

// skip comma
// next elem index

func squashValue(json string, i int) int { _ = "STUB: not implemented"; return 0 }

// true, null

// false

// string

// number

// array, object

// squashNumber reads b until it encounter a character
// than can follow a properly formatted number.
func squashNumber(json string, i int) int { _ = "STUB: not implemented"; return 0 }

// squashString reads b until it encounter a non-escaped
// double quote, which indicate the sep of the string.
func squashString(json string, i int) int {
	_ = "STUB: not implemented"
	// assume to be on opening quote
	return 0
}

// move to closing quote

// note: taken from https://github.com/tidwall/gjson
func squashObjectOrArray(json string, i int) int { _ = "STUB: not implemented"; return 0 }

// compact removes insignificant space characters from the
// input JSON byte slice and returns the compacted result.
func compact(json []byte) []byte { _ = "STUB: not implemented"; return nil }

// compactInPlace is similar to compact, but it reuses the input
// JSON buffer to avoid allocations.
func compactInPlace(json []byte) []byte { _ = "STUB: not implemented"; return nil }

func _compact(src, dst []byte) []byte { _ = "STUB: not implemented"; return nil }
