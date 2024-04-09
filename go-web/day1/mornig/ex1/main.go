package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func generatedId(s string) string {
	h := sha256.New()
	if _, err := h.Write([]byte(s)); err != nil {
		panic(err.Error())
	} else {
		bt := h.Sum(nil)
		return fmt.Sprintf("%x\n", bt)
	}
}

func main() {
	jsonW, err := json.MarshalIndent(Users, "", " ")
	if err != nil {
		panic("error create json")
	}
	err = ioutil.WriteFile("user.json", jsonW, 0644)
	if err != nil {
		panic("error write to a file")
	}
	jsonR, err := ioutil.ReadFile("user.json")
	if err != nil {
		panic("error read from file")
	}
	var usersR []User
	err = json.Unmarshal(jsonR, &usersR)
	if err != nil {
		panic("error ummarshal the json")
	}
	fmt.Println("this is what I  write:\n", Users)
	fmt.Println("This is waht I read after Write:\n", usersR)

}
