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

func TestMin(t *testing.T) {
	if Min(10, 2, 5).(int) != 2 {
		t.Error("Result of Min(10, 2, 5) is incorrect")
	}
	if Min(float32(10.3), float32(2.3), float32(3.5)).(float32) != 2.3 {
		t.Error("Result of Min(float32(10.3), float32(2.3), float32(3.5)) is incorrect")
	}
	if Min(float64(10.3), float64(2.3), float64(3.5)).(float64) != 2.3 {
		t.Error("Result of Min(float64(10.3), float64(2.3), float64(3.5)) is incorrect")
	}
	if Min("Fooo", "Bar").(string) != "Bar" {
		t.Error(`Result of Min("Fooo", "Bar") is incorrect`)
	}
}

func TestMax(t *testing.T) {
	if Max(10, 2, 5).(int) != 10 {
		t.Error("Result of Max(10, 2, 5) is incorrect")
	}
	if Max(float32(10.3), float32(2.3), float32(3.5)).(float32) != 10.3 {
		t.Error("Result of Max(float32(10.3), float32(2.3), float32(3.5)) is incorrect")
	}
	if Max(float64(10.3), float64(2.3), float64(3.5)).(float64) != 10.3 {
		t.Error("Result of Max(float64(10.3), float64(2.3), float64(3.5)) is incorrect")
	}
	if Max("Bar", "Fooo").(string) != "Fooo" {
		t.Error(`Result of Max("Fooo", "Bar") is incorrect`)
	}
}

func TestAvg(t *testing.T) {
	if Avg(3, 5, 2, 5).(int) != 3 {
		t.Error("Result of Avg(10, 2, 5) is incorrect")
	}
	if Avg(float32(3.5), float32(3.3), float32(3.4)).(float32) != 3.4000003 { // Division specifics (?)
		t.Error("Result of Avg(float32(3.5), float32(3.3), float32(3.4)) is incorrect")
	}
	if Avg(float64(3.5), float64(3.3), float64(3.4)).(float64) != 3.4 { // No division specifics (?)
		t.Error("Result of Avg(float32(3.5), float32(3.3), float32(3.4)) is incorrect")
	}
	if Avg("Bar", "Fooo").(float64) != 3.5 {
		t.Error(`Result of Avg("Fooo", "Bar") is incorrect`)
	}
}
