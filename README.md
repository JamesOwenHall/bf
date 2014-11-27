# bf

a [brainfuck](http://en.wikipedia.org/wiki/Brainfuck) interpreter

#### Building

bf is written in Go.  The entire interpreter is in a single `main.go` file.  Just run `go build` and you're done.

#### Usage

	./bf [-m memory-size] input-file

The default memory size is 30,000 bytes.

#### Language

The entire language is built on 8 single-character commands.

| Symbol | Effect                                            |
|:------:|:--------------------------------------------------|
|   `>`  | increment pointer                                 |
|   `<`  | decrement pointer                                 |
|   `+`  | increment data at pointer                         |
|   `-`  | decrement data at pointer                         |
|   `.`  | output byte at pointer                            |
|   `,`  | input byte to pointer                             |
|   `[`  | if data at pointer = 0, jump over matching `]`    |
|   `]`  | if data at pointer ≠ 0, jump back to matching `[` |

If `[`s and `]`s don't match, the program is invalid.