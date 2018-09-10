This is a simple interpreter for a toy scripting language I'm creating called PowerPlus. It takes PowerPlus as input and outputs PowerShell, essentially acting as a layer because the user and PowerShell giving PowerShell (what I believe to be) a nicer syntax.

The inspiration for creating comes out of my pure hatred for PowerShell and I hope that in creating this interpreter, I either learn how to use PowerShell better or I create a tool where I no longer need to use PowerShell.

I decided to write this interpreter in Go because despite how many years I spent programming in C#, I find it incredibly bloated and now that I use Linux as my main programming environment, it's too much of a hassle to use. I've found Go is a great mix of dynamic programming and a compiled, yet portable, language. This is my first project using Go and coming from a C#/JS background, it has definitely been a great learning experience.

Executing a PowerPlus Script:
Run the program and pass your script as the first argument

Syntax:
At the moment this language is very very basic, its current state is just the barebones lexer and runtime. But it serves as a great base to expand the language. The only current syntax is interpreting a PowerShell command. The syntax is as follows:

Command { Argument1: Value, Argument2: Value }

Which would be interpreted by the interpreter to:

Command -Argument1 Value -Argument2 Value

The same command can be expressed like the following:

Command {
    Argument1: Value,
    Argument2: Value
}

Since PowerShell also likes to use statements as expressions, all statements can be used as expressions through recursion:

Command { Argument1: Value, Argument2: Command1 { Argument3: Value } }
