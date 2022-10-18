package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bugsfunny/dictionary/dictionary"
)

func main() {

	action := flag.String("action", "list", "Action to perform on the dictionary")

	dict, err := dictionary.New("./badger")
	handleError(err)
	defer dict.CLose()

	flag.Parse()
	switch *action {
	case "list":
		actionList(dict)
	case "add":
		actionAdd(dict, flag.Args())
	case "define":
		actionDefine(dict, flag.Args())
	case "remove":
		actionRemove(dict, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	error := d.Remove(word)
	handleError(error)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, error := d.Get(word)
	handleError(error)
	fmt.Println(entry)
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	def := args[1]
	err := d.Add(word, def)
	handleError(err)
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleError(err)
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
