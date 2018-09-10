package runtime

import (
    "fmt"
   "../token"
    "../lexer"
)

var tokens *lexer.Tokens

func Run(tok *lexer.Tokens) {
    tokens = tok

    for {
        statement := runStatement()
        
        if statement == "" { break }
        
        fmt.Println(statement)
    }
}

func runStatement() string {
    tok := tokens.ReadToken()
    statement := ""
    
    if tok.Type == token.Eof { return "" }
    
    if tok.Type == token.Ident {
        ident := tok.Val
        tok = tokens.ReadToken()
        
        if tok.Type == token.LeftBrace {
            statement = runCommand(ident)
        } else {
            statement = ident
        }
    } else if tok.Type == token.String {
        return tok.Val
    }
    
    return statement
}

func runCommand(command string) string {
    for {
        arg := tokens.ReadToken()
        
        if arg.Type != token.Ident { panic("expected ident") }
        
        if arg.Val == "_" {
            command += " "
        } else {
            command += " -" + arg.Val + " "
        }
        
        if tokens.PeekToken().Type != token.Colon { panic("expected colon") }
        tokens.Pos++
        
        command += runStatement()
        
        if tokens.PeekToken().Type == token.Comma {
            tokens.Pos++
        } else if tokens.PeekToken().Type == token.RightBrace {
            tokens.Pos++
            break
        } else {
            panic ("expected comma")
        }
    }
    
    return command
}
