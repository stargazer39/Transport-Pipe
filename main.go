package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	log.SetOutput(NewColorWriter(os.Stderr))

	mode := flag.String("mode", "client", "Set the type (client/server)")
	address := flag.String("address", "127.0.0.1", "Enter an IP address / URL")
	bufSize := flag.String("b", "1M", "Enter a buffer size")
	flag.Parse()

	bufferSize := 1024 * 1024

	// Sanitize buffer size
	re := regexp.MustCompile("[0-9]+")
	log.Println(*bufSize)
	if strings.HasSuffix(*bufSize, "M") {
		value, valErr := strconv.Atoi(re.FindAllString(*bufSize, -1)[0])

		if valErr != nil {
			log.Panicln("Wrong -b")
		}

		bufferSize = value * 1024 * 1024
	} else if strings.HasSuffix(*bufSize, "K") {
		value, valErr := strconv.Atoi(re.FindAllString(*bufSize, -1)[0])

		if valErr != nil {
			log.Panicln("Wrong -b")
		}

		bufferSize = value * 1024
	} else {
		value, valErr := strconv.Atoi(re.FindAllString(*bufSize, -1)[0])

		if valErr != nil {
			log.Panicln("Wrong -b")
		}

		bufferSize = value
	}
	log.Printf("Buffer size %d", bufferSize)

	switch *mode {
	case "client":
		if cErr := StartClient(address, bufferSize); cErr != nil {
			log.Panicln(cErr)
		}
		log.Println("Successfully Read")
	case "server":
		if sErr := StartServer(bufferSize); sErr != nil {
			log.Panicln(sErr.Error())
		}

	default:
		log.Println("Wrong type (client/server)")
	}

}
