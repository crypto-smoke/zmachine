package main

import (
	"fmt"
	"os"

	zmachine "github.com/inhies/zmachine"
)

func main() {
	if len(os.Args) < 2 {
		panic("please specify a file to load")
	}

	err := make(chan error, 1)
	zm := zmachine.New(os.Args[1], os.Stdin, os.Stdout, err, false)

	go func() {
		for {
			e := <-err
			fmt.Println("ERROR:", e)
		}
	}()

	zm.Run()
	//zm := zmachine.NewMachine(os.Stdin, os.Stdout)
	//zm.Initialize(buffer, header)

	//for !zm.Done {
	//zm.InterpretInstruction()
	//}

}
