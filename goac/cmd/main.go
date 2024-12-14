package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

type flag string

const (
	YEAR flag = "year"
	DAY  flag = "day"
)

func newDay(args map[flag]string) error {
	year, ok := args[YEAR]
	if !ok {
		year = strconv.Itoa(time.Now().Year())
	}
	_, ok = args[DAY]
	if !ok {
		args[DAY] = strconv.Itoa(time.Now().Day())
	}

	day, err := strconv.Atoi(args[DAY])
	if err != nil {
		return fmt.Errorf("invalid value for flag %v: %v", DAY, args[DAY])
	}

	yearDir := path.Join(".", year)
	fi, err := os.Stat(yearDir)
	if err != nil {
		return fmt.Errorf("cannot create day for non-existent year")
	} else if !fi.IsDir() {
		return fmt.Errorf("year directory exists not as directory")
	}

	templatePath := path.Join(yearDir, "day00.go")
	dayStr := fmt.Sprintf("%02d", day)
	destPath := path.Join(yearDir, "day"+dayStr+".go")

	fi, err = os.Stat(destPath)
	if err != nil {
		if os.IsNotExist(err) {
			// expected
			copyGoTemplate(templatePath, destPath, dayStr)
		} else if os.IsNotExist(err) {
			return fmt.Errorf("day already exists")
		} else {
			return errors.New("unknown error occurred")
		}
	} else {
		fmt.Printf("%v already exists, will not overwrite...\n", fi.Name())
	}

	for _, toTouch := range []string{
		path.Join(yearDir, "example_data", "day"+dayStr+".txt"),
		path.Join(yearDir, "input_data", "day"+dayStr+".txt"),
	} {
		err = exec.Command("touch", toTouch).Run()
		if err != nil {
			return fmt.Errorf("unable to create %v", toTouch)
		}
	}

	return nil
}

func copyGoTemplate(template, dest, day string) error {
	read, err := os.ReadFile(template)
	if err != nil {
		return err
	}

	newContents := strings.ReplaceAll(string(read), "00", day)

	err = os.WriteFile(dest, []byte(newContents), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func parseArgs(args []string) (map[flag]string, error) {
	if len(args)%2 != 0 {
		return nil, errors.New("uneven number of input arguments")
	}

	parsedArgs := map[flag]string{}
	for i := 0; i < len(args); i += 2 {

		value := args[i+1]
		switch args[i] {
		case "-y", "--year":
			parsedArgs[YEAR] = value
		case "-d", "--day":
			parsedArgs[DAY] = value
		default:
			return nil, fmt.Errorf("invalid flag: '%v'", args[i])
		}
	}
	return parsedArgs, nil
}

func help() string {
	return `go advent of code CLI 🎄
functions
  new-day
    Create a new day using a template.

    args
      Optional
      --year, -y 	Year in which to create the files. Defaults to current year.
      --day, -d  	Day for which to create the files. Defaults to current date.
  help
    Print this help doc.

    args
      (None)
`
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(help())
		return
	}

	switch args[0] {
	case "new-day":
		input, err := parseArgs(args[1:])
		if err != nil {
			fmt.Printf("new day error: %v\n", err)
			return
		}
		err = newDay(input)
		if err != nil {
			fmt.Printf("new day error: %v\n", err)
			return
		}
	case "help":
		fmt.Println(help())
	default:
		fmt.Printf("invalid command: %v\n", args[0])
		fmt.Println(help())
	}
}
