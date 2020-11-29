package main

import (
	"fmt"
	"flag"
	"time"
	"sync"
	"math/big"
)


var bigOne = big.NewInt(1)


func gcdBetweenTwoPersons(privateInfoMap, publicInfoMap map[string][]string, person1, person2 string, 
									wg *sync.WaitGroup, semaphore *sync.Mutex) {
	publicInfoPerson1 := publicInfoMap[person1]
	var privateKeyPerson1, _ = new(big.Int).SetString(publicInfoPerson1[0], 10)
	
	publicInfoPerson2 := publicInfoMap[person2]
	var privateKeyPerson2, _ = new(big.Int).SetString(publicInfoPerson2[0], 10)

	q := big.NewInt(1)
	p := big.NewInt(1)
	
	q.GCD(nil, nil, privateKeyPerson1, privateKeyPerson2)
	if q.Cmp(bigOne) != 0 {
		semaphore.Lock()

		// person, q, p, m, public exponent
		privateInfoMap[person1] = []string{
			q.Text(10), p.Quo(privateKeyPerson1, q).Text(10), publicInfoPerson1[0], publicInfoPerson1[1]}
		privateInfoMap[person2] = []string{
			q.Text(10), p.Quo(privateKeyPerson2, q).Text(10), publicInfoPerson2[0], publicInfoPerson2[1]}

		semaphore.Unlock()
	} 
	wg.Done()
}


func crackPublicInfo(privateInfoMap, publicInfoMap map[string][]string) {
	var wg sync.WaitGroup
	var semaphore sync.Mutex

	var people []string
	for person := range publicInfoMap {
		people = append(people, person)
	}

	for i1 := range people[:len(people) - 1] {
		for i2 := i1 + 1; i2 < len(people); i2++ {
			wg.Add(1)
			go gcdBetweenTwoPersons(privateInfoMap, publicInfoMap, people[i1], people[i2], &wg, &semaphore)
		}
	}
	wg.Wait()
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
