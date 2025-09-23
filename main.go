package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	contacts := make(map[int]*Contact)
	id := 0

	ajouter := flag.Bool("ajouter", false, "Ajouter un contact")
	nom := flag.String("nom", "", "Nom du contact")
	email := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *ajouter {
		if *nom == "" || *email == "" {
			fmt.Println("Erreur: flags incorrects.")
			return
		}

		contacts[id] = &Contact{
			ID:    id,
			Nom:   *nom,
			Email: *email,
		}
		fmt.Printf("Contact ajoute ! ID: %v, Nom: %v, Email: %v\n", id, *nom, *email)
		id++
		return
	}

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
			c := &Contact{}
			c.ajouterContact(reader)
			contacts[id] = c
			fmt.Printf("Contact ajoute ! Id: %v, Nom: %v, Email: %v\n", id, c.Nom, c.Email)
			id++
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

func (c *Contact) ajouterContact(reader *bufio.Reader) {
	fmt.Print("Nom: ")
	nom, _ := reader.ReadString('\n')
	c.Nom = strings.TrimSpace(nom)

	fmt.Print("Mail: ")
	email, _ := reader.ReadString('\n')
	c.Email = strings.TrimSpace(email)
}

func afficherContacts(contacts map[int]*Contact) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact")
		return
	}

	fmt.Println("\nListe des contacts :")
	for id, c := range contacts {
		fmt.Printf("ID: %v, Nom: %v, Email: %v\n", id, c.Nom, c.Email)
	}
}

func supprimerContact(reader *bufio.Reader, contacts map[int]*Contact) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à supprimer")
		return
	}

	fmt.Print("ID à supprimer: ")
	idsup, _ := reader.ReadString('\n')
	idsup = strings.TrimSpace(idsup)

	id, err := strconv.Atoi(idsup)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if _, exists := contacts[id]; exists {
		delete(contacts, id)
		fmt.Println("Contact supprime !")
	} else {
		fmt.Println("ID inexistant")
	}
}

func mettreAJourContact(reader *bufio.Reader, contacts map[int]*Contact) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à mettre à jour")
		return
	}

	fmt.Print("ID à mettre à jour: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	c, exists := contacts[id]
	if !exists {
		fmt.Println("ID inexistant")
		return
	}

	fmt.Printf("Nom actuel: %s\n", c.Nom)
	fmt.Print("Nouveau nom: ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	if nom != "" {
		c.Nom = nom
	}

	fmt.Printf("Email actuel: %s\n", c.Email)
	fmt.Print("Nouvel email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		c.Email = email
	}

	fmt.Printf("Contact mis a jour ! ID: %v, Nom: %v, Email: %v\n", id, c.Nom, c.Email)
}
