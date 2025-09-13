# OOP-context

This repository contains example code and notes demonstrating core Object-Oriented Programming concepts in both Java and Go.

New in this repository
- Added a learning track for SOLID principles.
- Added an introductory Design Patterns section (with OOP-centered examples).
- Added Case Studies that apply OOP, SOLID and Design Patterns to small real-world scenarios.

Contents
- Object-Oriented-Programming/ — examples, context and FAQ for OOP
  - [example.java](Object-Oriented-Programming/example.java) — Java demonstration (public class [`OOPExample`](Object-Oriented-Programming/example.java))  
  - [example.go](Object-Oriented-Programming/example.go) — Go demonstration (package [`main`](Object-Oriented-Programming/example.go))  
  - [context.md](Object-Oriented-Programming/context.md) — explanatory notes and examples
  - [FAQ.md](Object-Oriented-Programming/FAQ.md) — short Q&A
  - Object-Oriented-Programming/SOLID/ — (planned) SOLID principle guides & examples
  - Object-Oriented-Programming/Design-Patterns/ — (planned) design pattern intro and OOP examples
  - Object-Oriented-Programming/Case-Studies/ — (planned) applied examples combining OOP, SOLID, and patterns

Quick run instructions (Linux)
- Java:
  1. Install JDK if needed: sudo apt update && sudo apt install openjdk-24-jdk -y
  2. Build & run:
     cd Object-Oriented-Programming
     javac example.java
     java OOPExample
- Go:
  1. Install Go if needed: sudo apt update && sudo apt install golang-go -y
  2. Run:
     cd Object-Oriented-Programming
     go run example.go
     (or: go build -o example example.go && ./example)

How this repo is organized
- OOP examples (Java & Go) show core concepts so you can compare language idioms: see [example.java](Object-Oriented-Programming/example.java) and [example.go](Object-Oriented-Programming/example.go).
- The contextual notes are in [context.md](Object-Oriented-Programming/context.md) and quick Q&A in [FAQ.md](Object-Oriented-Programming/FAQ.md).
- New folders (SOLID, Design-Patterns, Case-Studies) will contain short theory, minimal examples, and one applied case study each.

Contributing
- Add SOLID examples under Object-Oriented-Programming/SOLID/, one principle per file with short explanation and code.
- Add Design Patterns under Object-Oriented-Programming/Design-Patterns/, start with Factory, Strategy, Singleton, Adapter.
- Add Case Studies under Object-Oriented-Programming/Case-Studies/ showing how to refactor a small design using SOLID + a pattern.

See also
- [Object-Oriented-Programming/context.md](Object-Oriented-Programming/context.md)
- [Object-Oriented-Programming/FAQ.md](Object-Oriented-Programming/FAQ.md)
- [Object-Oriented-Programming/example.java](Object-Oriented-Programming/example.java)
- [Object-Oriented-Programming/example.go](Object-Oriented-Programming/example.go)
