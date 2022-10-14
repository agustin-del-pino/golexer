# golexer
Implementation for Golang of Lexer Decorator Pattern.

# What's Lexer Decorator Pattern?

This kind of pattern is based on **Decorator Pattern** oriented to a Lexer construction.

As the pattern establishes, for the Lexer construction you've to decorate **Lexer Functions**.
Those decorators can specify the *use case* (called *lex case*) of the *Lexer Function* (for numbers, strings, etc), 
and the function itself represent the behavior of the *lex case*.

When a *plain text* is processed by the Lexer, a **Cursor** is created. Think this **Cursor** as the cursor of the terminal, which has a current position, where that position is equivalent to a specific char.

Using the **Cursor** for read the *plain text*, the current char is evaluated, this evaluation is the **lex case**.

Depending on the *lex case*, it will be called the **Lexer Function** decorated for the particular *lex case*. 

The **Lexer Functions** will return a **Token** that represents the grabbed chars, given them a **Type and Value**.

## Lex Decorator
By default it must be five **Lex Decorators** for the following *lex cases*:
- Numbers
- Strings
- Identifier
- Comments
- Delimiters

### Number Lex Decorator
Its *lex case* is the number literals. For example:
```basic
1234
12.222
1212 444 444
```
It's not mandatory but assuming the digits from **0 to 9** are the expected numbers for this *lex case*, the common evaluation for this decorator is the set `[0, 9] = { 0 <= x <= 9 }`.

### String Lex Decorator
Its *lex case* is the string literals, assuming the literal strings as sequence of char surrounding by some particular char. For example:
```basic
"string"
'str ing'
`s t r i n g`
```
Following this thinking, the strings will be evaluated when the surrounding char is detected. It's responsibility of the implementation to allow more than one surrounding char.

### Identifier Lex Decorator
Its *lex case* is the identifiers, those are (for exhaustion) every sequence of char that belongs to the **Alpha-Numeric-Special** set. For example:
```basic
my_identifier
_my_identifier
MyIdentifier
Id3n71f13r
```
Because of this *lex case*, the *keywords* are inside of the identifier, in other words, the *keywords* are a subset of identifiers. For example, imagine a subset for write a letter: `{ addressee; issuer; title; header; body; footer }`.
```basic
addressee "Miss Abc"
issuer "Student Xyz"

title "Letter for Miss Abc"

header "Hi there Miss Abc,"
body "I'm writing this letter for communicate..."
footer "Regards"
```

### Comments Lex Decorator
Its *lex case* is for piece of text that can be ignore.
Like the *strings*, the comments must be start with an specific char for starts the evaluation. For example:
```
# This is a comment
#- 
This is a multi-line 
comment 
-#
```
*This concept comes from the comments in the programming languages*

Of course, for evaluate comments can be used other strategies, like *keywords* or whatever.

### Delimiter Lex Decorator
Its *lex use* is for evaluate a specific char that serve a delimiter of between Tokens. For example:

The delimiter is `+`.
```basic
123 + 456
```

The delimiter is `,`.
```basic
"A", "B", "C"
``` 

The delimiter are `( ) : ^`
```basic
f(x): x^2
``` 

## Cursor
As was mentioned before, the **Cursor** is the responsible for reading the *plain text*. 
Like a typewriter, every time a char is read the **Cursor** moves to the next char until there are no more chars for read.

It can be describe the **Cursor** as the following interface:

```cs
interface Cursor {
  void advance(); // Moves the cursor.
  char getChar(); // Retrieves the current char.
}
```

This interface can be extended for publish more information. For example:

```cs
interface Cursor {
  void advance();    // Moves the cursor.
  char getChar();    // Retrieves the current char.
  int getPosition(); // Retrieves the current position.
  int getLength();   // Retrieves the length of the reading text.
  bool isEOF();      // Verifies whether the end of reading text was reached.
}
```

## Token
The token is a subset of the plain text that has two attributes: **Type and Value**.

It can be describe as the following interface:
```cs
interface Token<T> {
  T getType();
  char[] getValue();
}
```

*Look out for the Generic!. The type object of Token Type can be whatever the implementation wants; an Enum, Number, String are just a few examples.*


### Token Type
Represents the kind of token's subset. This's aligning with the Lexer Decorator. 

*For each lexer decorator there will be at least one token type.*

### Token Value
Contains the subset of chars itself. For example:
Given `"I'm string"` as *plain text*, the **Token Value** will be: `I'm string`.


## Lexer Function
The interface of this function is the following:
```cs
Token<T> LexerFunction<T>(Cursor);
```
In human words, the function takes one parameter, the current Cursor, and returns the Token. The *generic type* stands for the solicited by the **Token Interface**.


# Golang Lexer Decorator Pattern Implementation
As everyone knows, Golang doesn't have any pretty way to implement the Decorator Pattern (in other case Python has the *Decorator & Decorated Functions*). So,this implementation has to be done by the old way.

`decorator_function(decorated_function)`

## Quick Start
*For deep information of how is the implementation, jump to the next topic.*

### Let's create a Number Lexer
First, create new instance of the Lexer Implementation.

```go
func main(){
  l, _ := lexer.NewLexer()
}
```

Add the range of the number to the lexer settings.
It can be done by using the ByteRange implementation `SingleByteRange`.

```go
func main(){
  l, _ := lexer.NewLexer(lexer.LexerSettings{
		Numbers: lexer.NewSingleByteRange(0x30, 0x39),
	})
}
```

