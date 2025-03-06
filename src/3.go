 package main
 import (
	 "fmt"
	 "math/rand"
	 "time"
 )

 func GenerateName() string {
 	rand.Seed(time.Now().UnixNano())
 	return fmt.Sprintf("%s-%d", randomString(8), rand.Intn(100))
 }

 func randomString(length int) string {
 	bytes := make([]byte, length)
 	for i := range bytes {
 		bytes[i] = byte(rand.Intn(25)+65)
 	}
 	return string(bytes)
 }

 func main() {
 	fmt.Println(GenerateName())
 }