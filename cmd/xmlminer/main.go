package main

import (
	"log"

	"github.com/ayuxsec/burp-xml-miner/internal/app/xmlminer/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
