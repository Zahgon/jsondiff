package jsondiff

import (
	"errors"
	"strings"
)

const (
	separator    = '/'
	escapeSlash  = "~1"
	escapeTilde  = "~0"
	emptyPointer = ""
)

var (
	// rfc6901Escaper is a replacer that escapes a JSON Pointer string
	// in compliance with the JavaScript Object Notation Pointer syntax.
	// https://tools.ietf.org/html/rfc6901
	rfc6901Escaper = strings.NewReplacer("~", escapeTilde, "/", escapeSlash)

	// rfc6901Unescaper is a replacer that unescape a JSON Pointer string.
	rfc6901Unescaper = strings.NewReplacer(escapeTilde, "~", escapeSlash, "/")
)

type segment struct {
	key string
	idx int
}

// pointer represents an RFC 6901 JSON Pointer.
type pointer struct {
	buf  []byte
	base segment
	prev segment
	sep  int
}

func (p *pointer) clone() pointer { _ = "STUB: not implemented"; return *new(pointer) }

func (p *pointer) copy() string { _ = "STUB: not implemented"; return "" }

func (p *pointer) string() string { _ = "STUB: not implemented"; return "" }

func (p *pointer) isRoot() bool { _ = "STUB: not implemented"; return false }

func (p *pointer) appendKey(key string) { _ = "STUB: not implemented"; return }

func (p *pointer) appendIndex(idx int) { _ = "STUB: not implemented"; return }

func (p *pointer) snapshot() { _ = "STUB: not implemented"; return }

func (p *pointer) rewind() { _ = "STUB: not implemented"; return }

func (p *pointer) reset() { _ = "STUB: not implemented"; return }

func (p *pointer) appendEscapeKey(k string) { _ = "STUB: not implemented"; return }

var (
	errLeadingSlash             = errors.New("no leading slash")
	errIncompleteEscapeSequence = errors.New("incomplete escape sequence")
	errInvalidEscapeSequence    = errors.New("invalid escape sequence")
)

func parsePointer(s string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Last char is a '/', next fragment is an empty string.

// End of string, accumulate from last separator.
