package betting

// BetType represents the type of bet to make in the market.
type BetType uint

const (
	// BetType_Back represents a back bet.
	BetType_Back = iota + 1
	// BetType_Lay represents a lay bet.
	BetType_Lay
)

// String returns the string representation of BetType.
func (bt BetType) String() string {
	return [...]string{"", "Back", "Lay"}[bt]
}
