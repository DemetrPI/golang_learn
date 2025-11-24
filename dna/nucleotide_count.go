package dna

// Histogram represents the count of each nucleotide.
type Histogram map[rune]int

// DNA represents a sequence of nucleotides.
type DNA string

// InvalidNucleotideError represents an error for an invalid nucleotide.
type InvalidNucleotideError struct {
	Nucleotide rune
}

// validNucleotides defines the set of valid nucleotide characters.
var validNucleotides = map[rune]bool{
	'A': true,
	'C': true,
	'G': true,
	'T': true,
}

// isValidNucleotide checks if a nucleotide is valid.
func isValidNucleotide(nucleotide rune) bool {
	_, valid := validNucleotides[nucleotide]
	return valid
}

// Error implements the error interface for InvalidNucleotideError.
func (e *InvalidNucleotideError) Error() string {
	return "invalid nucleotide: " + string(e.Nucleotide)
}

// Counts is a method on the DNA type that returns a histogram of nucleotide counts.
func (d DNA) Counts() (Histogram, error) {
	h := make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	if len(d) == 0 {
		return h, nil
	}
	for _, nucleotide := range d {
		if !isValidNucleotide(nucleotide) {
			return nil, &InvalidNucleotideError{Nucleotide: nucleotide}
		}
		h[nucleotide]++
	}
	return h, nil
}
