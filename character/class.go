package character

import "encoding/json"
import "io/ioutil"
//import "fmt"

type Classes map[string][]Class

type Class struct {
	Level           int
	BaseAttackBonus int
	FortitudeBonus  int
	ReflexBonus     int
	WillBonus       int
	//Special         string?
}

func LoadClasses(file string) Classes {
	o, err := ioutil.ReadFile(file)
	if err != nil { panic(err) }
	var classes Classes
	err = json.Unmarshal(o, &classes)
	if err != nil { panic(err) }
	return  classes
}



