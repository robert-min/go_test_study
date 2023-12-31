package mocking

import (
	"errors"
	"testing"
)

func TestThrowError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ThrowError(); (err != nil) != tt.wantErr {
				t.Errorf("DoSomeStuff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TEST
func TestDoSomeStuff(t *testing.T) {
	tests := []struct {
		name       string
		DoStuff    error
		ThrowError error
		wantErr    bool
	}{
		{"base-case", nil, nil, false},
		{"DoStuff error", errors.New("failed"), nil, true},
		{"ThrowError error", nil, errors.New("failed"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 모의 구조체를 사용해 인터페이스 테스트
			d := MockDoStuffer{}
			d.MockDoStuff = func(string) error { return tt.DoStuff }

			// mocking a function that is declared as a variable
			// will not work for func A(), must be var A = func()
			defer Patch(&ThrowError, func() error { return tt.ThrowError }).Restore()

			if err := DoSomeStuff(&d); (err != nil) != tt.wantErr {
				t.Errorf("DoSomeStuff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
