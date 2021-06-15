package main

import (
	"fmt"
	"go-training/cmd"
	"go-training/entity"
	"go-training/internal"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	s, err := cmd.NewStore(100)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer s.Free()

	_ = s.Bootstrap(10)
	id_1 := internal.UUID()

	if err := s.AddPeople(&entity.People{
		Id:      id_1,
		Name:    "Nguyen Van A",
		Age:     30,
		Company: "FPT Software",
		Address: "17 Duy Tan",
	}); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Print Asc Order")
	s.PrintAsOrder(internal.Desc)
}
