package strand

import "fmt"

// ToRNA transcribes a DNA strand into an RNA strand.
func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	for i, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		default:
			fmt.Printf("Invalid nucleotide: %c\n", nucleotide)
		}
	}
	return string(rna)
}