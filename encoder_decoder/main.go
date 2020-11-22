package main

import (
	"os"
	"fmt"
	"flag"
)


func main() {
	// arguments
	var encryptFlag bool
	var infoFileName string
	var interceptedMessagesFileName string
	var sender string
	var receiver string
	var messageOriginal string

	flag.BoolVar(&encryptFlag, "e", false, "Encrypt message")
	flag.StringVar(&infoFileName, "i", "info.txt", "File where to get the public and private info of each user")
	flag.StringVar(&interceptedMessagesFileName, "m", "intercepted.txt", "File where to get the intercepted messages (just valid for decrypt)")
	flag.StringVar(&sender, "sender", "Brian York", "Message sender (just valid for encrypt)")
	flag.StringVar(&receiver, "receiver", "Charlie Brown", "Message receiver (just valid for encrypt)")
	flag.StringVar(&messageOriginal, "message", "", "Message to encrypt (just valid for encrypt)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		println("The default behaviour is to run the decrypt of messages")
	}
	flag.Parse()

	
	// read info file
	infoMap := make(map[string][]string)
	readInfo(infoFileName, infoMap)


	// encrypt or decrypt
	if encryptFlag {
		encrypt(infoMap, "\"" + sender + "\"", "\"" + receiver + "\"", messageOriginal)
	} else {
		// read intercepted messages file
		var senders []string
		var receivers []string
		var encryptedMessages []string
		readInterceptedMessages(interceptedMessagesFileName, &senders, &receivers, &encryptedMessages)
		
		decrypt(senders, receivers, encryptedMessages, infoMap)
	}
}
