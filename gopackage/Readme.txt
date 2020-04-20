As of 1.11 you don't need to put everything in GOPATH

This gives example of creating custom packages and using dependency manager (go modules)

It is following this blog https://www.callicoder.com/golang-packages/

1. Create a go module  go mod init github.com/neilm/go-package-modules-example

>> this creates a go.mod file with following in it
   module github.com/neilm/gopackage

   go 1.14


2. Create some packages in the module
     /numbers
     /strings

3. Add the files prime.go and reverse.go under these folders

4. Create a nested package under the strings 
    /strings/greetings/texts.go

5. Create the main package
    /main.go

6. run go build


Notes:
1. ALthough the module path references github.com it actually looks for the package locally first


8. Third-party packages can be added by including in the import.

9. Do go build. This downloads the package and adds reference to go.mod and creates a go.sum
   >>>require rsc.io/quote v1.5.2

10. You can specify an output bin and filename by doing
       go build -o bin/my.exe