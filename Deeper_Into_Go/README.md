## Deeper Into Go



### Project Overview

- Creating a program to play Cards.
<img width=50% heigh=50% alt="image" src="https://user-images.githubusercontent.com/31771552/140613718-2fb14b43-2251-435b-849d-a8b58b8cb8dc.png">


### Variable Declarations

`var card string = "Ace of Spades"`

- `var` tells we are about to create new variable.
- `card` : After `var` we write the name of variable.
- `string` : Only a 'string' will be assigned to this variable.
- This line can also be written as : `card := "Ace of Spades"`

- We use `:=` only when we are defining and initializing new variables. 
- Use `=` only on re-assigning. Declaring a variable by using `c = 20` is invalid.

- `Go` is a statically-typed language. Others: `C++`, `Java`
- Dynamically-typed languages: `Javascript`, `Ruby`, `Python`


#### Declaring Global Variables

  - We can declare and initialize a variable outside a `func` by using this format.
  - Following are valid:
    - `var c int = 20`
    - `var c int`

- Following are invalid:
  - `c := 20`


### Basic Data Types in Go

- These are few fundamentals data types in Go. There are many more.

<img width=50% height=50% alt="image" src="https://user-images.githubusercontent.com/31771552/140631211-94dd6f70-6076-40de-80b7-fd43bebd9e19.png">

-----

### Function Declaration
