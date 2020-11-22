package main

import (
	"fmt"
	"time"
	"math/big"
	"math/rand"
)


func calculateMessageL(mL, initialMessage, dSender, mSender, mReceiver *big.Int, lenOriginalMessage int64) {
	r := make([]byte, 512 - lenOriginalMessage)
	rand.Read(r)

	message := big.NewInt(1)
	message.Add(
		initialMessage, new(big.Int).Mul(new(big.Int).SetBytes(r), new(big.Int).Exp(big.NewInt(256), 
		big.NewInt(lenOriginalMessage), nil)))	
	
	if message.Cmp(mSender) > 0 {
		calculateMessageL(mL, initialMessage, dSender, mSender, mReceiver, lenOriginalMessage)
	}
	println(message.Cmp(mSender))

	mL.Exp(message, dSender, mSender)
	println(mL.Cmp(mReceiver))
	
	if message.Cmp(mSender) > 0 {
		calculateMessageL(mL, initialMessage, dSender, mSender, mReceiver, lenOriginalMessage)
	}
}

func encrypt(infoMap map[string][]string) {
	rand.Seed(time.Now().UnixNano())

	dSender := big.NewInt(1)
	bigOne := big.NewInt(1)
	pMinus := big.NewInt(1)
	qMinus := big.NewInt(1)

	sender := "\"Brian York\""
	receiver := "\"Charlie Brown\""

	senderInfo := infoMap[sender]
	receiverInfo := infoMap[receiver]

	pSender, _ := new(big.Int).SetString(senderInfo[0], 10)
	qSender, _ := new(big.Int).SetString(senderInfo[1], 10)
	mSender, _ := new(big.Int).SetString(senderInfo[2], 10)
	eSender, _ := new(big.Int).SetString(senderInfo[3], 10)

	mReceiver, _ := new(big.Int).SetString(receiverInfo[2], 10)
	eReceiver, _ := new(big.Int).SetString(receiverInfo[3], 10)

	dSender.ModInverse(eSender, lcm(qMinus.Sub(qSender, bigOne), pMinus.Sub(pSender, bigOne)))
	messageOriginal := "Ola"
	messageOriginal += string(byte(0))
	initialMessage := big.NewInt(0)

	for i := 0; i < len(messageOriginal); i++ {
		initialMessage.Add(
			initialMessage, new(big.Int).Mul(big.NewInt(int64(byte(messageOriginal[i]))), 
			new(big.Int).Exp(big.NewInt(256), big.NewInt(int64(i)), nil)))
	}
	
	mL := big.NewInt(1)
	calculateMessageL(mL, initialMessage, dSender, mSender, mReceiver, int64(len(messageOriginal)))

	encryptedMessage := mL.Exp(mL, eReceiver, mReceiver)

	fmt.Printf("%s\t%s\t%s\n", sender, receiver, encryptedMessage.Text(10))
}
