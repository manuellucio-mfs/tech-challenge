package filemanager

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile(pathFile string) (string, error) {
	databytes, err := ioutil.ReadFile(pathFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// convertir el arreglo a string
	dataString := string(databytes)

	fmt.Println(dataString)
	return dataString, nil
}
