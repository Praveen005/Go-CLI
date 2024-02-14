package main

import (
	"errors"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		config
		err error
	}{
		{
			config: config{numTimes: 5},
		},
		//Not mentioning config here means it will take a default value i.e. zero here
		{
			err: errors.New("must specify a number greater than 0"),
		},
		{
			config: config{numTimes: -1},
			err: errors.New("must specify a number greater than 0"),
		},
	}

	for _, tc := range tests{
		err := validateArgs(tc.config)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, got: %v\n", err)
		}
		if tc.err != nil && tc.err.Error() != err.Error(){
			t.Fatalf("Expected: %v, Got: %v\n", tc.err, err)
		}
	}
}