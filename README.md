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
Go programs are organized into packages, Go's standar library provides different core packages for us to use such as "fmt". Run code by <code>go run {fileName}.go</code>

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

### Input
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