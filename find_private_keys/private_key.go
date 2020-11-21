package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"math/big"
)


func main() {

	publicModulusMap := make(map[string]string)
	publicExponentMap := make(map[string]string)
	factorsMap := make(map[string]string)

	// read files
	file, err := os.Open("rsa_public_info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		publicModulusMap[line[0]] = line[1]
		publicExponentMap[line[0]] = line[2]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	// read factors
	file2, err := os.Open("factors.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)

	for scanner2.Scan() {
		line := strings.Split(scanner2.Text(), "\t")
		factorsMap[line[0]] = line[2]
		factorsMap[line[1]] = line[2]
	}

	if err := scanner2.Err(); err != nil {
		log.Fatal(err)
	}


	q := big.NewInt(1)
	bigOne := big.NewInt(1)
	for name := range factorsMap {
		p, _ := new(big.Int).SetString(factorsMap[name], 10)
		privateKey, _ := new(big.Int).SetString(publicModulusMap[name], 10)

		if p.Cmp(bigOne) != 0 {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\n", 
				name, p.Text(10), q.Quo(privateKey, p).Text(10), publicModulusMap[name], publicExponentMap[name])
		}
	}
}
