package srgnt

type CommandFunction func()

type Command struct {
	Description string
	Callback    CommandFunction
}
