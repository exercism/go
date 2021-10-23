package counter

type Counter interface {
	AddString(string)

	Lines() int
	Letters() int
	Characters() int
}
