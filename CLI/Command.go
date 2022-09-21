package CLI

import "os"

type Command struct {
	Name       string
	Action     func()
	Flags      []Flag
	SubCommand []Command
}

func (c *Command) RegisterSubCommand(com Command) {
	c.SubCommand = append(c.SubCommand, com)
}

func (c *Command) GetFlagBool(name string) bool {
	// name can be option or optionshort make sure flag is of type flagbool
	for _, flag := range c.Flags {
		if flag.GetOption() == name || flag.GetOptionShort() == name {
			if flagBool, ok := flag.(*FlagBool); ok {
				return flagBool.Value
			}
		}
	}
	println("err:", name, "flag not found")
	os.Exit(0)
	return false
}

func (c *Command) GetFlagString(name string) string {
	// name can be option or optionshort make sure flag is of type flagstring
	for _, flag := range c.Flags {
		if flag.GetOption() == name || flag.GetOptionShort() == name {
			if flagString, ok := flag.(*FlagString); ok {
				return flagString.Value
			}
		}
	}
	println("err:", name, "flag not found")
	os.Exit(0)
	return ""
}

func (c *Command) GetFlagInt(name string) int {
	// name can be option or optionshort make sure flag is of type flagint
	for _, flag := range c.Flags {
		if flag.GetOption() == name || flag.GetOptionShort() == name {
			if flagInt, ok := flag.(*FlagInt); ok {
				return flagInt.Value
			}
		}
	}
	println("err:", name, "flag not found")
	os.Exit(0)
	return 0
}

func (c *Command) SetFlag(name string, value any) {
	//check if flag hold value type
	for _, flag := range c.Flags {
		if flag.GetOption() == name || flag.GetOptionShort() == name {
			switch flag.(type) {
			case *FlagBool:
				if _, ok := value.(bool); ok {
					flag.(*FlagBool).Value = value.(bool)
				} else {
					println("err:", name, "is a bool flag")
					os.Exit(0)
				}
			case *FlagString:
				if _, ok := value.(string); ok {
					flag.(*FlagString).Value = value.(string)
				} else {
					println("err:", name, "is a string flag")
					os.Exit(0)
				}
			case *FlagInt:
				if _, ok := value.(int); ok {
					flag.(*FlagInt).Value = value.(int)
				} else {
					println("err:", name, "is a int flag")
					os.Exit(0)
				}
			}
		}
	}
}
