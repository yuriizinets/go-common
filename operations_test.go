package common

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(10, 2, 3).(int) != 15 {
		t.Error("Result of Sum(10, 2, 3) is incorrect")
	}
	if Sum(float32(10.0), float32(3.3), float32(2.5)).(float32) != 15.8 {
		t.Error("Result of Sum(float32(10.0), float32(3.3), float32(2.5)) is incorrect")
	}
	if Sum(float64(10.0), float64(3.3), float64(2.5)).(float64) != 15.8 {
		t.Error("Result of Sum(float64(10.0), float64(3.3), float64(2.5)) is incorrect")
	}
	if Sum("Foo", "Bar").(string) != "FooBar" {
		t.Error(`Result of Sum("Foo", "Bar") is incorrect`)
	}
}

func TestSub(t *testing.T) {
	if Sub(10, 2, 3).(int) != 5 {
		t.Error("Result of Sub(10, 2, 3) is incorrect")
	}
	if Sub(float32(10.0), float32(3.3), float32(2.5)).(float32) != 4.2 {
		t.Error("Result of Sub(float32(10.0), float32(3.3), float32(2.5)) is incorrect")
	}
	if Sub(float64(10.0), float64(3.3), float64(2.5)).(float64) != 4.2 {
		t.Error("Result of Sub(float64(10.0), float64(3.3), float64(2.5)) is incorrect")
	}
	if Sub("FooBar", "Bar") != "Foo" {
		t.Error(`Result of Sub("FooBar", "Bar") is incorrect`)
	}
}

func TestMult(t *testing.T) {
	if Mult(10, 2, 3).(int) != 60 {
		t.Error("Result of Mult(10, 2, 3) is incorrect")
	}
	if Mult(float32(10.0), float32(3.3), float32(2.5)).(float32) != 82.5 {
		t.Error("Result of Mult(float32(10.0), float32(3.3), float32(2.5)) is incorrect")
	}
	if Mult(float64(10.0), float64(3.3), float64(2.5)).(float64) != 82.5 {
		t.Error("Result of Mult(float64(10.0), float64(3.3), float64(2.5)) is incorrect")
	}
	if Mult("Foo", 3) != "FooFooFoo" {
		t.Error(`Result of Mult("Foo", 3) is incorrect`)
	}
}

func TestDiv(t *testing.T) {
	if Div(10, 2).(int) != 5 {
		t.Error("Result of Div(10, 2) is incorrect")
	}
	if Div(float32(1.2), float32(2)).(float32) != 0.6 {
		t.Error("Result of Div(float32(1.2), float32(2)) is incorrect")
	}
	if Div(float64(1.2), float64(2)).(float64) != 0.6 {
		t.Error("Result of Div(float64(1.2), float64(2)) is incorrect")
	}
}
