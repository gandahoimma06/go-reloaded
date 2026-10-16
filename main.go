package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func capitaliser(mot string) string {
	r := []rune(strings.ToLower(mot))
	if len(r) == 0 {
		return mot
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func traiterLigne(ligne string) string {
	mots := strings.Fields(ligne)

	for i := 0; i < len(mots); i++ {
		if mots[i] == "(hex)" || mots[i] == "(bin)" {
			if i > 0 {
				base := 16
				if mots[i] == "(bin)" {
					base = 2
				}
				n, err := strconv.ParseInt(mots[i-1], base, 64)
				if err == nil {
					mots[i-1] = strconv.FormatInt(n, 10)
					mots[i] = ""
				}
			}
			continue
		}

		for _, tag := range []string{"up", "low", "cap"} {
			if mots[i] == "("+tag+")" && i > 0 {
				switch tag {
				case "up":
					mots[i-1] = strings.ToUpper(mots[i-1])
				case "low":
					mots[i-1] = strings.ToLower(mots[i-1])
				case "cap":
					mots[i-1] = capitaliser(mots[i-1])
				}
				mots[i] = ""
				break
			}

			if strings.HasPrefix(mots[i], "("+tag+",") {
				nombre := strings.TrimSuffix(
					strings.TrimPrefix(mots[i], "("+tag+","),
					")",
				)

				if nombre == "" && i+1 < len(mots) {
					nombre = strings.TrimSuffix(mots[i+1], ")")
				}

				n, err := strconv.Atoi(strings.TrimSpace(nombre))
				if err == nil && n > 0 {
					debut := i - n
					if debut < 0 {
						debut = 0
					}

					for j := debut; j < i; j++ {
						switch tag {
						case "up":
							mots[j] = strings.ToUpper(mots[j])
						case "low":
							mots[j] = strings.ToLower(mots[j])
						case "cap":
							mots[j] = capitaliser(mots[j])
						}
					}

					mots[i] = ""
					if i+1 < len(mots) && strings.HasSuffix(mots[i+1], ")") {
						mots[i+1] = ""
						i++
					}
				}
				break

			}
		}
	}

	texte := strings.Join(strings.Fields(strings.Join(mots, " ")), " ")
	// Coller la ponctuation au mot précédent.
	ponctuation := regexp.MustCompile(`\s+([.,!?;:]+)`)
	texte = ponctuation.ReplaceAllString(texte, "$1")
	ponctuationSuivante := regexp.MustCompile(`([.,!?;:]+)([^\s.,!?;:])`)
	texte = ponctuationSuivante.ReplaceAllString(texte, "$1 $2")

	// Coller les apostrophes aux mots qu'elles entourent.
	texte = regexp.MustCompile(`'\s+`).ReplaceAllString(texte, "'")
	texte = regexp.MustCompile(`\s+'`).ReplaceAllString(texte, "'")

	// « a » devient « an » devant une voyelle ou h.
	an := regexp.MustCompile(`(?i)\ba ([aeiouh]\w*)`)
	texte = an.ReplaceAllStringFunc(texte, func(s string) string {
		if strings.HasPrefix(s, "A ") {
			return "An " + s[2:]
		}
		return "an " + s[2:]
	})

	return strings.TrimSpace(texte)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Utilisation : go run . fichier_entree fichier_sortie")
		os.Exit(1)
	}

	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de lecture :", err)
		os.Exit(1)
	}

	lignes := strings.Split(string(contenu), "\n")
	for i := range lignes {
		lignes[i] = traiterLigne(lignes[i])
	}
	resultat := strings.Join(lignes, "\n")

	if err := os.WriteFile(os.Args[2], []byte(resultat), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur d'écriture :", err)
		os.Exit(1)
	}

	fmt.Println("Fichier créé avec succès !")
}
