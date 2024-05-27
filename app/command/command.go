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
	return newResult("+PONG\r\n"), nil
}

type CommandsCommand struct {
}

func NewCommandsCommand() *CommandsCommand {
	return &CommandsCommand{}
}

func (cc *CommandsCommand) Execute() (Result, error) {
	return newResult("*1\r\n*6\r\n$4\r\nping\r\n:1\r\n*1\r\n+readonly\r\n:0\r\n:0\r\n:0\r\n"), nil
}
