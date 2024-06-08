package command

type CommandHandler interface {
	Handle(request CommandRequest)
}
