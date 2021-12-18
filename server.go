package main

import (
	"context"
	"crypto/sha256"
	"io"
	"log"
	"net/http"
	"os"
)

func StartServer(bufferSize int) error {
	log.Println("Starting Server")
	server := http.Server{Addr: ":8899"}
	idleConnsClosed := make(chan bool)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("Request recieved from %s\n", r.RemoteAddr)

		// Send header
		rw.Header().Set("Special-Header", "XXSSXX")
		if value := r.Header.Get("Password-Header"); value == "" {
			log.Println("Header not set")
			return
		} else if value == "pass" {
			log.Println("Password Match")
		} else {
			log.Println("Password Dosen't Match")
			return
		}

		log.Printf("Sending Stream\n")
		hash := sha256.New()
		buffer := make([]byte, bufferSize)
		var total int64

		for {
			nRead, ioErr := io.ReadFull(os.Stdin, buffer)

			if ioErr == io.EOF {
				log.Println("EOF reached")
				break
			}

			if ioErr != io.ErrUnexpectedEOF && ioErr != nil {
				log.Println(ioErr)
				break
			}

			if nRead == bufferSize {
				rw.Write(buffer)
				hash.Write(buffer)
			} else {
				rw.Write(buffer[:nRead])
				hash.Write(buffer[:nRead])
			}

			total += int64(nRead)

			if ioErr == io.ErrUnexpectedEOF {
				log.Println("U EOF reached")
				break
			}

		}

		log.Printf("Total sent %dbytes  \n", total)
		log.Printf("Sum of the Hash - %x\n", hash.Sum(nil))
		idleConnsClosed <- true
	})

	waitToClose := make(chan bool)
	go func() {
		<-idleConnsClosed
		server.Shutdown(context.TODO())
		waitToClose <- true
	}()

	sErr := server.ListenAndServe()

	if sErr == http.ErrServerClosed {
		sErr = nil
	}

	<-waitToClose

	if sErr != nil {
		return sErr
	}

	return nil
}
