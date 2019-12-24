package _examples

import (
	"os"
	"text/tabwriter"

)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout,0,0,3, ' ', 0)

	opt := []string{"this", "is", "even"}

	gollum.Columnise(w, opt, 4)
	w.Flush()

}
