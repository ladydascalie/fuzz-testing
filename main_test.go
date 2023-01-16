package main

import (
	"os"
	"testing"
	"unicode/utf8"
)

// TODO(bc): Run this first to ensure we've got a clean slate!
func TestPrepare(_ *testing.T) {
	os.RemoveAll("testdata")
}

// Simply running this test will not provide much value.
// The test will pass, as the test cases are hard-coded, and very simple.
//
// However, if we run go test -fuzz=FuzzReverse, we should see that the test
// fails, as reversing some strings can produce invalid UTF-8!
//
// We can open the `testdata/fuzz/FuzzReverse` directory, and see that
// the test has generated a corpus of test cases, and has found a bug!
//
// Now if we run this test normally, we will see that the test fails, as the
// corpus generated in testdata will now be used.
//
// TODO: Run this:
// go test -fuzz=FuzzReverse
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// TODO: Run this:
// go test -fuzz=FuzzBetterReverse
//
// This will fail on the double reverse in the case where we receive invalid UTF-8.
// We can confirm this by logging the input and runes within the BetterReverse implementation
// then running this test again (as a regular test, not a fuzz test).
// TODO: Uncomment the logging in BetterReverse.
func FuzzBetterReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := BetterReverse(orig)
		doubleRev := BetterReverse(rev)
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// This will pass, as we are now preventing inputs that are not valid UTF-8.
//
// The default behavior of go test -fuzz is to run the test with the corpus
// until it fails, and then stop.
//
// As no error will happen here, we need to specify a maximum number of seconds
// to run the test for.
//
// TODO: Run this:
// go test -fuzz=FuzzCorrectReverse -fuzztime=30s
func FuzzCorrectReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := CorrectReverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := CorrectReverse(rev)
		if err2 != nil {
			return
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
