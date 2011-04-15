// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bloomfilter_test

import(
    "testing"
    . "github.com/bjarneh/bloomfilter"
)

var words = []string{"just", "a", "few", "words", "used", "for", "testing"}

func TestBloomfilter(t *testing.T) {
    filter := NewSize(3)
    for i := 0; i < len(words); i++ {
        filter.Add(words[i])
    }

    for i := 0; i < len(words); i++ {
        if ! filter.Marked(words[i]) {
            t.Fatalf("'%s' should be marked\n",words[i])
        }
    }

    filter.Clear()

    for i := 0; i < len(words); i++ {
        if filter.Marked(words[i]) {
            t.Fatalf("'%s' shouldn't be marked\n",words[i])
        }
    }
}

func BenchmarkBloomfilterHashFunc(b *testing.B) {

    s := "abcdefghijklmnopqrstuvwxyz"

    for i := 0; i < b.N; i++ {
        _ = HashFunc(s)
    }
}

func BenchmarkBuiltInHashFunc(b *testing.B) {

    m := make(map[string]interface{})
    s := "abcdefghijklmnopqrstuvwxyz"

    for i := 0; i < b.N; i++ {
        m[s] = nil
    }
}
