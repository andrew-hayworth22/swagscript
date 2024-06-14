package repl

import (
	"bufio"
	"fmt"
	"io"
	"swagscript/lexer"
	"swagscript/parser"
)

const PROMPT = ">> "

const DOLLAR_ART = `___________________________________
|#######====================#######|
|#(1)*UNITED STATES OF AMERICA*(1)#|
|#**          /===\   ********  **#|
|*# {G}      | (") |             #*|
|#*  ******  | /v\ |    O N E    *#|
|#(1)         \===/            (1)#|
|##=========ONE DOLLAR===========##|
------------------------------------
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" || line == "quit" {
			io.WriteString(out, "See ya brodie\n")
			return
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, DOLLAR_ART)
	io.WriteString(out, "WOAH EASY DAWG! That SwagScript was whack...\n")
	io.WriteString(out, " Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
