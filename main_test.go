package main

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
