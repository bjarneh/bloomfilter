// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bloomfilter

// this is the java.lang.String.hashCode (GPL)
// override HashFunc for a different hash function
var HashFunc func(string) uint32 = func(s string) uint32 {
    var val uint32 = 1
    for i := 0; i < len(s); i++ {
        val += (val * 37) + uint32(s[i])
    }
    return val
}

var mask [32]uint32

// will be called before module is used
func init() {
    mask[0] = 1
    for i := 1; i < len(mask); i++ {
        mask[i] = 2 * mask[i-1]
    }
}

type Filter struct {
    size uint32
    hits []uint32
}

func New() *Filter {
    return &Filter{3200000, make([]uint32, 100000)}
}

func NewSize(size int) *Filter {
    return &Filter{32 * uint32(size), make([]uint32, size)}
}

func (f *Filter) Add(w string) {
    h := (HashFunc(w) % f.size)
    c := h / 32 // cell
    o := h % 32 // offset
    f.hits[c] = f.hits[c] | mask[o]
}

func (f *Filter) Marked(w string) bool {
    h := (HashFunc(w) % f.size)
    c := h / 32 // cell
    o := h % 32 // offset
    return (f.hits[c] & mask[o]) > 0
}

func (f *Filter) Clear() {
    for i := 0; i < len(f.hits); i++ {
        f.hits[i] = 0
    }
}

