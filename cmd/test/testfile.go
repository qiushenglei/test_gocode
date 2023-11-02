package main

import (
	"fmt"
	"html"
)

func a() {
	a := html.EscapeString("<html> </html>")
	fmt.Println(a)
	fmt.Println(html.UnescapeString(a))
}
