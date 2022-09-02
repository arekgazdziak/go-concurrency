package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Config struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Id      int    `json:"id"`
}

func main() {
	cfg2 := Config{
		Name:    "John",
		Surname: "Williams",
	}

	log.Printf("name: %s, surname: %s, id: %d\n", cfg2.Name, cfg2.Surname, cfg2.Id)

	str := `{"name": "arek", "id": 44 }`
	reader := strings.NewReader(str)

	if err := json.NewDecoder(reader).Decode(&cfg2); err != nil {
		panic(err)
	}

	log.Printf("name: %s, surname: %s, id: %d\n", cfg2.Name, cfg2.Surname, cfg2.Id)
}
