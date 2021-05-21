package generator

import (
	"strings"

	"github.com/sap200/passwordManager/sampler"
)

type Generator struct {
	LowerCase         bool
	UpperCase         bool
	Digit             bool
	SpecialCharacters bool
	Tag               string
	Length            int
}

const LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"

var UppercaseLetters = strings.ToUpper(LowercaseLetters)

var Letters = LowercaseLetters + UppercaseLetters

const Digits = "0123456789"

const SpecialCharacters = "!@#$%^&*()-+_=|;:?"

var allInOne = Letters + Digits + SpecialCharacters

func New(lowerCase, upperCase, digit, specialCharacters bool, tag string, length int) Generator {
	return Generator{
		LowerCase:         lowerCase,
		UpperCase:         upperCase,
		Digit:             digit,
		SpecialCharacters: specialCharacters,
		Tag:               tag,
		Length:            length,
	}
}

func (g Generator) Generate() string {
	password := ""

	if g.LowerCase {
		password += sampler.Sample(LowercaseLetters, len(LowercaseLetters))
	}

	if g.UpperCase {
		password += sampler.Sample(UppercaseLetters, len(UppercaseLetters))
	}

	if g.Digit {
		password += sampler.Sample(Digits, len(Digits))
	}

	if g.SpecialCharacters {
		password += sampler.Sample(SpecialCharacters, len(SpecialCharacters))
	}

	if len(password) < g.Length {
		totalCharacters := g.Length - len(password)
		for i := 0; i < totalCharacters; i++ {
			password += sampler.Sample(allInOne, totalCharacters)
		}
	}

	return sampler.Shuffle(password, g.Length)
}
