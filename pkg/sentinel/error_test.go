// Copyright © 2017 Nelz
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package sentinel implements an error that can be used as a package-level
// constant, as per Dave Cheney's thought experiment:
// https://dave.cheney.net/2016/04/07/constant-errors
package sentinel

import (
	"fmt"
	"testing"
)

const one = Error("solo")
const two = Error("common")
const three = Error("common")

func TestEquality(t *testing.T) {
	// Nils are never equal to anything
	if one.Equals(nil) {
		t.Fatalf("one == nil")
	}

	// Identity, should be equal
	if !one.Equals(one) {
		t.Fatalf("one != one")
	}

	// Differing types, should NOT be equal
	if one.Equals(fmt.Errorf(one.Error())) {
		t.Fatalf("one == one")
	}

	// Completely different content, should NOT be equal
	if one.Equals(two) {
		t.Fatalf("one == two")
	}

	// Ideally we would like two distinct instances of
	// Error to evaluate to UNEQUAL, but since we are targeting
	// this package to work as consts, and the Go1
	// language doesn't allow memory address accesses for
	// constant expressions, therefore we are unable to
	// discriminate between two same-type-same-string variables
	// where each has separate memory addresses.
	if !two.Equals(three) {
		t.Fatalf("two != three")
	}
}

func TestMark(t *testing.T) {
	if one != one.Mark() {
		t.Fatalf("expected self return")
	}
}

func TestBehavior(t *testing.T) {
	if IsMarker(fmt.Errorf("fail case")) {
		t.Fatalf("non-sentinel errors are not Markers")
	}

	if !IsMarker(one) {
		t.Fatalf("expected sentinel to succeed")
	}
}
