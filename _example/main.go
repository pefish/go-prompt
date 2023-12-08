package main

import (
	"fmt"
	go_prompt "github.com/pefish/go-prompt"
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
}
