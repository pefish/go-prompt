package main

import (
	"fmt"
	go_prompt "github.com/pefish/go-prompt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	text, isExit := go_prompt.PromptInstance.Input("Please input name", []go_prompt.InputOption{
		{
			Text:        "test1",
			Description: "test1 description",
			IsDefault:   false,
		},
	})
	if isExit {
		return
	}
	fmt.Println(text)
	go func() {
		exitChan := make(chan os.Signal)
		signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
		for {
			select {
			case <-exitChan:
				fmt.Println(22)
			}
		}
	}()
	for {
		fmt.Println(1)
		time.Sleep(time.Second)
	}
}
