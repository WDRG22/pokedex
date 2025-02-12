// repl/commands_test.go
package repl

import (
    "testing"
)

func TestGetCommands(t *testing.T) {
    commands := getCommands()
    
    // Just test that we have all expected commands
    expectedCommands := []string{"exit", "help", "map", "mapb"}
    for _, name := range expectedCommands {
        if _, exists := commands[name]; !exists {
            t.Errorf("expected command %s to exist", name)
        }
    }
}

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "hello world",
            expected: []string{"hello", "world"},
        },
        {
            input:    "HELLO WORLD",
            expected: []string{"hello", "world"},
        },
        {
            input:    "   hello   world   ",
            expected: []string{"hello", "world"},
        },
    }

    for _, c := range cases {
        actual := CleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("length mismatch for input %q: got %d, want %d", 
                c.input, len(actual), len(c.expected))
            continue
        }
        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("mismatch at position %d for input %q: got %q, want %q", 
                    i, c.input, actual[i], c.expected[i])
            }
        }
    }
}
