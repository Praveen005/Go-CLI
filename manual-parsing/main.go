package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct{
	numTimes		int
	printUsage		bool

}


func parseArgs(args []string) (config, error){
	var numTimes int
	var err error

	c:= config{}

	if len(args) != 1{
		return c, errors.New("invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "--help"{
		c.printUsage =true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil{
		return c, err
	}
	c.numTimes = numTimes
	return c, nil
}

func validateArgs(c config) error{
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

// r is a Reader interface(ex. stdin)
// w is a Writer interface(ex. stdout)
func getName(r io.Reader, w io.Writer)(string, error){
	msg := "Your name please? Press Enter when done.\n"

	fmt.Fprint(w, msg)
	scanner := bufio.NewScanner(r) // a variable of scanner type is retured to read from r
	scanner.Scan() // scan the Reader for any input data, reads till newline char

	if err := scanner.Err(); err != nil{  // check if there was an error while using a scanner
		return "", err
	}

	name := scanner.Text() // returns data as a string

	if len(name) == 0{ // check if the user entered an empty name
		return "", errors.New("you didn't enter your name")
	}
	return name, nil
}

// stdin & stdout are satisfies of Reader & Writer interface but we don't em cz writting test will become difficult
// Input to hame test function se pass karna hoga na, naki stdin se.
func runCmd(r io.Reader, w io.Writer, c config) error{
	if c.printUsage{
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil{
		return err
	}

	greetUser(c, name, w)
	return nil
}

func greetUser(c config, name string, w io.Writer){
	msg := fmt.Sprintf("Nice to meet you, %s\n", name)

	for i := 0; i < c.numTimes; i++{
		fmt.Fprint(w, msg)
	}
}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|-help]

A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])

func printUsage(w io.Writer){
	fmt.Fprint(w, usageString)
}

func main(){
	c, err := parseArgs(os.Args[1:])
	if err != nil{
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil{
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil{
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
