package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

func main() {
	aDoc := Doctor{
		number:    43,
		actorName: "anfaas",
		companions: []string{
			"rayaan", "amaan",
		},
	}
	fmt.Println(aDoc.actorName)
}
