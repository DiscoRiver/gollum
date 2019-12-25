# Gollum
Columnised menu output for text/tabwriter.

##Example:

```
package main

import (
	"os"
	"text/tabwriter"
	"gollum"
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout,0,0,3, ' ', 0)

	opt := []string{"this", "is", "a", "columnised", "output", "for", "user", "reference"}

	gollum.Columnise(w, opt, 4)

	// Write the buffer
	w.Flush()
}
```

##Output:

```
this     is    a      columnised   
output   for   user   reference    
```


