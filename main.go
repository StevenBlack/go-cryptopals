package main

import (
	"conv"
	"fmt"
)

func main() {

	// challenge 1
	ret1, _ := conv.HexToBase64([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println(string(ret1))

	ret1a, _ := conv.HexStringToBase64String("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println(ret1a)

	// challenge 2
	h1 := "1c0111001f010100061a024b53535009181c"
	h2 := "686974207468652062756c6c277320657965"
	// expected := "746865206b696420646f6e277420706c6179"
	ret2, _ := conv.HexXOR(h1, h2)
	fmt.Println(ret2)
}
