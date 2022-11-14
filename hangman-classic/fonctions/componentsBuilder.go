package hangman

import (
	"github.com/nsf/termbox-go"
)

func CreateBox(xa, ya, xb, yb int, color termbox.Attribute){
	for i := xa; i <= xb; i++{
		termbox.SetCell(i, ya, '-', color, termbox.RGBToAttribute(0,0,0))
	}
	for j := ya+1; j < yb; j++{
		termbox.SetCell(xa, j, '|', color, termbox.RGBToAttribute(0,0,0))
		termbox.SetCell(xb, j, '|', color, termbox.RGBToAttribute(0,0,0))
	}
	for o := xa; o <= xb; o++{
		termbox.SetCell(o, yb, '-', color, termbox.RGBToAttribute(0,0,0))
	}
	termbox.Flush()
}

func WriteText(texte string, xa, ya int, color termbox.Attribute){
	for position, lettre := range texte {
		termbox.SetCell(xa+position, ya, lettre, color, termbox.RGBToAttribute(0,0,0))
	}
	termbox.Flush()
}

func WriteWord(word string, xa, ya int, color termbox.Attribute){
	for position, lettre := range word {
		termbox.SetCell(xa+position+position, ya, lettre, color, termbox.RGBToAttribute(0,0,0))
		termbox.SetCell(xa+position+position+1, ya, ' ', color, termbox.RGBToAttribute(0,0,0))
	}
	termbox.Flush()
}
