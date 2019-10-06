package complex

var testDataForUnOp = []struct {
	r, i float64
}{
	// Imaginary unit
	{0, 1},
	// Complex with real parts only
	{1, 0},
	{-1, 0},
	{5, 0},
	{7, 0},
	// Complex with imaginary parts only
	{0, 3.14},
	{0, 5},
	{0, -25},
	// Complex with real and imaginary parts
	{1, 2},
	{15.7, 3.69},
	{4, 5},
	{-21.3, 45},
	{507.14, -616.628},
}

var testDataForBinOp = []struct {
	aReal, aImag float64
	bReal, bImag float64
}{
	{ // Complex with real parts only
		5, 0,
		7, 0,
	},
	{ // Complex with imaginary parts only
		0, 1,
		0, 5,
	},
	{ // One complex with real part only and the other one with imaginary part only
		2, 0,
		0, 5,
	},
	// Complex with real and imaginary parts
	{
		1, 3,
		4, -5,
	},
	{
		10.5, -5,
		4, 3.333,
	},
	{
		42, -1,
		8, 12,
	},
	{
		37, 71,
		250, 120,
	},
	{
		42.42, 1789.6,
		-756, 1492.1,
	},
}
