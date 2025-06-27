package main

import (
	"log"
	"os"
)

// func main() {
// 	myfile, es := os.Create("helloo.txt")
// 	if es != nil {
// 		log.Fatal(es)
// 	}

// 	log.Println(myfile)
// 	myfile.Close()
// }

func main() {
	// fileName := "hellow.txt"
	// myfile, es := os.Stat(fileName)
	// if es != nil {
	// 	if os.IsNotExist(es) {
	// 		// log.Fatal("File not found")
	// 		os.Create(fileName)
	// 		log.Println("File not found so I created it for you.")
	// 	} else {log.Fatal(es)}
	// }
	// log.Println("File Exist")
	// log.Println("File Detail is:")
	// log.Println("Name is: ", myfile.Name())
	// log.Println("Size is: ", myfile.Size())

	if er := os.Mkdir("a", os.ModePerm) ; er != nil {
		log.Fatal(er)
	}
}