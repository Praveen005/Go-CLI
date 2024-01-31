package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T){
	tests := []struct{
		c		config
		input	string
		output	string
		err		error
	}{
		{
			c:			config{printUsage: true},
			output: 	usageString,
		},
		{
			c:			config{numTimes: 5},
			input: 		"",
			output: 	strings.Repeat("Your name please? Press Enter when done.\n", 1) ,
			err:		errors.New("you didn't enter your name"),
		},
		{
			c:			config{numTimes: 5},
			input: 		"Praveen",	
			output: 	"Your name please? Press the Enter key when done.\n" + strings.Repeat("Nice to meet you, Praveen\n", 5),
		},
	}

	byteBuf := new(bytes.Buffer)

	for _, tc := range tests {

		r := strings.NewReader(tc.input)
		err := runCmd(r, byteBuf, tc.c)
		if err != nil && tc.err == nil{
			t.Errorf("Expected nil error, Got: %s\n", err.Error())
		}
		if tc.err != nil && tc.err.Error() != err.Error(){
			t.Errorf("Expected: %s, Got: %s\n", tc.err.Error(), err.Error())
		}
		byteBuf.Reset()
	}
}