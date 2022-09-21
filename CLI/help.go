package CLI

import "fmt"

func helpCommand(commands []Command) Command {
	var help Command

	help.Name = "help"

	help.Action = func() {
		for _, command := range commands {
			printCommand("Command", command)
		}
	}

	return help
}

func printCommand(comType string, command Command) {
	fmt.Printf("%s: %s\n", comType, command.Name)
	for _, flag := range command.Flags {
		fmt.Printf("\tFlag: %s\n", flag.GetOption())
		fmt.Printf("\tType: %T\n", flag.Get())
		fmt.Printf("\tDescription: %s\n", flag.GetDescription())
		//print subcommands
		if command.SubCommand != nil {
			for _, subCommand := range command.SubCommand {
				printCommand("Subcommand", subCommand)
			}
		}
	}
}
