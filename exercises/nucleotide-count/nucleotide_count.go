package dna

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	return 0, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	return h, nil
}
