package main

import (
	"fmt"
	route2 "github.com/BaianorASR/ifc2-simulator/app/route"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	err := route.LoadPositions()
	if err != nil {
		panic(err)
	}
	stringJson, err := route.ExportJsonPositions()
	if err != nil {
		panic(err)
	}

	fmt.Println(stringJson[1])
}
