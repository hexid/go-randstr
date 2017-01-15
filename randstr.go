package randstr

import (
	"crypto/rand"
	"errors"
)

var (
	errNoChars = errors.New("Must provide at least one character to randomize")
	randChars  = []string{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",         // UpperASCII
		"abcdefghijklmnopqrstuvwxyz",         // LowerASCII
		"0123456789",                         // DigitASCII
		"!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", // SymbolASCII
		" ", // SpaceASCII
	}
)

// CharSet represents a set of characters that will
// be chosen from when generating a random string
type CharSet uint8

const (
	// UpperASCII is the set of uppercase alpha characters
	UpperASCII CharSet = 1 << iota
	// LowerASCII is the set of lowercase alpha characters
	LowerASCII
	// DigitASCII is the set of numbers
	DigitASCII
	// SymbolASCII is the basic set of ASCII symbols
	SymbolASCII
	// SpaceASCII is the ascii space
	SpaceASCII
)

const (
	// AlphaASCII is an alias for upper and lower case letters
	AlphaASCII CharSet = UpperASCII | LowerASCII
	// RandASCII is an alias for the full set of ASCII characters (w/o Space)
	RandASCII CharSet = AlphaASCII | DigitASCII | SymbolASCII
)

// String returns the string representation of a CharSet
func (cs CharSet) String() (chars string) {
	ind := 0
	for cs != 0 {
		if cs&1 > 0 {
			chars += randChars[ind]
		}
		cs = cs >> 1
		ind++
	}
	return
}

// Random returns a random string from a string of available characters
func Random(strSize uint, chars string) ([]byte, error) {
	charsLen := byte(len(chars))
	if charsLen == 0 {
		return nil, errNoChars
	}

	randBytes := make([]byte, strSize)
	rand.Read(randBytes)
	for k, v := range randBytes {
		randBytes[k] = chars[v%charsLen]
	}
	return randBytes, nil
}
