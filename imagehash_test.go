// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	_ "image/jpeg"
	"testing"
)

func TestNewImageHash(t *testing.T) {
	for _, tt := range []struct {
		datas    [][]uint8
		distance int
	}{
		{[][]uint8{{1, 0, 1, 1}, {0, 0, 0, 0}}, 3},
		{[][]uint8{{0, 0, 0, 0}, {0, 0, 0, 0}}, 0},
		{[][]uint8{{0, 0, 0, 0}, {0, 0, 0, 1}}, 1},
	} {
		data1 := tt.datas[0]
		data2 := tt.datas[1]
		var hash1 uint64
		var hash2 uint64

		for i := 0; i < len(data1); i++ {
			if data1[i] == 1 {
				hash1 |= 1 << uint(i)
			}
		}

		for i := 0; i < len(data2); i++ {
			if data2[i] == 1 {
				hash2 |= 1 << uint(i)
			}
		}

		dis := Distance(hash1, hash2)
		if dis != tt.distance {
			t.Errorf("Distance between %v and %v expected as %d but got %d", data1, data2, tt.distance, dis)
		}
	}
}
