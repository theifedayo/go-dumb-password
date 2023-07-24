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
validator, err := dumbpassword.DPValidator(passwordListPath)
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
The `Validate` method returns `true` if the provided password is not found in the list of common passwords loaded from the password list file. If the password is found in the list, it returns `false`, indicating that the password is common and should be avoided.


### Get Cache Key
To improve performance, the Dumb Password Validator utilizes caching. The GetCacheKey method is used to retrieve the cache key, which serves as an identifier for the cached data (the list of common passwords) associated with the password list file:
```go
cacheKey, err := validator.GetCacheKey()
if err != nil {
    // Handle the error if the cache key cannot be retrieved
} else {
    // Use cacheKey for any cache-related operations
}
```
The cache key is calculated based on the MD5 hash of the password list file's content. It helps the validator quickly determine if the cached data is still valid or if the file needs to be read again to update the cache.


## Contributing
If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request. Your contributions are highly appreciated!

## License
This library is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments
The go dumb password library is inspired by the [Laravel Password](https://github.com/unicodeveloper/laravel-password) package created by [Prosper Otemuyiwa](https://github.com/unicodeveloper/).
