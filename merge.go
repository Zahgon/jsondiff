package jsondiff

// MergePatch returns a JSON Merge Patch (RFC 7386)
// of the differences between the JSON representations
// of the given values.
func MergePatch(src, tgt interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MergePatchJSON compares the given JSON documents
// and returns the differences relative to the former
// as a JSON Merge Patch (RFC 7386)
func MergePatchJSON(src, tgt []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func mergePatch(src, tgt interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// If the target is not of the same type as the source,
// or both are not objects, the patch replaces the entire
// source with the target.
// https://datatracker.ietf.org/doc/html/rfc7386#section-2

// Null values in the merge patch are given
// special meaning to indicate the removal
// of existing values in the target.
// https://datatracker.ietf.org/doc/html/rfc7386#section-1
