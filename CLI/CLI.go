package CLI

import (
	"github.com/EVFUBS/commandgo/parser"
)

type Cli struct {
	Commands []Command
}

func New() Cli {
	cli := Cli{}
	return cli
}

func (cli *Cli) Run() {
	parser := parser.New()
	parsedCommand := parser.Parse()
	cli.RegisterCommand(helpCommand(cli.Commands))
	cli.executeCommand(parsedCommand, cli.Commands)
}

func (cli *Cli) executeCommand(parsedCommand []parser.ParseCommand, commands []Command) {
	for _, command := range commands {
		for _, pcommand := range parsedCommand {
			if pcommand.Command == command.Name {
				if pcommand.SubCommands != nil && command.SubCommand != nil {
					cli.executeCommand(pcommand.SubCommands, command.SubCommand)
				} else {
					// if parsed flag matches a command flag set the command flag value to the parsed value
					command = cli.assignFlagValue(pcommand.Flags, command)
					command.Action()
					break
				}
			}
		}
	}
}

func (cli *Cli) assignFlagValue(parsedFlags []parser.Flag, command Command) Command {
	for _, flag := range command.Flags {
		for _, pflag := range parsedFlags {
			if flag.GetOption() == pflag.Name || flag.GetOptionShort() == pflag.Name {
				command.SetFlag(flag.GetOption(), pflag.Value)
			}
		}
	}
	return command
}

func (cli *Cli) RegisterCommand(com Command) {
	cli.Commands = append(cli.Commands, com)
}
