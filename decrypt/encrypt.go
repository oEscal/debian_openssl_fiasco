package main

import (
	"fmt"
	"math/big"
	"math/rand"
)


func encrypt(privateKeysMap map[string][]string) {

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
