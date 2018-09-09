package parser

import (
    "../lexer"
    "../token"
    "../ast"
)

type Tree struct {
    Statements []ast.Stmt
    Pos int
}

func NewTree(lex *lexer.Tokens) *Tree {
    tree := &Tree { Statements: nil, Pos: 0 }
    
    for {
        tok := lex.ReadToken()
        stmt := ast.Stmt { Type: 0, Val: nil }
        
        if tok.Type == token.Eof {
            stmt = ast.Stmt { Type: token.Eof, Val: nil }
            tree.Statements = append(tree.Statements, stmt)
            break
        }
        
        if tok.Type == token.Ident {
            ident := tok.Val
            tok = lex.ReadToken()
            
            if tok.Type == token.LeftBrace {
                stmt = parseCommand(lex, ident, stmt)
            }
        }
        
        tree.Statements = append(tree.Statements, stmt)
    }
    
    return tree
}

func parseCommand(tokens *lexer.Tokens, command string, stmt ast.Stmt) ast.Stmt {
    stmt = ast.Stmt { Type: token.Command, Val: nil }
    exprCommand := ast.Expr { Type: token.Ident, Val: command }
    stmt.Val = append(stmt.Val, exprCommand)
    
    for {
        if tokens.PeekToken().Type == token.RightBrace {
            tokens.Pos++
            break
        }
        
        arg := tokens.ReadToken().Val
        
        if tokens.PeekToken().Type != token.Colon { panic("expected colon") }
        tokens.Pos++
        
        val := tokens.ReadToken()
        
        exprArg := ast.Expr { Type: token.Ident, Val: arg }
        exprVal := ast.Expr { Type: val.Type, Val: val.Val }
        stmt.Val = append(stmt.Val, exprArg)
        stmt.Val = append(stmt.Val, exprVal)
    }
    
    return stmt
}
