package jsondiff

import (
	"errors"
)

var (
	// ErrNonReversible is returned when a non-reversible
	// operation (not preceded by a test) is found in a patch.
	ErrNonReversible = errors.New("non-reversible operation")

	// ErrAmbiguousCopyOp is returned to signal that a copy
	// operation cannot be reversed as it is ambiguous.
	ErrAmbiguousCopyOp = errors.New("copy operation is ambiguous")
)

// A ErrTestPointer is returned by [Patch.Reverse] when the
// pointer of an operator does not match the preceding test
// operation pointer.
type ErrTestPointer struct {
	Op string
}

func (e ErrTestPointer) Error() string { _ = "STUB: not implemented"; return "" }

// Patch represents a series of JSON Patch operations.
type Patch []Operation

// Invert returns a patch that undo the modifications
// represented by this patch.
func (p Patch) Invert() (Patch, error) { _ = "STUB: not implemented"; return *new(Patch), nil }

func (p Patch) prevOp(i int) (*Operation, error) { _ = "STUB: not implemented"; return nil, nil }

func (p Patch) invertOp(i int) ([]Operation, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
