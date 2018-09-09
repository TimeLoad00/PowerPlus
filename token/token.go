package token

const (
    Illegal = 0
    Eof = 1
    
    String = 2
    Ident = 3
    Int = 4
    
    Add = 5
    DoubleAdd = 6
    Sub = 7
    DoubleSub = 8
    Mul = 9
    Div = 10
    
    Equal = 11
    DoubleEqual = 12
    NotEqual = 13
    AddEqual = 14
    SubEqual = 15
    MulEqual = 16
    DivEqual = 17
    Bigger = 18
    BiggerEqual = 19
    Smaller = 20
    SmallerEqual = 21
    
    Bang = 22
    Ampersand = 23
    DoubleAmpersand = 24
    Comma = 25
    Period = 26
    Colon = 27
    SemiColon = 28
    Backslash = 29
    LeftParan = 30
    RightParan = 31
    LeftBrace = 32
    DoubleLeftBrace = 33
    RightBrace = 34
    DoubleRightBrace = 35
    LeftBracket = 36
    RightBracket = 37
    
    Command = 38
)

type Token struct {
    Type int
    Val string
}
