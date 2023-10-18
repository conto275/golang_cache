package glowne

import (
	"fmt"
)

type Alrandom struct {
	name   string
	repeat string
}

func NewAlrandom(n string, r string) Alrandom {
	return Alrandom{
		name:   n,
		repeat: r,
	}
}
func (al Alrandom) Alrandomstart() (string, bool) { // возможно придется добавить блокировку
	switch al.name {
	case "Alesya":
		switch al.repeat {
		case "569884":
			return fmt.Sprintf("Dobrze, imie %s jest, haslo %s jest. Acces allowed", al.name, al.repeat), true
		default:
			return fmt.Sprintf("Stop, imie %s i hasla %s nie są w spisku. Nie ma prawa dostępa", al.name, al.repeat), false
		} ///////
	case "Vlad":
		switch al.repeat {
		case "62291881":
			return fmt.Sprintf("Dobrze, imie %s jest, haslo %s jest. Acces allowed", al.name, al.repeat), true
		default:
			return fmt.Sprintf("Stop, imie %s i hasla %s nie są w spisku. Nie ma prawa dostępa", al.name, al.repeat), false
		} ////
	case "Nika":
		switch al.repeat {
		case "VDSK56g":
			return fmt.Sprintf("Dobrze, imie %s jest, haslo %s jest. Acces allowed", al.name, al.repeat), true
		default:
			return fmt.Sprintf("Stop, imie %s i hasla %s nie są w spisku. Nie ma prawa dostępa", al.name, al.repeat), false
		} ////
	case "Maxim":
		switch al.repeat {
		case "7ft7ghfhr97g":
			return fmt.Sprintf("Dobrze, imie %s jest, haslo %s jest. Acces allowed", al.name, al.repeat), true
		default:
			return fmt.Sprintf("Stop, imie %s i hasla %s nie są w spisku. Nie ma prawa dostępa", al.name, al.repeat), false
		} ////
	default:
		return fmt.Sprintf("Stop, imie %s i hasla %s nie są w spisku. Nie ma prawa dostępa", al.name, al.repeat), false

	}

}

/*"Alesya", "Vlad", "Nika", "Maxim")
"1", "2", "3", "4", "5", "8", "7", "8", "9", "0")*/
