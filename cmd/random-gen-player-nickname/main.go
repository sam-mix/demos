package main

import (
	"fmt"
	"math/rand"
)

var adjectives = []string{"happy", "sad", "crazy", "cool", "angry", "silly", "grumpy", "funny", "bouncy"}
var nouns = []string{"cat", "dog", "bird", "fish", "monkey", "tiger", "dragon", "unicorn", "elephant"}

func main() {
	// // Seed the random number generator with the current time
	// rand.Seed(time.Now().Unix())

	// Generate a random number between 0 and the length of the adjectives array
	adjIndex := rand.Intn(len(adjectives))

	// Generate a random number between 0 and the length of the nouns array
	nounIndex := rand.Intn(len(nouns))

	// Generate the nickname by concatenating a random adjective and a random noun
	nickname := adjectives[adjIndex] + "_" + nouns[nounIndex]

	fmt.Println("Your random nickname is:", nickname)
}
