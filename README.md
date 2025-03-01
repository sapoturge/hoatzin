Hoatzin Programming Language
============================

The [Hoatzin] programming language is a Rust-inspired language that I am using to learn about programming
language design and implementation.

Objectives
----------

Hoatzin is primarily an educational project for me to learn about how languages work. If it turns out to
be useful, that will be a bonus.

Things I want to learn include how to work with tree-sitter, LLVM, and LSP. 

Because Hoatzin will include a language server, I'm also going to have it do "perpetual compilation":
after every change that leaves the program in a compilable state, the updates will be incorporated into
the compiled binary.

Like Rust, I want as much of Hoatzin's semantics as possible to be specified in the standard library
somewhere. Intrinsic functions will be defined for addition, subtraction, etc., with something like Rust's
lang items and traits being used to tie these functions to normal operators. (Connecting them to normal
operators is a lower priority.) I do not want panics to be baked into the language as much as they are
with Rust, and should be optional if they are added at all.

Syntax will be mostly Rust-like, with some Python-like additions or substitutions. Like Rust, almost every
construct will be an expression with a value, and functions will implicitly return the value of the final
expression in their body.

Further design decisions will be made as I need to make them, and will be affected by constraints imposed
by the technologies I am working with.



[Hoatzin]: (https://en.wikipedia.org/wiki/Hoatzin)
