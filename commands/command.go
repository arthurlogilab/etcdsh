package commands

type Command interface {
	Supports(string) bool
	Handle([]string)
	Verify([]string) error
	CommandString() string
}
