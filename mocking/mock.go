package mocking

// Simple interface for mocking test
type DoStuffer interface {
	DoStuff(input string) error
}
