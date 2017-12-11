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

// Error is just a synonym for the primitive string type so that
// consumers can assign it as a constant expression
type Error string

// Error fulfills the standard lib error interface
func (s Error) Error() string {
	return string(s)
}

// Equals provides our best guess of whether or not a given error is the same
// as our (const) sentinel error value.
//
// Just doing a "==" evaluation is not
// recommended for this package, as Error is just a synonym of string, and
// double-equals will just fall back to string comparison.
//
// Due to Go1 language limitations this method may respond with a false positive
// if you are comparing two instances of Error that have the same underlying
// string value.
func (s Error) Equals(e error) bool {
	if _, ok := e.(Error); ok {
		return s == e
	}
	return false
}

// Mark conforms to the Marker interface, and is provided to address desires
// for behavior-based identification of sentinel values (using "has a" rules,
// ranther than "is a" evaluation).
//
// See more on this at https://dave.cheney.net/2014/12/24/inspecting-errors
func (s Error) Mark() error {
	return s
}

// Marker is a simple interface to test errors to see if they are intended
// to be sentinel values
type Marker interface {
	Mark() error
}

// IsMarker tests whether a given error is intending to be used as a
// sentinel value.
func IsMarker(e error) bool {
	_, ok := e.(Marker)
	return ok
}
