// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

// Distance method returns a distance between two hashes.
func Distance(lhash, rhash uint64) int {
	hamming := lhash ^ rhash
	return popcnt(hamming)
}
