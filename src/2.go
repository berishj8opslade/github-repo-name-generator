package main

import "fmt"

func main() {
	repoName := fmt.Sprintf("github-repo-%s", randomString(6))
	description := generateDescription()

	// Print the repo name and description in JSON format
	fmt.Printf(`{"name": "%s", "description": "%s"}`, repoName, description)
}

func randomString(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphabet[b%byte(len(alphabet))]
	}
	return string(bytes)
}

func generateDescription() string {
	descriptions := []string{
		"A project to generate valid GitHub repo names and descriptions",
		"Generate a random GitHub repo name and description in JSON format",
		"Create a new GitHub repo with a unique name and description",
		"Randomly generate a GitHub repo name and description based on your input",
	}

	return descriptions[rand.Intn(len(descriptions))]
}
