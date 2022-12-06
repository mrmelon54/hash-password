package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"os"
)

func main() {
	var cost int
	flag.IntVar(&cost, "cost", 15, "cost of bcrypt hashing operation")
	flag.Parse()

	if term.IsTerminal(int(os.Stdin.Fd())) {
		_, _ = fmt.Fprint(os.Stderr, "Password: ")
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		_, _ = fmt.Fprintln(os.Stderr)
		_, _ = fmt.Fprintln(os.Stderr, "Hashing...")

		b, err := hashPassword(string(password), cost)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(b)
	} else {
		_, _ = fmt.Fprint(os.Stderr, "No terminal input found")
	}
}

func hashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
