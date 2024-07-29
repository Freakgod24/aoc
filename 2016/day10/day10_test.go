package main

import (
	"reflect"
	"testing"
)

func Test_parseInstruction(t *testing.T) {
	type args struct {
		instructionStr string
	}
	tests := []struct {
		name    string
		args    args
		want    Instruction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test 1", args{"value 5 goes to bot 2"}, Instruction{InstructionGoes, LocationInput, Index(5), LocationBot, Index(2), LocationNone, IndexNone}, false},
		{"Test 2", args{"bot 2 gives low to bot 1 and high to bot 0"}, Instruction{InstructionGives, LocationBot, Index(2), LocationBot, Index(1), LocationBot, Index(0)}, false},
		{"Test 3", args{"value 3 goes to bot 1"}, Instruction{InstructionGoes, LocationInput, Index(3), LocationBot, Index(1), LocationNone, IndexNone}, false},
		{"Test 4", args{"bot 1 gives low to output 1 and high to bot 0"}, Instruction{InstructionGives, LocationBot, Index(1), LocationOutput, Index(1), LocationBot, Index(0)}, false},
		{"Test 5", args{"bot 0 gives low to output 2 and high to output 0"}, Instruction{InstructionGives, LocationBot, Index(0), LocationOutput, Index(2), LocationOutput, Index(0)}, false},
		{"Test 6", args{"value 2 goes to bot 2"}, Instruction{InstructionGoes, LocationInput, Index(2), LocationBot, Index(2), LocationNone, IndexNone}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInstruction(tt.args.instructionStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
