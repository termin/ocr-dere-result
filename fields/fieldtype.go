package fields

import "fmt"

type FieldName string
type FieldType string

const (
	Lv         = FieldName("lv")
	Title      = FieldName("title")
	Difficulty = FieldName("difficulty")
	Perfect    = FieldName("perfect")
	Great      = FieldName("great")
	Nice       = FieldName("nice")
	Bad        = FieldName("bad")
	Miss       = FieldName("miss")
	Combo      = FieldName("combo")
	Score      = FieldName("score")

	TypeString = FieldType("string")
	TypeDigits = FieldType("digits")
)

// TODO: stringじゃなくてもっと格好良くしたい
func FieldTypeByName(name FieldName) (FieldType, error) {
	switch name {
	case Title:
		return TypeString, nil
	case Lv, Difficulty, Perfect, Great, Nice, Bad, Miss, Combo, Score:
		return TypeDigits, nil
	}

	return "", fmt.Errorf("TODO Error Description")
}