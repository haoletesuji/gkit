# Call your code from another module

1. From the command prompt in the app directory, run the following command:
```bash
go mod edit -replace github.com/thienhaole92/chat/common =../common
```
2. From the command prompt in the app directory, run the go mod tidy command to synchronize the `github.com/thienhaole92/chat/common` module's dependencies, adding those required by the code, but not yet tracked in the module.
```bash
go mod tidy
```
