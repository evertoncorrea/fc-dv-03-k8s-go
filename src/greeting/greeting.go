package greeting

import "fmt"

func Greeting(txt string) string {
	return fmt.Sprintf("<b>%s</b>", txt)
}
