![Monkey-face](./assets/monkey.jpg "Monkey-face")
# Monkey programming language

A powerful interpreted language written in Go



## Live demo
Try the monkey language [here](http://repl.it "here")

```sh
This is the Monkey programming language!
Feel free to type in commands
>> 
```

## Features
- Variable bindings
- Aritmetic expressions
- First-class and higher-order functions
- Closures
- Built-in functions
- Supported data structure: Integers, Strings, Booleans, Arrays, Hash Tables

## Getting started
### Variable declaration and arithmetic operations
Use `let` to declare a variable. Aritmetic operations supported are `+`, `-`, `*` and `\` 

#### Integers and Strings
```sh
>> let number = 1;
>> number + 9 * 11
100
>> let string1 = "Hi, this is the "
>> let string2 = "Monkey language!"
>> string1 + string2
Hi, this is the Monkey language!

```
#### Arrays
```sh
>> let array = [1, "Monkey", [2, 3]]
>> array[1]
Monkey
>> array[2]
[2, 3]
>> array[2][0]
2
```
#### Hash tables

```sh
>> let person = { "name": "Monkey", "age": 24}
>> person["name"]
Monkey
>> person["age"]
24
```
### Working with First-class and higher-order functions
```sh
>> let add = fn(x, y) { x + y}
>> let substract = fn(x, y) { x - y}
>> let multiply = fn(x, y) { x * y}
>> let divide = fn(x, y) { x / y}
>>
>> let calc = fn(operation, x, y) {
		return operation(x, y)
	}
>>calc(add, 6, 2)
8
>>calc(substract, 6, 2)
4
>>calc(multiply, 6, 2)
12
>>calc(divide, 6, 2)
3
```
### Making a Recursion call
#### Calculating Nth Fibonacci
```sh
>> let fibonacci = fn(x) {
		if (x == 0) {
			return 0
		} else {
			if (x == 1) {
				return 1
			} else {
				return fibonacci(x - 1) + fibonacci(x - 2);
			}
		}
	};
>>
>>fibonacci(10)
55
```
### Built-in functions
#### Currently there are 6 built-in functions supported:
**fn len(object)**
- Object: Only accept arrays and strings
- Return: If the obect type is an array, len() will return the number of elements in the array. If the object type is a string, len() will return the number of characters in the string.

**fn first(object)**
- Object: Only accept arrays
- Return: first() will return the first element of an array

**fn last(object)**
- Object: Only accept arrays
- Return: last() will return the first element of an array

**fn rest(object)**
- Object: Only accept arrays
- Return: rest() will return a copy of the the array passed in, except for the first element

**fn push(object, element)**
- Object: Only accept arrays
- element: Any type of object
- Return: push() will return a copy of the the array passed in and add the element passed in as the last element.

**fn print(object...)**
- Object: Any type of object
- Return: print() will print out all the arguments that are passed in

###### **Note:** All built-in functions does not mutate the original object

### Working with built-in functions
**Writing an array map function**
```sh
>> let map = fn(array, f) {
	let iter = fn(array, accumulated) {
		if (len(array) == 0) {
			return accumulated
		} else {
			iter(rest(array), push(accumulated, f(first(array))));
		}
	};
	iter(array, []);
};
>> let array = [1, 2, 3, 4];
>> let double = fn(x) { return x * 2 };
>> let doubledArray = map(array, double);
>> print(doubledArray)
[2, 4, 6, 8]
```

### What if I have some syntax errors?
A monkey face will pop out! xD
```sh
>> let array = [1, 2, 3
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
Woops! We ran into some monkey business here!
 parser errors:
        expected next token to be ], got EOF instead
```

## References
1.  [Writing An Interpreter In Go by Thorsten Ball](https://interpreterbook.com/)
2.  [Building Your Own Programming Language by Steve Kinney](https://frontendmasters.com/courses/programming-language/ "Building Your Own Programming Language by Steve Kinney")