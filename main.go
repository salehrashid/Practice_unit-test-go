package go_unit_test

func FuncHello(name string) string {
	return "Hello " + name
}

func CompleteName(firstName, middleName, lastName string) (string, string, string) {
	return firstName, middleName, lastName
}

func Add(a, b int) int {
	return a + b
}
