package format

import "fmt"

//MajorMinorFormat is visible outside scope package
func MajorMinorFormat(input string) string {
	return fmt.Sprintf("<< %s >>", input)
}

// emptyLine is visible only in the package
func emptyLine() string {
	return "______________________"
}

// MessageInABox pretty-prints a message
func MessageInABox(input string) {
	fmt.Println(fmt.Sprintf("%s\n%s\n%s", emptyLine(), MajorMinorFormat(input), emptyLine()))
}

func textSeparator() string {
	return "***"
}
