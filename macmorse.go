//Created by Yash Raj Singh(http://yashrajsingh.net/) on 16.9.2015
package main

import (
	"os"
	"fmt"
	"strings"
)

func Morse(arg string, length int) string {
	code := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",

		"Ä": ".-.-",
		"Á": ".--.-",
		"Å": ".--.-",
		"É": "..-..",
		"Ñ": "--.--",
		"Ö": "---.",
		"Ü": "..--",

		"0": "-----",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",

		"/": "-..-.",
		",": "--..--",
		".": ".-.-.-",
		"?": "..--..",
	    "=": "-...-",
		"+": ".-.-.",
		"#": "-.--.",
		"$": "...-.-",
		"@": ".--.-.",
		":": "---...",
		"-": "-....-",
		"[": "-.--.",
		"]": "-.--.-",
		"_": "..--.-",
		"'": ".----.",
		"(": "-.--.",
		")": "-.--.-",
		"!": "---.",
		" ": "    ",
	}

	a := ""
	for i := 0; i < length; i++ {
		val, ok := code[strings.ToLower(string(arg[i]))]

		if ok{
			a += val + " "
		}else{
			a = "Morse code for Character entered '"+ string(arg[i]) +"' was not found!"
			break
		}
	}
	if strings.HasSuffix(a, " ") {
		a = a[:len(a)-len(" ")]
    }
    return a
}

func main() {
	if len(os.Args) < 2 {
        fmt.Println("missing required argument:  input text to convert to morse code")
        return
	}
	arg := os.Args[1]
    m := Morse(arg, len(arg))
    fmt.Println(m)
}