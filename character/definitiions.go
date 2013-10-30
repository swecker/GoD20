package character

import "fmt"


type Character struct {
	Info      map[string]string
	Level     int
	Abilities Abilities
	Modifiers Modifiers
	Stats     map[string]Stat
	Skills	  map[string]Skill
	Classes	  map[string][]Class
}

type Abilities struct {
	Strength     Ability
	Dexterity    Ability
	Constitution Ability
	Intelligence Ability
	Wisdom       Ability
	Charisma     Ability
}

type Ability struct {
	Score     int
	BaseScore int
}


//ArmorClass / Touch / FlatFooted
//Initiative
//Saving Throws
//CMB, CMD?
type Stat struct {
	Score     int
	Modifiers []*int
}

type Modifiers struct {
// Abilities
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
// ArmorClass
	Size         int
	NaturalArmor int
	Deflection   int
	MiscArmor    int
	Dodge        int
// Saving Throws
	Fortitude    int
	Reflex       int
	Will         int
// Combat
	BaseAttack   int
	Initiative   int
}

type HitPoints struct {
	Current int
	Base    int
	Damage  int
	Subdual int
}

func loadData() (Classes, Skills){
	classes := LoadClasses("data/classes.json")
	skills := LoadSkills("data/skills.json")
	return classes, skills
}

func NewCharacter() Character {
	fmt.Println("makin a chaaracture")
	classData, skillData := loadData()
	info      := make(map[string]string)
	stats     := make(map[string]Stat)
	skills    := skillData
	classes   := make(map[string][]Class)
	mods      := Modifiers{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	abilities := Abilities{
		Ability{10,10},
		Ability{10,10},
		Ability{10,10},
		Ability{10,10},
		Ability{10,10},
		Ability{10,10},
	}
	c := Character{info, 0, abilities, mods, stats, skills, classes}
	c = c.addLevel(classData, "Fighter")
	c = c.addLevel(classData, "Fighter")
	return c
}

func (c Character) addLevel ( data Classes, className string ) Character{
	currentClass, ok := c.Classes[className]
	if ok {
		level := len(currentClass)
		c.Classes[className] = append(currentClass, data[className][level])
	} else {
		newClass := []Class{data[className][0]}
		c.Classes[className] = newClass
	}

	c.Level = 0
	c.Modifiers.Will = 0
	c.Modifiers.Reflex = 0
	c.Modifiers.Fortitude = 0
	c.Modifiers.BaseAttack = 0
	for i := range c.Classes {
		for x := range c.Classes[i] {
			c.Level ++
			c.Modifiers.Will += c.Classes[i][x].WillBonus
			c.Modifiers.Reflex += c.Classes[i][x].ReflexBonus
			c.Modifiers.Fortitude += c.Classes[i][x].FortitudeBonus
			c.Modifiers.BaseAttack += c.Classes[i][x].BaseAttackBonus
		}
	}

	return c
}
