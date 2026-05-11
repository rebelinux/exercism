package strand

import "strings"

type RNA map[string]string

func ToRNA(dna string) string {
	var str []string
	for _, v := range dna {
		RNAMap := RNA{
			"G": "C",
			"C": "G",
			"T": "A",
			"A": "U",
		}
		str = append(str, RNAMap[string(v)])
	}
	return strings.Join(str, "")
}
