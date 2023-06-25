# HasUpdates.exe
This is a simple command-line executable to identify which modules in a repository need to be updated and how signifant that update is.

# Installing
```
go install github.com/billgraziano/hasupdates`
```

# Features
`hasupdates.exe` will
* Identify any direct modules that need to be updated
* Based on the package semantic version differences
    * Major changes are listed in red
    * Minor changes are listed in yellow
    * No changes are listed in green (and hiddent by default)
* The `-v` flag will display modules without changes

![Screenshot](/docs/output.png)

## Under the Hood
* Run `go list -json -u -m all` and capture the output
* Parse into the `Module` struct from the `go list` documentation
* Format the output 
