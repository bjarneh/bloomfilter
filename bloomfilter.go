// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bloomfilter

/*
  A Bloomfilter is a negative filter which quickly
  determines whether a word is present in a large set
  of words, given that all words have been added to the
  filter.

  Basically it holds a set of boolean values, which may or
  may not have been hashed to, in a long int-array, where
  each int, or in this case uint32, holds 32 bits/bools.
  The larger the bloomfilter, the more accurate it becomes,
  but it also uses more memory naturally.

  Typical use:
  <pre>

     filter := bloomfilter.NewSize(10000)

     filter.Add("plenty")
     filter.Add("of")
     filter.Add("words")

     if ! filter.Marked("someword") {
        println("'someword' is not present")
     } else {
        println("'someword' may be here")
     }

   </pre>
*/


// this is the java.lang.String.hashCode()
// override for a different hash function
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

// new filter with default size (640000 booleans, in 20000 uint32)
func New() *Filter {
    return &Filter{640000, make([]uint32, 20000)}
}

func NewSize(size int) *Filter {
    if size > 32 {
        size = size/32
    }else{
        size = 1
    }
    return &Filter{32 * uint32(size), make([]uint32, size)}
}

// add word to filter
func (f *Filter) Add(w string) {
    h := (HashFunc(w) % f.size)
    c := h / 32 // cell
    o := h % 32 // offset
    f.hits[c] = f.hits[c] | mask[o]
}

// if this function returns false, w is not present
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

