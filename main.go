package main

import (
	"fmt"
	"math/rand"
	"time"
)

func passwordGenerator(length int, r *rand.Rand) string {
	// Define character sets
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	all := lower + upper + digits + special // Combine all character sets

	password := ""

	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(all)) // Generate random index
		password += string(all[randomIndex])

	}
	return password

}

func main() {
	var passwordLength int
	var passwordNumber int

	fmt.Print("Enter the desired password length (e.g., 12): ")

	fmt.Scanln(&passwordLength)

	fmt.Print("Enter the number of passwords to generate (e.g., 5): ")
	fmt.Scanln(&passwordNumber)

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator

	fmt.Println("\nPassword Generated:", passwordGenerator(passwordLength, r))

	fmt.Println("\nMore Options....")

	for i := 1; i <= passwordNumber; i++ {
		p := passwordGenerator(passwordLength, r)

		fmt.Printf("Password %d: %s\n", i, p)
	}

	fmt.Println("\nRe-run the program to generate new passwords.")
}
