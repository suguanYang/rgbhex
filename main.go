package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func rgbToHex(rgb string) string {
	values := strings.Split(rgb, ",")

	var hexOutPut string
	for _, v := range values {
		dec, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		hexOutPut += strconv.FormatInt(int64(dec), 16)
	}
	return hexOutPut
}

func hexToRgb(v string) string {
	gap := 2
	var decs []string
	integer := new(big.Int)

	for i := 0; i < len(v); i += gap {
		gapV := v[i : i+gap]
		integer.SetString(gapV, 16)
		decs = append(decs, integer.String())
	}

	return strings.Join(decs, ",")
}

func main() {
	text := strings.Join((os.Args[1:]), ",")

	var output string
	inputLength := len(text)
	if strings.Contains(text, ",") || strings.Contains(text, " ") {
		output = rgbToHex(text)
	} else if inputLength == 6 {
		output = hexToRgb(text)
	} else if inputLength == 3 {
		output = hexToRgb(text + text)
	} else {
		fmt.Println("Usage:\nrgbhex 255 255 255\nrgbhex 255,255,255\nrgbhex ffffff\nrgbhex fff")
	}

	fmt.Println(output)
}
