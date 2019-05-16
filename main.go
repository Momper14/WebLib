package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/kategorien"
	"github.com/Momper14/weblib/client/lernen"
	userspkg "github.com/Momper14/weblib/client/users"
)

func main() {
	{
		var kasten karteikaesten.Karteikasten
		karteikaesten := karteikaesten.New()
		{
			tmp, err := karteikaesten.KastenByID("6ea1e0c363e519c0b386b8d54c009d7d")
			kasten = tmp
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%+v\n", tmp)
			}
		}
		fmt.Println("")
		{
			tmp, err := karteikaesten.AnzahlOeffentlicherKaesten()
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Anzahl Kästen: ", tmp)
			}
		}
		fmt.Println("")
		{
			tmp, err := karteikaesten.AnzahlOeffentlicherKarten()
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Anzahl Karten: ", tmp)
			}
		}
		fmt.Println("")
		{
			tmp, err := karteikaesten.AnzahlKaestenUser("Max Mustermann")
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Anzahl Kästen von Max Mustermann: ", tmp)
			}
		}
		fmt.Println("")
		{
			tmp, err := karteikaesten.OeffentlicheKaestenByKategorie("Naturwissenschaften")
			if err != nil {
				log.Fatal(err)
			} else {
				for _, kasten := range tmp {
					fmt.Print(kasten, "\n\n")
				}
			}
		}
		fmt.Println("")
		{
			tmp, err := kasten.KartenProFach("Max Mustermann")
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Karten pro Fach: ", tmp)
			}
		}
		fmt.Println("")
		{
			tmp, err := kasten.Fortschritt("Max Mustermann")
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Fortschritt: ", tmp)
			}
		}
		fmt.Println("")
		{
			tmp, err := kasten.FachVonKarte("Max Mustermann", 3)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Fach von Karte: ", tmp)
			}
		}
	}
	fmt.Print("\n\n")
	{
		kategorien := kategorien.New()
		{
			tmp, err := kategorien.AlleKategorien()
			if err != nil {
				log.Fatal(err)
			} else {
				for _, kat := range tmp {
					fmt.Printf("%+v\n", kat)
				}
			}
		}
	}
	fmt.Print("\n\n")
	{
		users := userspkg.New()
		{
			tmp, err := users.AnzahlUsers()
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Anzahl User: ", tmp)
			}
		}
		fmt.Print("\n")
		{
			tmp, err := users.UserByID("Max Mustermann")
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%+v\n", tmp)
			}
		}
		fmt.Print("\n")
		{
			user := userspkg.User{
				Name:     "TestUser",
				Bild:     "",
				Email:    "test@email.com",
				Password: "123456",
				Seit:     time.Now().Unix(),
			}
			if err := users.UserErstellen(user); err != nil {
				fmt.Println(err)
			}
		}
		fmt.Print("\n")
		{
			if err := users.UserLoeschen("TestUser"); err != nil {
				fmt.Println(err)
			}
		}
		fmt.Print("\n")
		{
			if ok, err := users.CheckEmail("muster@example.com"); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ok)
			}

		}
		fmt.Print("\n")
		{
			if ok, err := users.CheckName("Max Mustermann"); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ok)
			}

		}
		fmt.Print("\n")
		{
			if ok, err := users.CheckName("Mustermann"); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ok)
			}

		}
	}
	fmt.Print("\n\n")
	{
		lernen := lernen.New()
		{
			tmp, err := lernen.LerneByUserAndKasten("Max Mustermann", "6ea1e0c363e519c0b386b8d54c00c942")
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%+v\n", tmp)
			}
		}
		fmt.Print("\n")
		{
			tmp, err := lernen.GelerntVonUser("Max Mustermann")
			if err != nil {
				log.Fatal(err)
			} else {
				for _, lern := range tmp {
					fmt.Printf("%+v\n", lern)
				}
			}
		}
	}
}
