## A Simple Start


`CodeFile-1` : [Hello World](https://github.com/alpha74/GL/tree/main)

### Contents

- [How do we Run Code in our project](https://github.com/alpha74/GL/tree/main)
- [What are Go Packages](https://github.com/alpha74/GL/tree/main)
- [Types of packages](https://github.com/alpha74/GL/tree/main)

----


### How do we Run Code in our project

<img width=40% height=40% alt="image" src="https://user-images.githubusercontent.com/31771552/140604718-a77a791b-5469-4e5a-b6b3-7139cedb3653.png">


------


### What are Go Packages?

- A package is a collection of common source files.
- A single application is like one package.


- A package can have multiple files `go` files. But all should start with line `package <package-name>`

<img width=40% height=40% alt="image" src="https://user-images.githubusercontent.com/31771552/140604890-b9c6381d-069f-4104-ada4-ce55da9ac013.png">


### Types of packages

#### Executable packages

- Generates a file that we can run, when compiled.
- The `go` file must always have a func `main()` in it. 

#### Reusable packages

- Code used as 'helpers'. Good place for resuable logic.
- Eg: Code dependencies, libraries.


#### How to determine the type?

- It the name of the package used at the start of file, that determines the type: executable or reusable.
- `package main` is used to make executable package. It creates an executable file on running `go build`.
- If any other name is used except `main`, it creates a reusable package on running `go build`.
