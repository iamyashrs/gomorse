//Created by Yash Raj Singh(http://yashrajsingh.net/) on 27.9.2015
package main

import(
	"testing"
	"fmt"
)

func TestMorse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello", ".... . .-.. .-.. ---"},
		{"0123456789", "----- .---- ..--- ...-- ....- ..... -.... --... ---.. ----."},
		{"GopherCon", "--. --- .--. .... . .-. -.-. --- -. .. -. -.. .. .-"},
		{"golang", "--. --- .-.. .- -. --."},
		{"@#$()!+_-", ".--.-. -.--. ...-.- -.--. -.--.- ---. .-.-. ..--.- -....-"},
		{"abcd", ".- -... -.-. -.."},
		{"", ""},
	}

	fmt.Println("Testing:\n========\n")

	for _, c := range cases {
		got := Morse(c.in, len(c.in))
		if got != c.want {
			t.Errorf("Error: Morse(%q) == %q, want %q", c.in, got, c.want)
		}
		fmt.Printf("OKAY: Morse('%s') == '%s'\n", c.in, got)
	}

	fmt.Println("\nAll tests passed!\n")
}