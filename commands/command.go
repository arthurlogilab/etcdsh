package commands

type AutoCompleteConfig struct {
	Available bool
}

type Command interface {
	Supports(string) bool
	Handle([]string)
	Verify([]string) error
	CommandString() string
	GetAutoCompleteConfig() AutoCompleteConfig
}


type OneArgumentAutoCompleteCommand struct {
}

type NoAutoCompleteCommand struct {

}

func (o *OneArgumentAutoCompleteCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available:true}
}

func (o *NoAutoCompleteCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available:false}
}