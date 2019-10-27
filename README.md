# demo-magician-errors

A quick demonstration of the new error related features available as of Go 1.13. The example demonstrates how to wrap errors to create an error chain and how to check for an specific error in this whole chain without having to traverse it manually. You can start reading the code from `cmd/demo-magician-errors/main.go`.

For more information regarding the usage of features check the official blog post at: [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)

![Screenshot](https://i.imgur.com/zqdRZOG.png)