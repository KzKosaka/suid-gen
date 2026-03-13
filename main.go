package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

func main() {
	var num = 0
	var uuidType = ""

	flag.IntVar(&num, "n", 100, "number of short-uuids generated")
	flag.StringVar(&uuidType, "t", "v4", "base uuid version (v1, v4, v6, v7)")
	flag.Parse()

	if num == 1 {
		fmt.Print(createShortUUIDWithType(uuidType))
	} else {
		for i := 0; i < num; i++ {
			fmt.Println(createShortUUIDWithType(uuidType))
		}
	}
}

func createShortUUIDWithType(uuidType string) string {
	var u uuid.UUID
	var err error
	switch uuidType {
	case "v1":
		u, err = uuid.NewUUID()
	case "v4":
		u = uuid.New()
	case "v6":
		u, err = uuid.NewV6()
	case "v7":
		u, err = uuid.NewV7()
	default:
		log.Fatalf("unsupported uuid type: %s", uuidType)
	}
	if err != nil {
		log.Fatal(err)
	}
	return shortuuid.DefaultEncoder.Encode(u)
}
