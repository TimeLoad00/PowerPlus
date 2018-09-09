package main

import(
    // "fmt"
    "io/ioutil"
    "os"
    "./lexer"
   // "./token"
    "./parser"
    "./runtime"
)

func main() {
    source, err := ioutil.ReadFile(os.Args[1])
    check(err)
    
    list := lexer.New(string(source))
    
    parser := parser.NewTree(list)
   // fmt.Println(parser.Statements[0].Type)
   // fmt.Println(parser.Statements[0].Val[0].Val)
   // fmt.Println(parser.Statements[0].Val[1].Val)
   // fmt.Println(parser.Statements[0].Val[2].Val)
    
   // fmt.Println(parser.Statements[1].Val[0].Val)
   // fmt.Println(parser.Statements[1].Val[1].Val)
   // fmt.Println(parser.Statements[1].Val[2].Val)
    
    runtime.Run(parser)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
