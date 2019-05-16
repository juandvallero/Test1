/********************************************************
 *
 * Autor: Juan del Valle
 *
 * Proyecto: Prueba código 1
 *
 * Fichero:  main.go
 *
 * Descripción:
 *      Lee y escribe un fichero empleando dos tareas independientes.
 *
 * *****************************************************/
package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// Datatype for data read
type TData struct {
	buffer []byte
	size   int64
}

func main() {
	// Get line commands arguments
	argsWithProg := os.Args
	if len(argsWithProg) != 3 {
		fmt.Printf("Usage: %s fileIn fileOut", argsWithProg[0])
		os.Exit(1)
	}

	// Define channels for thread synchronization
	writeEnd := make(chan bool)            // Write finish event
	bufferReaded := make(chan TData, 1000) // Buffer read event, with size for store 1000 data chunks

	// Launch Threads
	go writeFile("./out.jpg", writeEnd, bufferReaded)
	go readFile("./test.jpg", bufferReaded)

	// Wait until write operation finish
	<-writeEnd
	fmt.Println("Write File Finished")
}

// ReadFile function
func readFile(filename string, bufferReaded chan TData) {
	bufferSize := 102400

	// Open the file
	file, err := os.Open(filename)
	exitOnError(err)

	// Closing file on finish
	defer file.Close()

	// Read the content of the file
	for {
		// Allocate buffer
		buffer := make([]byte, bufferSize)

		// Read chunk
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				exitOnError(err)
			}
			break
		}

		fmt.Println("bytes Read: ", bytesRead)
		// Send data to write thread
		bufferReaded <- TData{size: int64(bytesRead), buffer: buffer}
	}

	fmt.Println("Read File closed")

	// Send read finish notification
	bufferReaded <- TData{size: int64(-1), buffer: nil}
}

// WriteFile function
func writeFile(filename string, writeEnd chan bool, bufferReaded chan TData) {
	// Open the file
	var file *os.File = nil
	var err error = nil
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var position int64 = 0
	for {
		// Wait for event
		data := <-bufferReaded
		if file == nil {
			file, err = os.Create(filename)
			exitOnError(err)

			// Closing file on finish
			defer file.Close()
		}

		// Check end of read
		if data.size == -1 {
			fmt.Println("Write File closed")
			writeEnd <- true
		} else {
			// Data Received
			n, writeErr := file.WriteAt(data.buffer[:data.size], position)
			exitOnError(writeErr)

			fmt.Println("bytes Writed: ", n)
			position += data.size
		}

		// Simulating Write delay between 0 and 500 ms
		time.Sleep(time.Duration(r.Intn(500)) * time.Millisecond)
	}
}

// Error function
func exitOnError(e error) {
	if e != nil {
		panic(e)
	}
}
