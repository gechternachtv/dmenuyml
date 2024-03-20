package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/go-yaml/yaml"
)

func errorHandling(err error, stderr string) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	if stderr != "" {
		fmt.Printf("stderr: %s\n", stderr)
		return
	}
}

func bar(options map[string]interface{}) {
	var optionsstr []string
	for key := range options {
		optionsstr = append(optionsstr, key)
	}

	cmd := exec.Command("dmenu")
	//cmd := exec.Command("rofi", "-dmenu")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("error obtaining stdin: %s\n", err)
		return
	}
	go func() {
		defer stdin.Close()
		fmt.Fprintf(stdin, "%s", strings.Join(optionsstr, "\n"))
	}()

	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error running dmenu: %s\n", err)
		return
	}

	filteredfoo := options[strings.TrimSpace(string(out))]

	if subOptions, ok := filteredfoo.(map[interface{}]interface{}); ok {
		subOptionsMap := make(map[string]interface{})
		for k, v := range subOptions {
			subOptionsMap[fmt.Sprintf("%v", k)] = v
		}
		bar(subOptionsMap)
	} else {
		if cmd, ok := filteredfoo.(string); ok {
			output, err := exec.Command(cmd).Output()
			if err != nil {
				fmt.Printf("error executing command: %s\n", err)
				return
			}
			fmt.Println(string(output))
			os.Exit(0)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no input file")
		return
	}

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file: %s\n", err)
		return
	}

	var foo map[string]interface{}
	err = yaml.Unmarshal(data, &foo)
	if err != nil {
		fmt.Printf("yaml error: %s\n", err)
		return
	}

	bar(foo)
}
