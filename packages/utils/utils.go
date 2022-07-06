package utils

import (
	"log"
	"os"
)

func GetCliArg(argNum int) string {
	args := os.Args
	if len(args) == argNum {
		return ""
	}
	arg := os.Args[argNum]
	log.Println("Found arg:", arg)
	return arg

}
