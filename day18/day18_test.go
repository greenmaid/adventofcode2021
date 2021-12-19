package day18

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	numbers := common.ReadFile("input.test1.txt")
	result := Step1(numbers)
	assert.Equal(t, 4140, result)
}

func TestPart2(t *testing.T) {
	numbers := common.ReadFile("input.test1.txt")
	result := Step2(numbers)
	assert.Equal(t, 3993, result)
}

func Test_getPair(t *testing.T) {
	tests := []struct {
		name      string
		number    string
		wantLeft  string
		wantRight string
	}{
		{name: "test1", number: "[1,2]", wantLeft: "1", wantRight: "2"},
		{name: "test2", number: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]", wantLeft: "[[[1,2],[3,4]],[[5,6],[7,8]]]", wantRight: "9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPair(tt.number)
			assert.Equal(t, tt.wantLeft, got[0])
			assert.Equal(t, tt.wantRight, got[1])
		})
	}
}

func Test_addRight(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		addValue int
		want     string
	}{
		{name: "test1", input: "0", addValue: 1, want: "1"},
		{name: "test2", input: "[1,3]]]", addValue: 5, want: "[6,3]]]"},
		{name: "test3", input: "[[[2,5]][1,9]]]", addValue: 9, want: "[[[11,5]][1,9]]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addRight(tt.input, tt.addValue)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_addLeft(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		addValue int
		want     string
	}{
		{name: "test1", input: "0", addValue: 1, want: "1"},
		{name: "test2", input: "[1,3]]]", addValue: 5, want: "[1,8]]]"},
		{name: "test3", input: "2,5]][1,9]]]", addValue: 2, want: "2,5]][1,11]]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addLeft(tt.input, tt.addValue)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_explodePair(t *testing.T) {
	type args struct {
		snailNumber string
		depth       int
	}
	tests := []struct {
		name          string
		args          args
		result        string
		leftReminder  int
		rightReminder int
	}{
		{name: "test1", args: args{snailNumber: "[1,2]", depth: 0}, result: "[1,2]", leftReminder: 0, rightReminder: 0},
		{name: "test2", args: args{snailNumber: "[1,2]", depth: 5}, result: "0", leftReminder: 1, rightReminder: 2},
		{name: "test3", args: args{snailNumber: "[[1,2],3]", depth: 3}, result: "[0,5]", leftReminder: 1, rightReminder: 0},
		{name: "test4", args: args{snailNumber: "[[[[[9,8],1],2],3],4]", depth: 0}, result: "[[[[0,9],2],3],4]", leftReminder: 9, rightReminder: 0},
		{name: "test5", args: args{snailNumber: "[7,[6,[5,[4,[3,2]]]]]", depth: 0}, result: "[7,[6,[5,[7,0]]]]", leftReminder: 0, rightReminder: 2},
		{name: "test6", args: args{snailNumber: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", depth: 0}, result: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", leftReminder: 0, rightReminder: 0},
		{name: "test6b", args: args{snailNumber: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", depth: 0}, result: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", leftReminder: 0, rightReminder: 2},
		{name: "test8", args: args{snailNumber: "[7,[6,[5,[4,[3,2]]]]]", depth: 0}, result: "[7,[6,[5,[7,0]]]]", leftReminder: 0, rightReminder: 2},
		{name: "test9", args: args{snailNumber: "[[6,[5,[4,[3,2]]]],1]", depth: 0}, result: "[[6,[5,[7,0]]],3]", leftReminder: 0, rightReminder: 0},
		{name: "test10", args: args{snailNumber: "[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]", depth: 0}, result: "[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]", leftReminder: 1, rightReminder: 0},
		{name: "test10", args: args{snailNumber: "[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]", depth: 0}, result: "[[[[3,0],[5,3]],[4,4]],[5,5]]", leftReminder: 0, rightReminder: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, leftReminder, rightReminder := explodePair(tt.args.snailNumber, tt.args.depth)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.leftReminder, leftReminder)
			assert.Equal(t, tt.rightReminder, rightReminder)
		})
	}
}

func Test_splitPair(t *testing.T) {
	tests := []struct {
		name        string
		snailNumber string
		want        string
	}{
		{name: "test1", snailNumber: "0", want: "0"},
		{name: "test2", snailNumber: "[12,3]", want: "[[6,6],3]"},
		{name: "test3", snailNumber: "[12,11]", want: "[[6,6],11]"},
		{name: "test4", snailNumber: "[[[[0,7],4],[15,[0,13]]],[1,1]]", want: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
		{name: "test5", snailNumber: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", want: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitPair(tt.snailNumber)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_addAndReducePair(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{num1: "[1,2]", num2: "[3,4]"}, want: "[[1,2],[3,4]]"},
		{name: "test2", args: args{num1: "[[[[4,3],4],4],[7,[[8,4],9]]]", num2: "[1,1]"}, want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addAndReducePair(tt.args.num1, tt.args.num2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calculateMagnitude(t *testing.T) {

	tests := []struct {
		name string
		num  string
		want int
	}{
		{name: "test1", num: "[9,1]", want: 29},
		{name: "test2", num: "[[9,1],[1,9]]", want: 129},
		{name: "test3", num: "[[1,2],[[3,4],5]]", want: 143},
		{name: "test4", num: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", want: 3488},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateMagnitude(tt.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_addSnailNumbers(t *testing.T) {
	input1 := []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}
	result1 := addSnailNumbers(input1)
	assert.Equal(t, "[[[[1,1],[2,2]],[3,3]],[4,4]]", result1)

	input1b := []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}
	result1b := addSnailNumbers(input1b)
	assert.Equal(t, "[[[[3,0],[5,3]],[4,4]],[5,5]]", result1b)

	input1c := []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}
	result1c := addSnailNumbers(input1c)
	assert.Equal(t, "[[[[5,0],[7,4]],[5,5]],[6,6]]", result1c)

	input2 := common.ReadFile("input.test2.txt")
	result2 := addSnailNumbers(input2)
	assert.Equal(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", result2)
}
