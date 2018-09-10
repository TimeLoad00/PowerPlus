package main

import(
    "io/ioutil"
    "os"
    "./lexer"
    "./runtime"
)

func main() {
    source, err := ioutil.ReadFile(os.Args[1])
    check(err)
    
    list := lexer.New(string(source))
    
    runtime.Run(list)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
