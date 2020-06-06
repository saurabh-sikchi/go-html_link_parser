package main

import (
	"fmt"
	"strings"

	link "github.com/saurabh-sikchi/go-html_link_parser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
	<a href="/other-page">A link to another page</a>
	<a href="/other-page">
		Bar
		Baz
	</a>
	<a href="/other-page">
		<span> <span> Foo </span> Bar </span>
		Baz
	</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
