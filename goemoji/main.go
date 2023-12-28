package main

import (
	"fmt"
	"strings"
)

func main() {
	emojis := make(map[string]string)
	emojis[":)"] = "ツ"

    originalString := "Original string :)"
    separated := strings.Split(originalString, " ")

    for i,w := range separated {
        if _, ok := emojis[w]; ok {
            fmt.Println(emojis[w])
            separated[i] = emojis[w]
        }
    }

    fmt.Println(originalString)
    fmt.Println(strings.Join(separated, " "))
}
