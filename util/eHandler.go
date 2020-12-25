package util

import (
	"io"
	"log"
	"os"
)

func HandleErrors(data interface{}, err error) interface{} {
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err == io.EOF {
		os.Exit(200)
	}
	return data
}
