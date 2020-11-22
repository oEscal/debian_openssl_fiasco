package main

import (
	"fmt"
	"math/big"
)


func decrypt(senders, receivers, encryptedMessages []string, privateKeysMap map[string][]string) {

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
