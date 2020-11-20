package main


import (
	"os"
	"bufio"
	"log"
	"strings"
	"math/big"
)

func main() {
	publicModulusMap := make(map[string]string)

	// read file
	file, err := os.Open("rsa_public_info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		publicModulusMap[line[0]] = line[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	bigOne := big.NewInt(1)
	gcd := big.NewInt(1)
	for person1 := range publicModulusMap {
		var num, _ = new(big.Int).SetString(publicModulusMap[person1], 10)
		
		for person2 := range(publicModulusMap) {
			go func(person1, person2 string) {
				if person1 != person2 {
					var num2, _ = new(big.Int).SetString(publicModulusMap[person2], 10)
					
					gcd.GCD(nil, nil, num, num2)
					if gcd.Cmp(bigOne) != 0 {
						println(person1, person2, gcd.Text(10))
					}
				}
			} (person1, person2)
		}
	}
}
