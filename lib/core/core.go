package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func CorrectAndRun(input, output string) string {
	cwd, err := os.Getwd()
	if err != nil {
		return "Directory problems"
	}
	input = path.Join(cwd, "programs", input)
	infd, err := os.Open(input)
	if err != nil {
		return fmt.Sprintf("Unable to open file %q\n", input)
	}
	defer infd.Close()

	syntax := []string{}
	scanner := bufio.NewScanner(infd)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return "Error while processing file"
		}
		line := scanner.Text()
		result := suggest(line)
		syntax = append(syntax, result)
	}

	output = path.Join(cwd, "syntax", output)
	outfd, err := os.Create(output)
	if err != nil {
		return fmt.Sprintf("Unable to open file %q\n", output)
	}
	defer outfd.Close()

	syntaxon := Syntax{
		Language: "Go",
		Sourcefile: input,
		Syntaxfile: output,
		Syntax: syntax,
	}
	json.NewEncoder(outfd).Encode(syntaxon)

	return "Success!"
}
