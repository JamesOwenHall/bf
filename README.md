# bf

a [brainfuck](http://en.wikipedia.org/wiki/Brainfuck) interpreter

**Usage**

	./bf [-m memory-size] input-file

The default memory size is 30,000 bytes.

**Instructions**

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

## Implementation

- The entire input file is read into memory.
- The program is checked for matching brackets.
- The memory buffer is allocated.
- The data pointer is set to zero.
- The instruction pointer is set to zero.