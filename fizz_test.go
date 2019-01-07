package unitegrate

import "testing"

func TestIsFizzInput15ShouldBeTrue(t *testing.T) {
	number := 15
	expected := true
	actual := isFizz(number)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}

func TestIsFizzInput30ShouldBeFalse(t *testing.T) {
	number := 13
	expected := false
	actual := isFizz(number)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}
