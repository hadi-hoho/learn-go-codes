package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	// when i wrote this code only me and god understood it, now only god understands it
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		var outline string
		act := string(line[0])
		ty := string(line[2])
		if act == "S" {
			str := line[4:]

			var word string
			var is_first_letter, is_first_word bool
			is_first_letter = true
			if ty == "C" {
				is_first_word = true
			}
			for _, ch := range str {
				sch := string(ch)
				if strings.ToLower(sch) != sch {
					if !is_first_letter {
						if is_first_word {
							outline = word
							is_first_word = false
						} else {
							outline = outline + " " + word
						}
					} else {
						is_first_letter = false
						if ty != "C" {
							outline = word
						}
					}
					word = strings.ToLower(sch)
				} else if sch == "(" {
					break
				} else {
					word = word + sch
				}
			}
			if len(outline) == 0 {
				outline = word
			} else {
				outline = outline + " " + word
			}
			fmt.Println(outline)
		} else if act == "C" {
			str := line[4:] + " "
			words := strings.Fields(str)
			var outline string
			if ty == "C" {
				for i := 0; i < len(words); i++ {
					outline = outline + strings.ToUpper(string(words[i][0])) + words[i][1:]
				}
			} else {
				outline = words[0]
				for i := 1; i < len(words); i++ {
					outline = outline + strings.ToUpper(string(words[i][0])) + words[i][1:]
				}
				if ty == "M" {
					outline = outline + "()"
				}
			}
			fmt.Println(outline)
		}
	}
}
