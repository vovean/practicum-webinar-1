package main

import (
	"log"
	"sandbox/cmd/dep_inv/repository/inmemory"
	"sandbox/cmd/dep_inv/service"
)

func main() {
	s := &service.Service{R: &inmemory.Repo{}}
	s.SetBalance(123, 100)
	money, err := s.GetBalance(123)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("money: %d\n", money)
}
