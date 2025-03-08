package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// GitHubRepo represents a GitHub repository with a name and description.
type GitHubRepo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	repo := generateRandomGitHubRepo()
	fmt.Println(repo)
}

func generateRandomGitHubRepo() GitHubRepo {
	rand.Seed(time.Now().UnixNano())

	name := randomString(10, true)
	description := randomString(20, true)

	return GitHubRepo{Name: name, Description: description}
}

func randomString(n int, vowels bool) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	if vowels {
		letterRunes = append(letterRunes, []rune("aeiou")...)
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
