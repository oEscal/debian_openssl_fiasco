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

	flag.BoolVar(&encryptFlag, "e", false, "Encrypt message")
	flag.StringVar(&infoFileName, "i", "info.txt", "File where to get the public and private info of each user")
	flag.StringVar(&interceptedMessagesFileName, "m", "intercepted.txt", "File where to get the intercepted messages")

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
		encrypt(infoMap)
	} else {
		// read intercepted messages file
		var senders []string
		var receivers []string
		var encryptedMessages []string
		readInterceptedMessages(interceptedMessagesFileName, &senders, &receivers, &encryptedMessages)
		
		decrypt(senders, receivers, encryptedMessages, infoMap)
	}
}
