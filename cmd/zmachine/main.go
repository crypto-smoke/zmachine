package main

import (
	"bufio"
	"fmt"
	"os"

	zmachine "github.com/inhies/zmachine"
)

func main() {
	if len(os.Args) < 2 {
		panic("please specify a file to load")
	}

	in := make(chan string, 1)
	out := make(chan string, 1)
	err := make(chan error, 1)
	zm := zmachine.New(os.Args[1], in, out, err)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in <- scanner.Text()
		}
	}()

	go func() {
		for {
			select {
			case e := <-err:
				fmt.Println("ERROR:", e)
			case txt := <-out:
				//fmt.Println(len(txt))
				fmt.Print(txt)
			}
		}
	}()

	zm.Run()
	//zm := zmachine.NewMachine(os.Stdin, os.Stdout)
	//zm.Initialize(buffer, header)

	//for !zm.Done {
	//zm.InterpretInstruction()
	//}

}
