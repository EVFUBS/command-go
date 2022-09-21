package CLI

type Flag interface {
	Get() any
	Set(value any)
	GetOption() string
	GetOptionShort() string
	GetDescription() string
}

type FlagBool struct {
	Option      string
	OptionShort string
	Description string
	Value       bool
}

func (f *FlagBool) Get() any {
	return f.Value
}

func (f *FlagBool) Set(value any) {
	f.Value = value.(bool)
}

func (f *FlagBool) GetOption() string {
	return f.Option
}

func (f *FlagBool) GetOptionShort() string {
	return f.OptionShort
}

func (f *FlagBool) GetDescription() string {
	return f.Description
}

type FlagString struct {
	Option      string
	OptionShort string
	Description string
	Value       string
}

func (f *FlagString) Get() any {
	return f.Value
}

func (f *FlagString) Set(value any) {
	f.Value = value.(string)
}

func (f *FlagString) GetOption() string {
	return f.Option
}

func (f *FlagString) GetOptionShort() string {
	return f.OptionShort
}

func (f *FlagString) GetDescription() string {
	return f.Description
}

type FlagInt struct {
	Option      string
	OptionShort string
	Description string
	Value       int
}

func (f *FlagInt) Get() any {
	return f.Value
}

func (f *FlagInt) Set(value any) {
	f.Value = value.(int)
}

func (f *FlagInt) GetOption() string {
	return f.Option
}

func (f *FlagInt) GetOptionShort() string {
	return f.OptionShort
}

func (f *FlagInt) GetDescription() string {
	return f.Description
}
