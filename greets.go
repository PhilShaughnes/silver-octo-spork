package main

var (
	thing string = "hello"
)

func hi(s string) string {
	greet := "Hello"
	if s != "" {
		greet += " " + s
	}
	return greet + "!"
}

func bye(s string) string {
	farewell := "Goodbye"
	return farewell + " " + s
}
