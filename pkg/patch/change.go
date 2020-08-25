package patch

type Change string

const (
	ChangeAdd    Change = "A"
	ChangeModify Change = "M"
	ChangeDelete Change = "D"
	ChangeDefault Change = ""
)

func (c Change) IsAddOrModify() bool {
	return c == ChangeAdd || c == ChangeModify || c == ChangeDefault
}
