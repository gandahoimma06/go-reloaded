package main

import "testing"

func TestCapitaliser(t *testing.T) {
	got := capitaliser("bonjour")
	want := "Bonjour"

	if got != want {
		t.Errorf("capitaliser() = %q, attendu %q", got, want)
	}
}

func TestTraiterLigne(t *testing.T) {
	tests := []struct {
		nom     string
		entree  string
		attendu string
	}{
		{
			nom:     "hexadecimal",
			entree:  "1E (hex) files",
			attendu: "30 files",
		},
		{
			nom:     "binaire",
			entree:  "10 (bin) years",
			attendu: "2 years",
		},
		{
			nom:     "majuscules",
			entree:  "go (up)",
			attendu: "GO",
		},
		{
			nom:     "minuscules",
			entree:  "SHOUTING (low)",
			attendu: "shouting",
		},
		{
			nom:     "capitalisation",
			entree:  "brooklyn bridge (cap)",
			attendu: "brooklyn Bridge",
		},
		{
			nom:     "majuscules sur plusieurs mots",
			entree:  "this is exciting (up, 2)",
			attendu: "this IS EXCITING",
		},

		{
			nom:     "up sur deux mots",
			entree:  "bonjour tout le monde (up, 2)",
			attendu: "bonjour tout LE MONDE",
		},
		{
			nom:     "low sur trois mots",
			entree:  "BONJOUR TOUT LE MONDE (low, 3)",
			attendu: "BONJOUR tout le monde",
		},
		{
			nom:     "cap sur trois mots",
			entree:  "welcome to the brooklyn bridge (cap, 3)",
			attendu: "welcome to The Brooklyn Bridge",
		},
		{
			nom:     "a devient an",
			entree:  "a amazing rock",
			attendu: "an amazing rock",
		},
		{
			nom:     "A devient An",
			entree:  "A amazing rock",
			attendu: "An amazing rock",
		},
		{
			nom:     "ponctuation",
			entree:  "hello ,world !",
			attendu: "hello, world!",
		},
		{
			nom:     "groupe de ponctuation",
			entree:  "thinking ... You",
			attendu: "thinking... You",
		},
		{
			nom:     "apostrophes",
			entree:  "' awesome '",
			attendu: "'awesome'",
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			got := traiterLigne(test.entree)

			if got != test.attendu {
				t.Errorf(
					"traiterLigne(%q) = %q, attendu %q",
					test.entree,
					got,
					test.attendu,
				)
			}
		})
	}
}

func TestCasLimites(t *testing.T) {
	tests := []struct {
		nom     string
		entree  string
		attendu string
	}{
		{
			nom:     "balise suivie de ponctuation",
			entree:  "times (up) , it",
			attendu: "TIMES, it",
		},
		{
			nom:     "capitalisation de plusieurs mots",
			entree:  "welcome to the brooklyn bridge (cap)",
			attendu: "welcome to the brooklyn Bridge",
		},
		{
			nom:     "six mots en majuscules",
			entree:  "one two three four five six (cap, 6)",
			attendu: "One Two Three Four Five Six",
		},
		{
			nom:     "apostrophes sur plusieurs mots",
			entree:  "' I am the most well-known person '",
			attendu: "'I am the most well-known person'",
		},
		{
			nom:     "ponctuation répétée",
			entree:  "hello !!",
			attendu: "hello!!",
		},
		{
			nom:     "mot commençant par h",
			entree:  "a honest person",
			attendu: "an honest person",
		},

		{
			nom:     "points de suspension",
			entree:  "Wait ... next",
			attendu: "Wait... next",
		},
		{
			nom:     "ponctuation combinee",
			entree:  "What !? really",
			attendu: "What!? really",
		},
		{
			nom:     "ponctuation sans espace avant",
			entree:  "Hello,world!",
			attendu: "Hello, world!",
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			got := traiterLigne(test.entree)
			if got != test.attendu {
				t.Errorf("obtenu %q, attendu %q", got, test.attendu)
			}
		})
	}
}
