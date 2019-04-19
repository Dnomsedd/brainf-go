# BrainF@ck Interpreter

And all of it is written in Golang so that you can install and use it.

## What is BrainF@ck

BrainF@ck is an esoteric programming language created by Urban Muller in 1993. Its interpreter is extremely easy to implement, however its syntax is as minimalistic as it gets... 

The general idea is that you have:

1. a string of code (that was hopefully written by you);
2. an instruction pointer that is initially set to 0 (points at the first instruction);
3. 30'000 memory locations that store integers and are also initially set to 0s;
4. a data pointer that points at some memory location (also set to 0 initially and points at the first memory cell).

Now, the only 9 characters that interpreter cares about are:

```
+ - > < . , [ ]
```

| Operator | Effect |
|:--------:|:-------|
| + -      | Increment / decrement the value stored in the memory cell pointed to by the data pointer |
| > <      | Shift data pointer right / left |
| ,        | Read one char from standard input and store it as an integer in the current memory cell |
| .        | Print ASCII character that corresponds to the integer value stored in the current memory cell |
| [ ]      | Loop while value of current memory cell is non-zero when ] is hit |
| !        | HALT command stops execution (it was not in the original BrainF@ck but I've added it to simplify things |

You can find a lot of resources about this language in the internet. You now also have your own compact and fast BrainF@ck interpreter ;)

## How to Use

Run these in your terminal and it will make you just a bit happier.

```
~$ go get github.com/sharpvik/brainf-go
~$ go install path/to/brainf-go
~$ path/to/brainf-go(.exe) path/to/file.bf
```

## Contact

For any personal or business enquiries:

+ Email: *sharp.vik@gmail.com*
+ [Twitter](https://twitter.com/sharp_vik)
+ [VK](https://vk.com/perigrinus)
+ [Instagram](https://www.instagram.com/viktooooor)
