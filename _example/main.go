package main

import (
	"fmt"
	go_shell "github.com/pefish/go-shell"
	"log"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	cmd := go_shell.NewCmd(`
#!/bin/bash

for i in {1..100}
do
 echo $i
done

gdrtg
`)
	resultChan := make(chan string)
	go func() {
		for {
			select {
			case r := <-resultChan:
				fmt.Println(r)
			}
		}
	}()
	err := go_shell.ExecForResultLineByLine(cmd, resultChan)
	if err != nil {
		return err
	}
	return nil
}
