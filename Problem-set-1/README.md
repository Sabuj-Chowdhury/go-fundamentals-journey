# Problem Set 1: Grade Calculator with Menu

A Go program that implements a menu-driven grade calculator with pass/fail checking.

## Problem Statement

Create a program that:
- Shows a menu with options: 1) Calculate grade, 2) Check pass/fail status, 3) Exit
- Use a switch statement for menu selection
- For option 1: Take a score (0-100) and use if-else to assign letter grades (A: 90-100, B: 80-89, C: 70-79, D: 60-69, F: below 60)
- For option 2: Use a switch to check if a score is pass (≥60) or fail
- Use a for loop to keep showing the menu until user chooses exit
- Include at least one anonymous function to display the menu

## What You'll Learn

### 1. Menu-Driven Programs
A loop that repeatedly shows options until the user chooses to exit.

```go
running := true
for running {
    displayMenu()
    fmt.Scan(&choice)
    // handle choice
}
```

### 2. Switch Statement
Go's switch statement for handling multiple cases cleanly.

```go
switch choice {
case 1:
    // calculate grade
case 2:
    // check pass/fail
case 3:
    // exit
default:
    // invalid option
}
```

### 3. Anonymous Function
Functions without a name, stored in a variable and called later.

```go
displayMenu := func() {
    fmt.Println("1) Calculate grade")
    fmt.Println("2) Check pass/fail status")
    fmt.Println("3) Exit")
}
```

### 4. If-Else Chain
Using if-else to assign letter grades based on score ranges.

```go
if score >= 90 && score <= 100 {
    return "A"
} else if score >= 80 && score <= 89 {
    return "B"
}
```

### 5. Functions with Return Values
Creating reusable functions that return a value (string in this case).

```go
func calculateGrade(score int) string {
    // logic here
    return "Grade"
}
```

## Example Output

```
Welcome to grade calculator
1) Calculate grade
2) Check pass/fail status
3) Exit the program
Choose an option: 1
Enter a score (0-100):85
You got B

Welcome to grade calculator
1) Calculate grade
2) Check pass/fail status
3) Exit the program
Choose an option: 2
Enter a score (0-100):55
You got Fail

Welcome to grade calculator
1) Calculate grade
2) Check pass/fail status
3) Exit the program
Choose an option: 3
exiting program ...
```

## How to Run

1. Make sure you have Go installed ([Download Go](https://go.dev/dl/))
2. Clone this repository
3. Run the program:

```bash
go run main.go
```

## Code Structure

```
main.go
├── displayMenu()    - Anonymous function to show menu
├── calculateGrade() - Returns letter grade based on score
└── checkPassFail()  - Returns Pass/Fail status
```

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Go by Example - Switch](https://gobyexample.com/switch)
- [Go by Example - Functions](https://gobyexample.com/functions)

## Author

Sabuj - Learning Go!
