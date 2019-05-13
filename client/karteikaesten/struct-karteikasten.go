package karteikaesten

import (
	"fmt"

	"github.com/Momper14/weblib/client/lernen"
)

// Karteikasten struct of a Karteikasten
type Karteikasten struct {
	ID             string        `json:"_id"`
	Rev            string        `json:"_rev"`
	Autor          string        `json:"Autor"`
	Kategorie      string        `json:"Kategorie"`
	Unterkategorie string        `json:"Unterkategorie"`
	Name           string        `json:"name"`
	Beschreibung   string        `json:"Beschreibung"`
	Public         bool          `json:"Public"`
	Karten         []Karteikarte `json:"Karten"`
	lerne          lernen.Lerne
}

// Karteikarte struct of a Karteikarte
type Karteikarte struct {
	Titel   string `json:"Titel"`
	Frage   string `json:"Frage"`
	Antwort string `json:"Antwort"`
}

// AnzahlKarten returns number of Karten
func (k Karteikasten) AnzahlKarten() int {
	return len(k.Karten)
}

// Fortschritt returns number of Karten
func (k Karteikasten) Fortschritt(userid string) (int, error) {
	faecher, err := k.KartenProFach(userid)
	if err != nil {
		return -1, err
	}

	var fortschritt int

	for i := 1; i <= 4; i++ {
		fortschritt += i * faecher[i]
	}
	fortschritt *= 100
	fortschritt /= 4 * k.AnzahlKarten()

	return fortschritt, nil
}

// KartenProFach returns an Array with number of Karten per Fach
func (k Karteikasten) KartenProFach(userid string) ([5]int, error) {
	var faecher [5]int

	lerne, err := k.getLerne(userid)
	if err != nil {
		return faecher, err
	}

	for _, v := range lerne.Karten {
		faecher[v]++
	}

	return faecher, nil
}

// FachVonKarte returns Fach of Karte
func (k Karteikasten) FachVonKarte(userid string, kartenindex int) (int, error) {
	lerne, err := k.getLerne(userid)
	if err != nil {
		return -1, err
	}

	if len(lerne.Karten) <= kartenindex {
		return -1, fmt.Errorf("Fehler: Karte %d für User %s in Kasten %s nicht gefunden", kartenindex, userid, k.ID)
	}
	return lerne.Karten[kartenindex], nil
}

func (k Karteikasten) getLerne(userid string) (lernen.Lerne, error) {
	var err error
	if k.lerne.ID != userid {
		k.lerne, err = lernen.New().LerneByUserAndKasten(userid, k.ID)
	}
	return k.lerne, err
}
