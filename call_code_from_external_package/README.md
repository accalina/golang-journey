# Steps

Create a simple program which call some code from external lib

```sh
> `1. Enable dependency tracking:`
$ go mod init github.com/accalina/golang-journey/call_code_from_external_package

> `2. Search the external lib on pkg.go`
$ firefox 'https://pkg.go.dev/search?q=quote'

> `3. Write your code`
> `4. Install the depedency`
$ go mod tidy
> `4. run your code`
$ go run .
> `5. build your code`
$ go build .
```
