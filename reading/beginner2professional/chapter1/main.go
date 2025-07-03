package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func hello(helloList []string, index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}

func main() {
	helloList := []string{
		"hap",
		"ui",
		"Ge",
		"omki",
	}

	msg, err := hello(helloList, 1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msg)
	}

	fmt.Println(len(helloList))
	fmt.Println(helloList[len(helloList)-1])
	// fmt.Println(helloList[len(helloList)]) // This will cause a panic

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Use the new rand instance
	n := r.Intn(5) + 1
	stars := strings.Repeat("*", n)
	fmt.Println(stars)

	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}

