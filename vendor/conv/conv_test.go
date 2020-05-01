package conv

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	t.Run("Valid hex passed", func(t *testing.T) {
		//These strings are pulled from the crypto pals challenge page
		hb := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
		expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
		ret, err := HexToBase64(hb)
		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}
		if string(ret) != expected {
			t.Errorf("expected: %s, got: %s", expected, ret)
		}
	})
}

func TestHexToBase64MixedCase(t *testing.T) {
	t.Run("Valid mixed-case hex passed", func(t *testing.T) {
		//These strings are pulled from the crypto pals challenge page
		hb := []byte("FfFfFfFfFfFf")
		expected := "////////"
		ret, err := HexToBase64(hb)
		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}
		if string(ret) != expected {
			t.Errorf("expected: %s, got: %s", expected, ret)
		}
	})
}

func TestHexToBase64Invalid(t *testing.T) {
	t.Run("Invalid hex passed", func(t *testing.T) {
		hb := []byte("Tobeornottobe")
		ret, err := HexToBase64(hb)
		if ret != nil {
			t.Errorf("expected: nil, got: %s", ret)
		}

		if err == nil {
			t.Error("expected: error not nil, got: nil")
		}
	})
}

func TestHexStringToBase64String(t *testing.T) {
	t.Run("Valid hex string passed", func(t *testing.T) {
		//These strings are pulled from the crypto pals challenge page
		hs := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
		expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
		ret, err := HexStringToBase64String(hs)
		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}
		if string(ret) != expected {
			t.Errorf("expected: %s, got: %s", expected, ret)
		}
	})
}

func TestFixedXOR(t *testing.T) {
	t.Run("Fixed-length, two hex XOR", func(t *testing.T) {
		//These strings are pulled from the crypto pals challenge page
		h1 := "1c0111001f010100061a024b53535009181c"
		h2 := "686974207468652062756c6c277320657965"
		expected := "746865206b696420646f6e277420706c6179"
		ret, err := HexXOR(h1, h2)
		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}
		if string(ret) != expected {
			t.Errorf("expected: %s, got: %s", expected, ret)
		}
	})
}
