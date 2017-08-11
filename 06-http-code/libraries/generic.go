package libraries

import (
	"log"
	"os"
	"io"
)

var file io.Writer
func init() {
	var err error
	file, err = os.OpenFile("error.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)
}

func CheckError(err error){
	if err != nil {
		log.Println(err.Error())
	}
}