package betting

type BetType uint

const (
	BetType_Back = iota
	BetType_Lay
)

func (bt BetType) String() string {
	return [...]string{"Back", "Lay"}[bt]
}
