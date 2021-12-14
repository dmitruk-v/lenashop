package greet

type DAO struct {
	DefaultMessage string
	BobMessage     string
	JuliaMessage   string
}

func (sdi DAO) GreetingForName(name string) (string, error) {
	switch name {
	case "Bob":
		return sdi.BobMessage, nil
	case "Julia":
		return sdi.JuliaMessage, nil
	default:
		return sdi.DefaultMessage, nil
	}
}
