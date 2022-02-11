package main
import "fmt"
func handlePanic() {
	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
		
	}
	fmt.Println("I am stopping here by closing all connections")
}
func entry(language *string, name *string) {
	defer handlePanic()       //2
	fmt.Println("I am first") //1
	if language == nil {
		panic("Error: Language Cannot be nil")
	}
	if name == nil {
		panic("Error name cannot be nil") //4
	}
	fmt.Println("language", language)
	fmt.Println("name", name)
}
func main() {
	defer fmt.Println("I am the main func")
	lang := "go lang"
	entry(&lang, nil)
	fmt.Println("I am good")
}
