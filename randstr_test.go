package randstr_test

import (
	"github.com/hexid/go-randstr"
	"testing"
)

func TestCharSet(t *testing.T) {
	testCharSet(t, randstr.LowerASCII|randstr.DigitASCII,
		"abcdefghijklmnopqrstuvwxyz0123456789")

	testCharSet(t, randstr.AlphaASCII,
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	testCharSet(t, randstr.RandASCII,
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"+
			"0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

	testCharSet(t, 0, "")
}

func TestRandom(t *testing.T) {
	var rnd []byte
	var err error

	if rnd, err = randstr.Random(0, randstr.RandASCII.String()); err != nil {
		t.Error(err)
	} else if len(rnd) != 0 {
		t.Errorf("Expected empty string")
	}

	if _, err = randstr.Random(10, ""); err == nil {
		t.Errorf("Expected error for empty character set")
	}

	if rnd, err = randstr.Random(20, randstr.SymbolASCII.String()); err != nil {
		t.Error(err)
	} else if len(rnd) != 20 {
		t.Errorf("Expected a random string of length 20, found %d", len(rnd))
	} else {
		t.Log(string(rnd))
	}
}

func testCharSet(t *testing.T, cs randstr.CharSet, testStr string) {
	if csStr := cs.String(); csStr != testStr {
		t.Errorf("Expected %s to equal %s", csStr, testStr)
	}
}
