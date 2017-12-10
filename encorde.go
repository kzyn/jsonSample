package main

import (
	"encoding/json"
	"os"
	"fmt"
	"bytes"
)

type Friends struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	ID         string   `json:"_id"`
	Index      int      `json:"index"`
	GUID       string   `json:"guid"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Picture    string   `json:"picture"`
	Age        int      `json:"age"`
	EyeColor   string   `json:"eyeColor"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	friends    Friends	`json:"friends"`
	Greeting      string `json:"greeting"`
	FavoriteFruit string `json:"favoriteFruit"`
}

const configFile = "C:\\Users\\kazuya\\Downloads\\encode.json"

func main () {
	person := Person{
			"1234567890",
			1,
			"hoge",
			true,
			"balance",
			"picture",
			31,
			"red",
			"nohara",
			"otoko",
			"saison",
			"aaa@hoge",
			"090000000",
			"ageo",
			"about",
			"aaaaa",
			12345,
			12345,
			Friends {
				10,
				"piko",
			},
			"greeting",
			"soccer",
	}

	fd, err := os.Create(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fd.Close()
/*
	err = json.NewEncoder(fd).Encode(&person)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
*/
	outputJson, err := json.Marshal(&person)
	if err != nil {
		os.Exit(1)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(outputJson), "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fd.WriteString(buf.String())
	os.Exit(0)
}