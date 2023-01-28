package ch

const (
	InitialLogSize = 10
)

type Log struct {
	Owner       User
	Cards       []Card
	CurrentSize uint
}

func NewLog(user User) Log {
	return Log{
		Owner:       user,
		Cards:       make([]Card, InitialLogSize),
		CurrentSize: InitialLogSize,
	}
}
