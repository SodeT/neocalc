// TODO:
// write better tests
// switch to a arbitrary precision math lbrary (math/big)

package main

import (
	"bufio"
	"fmt"
	"neocalc/src/ast"
	"neocalc/src/command"
	"neocalc/src/runtime"
	"neocalc/src/tokenizer"
	"neocalc/src/utils"
	"neocalc/src/validate"
	"os"
)

var (
	version = "1.0"
	fileAsInput = false

)

func main() {
	handleArgs(os.Args)
	if fileAsInput && len(os.Args) == 3 {
		file, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Println("Failed to load file:", os.Args[2])
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			runInput(scanner.Text())
		}
		return
	}

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		input = ""
		fmt.Print("; ")
		if scanner.Scan() {
			input = scanner.Text()
		}
		runInput(input)
	}
}

func handleArgs(args []string) {
	if len(args) <= 1 {
		return
	}
	switch args[1] {
	case "--help", "-h":
		fmt.Println(utils.HelpText)
		os.Exit(0)
	case "--version", "-v":
		fmt.Println(version)
		os.Exit(0)
	case "--input", "-i":
		fileAsInput = true
	}
}

func runInput(input string) float64 {
	ans := 0.0
	msg := utils.NilMsg

	msg = validate.Input(input)
	if msg.Level == utils.ERR_LOG {
		displayMsg(msg)
		return 0
	}

	if []rune(input)[0] == '#' {
		ans, msg = command.Run(input)
	} else {
		utils.Input = input
		toks := tokenizer.Tokenize(input)
		msg := validate.Tokens(toks)
		if msg.Level == utils.ERR_LOG {
			displayMsg(msg)
			return 0
		}

		ast := ast.Parse(toks)
		ans, msg = runtime.Execute(ast)
	}

	fmt.Println(">>>", ans)
	displayMsg(msg)
	return ans
}

func displayMsg(msg utils.Message) {
	switch msg.Level {
	case utils.NIL_LOG:
		return
	case utils.INFO_LOG:
		fmt.Println("INFO: ", msg.Message)
	case utils.WARN_LOG:
		fmt.Println("WARNING: ", msg.Message)
	case utils.ERR_LOG:
		fmt.Println("ERROR: ", msg.Message)
	}
}
