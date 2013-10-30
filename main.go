package main

import "fmt"
import "encoding/json"
import "github.com/swecker/d20/character"

func main() {
	sam := character.NewCharacter()
	b, err := json.Marshal(sam)
	if err != nil { panic(err) }
	fmt.Printf("----------------\n%s\n----------------\n", b)
}
