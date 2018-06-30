// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimagehash

import (
	"image"

	"github.com/corona10/goimagehash/etcs"
	"github.com/corona10/goimagehash/transforms"
	"github.com/nfnt/resize"
)

// AverageHash fuction returns a hash computation of average hash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
func AverageHash(img image.Image) uint64 {
	// Create 64bits hash.
	resized := resize.Resize(8, 8, img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	flattens := transforms.FlattenPixels(pixels, 8, 8)
	avg := etcs.MeanOfPixels(flattens)
	var hash uint64

	for idx, p := range flattens {
		if p > avg {
			hash |= 1 << uint(idx)
		}
	}

	return hash
}

// DifferenceHash function returns a hash computation of difference hash.
// Implementation follows
// http://www.hackerfactor.com/blog/?/archives/529-Kind-of-Like-That.html
func DifferenceHash(img image.Image) uint64 {
	resized := resize.Resize(9, 8, img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	idx := 0
	var hash uint64

	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[i])-1; j++ {
			if pixels[i][j] < pixels[i][j+1] {
				hash |= 1 << uint(idx)
			}
			idx++
		}
	}

	return hash
}

// PerceptionHash function returns a hash computation of phash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
func PerceptionHash(img image.Image) uint64 {
	resized := resize.Resize(64, 64, img, resize.Bilinear)
	pixels := transforms.Rgb2Gray(resized)
	dct := transforms.DCT2D(pixels, 64, 64)
	flattens := transforms.FlattenPixels(dct, 8, 8)
	median := etcs.MedianOfPixels(flattens)
	var hash uint64

	for idx, p := range flattens {
		if p > median {
			hash |= 1 << uint(idx)
		}
	}

	return hash
}
