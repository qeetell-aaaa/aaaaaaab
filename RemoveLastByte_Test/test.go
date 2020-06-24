package main

import (
	"fmt"
	pkge "github.com/qeetell-aaaa/aaaaaaab"
)

func main () {
	aome := []byte {1, 2}
	
	bome := pkge.RemoveLastByte (aome)
	come := pkge.RemoveLastByte (bome)
	dome := pkge.RemoveLastByte (come)
	eome := pkge.RemoveLastByte (dome)
	fome := pkge.RemoveLastByte (eome)

	fmt.Println (aome, bome, come, dome, eome, fome)
}
