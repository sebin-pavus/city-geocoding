package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("./../../input")

	if err != nil {
		log.Fatal(err)
	}

	ans := string(content)
	pri := strings.Split(ans, "\r")

	f, err := os.Create("./../../output")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i := 0; i < len(pri); i++ {
		_, err2 := f.WriteString(pri[i] + "\r")

		if err2 != nil {
			log.Fatal(err2)
		}

	}

	fmt.Println("done")
}
