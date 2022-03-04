package main

import "golang.org/x/crypto/bcrypt"

func main() {
	var a string
	flag.StringVar(&a, "a", "", "string to hash")
	flag.Parse()
	b, err := HashPassword(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
