package main

import (
	"fmt"
	"go-tutorial/tutorial/dict/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	dictionary.Search("first")
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("first", "First")
	fmt.Println(err)

	err = dictionary.Add("second", "Second")
	fmt.Println(err)

	err = dictionary.Update("second", "Update Second")
	fmt.Println(err)
}
