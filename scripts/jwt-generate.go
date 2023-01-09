package scripts

import (
	"crypto/rand"
	"fmt"
)

func GenerateSecret() {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		panic(err.Error())
	}
	res := string(key)
	fmt.Println(res)
}