package main

import (
"fmt"
"github.com/russross/blackfriday"

)


func mdparse(markdown string) string {
	input := []byte(markdown)
	output := string(blackfriday.MarkdownCommon(input))
	return output
}

func main() {
  fmt.Println(mdparse("# Heading"))


}



