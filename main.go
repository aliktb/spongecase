package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func toSpongeCase(text string) string {
  // Use strings.Builder for efficiency
  var result strings.Builder
  // Start with lowercase
  upper := false

  for _, char := range text {
    if unicode.IsLetter(char) {
      if upper {
        result.WriteRune(unicode.ToUpper(char))
      } else {
        result.WriteRune(unicode.ToLower(char))
      }
      // Toggle the case for the next letter
      upper = !upper
    } else {
      // Keep non-letter characters as they are
      result.WriteRune(char)
    }
  }

  // Convert strings.Builder to string
  return result.String()
}

func main() {
  // Join all arguments into a single string
  inputText := strings.Join(os.Args[1:], " ")

  if inputText == "" {
    fmt.Println("Please provide some text as input.")
    return
  }

  spongeCasedText := toSpongeCase(inputText)
  // Output the SpongeCase text
  fmt.Println(spongeCasedText)
}
