package client

import (
	"bufio"
	"crypto/sha256"
	"log"
	"net"
	"os"
)

func StartClient(address *string, bufferSize int, password string) error {
	conn, connErr := net.Dial("tcp", *address)

	if connErr != nil {
		return connErr
	}

	defer conn.Close()

	hash := sha256.New()
	buffer := make([]byte, bufferSize)
	var total int64

	stdOut := bufio.NewWriter(os.Stdout)
	connReader := bufio.NewReader(conn)

	defer stdOut.Flush()

	for {
		nRead, rErr := connReader.Read(buffer)

		if rErr != nil {
			log.Println(rErr)
			break
		}

		total += int64(nRead)

		if _, err := stdOut.Write(buffer[:nRead]); err != nil {
			log.Println(err)
			break
		}

		if _, err := hash.Write(buffer[:nRead]); err != nil {
			log.Println(err)
			break
		}
	}

	log.Printf("Total sent %dbytes  \n", total)
	log.Printf("Sum of the Hash - %x\n", hash.Sum(nil))

	return nil
}
