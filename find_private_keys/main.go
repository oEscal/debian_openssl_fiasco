package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"math/big"
)


type void struct{}
var dummyVoid void


func main() {
	publicModulusMap := make(map[string][]string)

	// read file
	file, err := os.Open("rsa_public_info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		publicModulusMap[line[0]] = make([]string, 2)
		publicModulusMap[line[0]][0] = line[1]
		publicModulusMap[line[0]][1] = line[2]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	bigOne := big.NewInt(1)
	foundPeople := make(map[string]void)
	for person1 := range publicModulusMap {
		publicInfoPerson1 := publicModulusMap[person1]
		var privateKeyPerson1, _ = new(big.Int).SetString(publicInfoPerson1[0], 10)

		flag := false

		for person2 := range(publicModulusMap) {
			if flag {
				go func(person1, person2 string) {
					q := big.NewInt(1)
					p := big.NewInt(1)
					
					publicInfoPerson2 := publicModulusMap[person2]
					var privateKeyPerson2, _ = new(big.Int).SetString(publicInfoPerson2[0], 10)
					
					q.GCD(nil, nil, privateKeyPerson1, privateKeyPerson2)
					if q.Cmp(bigOne) != 0 {
						// person, q, p, m, public exponent
						fmt.Printf("%s\t%s\t%s\t%s\t%s\n", 
							person1, q.Text(10), p.Quo(privateKeyPerson1, q).Text(10), publicInfoPerson1[0], publicInfoPerson1[1])
						fmt.Printf("%s\t%s\t%s\t%s\t%s\n", 
							person2, q.Text(10), p.Quo(privateKeyPerson2, q).Text(10), publicInfoPerson2[0], publicInfoPerson2[1])

						foundPeople[person1] = dummyVoid
						foundPeople[person2] = dummyVoid
					} 
				} (person1, person2)
			}

			if person1 == person2 {
				flag = true
			}
		}
	}

	// save the people we could not find
	for person := range publicModulusMap {
		if _, ok := foundPeople[person]; !ok {
			fmt.Printf("%s\t%d\t%d\t%s\t%s\n", person, 1, 1, publicModulusMap[person][0], publicModulusMap[person][1])
		}
	}
}
