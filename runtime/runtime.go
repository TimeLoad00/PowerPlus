package runtime

import (
    "fmt"
   "../ast"
   "../token"
    "../parser"
)

func Run(tree *parser.Tree) {
    for _, stmt := range tree.Statements {
        if stmt.Type == token.Eof { break }
        
        if stmt.Type == token.Command {
            RunCommand(stmt)
        }
    }
}

func RunCommand(stmt ast.Stmt) {
    command := stmt.Val[0].Val
    
    for i := 1; i < len(stmt.Val); i++ {
        arg := stmt.Val[i].Val
        val := stmt.Val[i + 1].Val
        
        if arg != "_" {
            command += " -" + arg
        }
        
        command += " " + val
        
        i += 1
    }
    
    fmt.Println(command)
}
