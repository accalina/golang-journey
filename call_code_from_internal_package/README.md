# Steps

Create a simple program which call some code from internal lib

```sh
> `1. Enable dependency tracking:`
$ go mod init github.com/accalina/golang-journey/call_code_from_internal_package

> `2. Write your code`

> `3. Replace the package to local path`
$ go mod edit -replace=github.com/accalina/golang-journey/create_go_module=../create_go_module

> `4. Install the library`
$ go mod tidy

> `5. Run the code`
$ go run .

> `5. Build the code`
$ go build .
```
