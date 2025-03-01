# golang101
 
### Initialize Project
Inside the project's directory run <code>go mod init {projectName}</code>

### Hello World
```go
package main
import "fmt"

func main(){
    fmt.Println("Hello world")
}
```
Go programs are organized into packages, Go's standar library provides different core packages for us to use such as "fmt". Run code by <code>go run {fileName}.go</code>. If you have more than one packages (yes must multiple line): 
```go
import (
    "fmt"
    "strings"
)
```

### Variables
If u declare but not use variable, red squiggly line will appear.
<code>var {variableName} = {value}</code>
<code>const {variableName} = {value}</code>
```go
var country = "Indonesia"
fmt.Println("I love ",country)
//using placeholders
fmt.Print("I love %v, it is so great\n",country)
```

Go is a statically typed language. When declaring a variable, data type is a must yet it can infer the type when we initialize with a value.
```go
//not valid
var username
username = "prabowo"

//valid
var username string
username = "prabowo"
```
You can also declare variables like this, but not for constant.
```go
//valid
username := "prabowo"

// invalid
const year int := 2024
```

### Package Level Variables (it's just global variables)
Declared outside main can be used globally, cannot be declared with <code>:=</code>
```go
var username = "fufufafa"
const phi float = 3.14
var islands = []string{}
func main(){
    fmt.Println("%v %v %v",username,phi,islands)
}
```

### Input (yes pointers are back)
```go
var year int
fmt.Scan(&year)
```

### Mathematical Operations
just like how u do it in any other language

### Arrays
```go
// declare array of string size 10, 2 values initialized
var candidates = [10]string{"Anies","Prabowo"}
candidates[2] = "Ganjar"

// declare with no values
var viceCandidates [10]string
```

```go
fmt.Printf("Full array: %v\n",candidates)
fmt.Printf("First: %v\n",candidates[0])
fmt.Printf("Array type: %T\n",candidates)
fmt.Printf("Array length: %v\n",len(candidates))
```

### Slices
Dynamically sized array, no need to declare
```go
var gryffindor []string
gryffindor = append(gryffindor, "Harry Potter")

slytherin := []string{"Draco Malfoy"}
slytherin = append(slytherin, "Lucius Malfoy")
```

```go
fmt.Printf("Full slice: %v\n",gyrffindor)
fmt.Printf("First: %v\n",gyrffindor[0])
fmt.Printf("Slice type: %T\n",gyrffindor)
fmt.Printf("Slice length: %v\n",len(gyrffindor))
```

### If Statements
```go
if condition {

}
else if condition {

}
else {

}
```
Can be used with <code>&&</code> <code>||</code> <code>!=</code> <code>==</code>

### I hate switch case
```go
island := "Java"
switch island{
    case "Sumatra":
        // code to exec if island is Sumatra
    case "Java":
        // code to exec if island is Java
    case "Kalimantan", "Sulawesi", "Papua":
        // code to exec if island is Kalimantan or Sulawesi or Papua
    default:
        fmt.Println("invalid island selected")
}
```

### Loop
```go
for {
    // this is an infinite loop
}
```
```go
for index, valueVar := range arrayVar{
    fmt.Printf("arrayVar[%v]: %v\n",index, valueVar)
}
```
```go
var gryffindors = []string{"Harry Potter","Ron Weasley","Hermione Granger"}
firstNames := []string{}
for _, gryffindor := range gryffindors{
    var words = strings.Fields(gryffindor)
    var firstName = words[0]
    firstNames = append(firstNames, firstName)
}
```
_ is a blank idetifier if you don't use index in the loop

```go
var remTickets = 50
for{
    var toBuy
    fmt.Printf("How many tickets: ")
    fmt.Scan(&toBuy)
    if toBuy > remTickets {
        fmt.Printf("Only have %v tickets, can't book that much, try again\n",remTickets)
        continue
    }
    else if remTickets <= toBuy {
        fmt.Printf("Sorry we r sold out: ")
        break
    }
    fmt.Printf("Succesfully bought %v tickets\n",toBuy)
}
```

### Functions
```go
func functionName(var1 type1, var2 type2, ...) retTypeIfAny{
    // codes to execute
    return varToReturnIfAny
}
functionName(value1, value2)
```

```go
func main(){
    firstNames := GetFirstNames(bookings)
}
func GetFirstNames(bookings []string) []string{
    firstNames := []string{}
    for _, booking := range bookings{
        var names = strings.Fields(booking)
        firstNames.append(firstNames, names[0])
    }
    return firstNames
}
```

U can return multiple variables from a function
```go
func main(){
    isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
}
func ValidateUserInput(firstName string, lastName, string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
    isValidName := len(firstName) >= 2 && len(lastName) >= 2
    isValidEmail := strings.Contains(email,"@")
    isValidTicketNumber := userTickets >0 && userTickets <= remainingTickets
    return isValidName, isValidEmail, isValidTicketNumber
}
```

### Packages
Go programs are organized into packages, a package is basically a collection of Go files. When creating a file, define which package it belongs to.
```go
package packageName
```
<code>go run file1.go file2.go fileX.go</code> or more easily <code>go run .</code>
Suppose this is your project structure
```
/projectDir
    /helper
        - helper.go
    - go.mod
    - main.go
```

In main.go you can import the package by using the module name in go.mod and the path to the package file.
```go
package main
import (
    "golang101/helper"
)

helper.CanBeUsed()
```

In helper.go you can define and choose which functions to be exported. Export just means that the function is publicly available in the file that imports it. To export it simply make sure that the name of the function starts with an uppercase.
```go
func CanBeUsed(){

}
func cannotBeUsed(){

}
```

Same thing with variable, just give it uppercase at first letter to make it publicly available.
```go

```

