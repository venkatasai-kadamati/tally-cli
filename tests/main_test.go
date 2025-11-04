package main

import "testing"
import "github.com/venkatasai-kadamati/tally-cli/algo"

// Test is the convention that lets go know on "go test" this function should be called
// Mandatory param -> (t *testing.T)
func TestCountWords(t *testing.T) {
	// test for what inputs: keyword used is "input"
	input := "one two three four five"

	// what are we expecting for the above input: keyword used is "wants"
	wants := 5

	result := algo.CountWords(input)

	// assertion -> compare wants with result. do they match?
	if result != wants {
		t.Fail()
	}
}
