package valueObject

import "reflect"

type Value interface {
	IsEqual(value Value) bool
	GetValue() int
}

const (
	todo = iota
	wait
	doing
	done
	cancel
)

type statusValue struct {
	value int
}

func (sv *statusValue) IsEqual(target Value) bool {
	return reflect.DeepEqual(sv, target)
}

func (sv *statusValue) GetValue() int {
	return sv.value
}

func (sv *statusValue) PreviousStatus() *statusValue {
	switch sv.value {
	case wait:
		return &TODO
	case doing:
		return &WAIT
	case done:
		return &DOING
	default:
		return &TODO
	}
}

func (sv *statusValue) NextStatus() *statusValue {
	switch sv.value {
	case todo:
		return &WAIT
	case wait:
		return &DOING
	case doing:
		return &DONE
	case done:
		return &DONE
	default:
		return &TODO
	}
}

var (
	TODO   = statusValue{value: todo}
	WAIT   = statusValue{value: wait}
	DOING  = statusValue{value: doing}
	DONE   = statusValue{value: done}
	CANCEL = statusValue{value: cancel}
)
