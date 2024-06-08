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

type AvailableCommandHandler int

const (
	pingCmdHandler AvailableCommandHandler = iota + 1
	commandCmdHandler
)

type CommandRequest interface {
	Type() string
}

type Command interface {
	Execute(request CommandRequest) (Result, error)
}

var AvailableCommandMap = map[AvailableCommandHandler]Command{
	pingCmdHandler:    NewPingCommand(),
	commandCmdHandler: NewCommandsCommand(),
}

type Ping struct {
}

func NewPingCommand() *Ping {
	return &Ping{}
}

func (p *Ping) Execute(request CommandRequest) (Result, error) {
	return newResult("+PONG\r\n"), nil
}

type CommandsCommand struct {
}

func NewCommandsCommand() *CommandsCommand {
	return &CommandsCommand{}
}

func (cc *CommandsCommand) Execute(request CommandRequest) (Result, error) {
	return newResult("*1\r\n*6\r\n$4\r\nping\r\n:1\r\n*1\r\n+readonly\r\n:0\r\n:0\r\n:0\r\n"), nil
}
