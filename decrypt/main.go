package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"math/big"
)


func lcm(m, n *big.Int) *big.Int {
	result := big.NewInt(1)
	mult := big.NewInt(1)
	gcd := big.NewInt(1)
	
	mult.Mul(m, n)
	gcd.GCD(nil, nil, m, n)
	result.Div(mult, gcd)

	return result
}


func main() {
	var senders []string
	var receivers []string
	var encryptedMessages []string

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

	dReceiver := big.NewInt(1)
	bigOne := big.NewInt(1)
	pMinus := big.NewInt(1)
	qMinus := big.NewInt(1)
	for i := range senders {
		sender := senders[i]
		receiver := receivers[i]

		receiverInfo := privateKeysMap[receiver]
		if receiverInfo[0] != "1" {
			pReceiver, _ := new(big.Int).SetString(receiverInfo[0], 10)
			qReceiver, _ := new(big.Int).SetString(receiverInfo[1], 10)
			mReceiver, _ := new(big.Int).SetString(receiverInfo[2], 10)
			eReceiver, _ := new(big.Int).SetString(receiverInfo[3], 10)
			
			fmt.Printf("From %s to %s: ", sender, receiver)

			dReceiver.ModInverse(eReceiver, lcm(qMinus.Sub(qReceiver, bigOne), pMinus.Sub(pReceiver, bigOne)))
	
			encryptedMessage, _ := new(big.Int).SetString(encryptedMessages[i], 10)
			mL := big.NewInt(1)
			mL.Exp(encryptedMessage, dReceiver, mReceiver)

			senderInfo := privateKeysMap[sender]
			
			mSender, _ := new(big.Int).SetString(senderInfo[2], 10)
			eSender, _ := new(big.Int).SetString(senderInfo[3], 10)
			message := big.NewInt(1)
			message.Exp(mL, eSender, mSender)
			messageBytes := message.Bytes()
			for mi := len(messageBytes) - 1; mi >= 0; mi-- {
				if messageBytes[mi] == 0 {
					break
				}
				print(string(messageBytes[mi]))
			} 
			println("\n")
		}
	}
}
