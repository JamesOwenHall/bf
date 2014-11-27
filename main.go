package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Program holds all data for a single instance of a Brainfuck program
type Program struct {
	Input          []byte
	Memory         []byte
	InstructionPtr int
	DataPtr        int
}

// Run executes the program
func (p *Program) Run() {
	for p.InstructionPtr < len(p.Input) {
		// Skip non-instructions
		if !p.isInstruction(p.Input[p.InstructionPtr]) {
			p.InstructionPtr++
			continue
		}

		switch p.Input[p.InstructionPtr] {
		case '>':
			p.DataPtr++
		case '<':
			p.DataPtr--
		case '+':
			p.Memory[p.DataPtr]++
		case '-':
			p.Memory[p.DataPtr]--
		case '.':
			fmt.Printf("%c", p.Memory[p.DataPtr])
		case ',':
			_, err := fmt.Scanf("%c\n", &p.Memory[p.DataPtr])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		case '[':
			if p.Memory[p.DataPtr] == 0 {
				p.InstructionPtr = p.correspondingClosingBracket()
			}
		case ']':
			if p.Memory[p.DataPtr] != 0 {
				p.InstructionPtr = p.correspondingOpeningBracket()
			}
		}

		p.InstructionPtr++
	}
}

// isInstruction is used to filter out the comments
func (p *Program) isInstruction(b byte) bool {
	return b == '>' ||
		b == '<' ||
		b == '+' ||
		b == '-' ||
		b == '.' ||
		b == ',' ||
		b == '[' ||
		b == ']'
}

// correspondingClosingBracket assumes the instruction pointer is currently on
// an open bracket and returns the location of its corresponding closing
// bracket.
func (p *Program) correspondingClosingBracket() int {
	ptr := p.InstructionPtr + 1
	count := 0

	for {
		if p.Input[ptr] == ']' {
			if count == 0 {
				break
			} else {
				count--
			}
		} else if p.Input[ptr] == '[' {
			count++
		}
		ptr++
	}

	return ptr
}

// correspondingOpeningBracket is just like correspondingClosingBracket except
// that it returns the corresponding opening bracket.
func (p *Program) correspondingOpeningBracket() int {
	ptr := p.InstructionPtr - 1
	count := 0

	for {
		if p.Input[ptr] == '[' {
			if count == 0 {
				break
			} else {
				count--
			}
		} else if p.Input[ptr] == ']' {
			count++
		}
		ptr--
	}

	return ptr
}

func main() {
	// Command line parsing

	memSize := flag.Int("m", 3e4, "the number of bytes of memory available to the program")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "error: you must pass an input file")
		return
	} else if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: you must only pass a single input file")
		return
	}

	filename := flag.Arg(0)

	// Read the input file

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: unable to read file", filename)
		return
	}

	// Syntax checking
	if !hasValidBracketing(input) {
		fmt.Fprintln(os.Stderr, "error: program has invalid bracketing")
		return
	}

	prog := Program{Input: input, Memory: make([]byte, *memSize)}
	prog.Run()
}

// hasValidBracketing returns false if any of the brackets in the input are
// left unmatched.
func hasValidBracketing(input []byte) bool {
	count := 0
	result := true

	for _, b := range input {
		if b == '[' {
			count++
		} else if b == ']' {
			count--
		}

		if count < 0 {
			result = false
		}
	}

	if count != 0 {
		return false
	}

	return result
}
