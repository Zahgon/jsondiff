package jsondiff

type invalidJSONTypeError struct {
	t any
}

// jsonValueType represents the type of JSON value.
// It follows the types of values stored by json.Unmarshal
// in interface values.
type jsonValueType uint

const (
	jsonInvalid jsonValueType = iota
	jsonNull
	jsonString
	jsonBoolean
	jsonNumberFloat
	jsonNumberString
	jsonArray
	jsonObject
)

// jsonTypeSwitch returns the JSON type of the value
// held by the interface using a type switch statement.
func jsonTypeSwitch(i interface{}) jsonValueType {
	_ = "STUB: not implemented"
	return *new(jsonValueType)
}

// areComparable returns whether the interface values
// i1 and i2 can be compared. The values are comparable
// only if they are both non-nil and share the same kind.
func areComparable(i1, i2 interface{}) bool { _ = "STUB: not implemented"; return false }

func deepEqual(src, tgt interface{}) bool { _ = "STUB: not implemented"; return false }

// Fast path.

func deepEqualValue(src, tgt interface{}) bool { _ = "STUB: not implemented"; return false }

// Key not found in target.

var jsonTypeNames = []string{
	jsonInvalid:      "Invalid",
	jsonBoolean:      "Boolean",
	jsonNumberFloat:  "Number",
	jsonNumberString: "json.Number",
	jsonString:       "String",
	jsonNull:         "Null",
	jsonObject:       "Object",
	jsonArray:        "Array",
}

// String implements fmt.Stringer for jsonValueType.
func (t jsonValueType) String() string { _ = "STUB: not implemented"; return "" }
