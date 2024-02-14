package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestParseArgs(t *testing.T){
	tests := []struct{
		args		[]string
		err			error
		config
		output		string
	}{
		{
			args: []string{"-h"},
			output: `
			A greeter application which prints the name you entered a specified number of times.

			Usage of greeter: <options> [name]

			Options: 
			-n int
					Number of times to greet
			`,
			err: errors.New("flag: help requested"),
			config: config{numTimes: 0},
		},
		{
			args: []string{"-n", "10"},
			err: nil,
			config: config{numTimes: 10},
		},
		{
			args: []string{"-n", "abc"},
			err: errors.New("invalid value \"abc\" for flag -n: parse error"),
			config: config{numTimes: 0},
		},
		{
			args: []string{"-n", "5", "praveen", "kumar"},
			err: errors.New("more than one positional argument specified"),
			config: config{numTimes: 5},
		},
	}

	byteBuf := new(bytes.Buffer)

	for _, tc := range tests{
		c, err := parseArgs(byteBuf, tc.args)
		if tc.err == nil && err != nil{
			t.Errorf("Expected nil error, Got: %v\n", err)
		}
		// if non-nil is there, means the value of errors should match, if not throw it
		if tc.err != nil && err.Error() != tc.err.Error(){
			t.Errorf("Expected %v, Got: %v\n", tc.err, err)
		}
		if c.numTimes != tc.numTimes{
			t.Errorf("Expected: %d, Got: %d\n", tc.numTimes, c.numTimes)
		}
		message := byteBuf.String()
		if len(tc.output) != 0 && message != tc.output{
			t.Errorf("Expected: %#v, Got: %#v\n", tc.output, message)
		}
		byteBuf.Reset()
	}

}