package example

import "fmt"

func ExampleDemo_Hello_spanish() {
	greetings, err := Demo{}.Hello("Anderson")
	if err != nil {
		panic(err)
	}
	fmt.Println(greetings)

	// Output: Hello, Anderson
}
