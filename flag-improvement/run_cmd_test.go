package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T){
	// will have two outputs: err or greeting message
	// will take config.

	tests := []struct{
		config
		err 		error
		input		string
		output		string
	}{
		{
			config: config{numTimes: 6},
			output: "Your name please? Press Enter when done\n",
			err: errors.New("you didn't enter your name"),
		},
		{	
			config: config{numTimes: 5, name: "Praveen"},
			input: "",
			output: strings.Repeat("Nice to meet you, Praveen\n", 5),
		},
		{
			config: config{numTimes: 3},
			input: "Prateek",
			output: "Your name please? Press Enter when done\n" + strings.Repeat("Nice to meet you, Prateek\n", 3),

		},
	}

	byteBuf := new(bytes.Buffer)
	for _ , tc := range tests{
		r := strings.NewReader(tc.input)
		err := runCmd(r, byteBuf, tc.config)
		
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, Got: %v\n", err)
		}
		if tc.err != nil && tc.err.Error() != err.Error(){
			t.Fatalf("Expected: %v, Got: %v\n", tc.err, err)
		}

		message:= byteBuf.String()

		if message != tc.output {
			t.Fatalf("Expected: %v \nGot: %v\n", tc.output, message)
		}
		byteBuf.Reset()
	}
}