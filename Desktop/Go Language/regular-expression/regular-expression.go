package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a - z]+)ch","peach");

	fmt.Println(match)

	r, _ := regexp.Compile("p([a - z]+)ch");

	fmt.Println(r.MatchString("peach"));

	fmt.Println(r.FindString("peach punch"));
	

}