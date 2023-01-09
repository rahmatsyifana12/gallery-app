package scripts

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"strings"
)

func GenerateSecret() (error) {
	key := make([]byte, 64)

	_, err := rand.Read(key)
	if err != nil {
		return err
	}

	str_key := fmt.Sprintf("%x", key)
	
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		return err
	}

	new_val_str := ""

	str_arr := strings.Split(string(b), "\n")
	for _, el := range str_arr {
		if el == "JWT_ACCESS_SECRET=" {
			el += str_key
		}
		new_val_str += el + "\n"
	}

	new_val := []byte(new_val_str)
	err = ioutil.WriteFile(".env", new_val, 0644)
	if err != nil {
		return err
	}

	return nil
}