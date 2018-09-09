package lexer

import (
    "../token"
)

type Lexer struct {
    source string
    pos int
}

type Tokens struct {
    list []token.Token
    Pos int
}

func New(src string) *Tokens {
    lex := &Lexer { source: src, pos: 0 }
    tokens := &Tokens { list: nil, Pos: 0 }
    
    for {
        tok := lex.readToken()
        
        if tok.Type == token.Eof {
            tokens.list = append(tokens.list, tok)
            break
        } else {
            tokens.list = append(tokens.list, tok)
        }
    }
    
    return tokens
}

func (tokens *Tokens) ReadToken() token.Token {
    tok := tokens.list[tokens.Pos]
    tokens.Pos++
    return tok
}

func (tokens *Tokens) PeekToken() token.Token {
    return tokens.list[tokens.Pos]
}

func (lex *Lexer) readToken() token.Token {
    var tok token.Token
    var val string
    
    for {
        if lex.pos == len(lex.source) {
            tok = token.Token { Type: token.Eof, Val: val }
            break
        }
        
        ch := lex.peekCharacter()
        
        if ch == '\n' || ch == ' ' {
            lex.pos++
        } else if ch == '"' {
            tok = lex.readString()
            break
        } else if isLetter(ch) {
            tok = lex.readIdent()
            break
        } else if isDigit(ch) {
            tok = lex.readDigit()
            break
        } else {
            tok = lex.readSymbol()
            break
        }
    }
    
    return tok
}

func (lex *Lexer) peekCharacter() byte {
    return lex.source[lex.pos]
}

func (lex *Lexer) readCharacter() byte {
    ch := lex.source[lex.pos]
    lex.pos++
    return ch
}

func (lex *Lexer) readString() token.Token {
    tok := token.Token { Type: token.String, Val: "" }
    lex.pos++
    
    for {
        if lex.pos == len(lex.source) { break }
        
        if lex.peekCharacter() == '\\' {
            lex.pos++
            tok.Val += string(lex.readCharacter())
        } else if lex.peekCharacter() != '"' {
            tok.Val += string(lex.readCharacter())
        } else {
            lex.pos++
            break
        }
    }
    
    return tok
}

func (lex *Lexer) readIdent() token.Token {
    tok := token.Token { Type: token.Ident, Val: string(lex.readCharacter()) }
    
    for {
        if lex.pos == len(lex.source) { break }
        
        if isLetter(lex.peekCharacter()) || isDigit(lex.peekCharacter()) {
            tok.Val += string(lex.readCharacter())
        } else {
            break
        }
    }
     
    return tok
}

func (lex *Lexer) readDigit() token.Token {
    tok := token.Token { Type: token.Ident, Val: string(lex.readCharacter()) }

    for {
        if lex.pos == len(lex.source) { break }
        
        if isDigit(lex.peekCharacter()) {
            tok.Val += string(lex.readCharacter())
        } else {
            break
        }
    }

    return tok
}

func (lex *Lexer) readSymbol() token.Token {
    tok := token.Token { Type: token.Illegal, Val: "" }
    symbol := lex.readCharacter()
    symbolPeek := lex.peekCharacter()
    
    if symbol == '+' {
        if symbolPeek == '+' {
            tok = token.Token { Type: token.DoubleAdd, Val: "" }
        } else if symbolPeek == '=' {
            tok = token.Token { Type: token.AddEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Add, Val: "" }
        }
    } else if symbol == '-' {
        if symbolPeek == '-' {
            tok = token.Token { Type: token.DoubleSub, Val: "" }
        } else if symbolPeek == '=' {
            tok = token.Token { Type: token.SubEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Sub, Val: "" }
        }
    } else if symbol == '*' {
        if symbolPeek == '=' {
            tok = token.Token { Type: token.MulEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Mul, Val: "" }
        }
    } else if symbol == '/' {
        if symbolPeek == '=' {
            tok = token.Token { Type: token.DivEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Div, Val: "" }
        }
    } else if symbol == '=' {
        if symbolPeek == '=' {
            tok = token.Token { Type: token.DoubleEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Equal, Val: "" }
        }
    } else if symbol == '>' {
        if symbolPeek == '=' {
            tok = token.Token { Type: token.BiggerEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Bigger, Val: "" }
        }
    } else if symbol == '<' {
        if symbolPeek == '=' {
            tok = token.Token { Type: token.SmallerEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Smaller, Val: "" }
        }
    } else if symbol == '!' {
        if symbolPeek == '=' {
            tok = token.Token { Type: token.NotEqual, Val: "" }
        } else {
            tok = token.Token { Type: token.Bang, Val: "" }
        }
    } else if symbol == '&' {
        if symbolPeek == '&' {
            tok = token.Token { Type: token.DoubleAmpersand, Val: "" }
        } else {
            tok = token.Token { Type: token.Ampersand, Val: "" }
        }
    } else if symbol == ',' {
        tok = token.Token { Type: token.Comma, Val: "" }
    } else if symbol == '.' {
        tok = token.Token { Type: token.Period, Val: "" }
    } else if symbol == ':' {
        tok = token.Token { Type: token.Colon, Val: "" }
    } else if symbol == ';' {
        tok = token.Token { Type: token.SemiColon, Val: "" }
    } else if symbol == '(' {
        tok = token.Token { Type: token.LeftParan, Val: "" }
    } else if symbol == ')' {
        tok = token.Token { Type: token.RightParan, Val: "" }
    } else if symbol == '{' {
        if symbolPeek == '{' {
            tok = token.Token { Type: token.DoubleLeftBrace, Val: "" }
        } else {
            tok = token.Token { Type: token.LeftBrace, Val: "" }
        }
    } else if symbol == '}' {
        if symbolPeek == '}' {
            tok = token.Token { Type: token.DoubleRightBrace, Val: "" }
        } else {
            tok = token.Token { Type: token.RightBrace, Val: "" }
        }
    } else if symbol == '[' {
        tok = token.Token { Type: token.LeftBracket, Val: "" }
    } else if symbol == ']' {
        tok = token.Token { Type: token.RightBracket, Val: "" }
    }

    return tok
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '-'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func isIdentBreak(ch byte) bool {
    b := []byte { ' ', '\n' }
    
    for _, c := range b {
        if c == ch {
            return true
        }
    }
    
    return false
}
