package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	animals := []string{"cat", "dog", "zebra"}

	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Randomly select an animal from the slice.
	targetAnimal := animals[rand.Intn(len(animals))]

	fmt.Println("Guess the animal!")

	var guess string

	// Loop until the guess is correct.
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		// Convert guess to lowercase to make the comparison case-insensitive.
		guess = strings.ToLower(guess)

		if guess == targetAnimal {
			fmt.Println("Congratulations! You've guessed the animal correctly.")
			break
		} else {
			fmt.Println("Wrong! Try again.")
		}
	}
}
