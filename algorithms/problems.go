package algorithms

func RunAlgorithms(n int) interface{} {
	switch n {
	case 1:
		return algo1("AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC")
	case 2:
		return algo2("GAUGGAACUUGACUACGUAAAUU")
	default:
		return ""
	}
}