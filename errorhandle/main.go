package main

import (
	"fmt"
	"log"
)

func main() {
	myDB, err := NewMyDB("root", "pass", "localhost", 3306, "test")
	if err != nil {
		log.Fatal("db connection occur error : ", err)
	}
	u, err := myDB.Get(2)
	if err != nil {
		log.Fatal("db query occur error : ", err)
	}
	fmt.Println(u)
	myDB.Close()
}
