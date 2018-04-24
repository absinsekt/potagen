package main

import (
	"fmt"
	"github.com/absinsekt/potagen/core"
	"github.com/snksoft/crc"
	"strings"
)

var (
	PosCounter int = 1
	ITN        int = 770123456789
	PN         int = 123456789
)

func main() {
  fmt.Println()
	core.AnswerQuestionInt("Enter POS inventory number", &PosCounter)
	core.AnswerQuestionInt("Enter the ITN", &ITN)
	core.AnswerQuestionInt("Enter POS part number", &PN)

	src := strings.Join([]string{
		core.GetAlignedString(PosCounter, 10, "0"),
		core.GetAlignedString(ITN, 12, "0"),
		core.GetAlignedString(PN, 20, "0"),
	}, "")

	sum := crc.CalculateCRC(crc.CCITT, []byte(src))
	result := strings.Join([]string{
		core.GetAlignedString(PosCounter, 10, "0"),
		core.GetAlignedString(int(sum), 6, "0"),
	}, "")

	fmt.Printf("\nPOS registration number: %s\n\n", result)
}
