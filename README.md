# BrainF@ck Interpreter

And all of it is written in Golang so that you can install and use it.

## What is BrainF@ck

BrainF@ck is an esoteric programming language created by Urban Muller in 1993. Its interpreter is extremely easy to implement, however its syntax is as minimalistic as it gets... 

The general idea is that you have:

1. a string of code (that was hopefully written by you);
2. an instruction pointer that is initially set to 0 (points at the first instruction);
3. 30'000 memory locations that store integers and are also initially set to 0s;
4. a data pointer that points at some memory location (also set to 0 initially and points at the first memory cell).

Now, the only 8 characters that interpreter cares about are:

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

You can find a lot of resources about this language in the internet. You now also have your own compact and fast BrainF@ck interpreter ;)

## How to Use

Run these in your terminal and it will make you just a bit happier.

```bash
~$ go get github.com/sharpvik/brainf-go
~$ go install path/to/brainf-go
```

After running the above commands you will find a binary file in the "bin" folder of your go environment, rename it to something short and memorable like "brainf". If you are a Windows user, you can add path/to/the/bin into the PATH environmental variable, then restart your PC. If you are on UNIX/Linux, I think you can copy the binary file into the /usr/bin folder. You need to do this to be able to use this BrainF@ck interpreter like you normally use Python for example.

```bash
brainf helloworld.bf 
# above command runs the brainf script stored in the helloworld.bf file and 
# produces the following output
hello world
```

The cool thing about this interpreter is that it doesn't care about symbols that are not used by the BrainF@ck language, so you can write all sorts of comments for yourself while coding. For example, *helloworld.bf* file looks like this:

```brainf
+++++ +++++ [
    > +++++ +++++
    < -
]

id: 0 1   2 3 4 5
vl: 0 100 0 0 0 0
pt: ^

> ++++ .                        // 'h'
--- .                           // 'e'
+++++ ++ ..                     // 'll'
+++ .                           // 'o'

id: 0 1   2 3 4 5
vl: 0 111 0 0 0 0
pt:   ^

< +++++ +++++  
  +++++ +++++  
  +++++ +++++ ++ . >            // ' '

id: 0  1   2 3 4 5
vl: 32 111 0 0 0 0
pt:    ^

+++++ +++ .                     // 'w'
----- --- .                     // 'o'
+++ .                           // 'r'
----- - .                       // 'l'
----- --- .                     // 'd'

< ----- -----
  ----- ----- -- . >

id: 0  1   2 3 4 5
vl: 10 100 0 0 0 0
pt:    ^
```

It runs fine, since slashes, letters, ans spacing are ignored. However, you can also use the *minbf.py* script to minify that file and produce *helloworld.min.bf* as follows:

```bash
python3 minbf.py path/to/helloworld.bf /path/to/helloworld.min.bf
```

This will write the following **minified** brainf code into the file specified by you (in my example it was *helloworld.min.bf*).

```brainf
++++++++++[>++++++++++<-]>++++.---.+++++++..+++.<++++++++++++++++++++++++++++++++.>++++++++.--------.+++.------.--------.<----------------------.>
```

## Contact

For any personal or business enquiries:

+ Email: *sharp.vik@gmail.com*
+ [Twitter](https://twitter.com/sharp_vik)
+ [VK](https://vk.com/perigrinus)
+ [Instagram](https://www.instagram.com/viktooooor)
