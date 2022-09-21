package parser

import (
	"os"
	"strconv"
	"strings"
)

type Arg interface {
	String() string
}

type ParseCommand struct {
	Command     string
	Flags       []Flag
	SubCommands []ParseCommand
}

func (c *ParseCommand) String() string {
	var output string

	output += "ParseCommand("
	output += c.Command
	for _, flag := range c.Flags {
		output += flag.String()
	}
	for _, command := range c.SubCommands {
		output += command.String()
	}
	output += ")"

	return output
}

type Flag struct {
	Name  string
	Value any
}

func (f *Flag) String() string {
	return f.Value.(string)
}

type Parser struct {
	Args    []string
	curArg  int
	peekArg int
}

func New() *Parser {
	parser := &Parser{
		Args:    os.Args[1:],
		curArg:  0,
		peekArg: 1,
	}
	return parser
}

func (p *Parser) AdvanceArgs() {
	p.curArg += 1
	p.peekArg += 1
}

func (p *Parser) currentArg() string {
	return p.Args[p.curArg]
}

func (p *Parser) Parse() []ParseCommand {
	var parsedCommand []ParseCommand
	parsedCommand = append(parsedCommand, p.ParseCommand())
	return parsedCommand
}

func (p *Parser) ParseCommand() ParseCommand {
	var command ParseCommand
	if !strings.Contains(p.currentArg(), "--") {
		command.Command = p.currentArg()
		p.AdvanceArgs()

		for p.curArg < len(p.Args) {
			if strings.Contains(p.currentArg(), "--") {
				command.Flags = p.ParseFlags()
			} else {
				break
			}
		}

		for p.curArg < len(p.Args) {
			if !strings.Contains(p.currentArg(), "--") {
				command.SubCommands = append(command.SubCommands, p.ParseCommand())
			} else {
				break
			}
		}
	}
	return command
}

func (p *Parser) ParseFlags() []Flag {
	var flags []Flag
	flags = append(flags, p.ParseFlag())
	p.AdvanceArgs()
	for p.curArg < len(p.Args) {
		if strings.Contains(p.currentArg(), "--") {
			flags = append(flags, p.ParseFlag())
			p.AdvanceArgs()
		} else {
			break
		}
	}
	return flags
}

func (p *Parser) ParseFlag() Flag {
	arg := strings.Replace(p.currentArg(), "--", "", 1)

	//regex for quote with string between
	if strings.Contains(arg, "=") {
		// split string between =
		split := strings.Split(arg, "=")

		// covert split[1] to bool
		if split[1] == "true" {
			return Flag{
				Name:  split[0],
				Value: true,
			}
		} else if split[1] == "false" {
			return Flag{
				Name:  split[0],
				Value: false,
			}
		}

		// convert split[1] to int
		if i, err := strconv.Atoi(split[1]); err == nil {
			return Flag{
				Name:  split[0],
				Value: i,
			}
		}

		return Flag{
			Name:  split[0],
			Value: split[1],
		}
	} else {
		return Flag{
			Name:  arg,
			Value: true,
		}
	}

}
