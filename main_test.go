package main

import "testing"

// TestUnknownCommand verifies there is an error and it has expected value when a command is unknown
func TestUnknownCommand(t *testing.T) {
	commandErr := executeCommand("llamallamaredpajama")
	if commandErr == nil {
		t.Error("Expected error but got nil")
	}
	expectedErr := "command: llamallamaredpajama is not a valid command"
	if commandErr.Error() != expectedErr {
		t.Errorf("Expected error '%s' does not equal actual error '%s'", expectedErr, commandErr)
	}
}
