package main

import "fmt"
import "math/rand"

func GenerateRepoName() string {
	return fmt.Sprintf("repo%d", rand.Intn(100))
}

func GenerateRepoDescription() string {
	return fmt.Sprintf("This is a test description %d", rand.Intn(100))
}

func main() {
	fmt.Println(GenerateRepoName())
	fmt.Println(GenerateRepoDescription())
}