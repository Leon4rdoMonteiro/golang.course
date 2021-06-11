package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Leon4rdoMonteiro/command-line-app/app"
)

func main() {
	fmt.Printf("Searching for results...\n\n")

	time.Sleep(1 * time.Second)

	application := app.Generate()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
