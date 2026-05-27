package jsondiff

// apply applies the patch to the given source document.
// If valid is true, the document is validated prior to
// the application of the patch.
//
// Note for readers, this method will **NEVER** be exported,
// as it is only used for internal tests, and is feature-wise
// out of scope of the project.
// See https://github.com/wI2L/jsondiff/issues/28#issuecomment-2360883098
func (p Patch) apply(src []byte, valid bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make a copy of the source document which
// will receive the patch mutations.

// First fetch the value from the source path,
// and then add it to the destination path.

// If the operation is a move, delete the
// source value before adding it at its new
// position, to preserve array index position.

// bail out to interpret error

// Special case when the path inserts an
// element at the end of an array, we revert
// the array and test the first element.

func replace(tgt []byte, path string, val interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func add(tgt []byte, path string, val interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unsupported by the sjson package.
		// Since an empty path represent the root
		// document, we can simply marshal the value
		// and return it as-is.
		nil
}

// If we're dealing with an array indices, we want to
// "insert" the element instead of replacing it.
// We insert a null value manually where the new element
// is supposed to be (before the current element), and
// finally replace the placeholder with the new value.

func isArrayIndex(path string) bool { _ = "STUB: not implemented"; return false }

// dotPath converts the given JSON Pointer string to the
// dot-path notation used by sjson package.
// The source document is required in order to distinguish
// numeric object keys from array indices
func toDotPath(path string, src []byte) (string, error) {
	_ = "STUB: not implemented"

	// @this returns the current element.
	// It is used to retrieve the root element.
	return "", nil
}

// The fragment starts with a digit, which
// indicate that it might be a number.

// The number is valid, but it could either be an
// array indices or an object key.
// Since the JSON Pointer RFC does not differentiate
// between the two, we have to look up the value to
// know what we're dealing with.

// Write array indices as-is.

// Force the number as an object key, by
// preceding it with a colon character.

// If the last fragment is the "-" character,
// it indicates that the value is a nonexistent
// element to append to the array.

// Add separator character
