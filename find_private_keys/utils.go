package main

import(
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)


func readPublicInfo(fileName string, publicModulusMap map[string][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		publicModulusMap[line[0]] = make([]string, 2)
		publicModulusMap[line[0]][0] = line[1]
		publicModulusMap[line[0]][1] = line[2]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func saveResults(fileName string, publicModulusMap, foundPeople map[string][]string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// save the people we could not find
	for person := range publicModulusMap {
		info, ok := foundPeople[person]
		if !ok {
			fmt.Fprintf(file, "%s\t%d\t%d\t%s\t%s\n", person, 1, 1, publicModulusMap[person][0], publicModulusMap[person][1])
		} else {
			fmt.Fprintf(file, "%s\t%s\t%s\t%s\t%s\n", person, info[0], info[1], info[2], info[3])
		}
	}
}
