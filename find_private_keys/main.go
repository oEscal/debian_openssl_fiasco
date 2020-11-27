package main

import (
	"fmt"
	"flag"
	"time"
	"sync"
	"math/big"
)


func crackPublicInfo(privateInfoMap, publicInfoMap map[string][]string) {
	var semaphore sync.RWMutex

	bigOne := big.NewInt(1)
	for person1 := range publicInfoMap {
		publicInfoPerson1 := publicInfoMap[person1]
		var privateKeyPerson1, _ = new(big.Int).SetString(publicInfoPerson1[0], 10)

		flag := false

		for person2 := range(publicInfoMap) {
			if flag {
				go func(person1, person2 string) {
					q := big.NewInt(1)
					p := big.NewInt(1)
					
					publicInfoPerson2 := publicInfoMap[person2]
					var privateKeyPerson2, _ = new(big.Int).SetString(publicInfoPerson2[0], 10)
					
					q.GCD(nil, nil, privateKeyPerson1, privateKeyPerson2)
					if q.Cmp(bigOne) != 0 {
						semaphore.Lock()
						defer semaphore.Unlock()

						// person, q, p, m, public exponent
						privateInfoMap[person1] = []string{
							q.Text(10), p.Quo(privateKeyPerson1, q).Text(10), publicInfoPerson1[0], publicInfoPerson1[1]}
						privateInfoMap[person2] = []string{
							q.Text(10), p.Quo(privateKeyPerson2, q).Text(10), publicInfoPerson2[0], publicInfoPerson2[1]}
					} 
				} (person1, person2)
			}

			if person1 == person2 {
				flag = true
			}
		}
	}
}


func main() {
	// arguments
	var publicInfoFileName string
	var resultsFileName string

	flag.StringVar(&publicInfoFileName, "p", "rsa_public_info.txt", "File where are saved the public RSA info")
	flag.StringVar(&resultsFileName, "r", "info.txt", "File where to save the resultant private and public information")
	flag.Parse()


	// privateInfoMap is the resultant private info cracked; the publicInfoMap is the public info obtained from the document passed as argument
	privateInfoMap := make(map[string][]string)
	publicInfoMap := make(map[string][]string)


	// read file
	readPublicInfo(publicInfoFileName, publicInfoMap)
	

	// crack the public info
	init := time.Now()
	crackPublicInfo(privateInfoMap, publicInfoMap)
	fmt.Printf("Done in %f s!\n", time.Now().Sub(init).Seconds())


	// save the results 
	saveResults(resultsFileName, publicInfoMap, privateInfoMap)	
}
