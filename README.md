# HasUpdates.exe
A command-line executable that lists the direct modules with updates and how signifant the update is.

# Installing
```
go install github.com/billgraziano/hasupdates
```

# Features
`hasupdates.exe` will
* Identify any direct modules that need to be updated.
* Major changes are listed in red.  _Note: `go list` doesn't list most major changes.  But if a change looks major, it is in red._
* Minor changes are listed in yellow.
* No changes are listed in green (and hiddent by default).
* The `-v` flag will display modules that are current.

![Screenshot](/docs/output.png)

## Under the Hood
* Run `go list -json -u -m all` and capture the output
* Parse into the `Module` struct from the `go list` documentation
* Format the output 


