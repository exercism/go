package piecingittogether

type PuzzleDetails struct {
	Pieces      int
	Border      int
	Inside      int
	Rows        int
	Columns     int
	AspectRatio float64
	Format      string
}

func JigsawData(details PuzzleDetails) (PuzzleDetails, error) {
	panic("Please implement the JigsawData function")
}
