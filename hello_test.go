package git_academy_2

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Bob")
	if result != "Hello Bob" {
		// error
		panic("Expected 'Hello Bob', got " + result)
	}
}

func TestLogin(t *testing.T) {
	result := Login("bob", "bob")
	if !result {
		panic("Invalid Login")
	}
}
