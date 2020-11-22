package main

import(
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)


func readPublicInfo(fileName string, publicInfoMap map[string][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		publicInfoMap[line[0]] = make([]string, 2)
		publicInfoMap[line[0]][0] = line[1]
		publicInfoMap[line[0]][1] = line[2]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func saveResults(fileName string, publicInfoMap, privateInfoMap map[string][]string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// save the people we could not find
	for person := range publicInfoMap {
		info, ok := privateInfoMap[person]
		if !ok {
			fmt.Fprintf(file, "%s\t%d\t%d\t%s\t%s\n", person, 1, 1, publicInfoMap[person][0], publicInfoMap[person][1])
		} else {
			fmt.Fprintf(file, "%s\t%s\t%s\t%s\t%s\n", person, info[0], info[1], info[2], info[3])
		}
	}
}
