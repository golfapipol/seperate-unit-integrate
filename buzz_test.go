package unitegrate

import "testing"

func TestIsBuzzInput15ShouldBeTrue(t *testing.T) {
	number := 15
	expected := true
	actual := isBuzz(number)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}

func TestIsBuzzInput30ShouldBeFalse(t *testing.T) {
	number := 13
	expected := false
	actual := isBuzz(number)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}
