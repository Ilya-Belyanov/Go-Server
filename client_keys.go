package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ClientKeys struct {
	TYPE   string
	NAME   string
	PASS   string
	ID     string
	TAGS   string
	AUTHOR string
}

func getClientKeys() ClientKeys {
	keys := ClientKeys{}
	jsonFile, err := os.Open("clientkey.json")
	if err != nil {
		fmt.Print("Don't found client keys")
		return keys
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &keys)
	if err != nil {
		fmt.Print("Don't parse client keys")
	}
	return keys
}
