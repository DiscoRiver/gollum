package gollum

import (
	"fmt"
	"text/tabwriter"
)

// Columnise formats the values (opt) into specified number of columns (cc) and writes to a tabwriter.Writer interface (w).
func Columnise(w *tabwriter.Writer, opt []string, cc int) (err error){
	mod := len(opt) % cc
	divList := opt[:len(opt)-mod]
	modList := opt[len(divList):]

	for i := 0; i < len(divList); i=i+cc {
		var row string

		for t := 0; t < cc; t++ {
			row = row + fmt.Sprintf("%s\t", divList[i+t])
		}
		fmt.Fprintln(w, row)
	}

	if len(modList) != 0 {
		var row string
		for i := range modList {
			row = row + fmt.Sprintf("%s\t", modList[i])
		}
		fmt.Fprintln(w, row)
	}
}
