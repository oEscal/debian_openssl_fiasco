package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"math/big"
	"math/rand"
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
	file, err := os.Open("intercepted.txt")
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

	dSender := big.NewInt(1)
	bigOne := big.NewInt(1)
	pMinus := big.NewInt(1)
	qMinus := big.NewInt(1)

	sender := "\"Brian York\""
	receiver := "\"Charlie Brown\""

	senderInfo := privateKeysMap[sender]
	receiverInfo := privateKeysMap[receiver]

	pSender, _ := new(big.Int).SetString(senderInfo[0], 10)
	qSender, _ := new(big.Int).SetString(senderInfo[1], 10)
	mSender, _ := new(big.Int).SetString(senderInfo[2], 10)
	eSender, _ := new(big.Int).SetString(senderInfo[3], 10)

	mReceiver, _ := new(big.Int).SetString(receiverInfo[2], 10)
	eReceiver, _ := new(big.Int).SetString(receiverInfo[3], 10)

	dSender.ModInverse(eSender, lcm(qMinus.Sub(qSender, bigOne), pMinus.Sub(pSender, bigOne)))
	messageOriginal := "Ola"
	messageOriginal += string(byte(0))
	message := big.NewInt(0)

	for i := 0; i < len(messageOriginal); i++ {
		message.Add(message, big.NewInt(1).Mul(big.NewInt(int64(byte(messageOriginal[i]))), big.NewInt(1).Exp(big.NewInt(256), big.NewInt(int64(i)), nil)))
	}

	r := make([]byte, 512 - len(messageOriginal))
	rand.Read(r)
	
	message.Add(message, big.NewInt(1).Mul(new(big.Int).SetBytes(r), big.NewInt(1).Exp(big.NewInt(256), big.NewInt(int64(len(messageOriginal))), nil)))
	println(message.Cmp(mSender))

	mL := big.NewInt(1)
	mL.Exp(message, dSender, mSender)
	println(mL.Cmp(mReceiver))

	encryptedMessage := mL.Exp(mL, eReceiver, mReceiver)

	fmt.Printf("%s\t%s\t%s\n", sender, receiver, encryptedMessage.Text(10))
}
