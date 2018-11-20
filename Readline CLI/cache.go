//
// cache.go
//
// A basic cache written in Go.
// This is an exercise where we will implement a simple key value store written in Go.
// We will use a simple readline interface and two commands: PUT and GET.
//
// Requirements:
//
// 1. PUT key value     Set a value in the cache.
// 2. GET key           Get a value stored in the cache.
// 3. EXIT/QUIT         Exits the interactive prompt (can also be done with Ctrl-d thanks to the readline pkg).
// 4. Use only packages from the stdlib (except for the readline package already imported below).
//
package main

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
)

var m = make(map[string]string)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("put"),
	readline.PcItem("get"),
	readline.PcItem("exit"),
)

func main() {
	prompt, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer prompt.Close()

	for {
		line, err := prompt.Readline()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "put"):
			line := strings.TrimSpace(line[3:])
			index := strconv.Itoa(len(m) + 1)
			m[index] = line
			log.Printf("value %v inserted at: %v", line, index)
		case strings.HasPrefix(line, "get"):
			line := strings.TrimSpace(line[3:])
			log.Printf("value %v received from: %v", m[line], line)
		case line == "exit":
			goto exit
		case line == "":
		}
	}
exit:
}
