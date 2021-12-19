package server

import (
	"bufio"
	"crypto/sha256"
	"log"
	"net"
	"os"
)

func StartServer(bufferSize int, password string) error {
	log.Println("Starting Server")

	serv, tcpErr := net.Listen("tcp", ":8899")

	if tcpErr != nil {
		return tcpErr
	}

	defer serv.Close()

	conn, connErr := serv.Accept()

	if connErr != nil {
		return connErr
	}

	connWriter := bufio.NewWriter(conn)
	defer connWriter.Flush()

	hash := sha256.New()
	buffer := make([]byte, bufferSize)
	var total int64

	for {
		nRead, rErr := os.Stdin.Read(buffer)

		if rErr != nil {
			log.Println(rErr)
			break
		}

		total += int64(nRead)

		if _, err := connWriter.Write(buffer[:nRead]); err != nil {
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
