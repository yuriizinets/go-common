package common

import (
	"testing"
)

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

func TestIn(t *testing.T) {
	if In(3, []int{4, 3, 5}) == false {
		t.Error("Result of In(3, []int{4, 3, 5}) is incorrect")
	}
	if In(3.5, []float64{3.6, 3.5, 3.7}) == false {
		t.Error("Result of In(3.5, []float64{3.6, 3.5, 3.7}) is incorrect")
	}
	if In("foo", []string{"foo", "bar"}) == false {
		t.Error(`Result of In("foo", []string{"foo", "bar"}) is incorrect`)
	}
}

func TestRemove(t *testing.T) {
	if len(Remove(3, []int{4, 3, 5}).([]int)) != 2 {
		t.Error("Result of Remove(3, []int{4, 3, 5}) is incorrect")
	}
	if len(Remove("foo", []string{"foo", "bar"}).([]string)) != 1 {
		t.Error(`Result of Remove("foo", []string{"foo", "bar"}) is incorrect`)
	}
}
