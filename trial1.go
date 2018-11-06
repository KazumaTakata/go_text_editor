package main

import (
	"fmt"
	"io/ioutil"

	termbox "github.com/nsf/termbox-go"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getcharacterposition(userinput []rune, cursorPosition map[string]int) int {

	offsetx := 0
	offsety := 0

	for i, ch := range userinput {

		// if offsety == cursorPosition["y"] && offsetx == cursorPosition["x"] {
		// 	return i + 1
		// }

		if ch == '\n' {
			offsety += 1
			offsetx = -1
		}
		offsetx += 1

		if offsety == cursorPosition["y"] && offsetx == cursorPosition["x"] {
			return i + 1
		}

	}

	return 0
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	userinput := []rune{}
	cursorPosition := map[string]int{}
	cursorPosition["x"] = 0
	cursorPosition["y"] = 0

	defer termbox.Close()
	termbox.SetCell(1, 2, 'e', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(cursorPosition["x"], cursorPosition["y"])
	termbox.Flush()

mainloop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				err := ioutil.WriteFile("input.txt", []byte(string(userinput)), 0644)
				check(err)
				break mainloop
			case termbox.KeyArrowRight:
				cursorPosition["x"] += 1
				termbox.SetCursor(cursorPosition["x"], cursorPosition["y"])
			case termbox.KeyArrowLeft:
				cursorPosition["x"] -= 1
				termbox.SetCursor(cursorPosition["x"], cursorPosition["y"])
			case termbox.KeyArrowDown:
				cursorPosition["y"] += 1
				termbox.SetCursor(cursorPosition["x"], cursorPosition["y"])
			case termbox.KeyBackspace:

			case termbox.KeyEnter:
				insertIndex := getcharacterposition(userinput, cursorPosition)
				userinput = append(userinput[:insertIndex], append([]rune{'\n'}, userinput[insertIndex:]...)...)
				// userinput = append(userinput, '\n')
				cursorPosition["x"] += 1
				termbox.SetCursor(cursorPosition["x"], cursorPosition["y"])
			default:
				if ev.Ch != 0 {
					insertIndex := getcharacterposition(userinput, cursorPosition)
					// userinput = append(userinput, ev.Ch)
					userinput = append(userinput[:insertIndex], append([]rune{ev.Ch}, userinput[insertIndex:]...)...)
					cursorPosition["x"] += 1
					termbox.SetCursor(cursorPosition["x"], cursorPosition["y"])
				}

			}
		}
		offsetx := 0
		offsety := 0
		for _, ch := range userinput {
			termbox.SetCell(offsetx, offsety, ch, termbox.ColorDefault, termbox.ColorDefault)
			offsetx += 1
			fmt.Sprintf(string(ch))
			if ch == '\n' {
				offsety += 1
				offsetx = 0
			}
		}

		termbox.Flush()
	}
}
