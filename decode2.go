package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type config struct {
	Id			string	`json:"id"`
	IsActive	bool	`json:"isActive"`
	Age			int 	`json:"age"`
}

const configFile = "C:\\Temp\\generated.json"

func main() {
	cfg, err := readConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("id =", cfg.Id)
	fmt.Println("age =", cfg.Age)
	fmt.Println("active =", cfg.IsActive)
	os.Exit(0)
}

func readConfig() (*config, error) {
	var cfg config

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
	return &cfg, err
}