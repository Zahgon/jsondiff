package jsondiff

import (
	"hash/maphash"
)

type hasher struct {
	mh maphash.Hash
}

func (h *hasher) digest(val interface{}, sort bool) uint64 { _ = "STUB: not implemented"; return 0 }

func (h *hasher) hash(i interface{}, sort bool) { _ = "STUB: not implemented"; return }

// Extract keys first, and sort them
// in lexicographical order.

func (h *hasher) sortArray(a []interface{}) { _ = "STUB: not implemented"; return }
