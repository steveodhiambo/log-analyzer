package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main()  {

	path := flag.String("path", "app.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR and CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal("The Log file does not exist")
	}
	 defer f.Close()

	//bufio - is designed to read an input stream from a network connection
	r := bufio.NewReader(f)

	for  {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level){
			fmt.Println(s)
		}
	}


}