Use the `LexNumber` method of the instanced lexer for add a **Lex Function**.

```go
func main(){
  l, _ := lexer.NewLexer(lexer.LexerSettings{
		Numbers: lexer.NewSingleByteRange(0x30, 0x39),
	})

  l.LexNumber(lexNumbs)
}
```

Just for finish up the main function, call the `Tokenize` method and pass a number.

```go
func main(){
  l, _ := lexer.NewLexer(lexer.LexerSettings{
		Numbers: lexer.NewSingleByteRange(0x30, 0x39),
	})

  l.LexNumber(lexNumbs)
  t, _ := l.Tokenize([]byte("0123456789"))  
	fmt.Printf("tokens: %v\n", t)
}
```

Now the fun part, let's build the **Lex Function**.

```go
func lexNumbs(c *lexer.Cursor) lexer.Token[lexer.TokenType] {
}
```

The digits need to be storage in a buffer. Using a byte array may be enough. Also, add the current char to the buffer and don't forget to move the cursor.

```go
func lexNumbs(c *lexer.Cursor) lexer.Token[lexer.TokenType] {
  var b []byte
  
  b = append(b, (*c).GetChar())
  
  (*c).Advance()
}
```

For now return a new token. Of course, the type must be **Number** and the value will be the buffer.


```go
func lexNumbs(c *lexer.Cursor) lexer.Token[lexer.TokenType] {
  var b []byte
  
  b = append(b, (*c).GetChar())
  
  (*c).Advance()
  return lexer.NewToken(lexer.Number, b)
}
```

With all of this done, just run the program.

```shell
go run main.go
```

As you can notice, everything works fine but the lexer is detecting each number as separated tokens without any relation to the previous or next numbers. 

### Let's improve the Number Lexer

A little modification must be done before start. Pass the `ByteRange` to a global variable. Also, add some error handling.

```go
var NUMBS = lexer.NewSingleByteRange(0x30, 0x39)

func lexNumbs(c *lexer.Cursor) lexer.Token[lexer.TokenType] {
	var b []byte

	b = append(b, (*c).GetChar())

	(*c).Advance()

	return lexer.NewToken(lexer.Number, b)
}

func main() {
	l, lerr := lexer.NewLexer(lexer.LexerSettings{
		Numbers: NUMBS,
	})

	if lerr != nil {
		fmt.Printf("lerr: %v\n", lerr)
		return
	}

	l.LexNumber(lexNumbs)

	if t, terr := l.Tokenize([]byte("0123456789")); terr != nil {
		fmt.Printf("terr: %v\n", terr)
	} else {
		fmt.Printf("t: %v\n", t)
	}
}
```

So far, so good. 

From now, let's focus into grab more digits for creates a one single token. For do that it such simple as moving along the cursor and buff the current char until a condition. 
What condition?. Well, the current char must be a digit, in other words, must be in range of `NUMBS`, otherwise the current char is not a digit, in this case the loop breaks and the token can be created with the buffered chars.

```go
func lexNumbs(c *lexer.Cursor) lexer.Token[lexer.TokenType] {
	var b []byte

	b = append(b, (*c).GetChar())

	(*c).Advance()

  for NUMBS.IsInRange((*c).GetChar()) {
    b = append(b, (*c).GetChar())
		(*c).Advance()
  }

	return lexer.NewToken(lexer.Number, b)
}
```

With this, only one thing left to do: *read multiple number tokens*.

This last step is very simple, just add the white spaces as ignorable char, so when space appears, the lexer will ignore it and moves to the next char.

```go
var NUMBS = lexer.NewSingleByteRange(0x30, 0x39)
var IGNORE = lexer.NewBPoints(0x20)

func lexNumbs(c *lexer.Cursor) lexer.Token[lexer.TokenType] {
	var b []byte

	b = append(b, (*c).GetChar())

	(*c).Advance()

	for NUMBS.IsInRange((*c).GetChar()) {
		b = append(b, (*c).GetChar())
		(*c).Advance()
	}

	return lexer.NewToken(lexer.Number, b)
}

func main() {
	l, lerr := lexer.NewLexer(lexer.LexerSettings{
		Numbers: NUMBS,
		Ignore:  IGNORE,
	})

	if lerr != nil {
		fmt.Printf("lerr: %v\n", lerr)
		return
	}

	l.LexNumber(lexNumbs)

	if t, terr := l.Tokenize([]byte("001100 987654321 123 456 789")); terr != nil {
		fmt.Printf("terr: %v\n", terr)
	} else {
		fmt.Printf("t: %v\n", t)
	}
}
```

Congratulation, the Number Lexer is done.

## Deep Understanding

The implementation is quite simple, but is useful knows how it's structured for make better use of.

### The Lexer Implementation
All **Lex Decorators** are prepared for take more than one **Lex Function**, but it's a feature for the future, right now doesn't have any useful use. 

When the tokenize is executed the order of is:

1. The Ignore Bytes.
2. The Lex Number.
3. The Lex String.
4. The Lex Identifier.
5. The Lex Comment.
6. The Lex Delimiter.

In case the current char doesn't match with any of the *lex uses* described, an error will throw.

### The Cursor Implementation
The start values are:
```go
{
	char:     EOF,
	position: 0,
	content:  b,
	length:   len(b),
}
```
The `EOF` constant is the `0x00` (`NULL`) byte.

### The Token Implementation
The token types are:

- Number
- String
- Comment
- Identifier
- Keyword
- Delimiter
