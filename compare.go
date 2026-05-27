package jsondiff

// Compare compares the JSON representations of the
// given values and returns the differences relative
// to the former as a list of JSON Patch operations.
func Compare(source, target interface{}, opts ...Option) (Patch, error) {
	_ = "STUB: not implemented"
	return *new(Patch), nil
}

// CompareJSON compares the given JSON documents and
// returns the differences relative to the former as
// a list of JSON Patch operations.
func CompareJSON(source, target []byte, opts ...Option) (Patch, error) {
	_ = "STUB: not implemented"
	return *new(Patch), nil
}

// CompareWithoutMarshal is similar to Compare, but it assumes
// that the given interface values consists only of primitives
// Go types that are recognized by the json.Unmarshal function,
// and therefore does not marshal/unmarshal before comparison.
func CompareWithoutMarshal(source, target interface{}, opts ...Option) (patch Patch, err error) {
	_ = "STUB: not implemented"
	return *new(Patch), nil
}

func compare(d *Differ, src, tgt interface{}) (Patch, error) {
	_ = "STUB: not implemented"
	return *new(Patch), nil
}

func compareJSON(d *Differ, src, tgt []byte, unmarshal unmarshalFunc) (Patch, error) {
	_ = "STUB: not implemented"
	return *new(Patch), nil
}

// marshalUnmarshal returns the result of unmarshaling
// the JSON representation of the given interface value.
func marshalUnmarshal(v any, opts options) (interface{}, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
