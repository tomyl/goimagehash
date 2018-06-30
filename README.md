[![Build Status](https://travis-ci.org/tomyl/goimagehash.svg?branch=master)](https://travis-ci.org/tomyl/goimagehash)
[![GoDoc](https://godoc.org/github.com/tomyl/goimagehash?status.svg)](https://godoc.org/github.com/tomyl/goimagehash)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomyl/goimagehash)](https://goreportcard.com/report/github.com/tomyl/goimagehash)
[![Coverage Status](https://coveralls.io/repos/github/tomyl/goimagehash/badge.svg)](https://coveralls.io/github/tomyl/goimagehash)

This is a fork of [corona10/goimagehash](https://github.com/corona10/goimagehash). Changes:
* Somewhat more straightforward API.
* A tiny bit faster.

# goimagehash
> Inspired by [imagehash](https://github.com/JohannesBuchner/imagehash)

A image hashing library written in Go. ImageHash supports:
* [Average hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
* [Difference hashing](http://www.hackerfactor.com/blog/index.php?/archives/529-Kind-of-Like-That.html)
* [Perception hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
* [Wavelet hashing](https://fullstackml.com/wavelet-image-hash-in-python-3504fdd282b5) [TODO]

**Only support 64bits hash.**

## Installation
```
go get github.com/corona10/goimagehash
```
## Special thanks to
* [Haeun Kim](https://github.com/haeungun/)

## Usage

``` Go
func main() {
        file1, _ := os.Open("sample1.jpg")
        file2, _ := os.Open("sample2.jpg")
        defer file1.Close()
        defer file2.Close()

        img1, _ := jpeg.Decode(file1)
        img2, _ := jpeg.Decode(file2)
        hash1 := goimagehash.AverageHash(img1)
        hash2 := goimagehash.AverageHash(img2)
        distance := goimagehash.Distance(hash1, hash2)
        fmt.Printf("Distance between images: %v\n", distance)

        hash1 = goimagehash.DifferenceHash(img1)
        hash2 = goimagehash.DifferenceHash(img2)
        distance = goimagehash.Distance(hash1, hash2)
        fmt.Printf("Distance between images: %v\n", distance)
}
```

## Release Note

### v0.2.0
- Perception Hash is updated.
- Fix a critical bug of finding median value.

### v0.1.0
- Support Average hashing
- Support Difference hashing
- Support Perception hashing
- Use bits.OnesCount64 for computing Hamming distance by @dominikh
- Support hex serialization methods to ImageHash by @brunoro
