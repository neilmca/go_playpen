package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/neilm/gopackage/numbers"
	"github.com/neilm/gopackage/strings"
	"github.com/neilm/gopackage/strings/greetings" // Importing a nested package
	"rsc.io/quote"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))

	fmt.Println(quote.Go())
}