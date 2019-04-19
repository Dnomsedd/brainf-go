/* By Viktor A. Rozenko Voitenko (2019)
 *
 * BRAINFUCK SYNTAX
 *
 *  +  =>  memory[dp]++
 *  -  =>  memory[dp]--
 *  >  =>  dp++
 *  <  =>  dp--
 *  .  =>  fmt.Print(memory[dp])
 *  ,  =>  reader.ReadByte()
 *  [  =>  jump to ] if memory[dp] == 0
 *  ]  =>  jump back to [ if memory[dp] != 0
 *
 */



package main


import (
    "fmt"
    "io/ioutil"
    "bufio"
    "os"
)



type BVM struct {
    code    []byte
    ip      int

    memory  [30000]int
    dp      int

    reader  *bufio.Reader

    running bool
    err     bool
}


func New(src string) *BVM {
    // read code
    data, err := ioutil.ReadFile(src)
    if err != nil {
        fmt.Println(err)
    }

    // return instance
    return &BVM {
        code    : data,
        reader  : bufio.NewReader(os.Stdin),
        running : true,
    }
}


func (vm *BVM) Exec() {
    i := vm.code[vm.ip]
    switch i {
        case '+':
            vm.memory[vm.dp]++
            vm.ip++

        case '-':
            vm.memory[vm.dp]--
            vm.ip++
        case '>':
            vm.dp++
            if vm.dp == 30000 {
                vm.err = true
                fmt.Printf("\n#%d: Data Pointer out of bounds.\n", vm.ip)
            }
            vm.ip++

        case '<':
            vm.dp--
            if vm.dp == -1 {
                vm.err = true
                fmt.Printf("\n#%d: Data Pointer out of bounds.\n", vm.ip)
            }
            vm.ip++

        case '.':
            c := vm.memory[vm.dp]
            if 0 <= c && c <= 128 {
                fmt.Printf( string(c) )
            } else {
                vm.err = true
                fmt.Printf(`\n#%d: Invalid int value for ASCII char %d
                                   at memory location %d.\n`,
                           vm.ip, c, vm.dp)
            }
            vm.ip++

        case ',':
            fmt.Print(">> ")
            b, err := vm.reader.ReadByte()
            vm.reader.Discard( vm.reader.Buffered() )
            if err != nil {
                vm.err = true
                fmt.Print("\n")
                fmt.Println(err)
                fmt.Print("\n")
            } else {
                vm.memory[vm.dp] = int(b)
            }
            vm.ip++

        case '[':
            if vm.memory[vm.dp] == 0 {
                for vm.code[vm.ip] != '!' {
                    vm.ip++
                    if vm.code[vm.ip] == ']' {
                        vm.ip++
                        return
                    }
                }
                vm.err = true
                fmt.Printf("\nSyntax Error: Did not find matching ']'.\n")
            } else {
                vm.ip++
            }

        case ']':
            if vm.memory[vm.dp] != 0 {
                for vm.ip >= 0 {
                    vm.ip--
                    if vm.code[vm.ip] == '[' {
                        vm.ip++
                        return
                    }
                }
                vm.err = true
                fmt.Printf("\nSyntax Error: Did not find matching '['.\n")
            } else {
                vm.ip++
            }

        case '!':
            vm.running = false

        default:
            vm.ip++ // don't pay attention to unknown command
    }
}



func main() {
    src := os.Args[1] // grab the source code filename
    vm := New(src)

    for vm.running && !vm.err {
        vm.Exec()
    }

    if vm.err {
        fmt.Println("\nBrainfuck exited with error.\n")
    }
}

