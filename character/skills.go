package character

import "encoding/json"
import "io/ioutil"
//import "fmt"

type Skills map[string]Skill

type Skill struct {
	Name       string
	Ability    string
	Score      int
	ClassSkill bool
	Restricted bool
	Trainable  bool
	Ranks      int
	AbilityMod *int
	MiscBonus  int
}

func LoadSkills(file string) Skills {
	o, err := ioutil.ReadFile(file)
	if err != nil { panic(err) }
	var skills Skills
	err = json.Unmarshal(o, &skills)
	if err != nil { panic(err) }
	return  skills
}

