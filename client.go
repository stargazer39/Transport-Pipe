package main

import (
	"bufio"
	"crypto/sha256"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func StartClient(address *string, bufferSize int, password string) error {
	log.Printf("Starting client on %s...", *address)
	// tr := http.Transport{
	// 	IdleConnTimeout:       time.Hour * 999999,
	// 	ResponseHeaderTimeout: time.Hour * 999999,
	// 	TLSHandshakeTimeout:   time.Hour * 999999,
	// 	ExpectContinueTimeout: time.Hour * 999999,
	// }
	client := http.Client{
		// Transport: &tr,
	}

	var resp *http.Response
	var resErr error
	for {
		req, rErr := http.NewRequest("GET", *address, nil)

		if rErr != nil {
			return rErr
		}

		req.Header.Set("Password-Header", password)
		resp, resErr = client.Do(req)

		if resErr != nil {
			log.Println(resErr)
			time.Sleep(time.Second)
			continue
		}
		break
	}

	defer resp.Body.Close()

	hash := sha256.New()
	buffer := make([]byte, bufferSize)
	var total int64

	stdOut := bufio.NewWriter(os.Stdout)

	defer stdOut.Flush()

	for {
		nRead, ioErr := io.ReadFull(resp.Body, buffer)

		if ioErr == io.EOF {
			log.Println("EOF reached")
			break
		}

		if ioErr != io.ErrUnexpectedEOF && ioErr != nil {
			log.Println(ioErr)
			break
		}

		if nRead == bufferSize {
			stdOut.Write(buffer)
			hash.Write(buffer)
		} else {
			stdOut.Write(buffer[:nRead])
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

	return nil
}
