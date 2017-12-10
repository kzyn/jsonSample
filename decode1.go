package main

import (
	"encoding/json"
	"os"
	"fmt"
)

const configFile = "C:\\Temp\\generated.json"

func main () {
	var cfg interface{}

	fd, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fd.Close()

	err = json.NewDecoder(fd).Decode(&cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	age, _ := cfg.(map[string]interface{})["person"].(map[string]interface{})["age"]
	fmt.Println(age)
	os.Exit(0)
}
