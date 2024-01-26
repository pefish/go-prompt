package go_prompt

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/pkg/term/termios"
	"syscall"
)

type Prompt struct {
}

type InputOption struct {
	Text        string
	Description string
	IsDefault   bool
}

var PromptInstance = NewPrompt()

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) Input(
	inputTip string,
	options []InputOption,
) (inputText string, isExit bool) {
	// store
	fd, err := syscall.Open("/dev/tty", syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	originalTermios, err := termios.Tcgetattr(uintptr(fd))
	if err != nil {
		panic(err)
	}

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
	pInstance := prompt.New(
		func(s string) {
			//fmt.Println(s)
		},
		func(document prompt.Document) []prompt.Suggest {
			return prompt.FilterHasPrefix(
				suggests,
				document.GetWordBeforeCursor(),
				true,
			)
		},
		prompt.OptionPrefix(">>> "),
		prompt.OptionInitialBufferText(defaultText),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionOnDown(),
	)
	//pInstance.Run()
	inputText = pInstance.Input()

	// recover
	err = termios.Tcsetattr(uintptr(fd), termios.TCSANOW, originalTermios)
	if err != nil {
		panic(err)
	}

	if inputText == "exit" {
		return inputText, true
	}

	return inputText, false
}
