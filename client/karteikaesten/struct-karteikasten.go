package karteikaesten

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Momper14/weblib/client"
	"github.com/Momper14/weblib/client/lernen"
)

// Karteikasten struct of a Karteikasten
type Karteikasten struct {
	ID             string        `json:"_id,omitempty"`
	Rev            string        `json:"_rev,omitempty"`
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

// KarteHinzufuegen fügt dem Kasten die neue Karte hinzu
func (k Karteikasten) KarteHinzufuegen(karte Karteikarte) error {
	lernen := lernen.New()

	k.Karten = append(k.Karten, karte)

	lArr, err := lernen.AlleLerneZuKasten(k.ID)
	if err != nil {
		return err
	}

	for i := 0; i < len(lArr); i++ {
		lArr[i].Karten = append(lArr[i].Karten, 0)
	}

	return lernen.AktualisiereAlleLerne(lArr)
}

// KarteAktualisieren ersetzt die Karte am gegebenen Index
func (k Karteikasten) KarteAktualisieren(index int, karte Karteikarte) error {
	lernen := lernen.New()

	if index < 0 || index >= len(k.Karten) {
		return client.IndexOutOfRangeError{Msg: fmt.Sprintf("%d ist keine gültige Karte", index)}
	}

	k.Karten[index] = karte

	lArr, err := lernen.AlleLerneZuKasten(k.ID)
	if err != nil {
		return err
	}

	for i := 0; i < len(lArr); i++ {
		lArr[i].Karten[index] = 0
	}

	return lernen.AktualisiereAlleLerne(lArr)
}

// KarteLoeschen löscht die Karte am gegebenen Index
func (k Karteikasten) KarteLoeschen(index int) error {
	lernen := lernen.New()

	if index < 0 || index >= len(k.Karten) {
		return client.IndexOutOfRangeError{Msg: fmt.Sprintf("%d ist keine gültige Karte", index)}
	}

	k.Karten = append(k.Karten[:index], k.Karten[index+1:]...)

	lArr, err := lernen.AlleLerneZuKasten(k.ID)
	if err != nil {
		return err
	}

	for i := 0; i < len(lArr); i++ {
		lArr[i].Karten = append(lArr[i].Karten[:index], lArr[i].Karten[index+1:]...)
	}

	return lernen.AktualisiereAlleLerne(lArr)
}

// AnzahlKarten gibt die Anzahl von Karten zurück
func (k Karteikasten) AnzahlKarten() int {
	return len(k.Karten)
}

// Fortschritt gibt den Fortschritt des angegebenen Users für diesen Karteikasten zurück
func (k Karteikasten) Fortschritt(userid string) (int, error) {
	if k.AnzahlKarten() == 0 {
		return 0, nil
	}

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

// KartenProFach gibt einen Array mit der Anzahl an Karten pro Fach zurück
// Index = Fach
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

// FachVonKarte gibt das Fach der angegebenen Karte für den angegebenen User zurück
func (k Karteikasten) FachVonKarte(userid string, kartenindex int) (int, error) {
	lerne, err := k.getLerne(userid)
	if err != nil {
		return -1, err
	}

	if len(lerne.Karten) <= kartenindex {
		return -1, client.NotFoundError{Msg: fmt.Sprintf("Error: Karte %d für User %s in Kasten %s nicht gefunden", kartenindex, userid, k.ID)}
	}
	return lerne.Karten[kartenindex], nil
}

// getLerne gibt den Lernstand des angegebenen Users für diesen Kasten zurück
// speichert diesen zwichen für wiederholte Anfragen
func (k Karteikasten) getLerne(userid string) (lernen.Lerne, error) {
	var err error
	if k.lerne.User != userid {
		k.lerne, err = lernen.New().LerneByUserAndKasten(userid, k.ID)
	}

	if _, ok := err.(client.NotFoundError); ok {
		if k.Public {
			lerne := lernen.Lerne{
				User:   userid,
				Kasten: k.ID,
				Karten: make([]int, k.AnzahlKarten()),
			}
			return lerne, lernen.New().NeuesLerne(lerne)
		}
		return lernen.Lerne{}, client.ForbiddenError{Msg: fmt.Sprintf("User %s is not allowed to lern %s", userid, k.ID)}
	}

	return k.lerne, err
}

// Zufallskarte gibt eine zufällige Karte zurück
func (k Karteikasten) Zufallskarte(userid string) (int, Karteikarte, error) {

	if k.AnzahlKarten() == 0 {
		return -1, Karteikarte{}, client.NotFoundError{Msg: "Keine Karte im Kasten vorhanden!"}
	}

	var faecher [5][]int
	lerne, err := k.getLerne(userid)
	if err != nil {
		return 0, Karteikarte{}, err
	}

	for i, v := range lerne.Karten {
		faecher[v] = append(faecher[v], i)
	}

	for {
		fach := zufallsfach()
		if len(faecher[fach]) == 0 {
			continue
		}
		index := randomInt(len(faecher[fach]))
		index = faecher[fach][index]
		return index, k.Karten[index], nil
	}

}

// zufallsfach bestimmt ein zufälliges fach
// Algorithmus wie vorgegeben
func zufallsfach() int {
	switch randomInt(15) {
	case 0:
		return 4
	case 1, 2:
		return 3
	case 3, 4, 5:
		return 2
	case 6, 7, 8, 9:
		return 1
	case 10, 11, 12, 13, 14:
		return 0
	default:
		return -1
	}
}

// randomInt kapselung von rand.Intn()
func randomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
