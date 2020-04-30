package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	ret, _ := HexToBase64([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println(string(ret))
}

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

//HexToBase64 takes a slice of hex encoded bytes, and returns a slice of base64 encoding bytes
func HexToBase64(hexBytes []byte) ([]byte, error) {
	b, err := decodeHexBytes(hexBytes)
	if err != nil {
		return nil, err
	}

	ret := encodeToBase64(b)
	return ret, nil
}
