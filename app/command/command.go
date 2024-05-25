package command

type Result struct {
	output string
}

func newResult(output string) Result {
	return Result{output: output}
}

func (r Result) Output() string {
	return r.output
}

type Command interface {
	Execute() (Result, error)
}

type Ping struct {
}

func NewPingCommand() *Ping {
	return &Ping{}
}

func (p *Ping) Execute() (Result, error) {
	return newResult("PONG\n"), nil
}

type UnknownCommand struct {
}

func NewUnknownCommand() *UnknownCommand {
	return &UnknownCommand{}
}

func (uc *UnknownCommand) Execute() (Result, error) {
	return newResult("Unknown command. Please, consult help\n"), nil
}
