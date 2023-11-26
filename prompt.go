package go_prompt

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

type Prompt struct {
}

type InputOption struct {
	Text        string
	Description string
	IsDefault   bool
}

func (p *Prompt) Input(
	inputTip string,
	options []InputOption,
) (inputText string, isExit bool) {
	defaultText := ""
	suggests := make([]prompt.Suggest, 0)
	for _, o := range options {
		if o.IsDefault && defaultText == "" {
			defaultText = o.Text
		}
		suggests = append(suggests, prompt.Suggest{
			Text:        o.Text,
			Description: o.Description,
		})
	}

	fmt.Println(inputTip)
	inputText = prompt.Input(
		">>> ",
		func(d prompt.Document) []prompt.Suggest {
			return prompt.FilterHasPrefix(
				suggests,
				d.GetWordBeforeCursor(),
				true,
			)
		},
		prompt.OptionInitialBufferText(defaultText),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionOnDown(),
	)
	if inputText == "exit" {
		return inputText, true
	}
	return inputText, false
}
