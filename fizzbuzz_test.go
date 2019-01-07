// +build integration

package unitegrate

import "testing"

func TestFizzBuzzInput15ShouldBeFizzBuzz(t *testing.T) {
	number := "15"
	expected := "FizzBuzz"
	actual := FizzBuzz(number)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func TestFizzBuzzInput30ShouldBeFizzBuzz(t *testing.T) {
	number := "30"
	expected := "FizzBuzz"
	actual := FizzBuzz(number)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func TestFizzBuzzInput31ShouldBe31(t *testing.T) {
	number := "31"
	expected := "31"
	actual := FizzBuzz(number)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}
