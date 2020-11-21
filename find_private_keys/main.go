package main

import (
	"fmt"
	"time"
	"math/big"
)


type void struct{}
var dummyVoid void


func main() {
	publicModulusMap := make(map[string][]string)

	// read file
	readPublicInfo("rsa_public_info.txt", publicModulusMap)
	

	init := time.Now()

	bigOne := big.NewInt(1)
	foundPeople := make(map[string][]string)
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
						foundPeople[person1] = []string{
							q.Text(10), p.Quo(privateKeyPerson1, q).Text(10), publicInfoPerson1[0], publicInfoPerson1[1]}
						foundPeople[person2] = []string{
							q.Text(10), p.Quo(privateKeyPerson2, q).Text(10), publicInfoPerson2[0], publicInfoPerson2[1]}
					} 
				} (person1, person2)
			}

			if person1 == person2 {
				flag = true
			}
		}
	}

	fmt.Printf("Done in %f s!\n", time.Now().Sub(init).Seconds())


	// save the results 
	saveResults("results2.txt", publicModulusMap, foundPeople)	
}
