package foo

import "os"

func main() {
	os.Stderr.WriteString("your message here")
}
