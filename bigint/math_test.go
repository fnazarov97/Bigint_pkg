package bigint

import (
	"testing"
)

type Test struct {
	arg1, arg2, expected Bigint
}

func TestAdd(t *testing.T) {
	var addTests = []Test{
		{Bigint{Value: "6565"}, Bigint{"13"}, Bigint{"6578"}},
		{Bigint{Value: "13"}, Bigint{"6565"}, Bigint{"6578"}},
		{Bigint{Value: "13"}, Bigint{"13"}, Bigint{"26"}},
		{Bigint{Value: "13"}, Bigint{"-13"}, Bigint{"0"}},
		{Bigint{Value: "-1236"}, Bigint{"-123"}, Bigint{"-1359"}},
		{Bigint{Value: "23"}, Bigint{"-333"}, Bigint{"-310"}},
		{Bigint{Value: "-26653"}, Bigint{"13"}, Bigint{"-26640"}},
	}
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}

	}
}
func TestSub(t *testing.T) {
	var subTests = []Test{
		{Bigint{Value: "6565"}, Bigint{"13"}, Bigint{"6552"}},
		{Bigint{Value: "13"}, Bigint{"6565"}, Bigint{"-6552"}},
		{Bigint{Value: "13"}, Bigint{"22"}, Bigint{"-9"}},
		{Bigint{Value: "13"}, Bigint{"13"}, Bigint{"0"}},
		{Bigint{Value: "-1236"}, Bigint{"-123"}, Bigint{"-1113"}},
		{Bigint{Value: "223"}, Bigint{"-333"}, Bigint{"556"}},
		{Bigint{Value: "-26653"}, Bigint{"13"}, Bigint{"-26666"}},
	}
	for _, test := range subTests {
		if output := Sub(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}

	}
}

func TestMultiply(t *testing.T) {
	var mulTests = []Test{
		{Bigint{Value: "6565"}, Bigint{"13"}, Bigint{"85345"}},
		{Bigint{Value: "0"}, Bigint{"13"}, Bigint{"0"}},
		{Bigint{Value: "13"}, Bigint{"6565"}, Bigint{"85345"}},
		{Bigint{Value: "13"}, Bigint{"22"}, Bigint{"286"}},
		{Bigint{Value: "13"}, Bigint{"13"}, Bigint{"169"}},
		{Bigint{Value: "-1236"}, Bigint{"-123"}, Bigint{"152028"}},
		{Bigint{Value: "223"}, Bigint{"-333"}, Bigint{"-74259"}},
		{Bigint{Value: "-26653"}, Bigint{"13"}, Bigint{"-346489"}},
	}
	for _, test := range mulTests {
		if output := Multiply(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}

	}
}

func TestMod(t *testing.T) {
	var modTests = []Test{
		{Bigint{Value: "6565"}, Bigint{"13"}, Bigint{"0"}},
		{Bigint{Value: "10"}, Bigint{"564"}, Bigint{"10"}},
		{Bigint{Value: "13"}, Bigint{"5"}, Bigint{"3"}},
	}
	for _, test := range modTests {
		if output := Mod(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}

	}
}
func TestAbs(t *testing.T) {
	input1 := Bigint{Value: "-555565644"}
	expected := "555565644"
	input2 := Bigint{Value: "555565644"}
	if output := input1.Abs(); output.Value != "555565644" {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
	if output := input2.Abs(); output.Value != "555565644" {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}
func TestSet(t *testing.T) {
	input1 := Bigint{Value: "000555565644"}
	input2 := Bigint{Value: "+555565644"}
	input1.Set("000555565644")
	if input1.Value != "555565644" {
		t.Errorf("Output not equal to expected ")
	}
	input2.Set("000555565644")
	if input1.Value != "555565644" {
		t.Errorf("Output not equal to expected ")
	}
}
