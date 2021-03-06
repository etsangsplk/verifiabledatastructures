/*

Copyright 2017 Continusec Pty Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package merkle

import "crypto/sha256"

// ConstructMapKeyPath returns the path in the tree for a given key. Specifically it takes
// the SHA256 hash of the key, and then returns a big-endian slice of booleans representing
// the equivalent path in the tree.
func ConstructMapKeyPath(key []byte) []bool {
	h := sha256.Sum256(key)
	nm := len(h) * 8
	rv := make([]bool, nm)
	for i, b := range h {
		for j := uint(0); j < 8; j++ {
			if ((b >> j) & 1) == 1 {
				rv[(uint(i)<<3)+7-j] = true
			}
		}
	}
	return rv
}

var defaultLeafValues = GenerateMapDefaultLeafValues()

// GenerateMapDefaultLeafValues returns a copy of the default leaf values for any empty nodes
// in a proof. This can be useful for implementations that verify inclusion proofs of Map Values.
func GenerateMapDefaultLeafValues() [][]byte {
	rv := make([][]byte, 257)
	rv[256] = LeafHash(nil)
	for i := 255; i >= 0; i-- {
		rv[i] = NodeHash(rv[i+1], rv[i+1])
	}
	return rv
}

// NodeHash is a utility function for calculating the Merkle Tree Hash for a node.
func NodeHash(l, r []byte) []byte {
	h := sha256.New()
	h.Write([]byte{1})
	h.Write(l)
	h.Write(r)
	return h.Sum(nil)
}

// LeafHash is a utility function for calculating the Merkle Tree Hash for a leaf.
func LeafHash(b []byte) []byte {
	h := sha256.New()
	h.Write([]byte{0})
	h.Write(b)
	return h.Sum(nil)
}

// IsPow2 returns true if n is a power of 2
func IsPow2(n int64) bool {
	return CalcK(n+1) == n
}

// CalcK returns the largest value that is a power of 2, less than n
func CalcK(n int64) int64 {
	k := int64(1)
	for (k << 1) < n {
		k <<= 1
	}
	return k
}

// Path returns indices as per RFC6962
func Path(m, startN, endN int64) [][2]int64 {
	n := endN - startN
	if n == 1 {
		rv := [0][2]int64{}
		return rv[:]
	}
	k := CalcK(n)
	if m < k {
		rv := Path(m, startN, startN+k)
		rv = append(rv, [2]int64{startN + k, endN})
		return rv
	}
	rv := Path(m-k, startN+k, endN)
	rv = append(rv, [2]int64{startN, startN + k})
	return rv
}

// SubProof returns indices as per RFC6962
func SubProof(m, startN, endN int64, b bool) [][2]int64 {
	n := endN - startN
	if m == n {
		if b {
			return nil
		}
		return [][2]int64{{startN, endN}}
	}
	k := CalcK(n)
	if m <= k {
		return append(SubProof(m, startN, startN+k, b), [2]int64{startN + k, endN})
	}
	return append(SubProof(m-k, startN+k, endN, false), [2]int64{startN, startN + k})
}
