package mocking

import "reflect"

// Restorer can be used to restore some previous state.
type Restorer func()

// Restore restores some previous state.
func (r Restorer) Restore() {
	r()
}

// Patch sets the value pointed to by the given dest to the given value,
// and returns a function to restore it to its original value.
// The value must be assginable to the element type of the destination
func Patch(dest, value interface{}) Restorer {
	// reflect.Valueof : 어떤 인터페이스로부터, 가지고 있는 Value를 꺼내올 수 있음
	// .Elem : Value가 가지고 있는 값을 가져옴
	destv := reflect.ValueOf(dest).Elem()
	oldv := reflect.New(destv.Type()).Elem()
	oldv.Set(destv)
	valuev := reflect.ValueOf(value)
	if !valuev.IsValid() {
		// destv 타입이 nilable이 아닌 경우
		valuev = reflect.Zero(destv.Type())
	}
	destv.Set(valuev)
	return func() {
		destv.Set(oldv)
	}
}
