package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	REQ_DISC     = -1
	PROOF_DISC   = 0
	REG          = 1
	AUTH         = 2
	ALL_SONGS    = 3
	ADD_SONG     = 4
	ALL_TAGS     = 5
	MY_SONGS     = 6
	BUYING_SONGS = 7
	BUY_SONG     = 8
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
