package main


import (
	"os"
	"bufio"
	"log"
)

func main() {

	// read file
	file, err := os.Open("rsa_public_info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}



}
