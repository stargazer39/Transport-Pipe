package main

import (
	"crypto/sha256"
	"flag"
	"log"
	"os"
	"regexp"
	"stargazer/transport-pipe/client"
	"stargazer/transport-pipe/color"
	"stargazer/transport-pipe/server"
	"strconv"
)

func main() {
	log.SetOutput(color.NewWriter(os.Stderr))

	mode := flag.String("mode", "client", "Set the type (client/server)")
	address := flag.String("address", "127.0.0.1", "Enter an IP address / URL")
	bufSize := flag.String("b", "1M", "Enter a buffer size")
	password := flag.String("password", "pass", "Server / Client password")
	flag.Parse()

	_password := *password
	// Sanitize buffer size
	re := regexp.MustCompile("[0-9]+")
	units := []string{"K", "M", "G"}

	unit := (*bufSize)[len(*bufSize)-1:]
	value, valErr := strconv.Atoi(re.FindAllString(*bufSize, -1)[0])

	if valErr != nil {
		log.Panicln("Wrong -b")
	}

	bufferSize := value

	if !re.Match([]byte(unit)) {
		for _, v := range units {
			bufferSize *= 1024
			if v == unit {
				break
			}
		}
	}

	log.Printf("Buffer size %d", bufferSize)

	switch *mode {
	case "client":
		if cErr := client.StartClient(address, bufferSize, _password); cErr != nil {
			log.Panicln(cErr.Error())
		}
		log.Println("Successfully Read")
	case "server":
		if sErr := server.StartServer(bufferSize, _password); sErr != nil {
			log.Panicln(sErr.Error())
		}

	default:
		log.Println("Wrong type (client/server)")
	}

}

func String2Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return string(h.Sum(nil))
}
