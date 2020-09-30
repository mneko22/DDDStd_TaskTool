package valueObject

import "testing"

func TestStatusValue_IsEqual(t *testing.T) {
	if result := DOING.IsEqual(&DOING); result != true {
		t.Fatal("Not Match")
	}
	if result := WAIT.IsEqual(&DOING); result != false {
		t.Fatal("Not Match")
	}
}

func TestStatusValue_NextStatus(t *testing.T) {
	if nextStatus := TODO.NextStatus(); !nextStatus.IsEqual(&WAIT) {
		t.Fail()
	}

	if nextStatus := WAIT.NextStatus(); !nextStatus.IsEqual(&DOING) {
		t.Fail()
	}

	if nextStatus := DOING.NextStatus(); !nextStatus.IsEqual(&DONE) {
		t.Fail()
	}

	if nextStatus := DONE.NextStatus(); !nextStatus.IsEqual(&DONE) {
		t.Fail()
	}

	if nextStatus := CANCEL.NextStatus(); !nextStatus.IsEqual(&TODO) {
		t.Fail()
	}
}

func TestStatusValue_PreviousStatus(t *testing.T) {
	if prevStatus := TODO.PreviousStatus(); !prevStatus.IsEqual(&TODO) {
		t.Fail()
	}

	if prevStatus := WAIT.PreviousStatus(); !prevStatus.IsEqual(&TODO) {
		t.Fail()
	}

	if prevStatus := DOING.PreviousStatus(); !prevStatus.IsEqual(&WAIT) {
		t.Fail()
	}

	if prevStatus := DONE.PreviousStatus(); !prevStatus.IsEqual(&DOING) {
		t.Fail()
	}

	if prevStatus := CANCEL.PreviousStatus(); !prevStatus.IsEqual(&TODO) {
		t.Fail()
	}
}
