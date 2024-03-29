package entities

type Link struct {
	ID       uint
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

func (l *Link) AddClick() {
	l.Clicks++
}
