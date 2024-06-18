package rewrites

type MatchType string

func (m MatchType) String() string {
	return string(m)
}

const (
	MatchTypeStrict MatchType = "strict"
	MatchTypeSuffix MatchType = "suffix"
)
