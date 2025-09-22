package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	contacts := make(map[int]map[string]string)
	id := 0

	for {
		fmt.Println("\nDifferentes actions:")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Afficher tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre a jour un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Choix: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			id = ajouterContact(reader, contacts, id)
		case "2":
			afficherContacts(contacts)
		case "3":
			supprimerContact(reader, contacts)
		case "4":
			mettreAJourContact(reader, contacts)
		case "5":
			fmt.Println("Au revoir")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func ajouterContact(reader *bufio.Reader, contacts map[int]map[string]string, id int) int {
	fmt.Print("Nom: ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)

	fmt.Print("Mail: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	contacts[id] = map[string]string{
		"nom":   nom,
		"email": email,
	}
	fmt.Printf("Contact ajoute ! Id: %v, Nom: %v, Email: %v\n", id, nom, email)
	return id + 1
}

func afficherContacts(contacts map[int]map[string]string) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact")
		return
	}

	fmt.Println("\nListe des contacts :")
	for id, c := range contacts {
		fmt.Printf("ID: %v, Nom: %v, Email: %v\n", id, c["nom"], c["email"])
	}
}

func supprimerContact(reader *bufio.Reader, contacts map[int]map[string]string) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact a supprimer")
		return
	}

	fmt.Print("ID a supprimer: ")
	idsup, _ := reader.ReadString('\n')
	idsup = strings.TrimSpace(idsup)

	id, err := strconv.Atoi(idsup)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	_, res := contacts[id]
	if res {
		delete(contacts, id)
		fmt.Println("Contact supprime !")
	} else {
		fmt.Println("ID inexistant")
	}

}

func mettreAJourContact(reader *bufio.Reader, contacts map[int]map[string]string) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact a mettre a jour")
		return
	}

	fmt.Print("ID a mettre à jour: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	contact, res := contacts[id]
	if res {
		fmt.Printf("Nom actuel: %s\n", contact["nom"])
		fmt.Print("Nouveau nom: ")
		nom, _ := reader.ReadString('\n')
		nom = strings.TrimSpace(nom)
		if nom != "" {
			contact["nom"] = nom
		}

		fmt.Printf("Email actuel: %s\n", contact["email"])
		fmt.Print("Nouvel email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if email != "" {
			contact["email"] = email
		}

		contacts[id] = contact
		fmt.Printf("Contact mis à jour ! ID: %v, Nom: %v, Email: %v\n", id, contact["nom"], contact["email"])
	} else {
		fmt.Println("ID inexistant")
	}
}
