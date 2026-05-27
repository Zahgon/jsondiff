package jsondiff

// JSON Patch operation types.
// These are defined in RFC 6902 section 4.
// https://datatracker.ietf.org/doc/html/rfc6902#section-4
const (
	OperationAdd     = "add"
	OperationReplace = "replace"
	OperationRemove  = "remove"
	OperationMove    = "move"
	OperationCopy    = "copy"
	OperationTest    = "test"
)

const (
	fromFieldLen  = 10 // ,"from":""
	valueFieldLen = 9  // ,"value":
	opBaseLen     = 19 // {"op":"","path":""}
)

// null represents a JSON null value.
type null struct{}

// Operation represents a single JSON Patch (RFC6902) operation.
type Operation struct {
	Value    interface{} `json:"value,omitempty"`
	OldValue interface{} `json:"-"`
	Type     string      `json:"op"`
	From     string      `json:"from,omitempty"`
	Path     string      `json:"path"`
	valueLen int
}

// MarshalJSON implements the json.Marshaler interface.
func (null) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// String implements the fmt.Stringer interface.
		nil
}

func (o Operation) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implements the json.Marshaler interface.
func (o Operation) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Generic check that works for nil
// and typed nil interface values.

// jsonLength returns the length in bytes that the
// operation would occupy when marshaled to JSON.
func (o Operation) jsonLength() int { _ = "STUB: not implemented"; return 0 }

func (o Operation) hasFrom() bool { _ = "STUB: not implemented"; return false }

func (o Operation) marshalWithValue() bool { _ = "STUB: not implemented"; return false }

func (p *Patch) remove(idx int) Patch { _ = "STUB: not implemented"; return *new(Patch) }

func (p *Patch) append(typ string, from, path string, src, tgt interface{}, vl int) Patch {
	_ = "STUB: not implemented"
	return *new(Patch)
}

func (p *Patch) insert(pos int, typ string, from, path string, src, tgt interface{}, vl int) Patch {
	_ = "STUB: not implemented"
	return *new(Patch)
}

func (p *Patch) jsonLength() int { _ = "STUB: not implemented"; return 0 }

// Count comma-separators if the patch
// has more than one operation.

// String implements the fmt.Stringer interface.
func (p *Patch) String() string { _ = "STUB: not implemented"; return "" }
