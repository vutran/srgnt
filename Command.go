package srgnt

type CommandFunction func()

type Command struct {
	Callback CommandFunction
}
