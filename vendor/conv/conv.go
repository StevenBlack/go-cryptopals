// Package conv contains various conversion functions.
package conv

import (
	"encoding/base64"
	"encoding/hex"
)

func decodeHexBytes(hexBytes []byte) ([]byte, error) {
	ret := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(ret, hexBytes)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func encodeToBase64(input []byte) []byte {
	ret := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(ret, input)
	return ret
}

// HexToBase64 takes a slice of hex encoded bytes, and
// returns a slice of base64 encoding bytes
func HexToBase64(hexBytes []byte) ([]byte, error) {
	b, err := decodeHexBytes(hexBytes)
	if err != nil {
		return nil, err
	}
	ret := encodeToBase64(b)
	return ret, nil
}

// HexStringToBase64String takes a string representation of Hex,
// returns a string representation of base64.
func HexStringToBase64String(hexString string) (string, error) {
	b, err := HexToBase64([]byte(hexString))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// HexXOR takes two equal length strings and BitwiseXOR's them together
// returns the result as string. Presumes both input strings are equal length.
func HexXOR(hexString1, hexString2 string) (string, error) {
	b1, err := decodeHexBytes([]byte(hexString1))
	if err != nil {
		return "", err
	}
	b2, err := decodeHexBytes([]byte(hexString2))
	if err != nil {
		return "", err
	}
	out := []byte{}
	for i, b1b := range b1 {
		out = append(out, b1b^b2[i])
	}
	ret := hex.EncodeToString(out)
	return ret, nil
}
