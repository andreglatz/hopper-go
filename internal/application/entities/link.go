package entities

type Link struct {
	Id       uint
	Original string
	Short    string
	Clicks   uint
}

func NewLink(original string, short string) Link {
	return Link{
		Original: original,
		Short:    short,
		Clicks:   0,
	}
}
