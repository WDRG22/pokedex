package commands


type CliCommand struct {
        Name            string
        Description     string
        Callback        func(*Config) error
}

var CommandRegistry map[string]CliCommand
