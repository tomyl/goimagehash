// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"image"
	"image/jpeg"
	"os"
	"testing"
)

func TestHashCompute(t *testing.T) {
	for _, tt := range []struct {
		img1     string
		img2     string
		method   func(img image.Image) uint64
		name     string
		distance int
	}{
		{"_examples/sample1.jpg", "_examples/sample1.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", AverageHash, "AverageHash", 0},
		{"_examples/sample1.jpg", "_examples/sample2.jpg", AverageHash, "AverageHash", 42},
		{"_examples/sample1.jpg", "_examples/sample3.jpg", AverageHash, "AverageHash", 4},
		{"_examples/sample1.jpg", "_examples/sample4.jpg", AverageHash, "AverageHash", 38},
		{"_examples/sample2.jpg", "_examples/sample3.jpg", AverageHash, "AverageHash", 40},
		{"_examples/sample2.jpg", "_examples/sample4.jpg", AverageHash, "AverageHash", 6},
		{"_examples/sample1.jpg", "_examples/sample1.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample1.jpg", "_examples/sample2.jpg", DifferenceHash, "DifferenceHash", 43},
		{"_examples/sample1.jpg", "_examples/sample3.jpg", DifferenceHash, "DifferenceHash", 0},
		{"_examples/sample1.jpg", "_examples/sample4.jpg", DifferenceHash, "DifferenceHash", 37},
		{"_examples/sample2.jpg", "_examples/sample3.jpg", DifferenceHash, "DifferenceHash", 43},
		{"_examples/sample2.jpg", "_examples/sample4.jpg", DifferenceHash, "DifferenceHash", 16},
		{"_examples/sample1.jpg", "_examples/sample1.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample2.jpg", "_examples/sample2.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample3.jpg", "_examples/sample3.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample4.jpg", "_examples/sample4.jpg", PerceptionHash, "PerceptionHash", 0},
		{"_examples/sample1.jpg", "_examples/sample2.jpg", PerceptionHash, "PerceptionHash", 32},
		{"_examples/sample1.jpg", "_examples/sample3.jpg", PerceptionHash, "PerceptionHash", 2},
		{"_examples/sample1.jpg", "_examples/sample4.jpg", PerceptionHash, "PerceptionHash", 30},
		{"_examples/sample2.jpg", "_examples/sample3.jpg", PerceptionHash, "PerceptionHash", 34},
		{"_examples/sample2.jpg", "_examples/sample4.jpg", PerceptionHash, "PerceptionHash", 20},
	} {
		file1, err := os.Open(tt.img1)
		if err != nil {

		}
		defer file1.Close()

		file2, err := os.Open(tt.img2)
		if err != nil {
			t.Errorf("%s", err)
		}
		defer file2.Close()

		img1, err := jpeg.Decode(file1)
		if err != nil {
			t.Errorf("%s", err)
		}

		img2, err := jpeg.Decode(file2)
		if err != nil {
			t.Errorf("%s", err)
		}

		hash1 := tt.method(img1)
		hash2 := tt.method(img2)
		dis1 := Distance(hash1, hash2)
		dis2 := Distance(hash2, hash1)

		if dis1 != dis2 {
			t.Errorf("Distance should be identical %v vs %v", dis1, dis2)
		}

		if dis1 != tt.distance {
			t.Errorf("%s: Distance between %v and %v is expected %v but got %v", tt.name, tt.img1, tt.img2, tt.distance, dis1)
		}
	}
}

func BenchmarkDistanceIdentical(b *testing.B) {
	h1 := uint64(0xe48ae53c05e502f7)
	h2 := uint64(0xe48ae53c05e502f7)

	for i := 0; i < b.N; i++ {
		Distance(h1, h2)
	}
}

func BenchmarkDistanceDifferent(b *testing.B) {
	h1 := uint64(0xe48ae53c05e502f7)
	h2 := uint64(0x678be53815e510f7) // 8 bits flipped

	for i := 0; i < b.N; i++ {
		Distance(h1, h2)
	}
}
