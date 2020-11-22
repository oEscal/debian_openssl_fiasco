package main


import (
	"os"
	"log"
	"bufio"
	"strings"
	"math/big"
)


func readInfo(fileName string, infoMap map[string][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		infoMap[line[0]] = make([]string, 4)
		infoMap[line[0]][0] = line[1]
		infoMap[line[0]][1] = line[2]
		infoMap[line[0]][2] = line[3]
		infoMap[line[0]][3] = line[4]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func readInterceptedMessages(fileName string, senders, receivers, encryptedMessages *[]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		*senders = append(*senders, line[0])
		*receivers = append(*receivers, line[1])
		*encryptedMessages = append(*encryptedMessages, line[2])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func lcm(m, n *big.Int) *big.Int {
	result := big.NewInt(1)
	mult := big.NewInt(1)
	gcd := big.NewInt(1)
	
	mult.Mul(m, n)
	gcd.GCD(nil, nil, m, n)
	result.Div(mult, gcd)

	return result
}
