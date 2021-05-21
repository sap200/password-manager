package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sap200/passwordManager/generator"
)

func main() {
	uppercase := flag.Bool("u", true, "Include Uppercase Letters A-Z in Password, default=true")
	lowercase := flag.Bool("l", true, "Include Lowercase Letters a-z in Password, default=true")
	digits := flag.Bool("d", true, "Include Digits 0-9 in Password, defaults=true")
	specialChars := flag.Bool("s", true, "Include Special Chars in "+generator.SpecialCharacters+"Password, default=true")
	tag := flag.String("tag", "UNSPECIFIED", "Tag for password, default=UNSPECIFIED")
	length := flag.Int("len", 14, "Length of password generated, default=14")
	list := flag.Bool("list", false, "List the passwords")

	flag.Parse()

	if *list {
		// perform
		fmt.Println("\n-----------------------------------------PASSWORDS-------------------------------------------")
		fmt.Println()
		dat, err := ioutil.ReadFile("data.txt")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(dat))
		return
	}

	g := generator.New(*lowercase, *uppercase, *digits, *specialChars, *tag, *length)
	password := g.Generate()

	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	pass := "{password: \"" + password + "\" Tag: \"" + g.Tag + "\"}"
	f.Write([]byte(pass))
	f.WriteString("\n\n")

	fmt.Println("Password generated successfully")
	fmt.Println("Password:", password, "\nTag:", g.Tag)
}
