# Go Dumb Password

`Go Dumb Password` is a Go library that provides a simple and efficient way to check if a given password is common or "dumb." It loads a list of common passwords from a file and allows you to validate passwords against this list to encourage users to choose stronger and less predictable passwords.

## Installation
You can install the library using the `go get` command:
```bash
go get github.com/theifedayo/go-dumb-password
```

## Usage
To use this library in your Go application, you can simply import it as follows:

```go
import (
    "github.com/theifedayo/go-dumb-password"
)
```

### Creating a DumbPassword Validator
To create a new instance of the DumbPasswordValidator, you need to provide the path to the `passwordlist.txt` file containing the list of common passwords:
```go
passwordListPath := "/path/to/passwordlist.txt"
validator, err := dumbpassword.NewDumbPasswordValidator(passwordListPath)
if err != nil {
    // Handle the error if the file cannot be read or other issues occur during initialization.
}
```

### Validating Passwords
You can use the `Validate` method of the DumbPasswordValidator to check if a given password is common or "dumb":
```go
password := "password123"
if validator.Validate(password) {
    fmt.Println("Password is not common.")
} else {
    fmt.Println("Password is common. Please try another!")
}
```

### Caching
The library supports caching of the password list to avoid reading the file on every validation request. By default, it uses a simple in-memory cache (`sync.Map`). If you prefer to use a different caching mechanism, you can customize it by implementing your own cache functions.
