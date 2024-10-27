package main

import (
	"fmt"
	"log"

	go_shell "github.com/pefish/go-shell"
)

func main() {
	err := do1()
	if err != nil {
		log.Fatal(err)
	}
}

func do1() error {
	cmd := go_shell.NewCmd("sudo docker logs redis --tail 20")
	result, err := go_shell.ExecForResult(cmd)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func do() error {
	cmd := go_shell.NewCmd(`
#!/bin/bash
set -euxo pipefail

for i in {1..100}
do
 echo $i
done

# cd ~/svdfs/serg
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
