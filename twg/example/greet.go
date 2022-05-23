package example

import "fmt"

type Demo struct{}

func (d Demo) Hello(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}

func NoneHello(name string) {
	fmt.Println("Hello,", name)
}

func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in.", name)
		}
	}
}
