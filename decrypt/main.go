package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"flag"
	"strings"
)


func main() {
	var senders []string
	var receivers []string
	var encryptedMessages []string

	// arguments
	var encryptFlag bool
	flag.BoolVar(&encryptFlag, "e", false, "Encrypt message")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		println("The default behaviour is to run the decrypt of messages")
	}
	flag.Parse()

	// read file
	file, err := os.Open("my_messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		senders = append(senders, line[0])
		receivers = append(receivers, line[1])
		encryptedMessages = append(encryptedMessages, line[2])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}



	privateKeysMap := make(map[string][]string)

	// read files
	file2, err := os.Open("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)

	for scanner2.Scan() {
		line := strings.Split(scanner2.Text(), "\t")
		privateKeysMap[line[0]] = make([]string, 4)
		privateKeysMap[line[0]][0] = line[1]
		privateKeysMap[line[0]][1] = line[2]
		privateKeysMap[line[0]][2] = line[3]
		privateKeysMap[line[0]][3] = line[4]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(encryptFlag)
	if encryptFlag {
		encrypt(privateKeysMap)
	} else {
		decrypt(senders, receivers, encryptedMessages, privateKeysMap)
	}
}
