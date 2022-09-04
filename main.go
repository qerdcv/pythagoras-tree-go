package main

import "log"

func main() {
	if err := newGame().Run(); err != nil {
		log.Fatal(err.Error())
	}
}
