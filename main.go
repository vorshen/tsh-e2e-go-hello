package main

import "fmt"

import "github.com/google/uuid"

func main() {
	id := uuid.MustParse("4f51c3d2-d7c7-4ef8-98d0-1ca621f45595")
	if id.Version() != 4 {
		panic("fixture UUID is not version 4")
	}
	fmt.Println("tsh-e2e-go-modules-ok")
}
