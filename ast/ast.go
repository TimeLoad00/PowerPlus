package ast

type Stmt struct {
    Type int
    Val []Expr
}

type Expr struct {
    Type int
    Val string
}
