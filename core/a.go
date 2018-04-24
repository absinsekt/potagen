package core

import (
  "fmt"
  "strconv"
  "strings"
)


func AnswerQuestionInt(question string, variable *int) {
  fmt.Printf("%s [%d]: ", question, *variable)
  fmt.Scanln(variable)
}

func GetAlignedString(srcNumber int, places int, placeholder string) string {
  result := strconv.Itoa(srcNumber)
  return strings.Repeat(placeholder, places-len(result)) + result
}
