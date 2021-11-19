## Deeper Into Go



### Project Overview

- [CardsApp code repo](https://github.com/alpha74/CardsApp/tree/main)
- Creating a program to play Cards.


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


`CodeFile-1` [Function Calls and Returns](https://github.com/alpha74/GL/blob/main/Deeper_Into_Go/func_calls.go)

`func newCard()`
- `Go` expects that this function will return `void`.
- Hence, it will give error will we try to `return` or assign the returned value to some var.


`func newCard() string`
- This tells `Go` that func `newCard()` will return data of `string` type.


-----

### Arrays in Go

`Go` has two data types for handling list of records:

`CodeFile-2` [Using Arrays and Slices](https://github.com/alpha74/GL/blob/main/Deeper_Into_Go/using_arrays_slices.go)

#### 1. Array

- Fixed length list of strings.
- Each array needs to declared with a data type.

#### 2. Slice

- An Array that can grow or shrink.
- Each slice needs to declared with a data type.


------

### Custom Type Declarations

`type Deck []string`

- `Deck` now extends functionality if `[]string`
- `type` is analogous to `typedef` in `C`.

- `cards := Deck{"Ace of Diamonds"}` means same as `cards := []string{"Ace of Diamonds"}`

#### Receiver on function

- We can define a receiver on function to work with `Deck`
- Defined in `deck.go`

```
// A receiver on func for printing list of cards in 'Deck'
func (D Deck) print() {
	for i, card := range D {
		fmt.Println(i, card)
	}
}
```

- This function is analogous to member function of a class in `C++`
- Any member of type `Deck` has access to `printDeck()`
- Usage is done in `main.go`

```
// Using receiver func instead of directly printing
cards.printDeck()
```

### Type Conversions

- Convert a string to `byte slice`

```
msg := "Hello"
fmt.Println([]byte(msg))
```

- `[]byte` is a representation in terms if `ASCII` numbers.



### File I/O

- Use `ioutil.WriteFile` for writing to a file.
- It returns an object of `error` type

```
// Func to write Deck to file
func (D Deck) saveToFile(filename string) error {
	// 0666 : Anyone can read/write
	return ioutil.WriteFile(filename, []byte(D.toString()), 0666)
}
```
