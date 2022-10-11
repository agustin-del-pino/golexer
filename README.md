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

