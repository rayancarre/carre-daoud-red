//CONCERNANT MON CODE J AI PREFERE TOUT FAIRE SUR UN SEUL FICHIER,J'AI EXPLIQUE CHAQUE FONCTION A QUOI ELLE SERVAIT,
//TOUT EST TRES BIEN DETAILLE.
//POUR RETROUVER TOUTES LES FONCTIONS VOICI VOUS AURIEZ JUSTE A (CONTROLE F) TOUT EST BIEN GUIDE ET BIEN EXPLIQUE.

// ETANT DONNE QUE NOUS AVONS TRAVAILLE SUR UN SEUL FICHIER EN MEME TEMPS JE VAIS VOUS EXPLIQUER QUI C'EST OCCUPE DE QUEL TACHE
//ALORS DEJA LES TACHES DE 1 A 6 BIEN ENTENDU NOUS L AVIONS FAIT SEPARREMENT SUITE A CELA MON CAMARADE EVAN DAOUD
// S'EST OCCUPE DES TACHES SUIVANTES :

// - TACHE 8 : WASTED !
// - TACHE 14 : I SAW IN THE MIRROR
// - MISSION 3 : INITIATUVE
// - TACHE 16 : LA CHOSE

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// ////////////////////////////////////////////////////////////////////////////////////////////PLAYER/////////////////////////////////////////////////////////////////////////////////
type Player struct {
	name               string
	class              string
	level              int
	lifemax            int
	life               int
	inventory          map[string]int
	money              int
	skill              []string
	equipment          Equipment
	inventoryUpgrades  int
	initiative         int
	ExperienceActuelle int
	ExperienceMaximale int
	Niveau             int
}

///////////////////////////////////////////////////////////////////////////////////////////////Equipement//////////////////////////////////////////////////////////////////////////

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

// /////////////////////////////////////////////////////////////MONSTRE////////////////////////////////////////////////////////////////////////
type Monster struct {
	nom        string
	attaque    int
	pvMax      int
	pvActuels  int
	initiative int
}

// ////////////////////////////////////////////////////////////tous les perso qui exisent
var p1 Player
var p2 Player
var p3 Player
var p4 Player

///////////////////////////////////////////////////////FONCTION MAIN///////////////////////////////////////////////////

func main() {
	RaconterHistoire()
	StartScreen()
	// Crée un gobelin d'entraînement en utilisant la fonction InitGoblin
	gobelin := InitGoblin()

	// Affiche les caractéristiques du gobelin
	fmt.Printf("Un %s apparaît !\n", gobelin.nom)
	fmt.Printf("Points de vie : %d/%d\n", gobelin.pvActuels, gobelin.pvMax)
	fmt.Printf("Points d'attaque : %d\n", gobelin.attaque)

	map1 := map[string]int{"Potions de vie": 3}
	p1.Init("Rayan", "Humain", 1, 100, 40, map1, 100, "Coup de poing", p1.equipment)

	p2.Init("Evan", "Elfe", 1, 80, 50, map1, 100, "Coup de poing", p1.equipment)

	p3.Init("Olivier", "Nain", 1, 120, 50, map1, 100, "Coup de poing", p1.equipment)

	StartScreen()
}

////////////////////////////////////////////EN PREMIER NOUS ALLONS RACONTER L HISTOIRE//////////////////////////////////////////////////////////

func RaconterHistoire() {
	ClearConsole()

	SpeedMsg("\t   ____             U _____ u_   _  __     __ U _____ u_   _       _   _U _____ u    \n\tU | __\")u    ___    \\| ___\"|/ \\ |\"| \\ \\   /\"/u\\| ___\"|/ \\ |\"|   U |\"|u| \\| ___\"|/    \n\t \\|  _ \\/   |_\"_|    |  _|\"<|  \\| |> \\ \\ / //  |  _|\"<|  \\| |>   \\| |\\| ||  _|\"      \n\t  | |_) |    | |     | |___U| |\\  |u /\\ V /_,-.| |___U| |\\  |u    | |_| || |___      \n\t  |____/   U/| |\\u   |_____||_| \\_| U  \\_/-(_/ |_____||_| \\_|    <<\\___/ |_____|     \n\t _|| \\_.-,_|___|_,-.<<   >>||   \\,-.//       <<   >>||   \\,-.(__) )(  <<   >>  \n\t(__) (__)\\_)-' '-(_/(__) (__|_\"  (/_/(__)     (__) (__|_\")  (_/     (__)(__) (__)    \n", 10, "cyan")
	fmt.Println("                                                           -                                                        ")
	fmt.Println("                                                           -                                                        ")
	fmt.Println("                                                           -                                                        ")
	SpeedMsg("----------Il y a longtemps, dans un royaume lointain, le maléfique dragon Smaug a volé l'Orbe de la Destinée----------\n", 20, "green")
	SpeedMsg("----------une puissante relique capable de contrôler la magie elle-même. Le royaume est plongé dans les ténèbres\n", 20, "green")
	SpeedMsg("----------et le désespoir règne parmi le peuple.\n", 20, "red")
	SpeedMsg("\n\t       _______________ ______________________ ________________________ \n\t /  _  \\   \\ /   /\\_   _____/ \\      \\__    ___/    |   \\______   \\_   _____/ \n\t/  /_\\  \\   Y   /  |    __)_  /   |   \\|    |  |    |   /|       _/|    __)_\n\t/    |    \\     /   |        \\/    |    \\    |  |    |  / |    |   \\|        \\ \n\t\\____|__  /\\___/   /_______  /\\____|__  /____|  |______/  |____|_  /_______  / \n\t\\/                 \\/         \\/                         \\/        \\/           ", 10, "cyan")
	fmt.Println()
	SpeedMsg("----------Vous êtes un jeune aventurier courageux, prêt à tout pour sauver votre royaume. Vous avez entendu parler----------\n", 20, "green")
	SpeedMsg("----------d'une prophétie ancienne qui dit que seul un héros digne pourra reprendre l'Orbe de la Destinée et----------\n", 20, "green")
	SpeedMsg("----------défaire Smaug. C'est votre destinée qui vous appelle, et vous partez pour une aventure épique.----------\n", 20, "green")
	fmt.Println()
	SpeedMsg("----------Votre voyage commence ici. Choisissez avec sagesse votre personnage et préparez-vous à affronter des----------\n", 20, "green")
	SpeedMsg("----------dangers inimaginables. L'avenir du royaume repose entre vos mains.----------\n", 20, "green")
	fmt.Println()
	SpeedMsg("-------------------------------( Appuyez sur Entrée pour commencer votre quête...) -------------------------------\n", 20, "blue")
	fmt.Scanln() // Attendez l'entrée de l'utilisateur pour passer au jeu.
}

/////////////////////////////////////////////////////////////////////////POUR CELA NOUS ALLONS DONC CREER UNE FONCTION QUI VA PERMETTRE DE SCROLL////////////////////////

// ///////////////////////                     ///////////////////////////////////////////////SCROLL LE TEXTE/////////////////////////////////////////////////////////////////
func SpeedMsg(message string, speed int, colorName string) {
	defaultColor := color.New(color.FgWhite)
	Green := color.New(color.FgGreen)
	Red := color.New(color.FgRed)
	Blue := color.New(color.FgBlue)
	Cyan := color.New(color.FgCyan)
	Yellow := color.New(color.FgYellow)

	var selectedColor *color.Color

	switch colorName {
	case "green":
		selectedColor = Green
	case "red":
		selectedColor = Red
	case "blue":
		selectedColor = Blue
	case "cyan":
		selectedColor = Cyan
	case "yellow":
		selectedColor = Yellow
	default:
		selectedColor = defaultColor
	}

	for _, char := range message {
		selectedColor.Print(string(char))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

/////////////////////////////////////////////////////////////////////NOUS ALLONS RAJOUTER LA FONCTION CLEARCONSOLE POUR BIEN AFFICHER SUR LE TERMINAL LE JEU APRES CHAQUE FONCTION

func ClearConsole() {
	switch runtime.GOOS {
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Système d'exploitation non pris en charge. La console ne peut pas être effacée.")
	}
}

// ////////////////////////////////////////////////////////////////////////////AFFICHER L ECRAN INITIAL/////////////////////////////////////////////////////////////////////////////////////////////////
func StartScreen() {
	ClearConsole()
	fmt.Println("Bienvenue dans le jeu !")
	fmt.Println("Options :")
	color.Green("1. Commencer la partie")
	color.Red("2. Quitter la partie")
	fmt.Print("Choisissez une option : ")

	var choix int
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	switch choix {
	case 1:
		CreateCharacter() // Lancer la création du personnage
		Menu()            // Après la création du personnage, afficher le menu principal
	case 2:
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide.")
		StartScreen()
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////Choisir personnage////////////////////////////////////////////////////////////////

func CreateCharacter() {
	ClearConsole()
	fmt.Println("Création de personnage :")

	var nom string
	fmt.Print("Choisissez le nom de votre personnage : ")
	_, err := fmt.Scan(&nom)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	var classe int
	fmt.Println("Choisissez la classe de votre personnage :")
	fmt.Println("1. Humain")
	fmt.Println("2. Elfe")
	fmt.Println("3. Nain")

	fmt.Print("Choisissez une classe (1/2/3) : ")
	_, err = fmt.Scan(&classe)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	var newCharacter Player
	switch classe {
	case 1:
		newCharacter.Init(nom, "Humain", 1, 100, 40, map[string]int{"potions": 1}, 100, "Coup de poing", Equipment{})
	case 2:
		newCharacter.Init(nom, "Elfe", 1, 80, 50, map[string]int{"potions": 1}, 100, "Coup de poing", Equipment{})
	case 3:
		newCharacter.Init(nom, "Nain", 1, 120, 50, map[string]int{"potions": 1}, 100, "Coup de poing", Equipment{})
	default:
		fmt.Println("Classe invalide.")
		CreateCharacter()
		return
	}

	if p1.name == "" {
		p1 = newCharacter
	} else if p2.name == "" {
		p2 = newCharacter
	} else if p3.name == "" {
		p3 = newCharacter
	} else if p4.name == "" {
		p4 = newCharacter
	} else {
		fmt.Println("Tous les emplacements de personnage sont occupés.")
		ReturnMenu()
		return
	}

	fmt.Println("Personnage créé avec succès.")
	ReturnMenu()
}

// ///////////////////////////////////////////////////////////////////FONCTION POUR LE RETOUR MENU/////////////////////////////////////////////////////////////////////////////////
func ReturnMenu() {
	fmt.Println("Retourner au menu (y/n)")

	var choice string
	_, Erreur := fmt.Scan(&choice)
	if Erreur != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	switch choice {
	case "y":
		Menu()
	case "n":
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////choisir personnage apres avoirr cliquer sur commencer////////////////
func (p *Player) ChoiceClass() {
	ClearConsole()
	fmt.Println("Choisis la classe de ton personnage :")
	fmt.Println("1. Humain")
	fmt.Println("2. Elfe")
	fmt.Println("3. Nain")
	fmt.Print("Choisissez une classe : ")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	switch choice {
	case 1:
		p.class = "Humain"
		p.lifemax = 100
	case 2:
		p.class = "Elfe"
		p.lifemax = 80
	case 3:
		p.class = "Nain"
		p.lifemax = 120
	default:
		fmt.Println("Option invalide.")
		p.ChoiceClass()
	}

	fmt.Printf("Ton personnage %s est un %s.\n", p.name, p.class)
	fmt.Println("Appuyez sur Entrée pour continuer...")
	fmt.Scanln() // Attendez l'entrée de l'utilisateur
}

// //////////////////////////////////////////////////////////////////////////////////////FONCTION INIT////////////////////////////////////////////////////////////////////////////
func (p *Player) Init(
	name string,
	class string,
	level int,
	lifemax int,
	life int,
	inventory map[string]int,
	money int,
	skill string,
	equipment Equipment, // Ajoutez cet argument
) {
	p.name = name
	p.class = class
	p.level = level
	p.lifemax = lifemax
	p.life = life
	p.inventory = map[string]int{"potions": 1, "potion de poison": 1}
	p.money = money
	p.skill = []string{"Coup de poing"}

	// Ajoutez l'équipement initial à la structure Equipment
	p.equipment = equipment

	if skill != "" {
		p.skill = append(p.skill, skill)
	}
}

// ///////////////////////////////////////////////////////////////////////////////GROS GOBELIN///////////////////////////////////////////////
func InitGoblin() Monster {
	gobelin := Monster{
		nom:       "Gobelin d'entraînement",
		pvMax:     40,
		pvActuels: 40, // Les points de vie actuels sont égaux aux points de vie maximum au début
		attaque:   5,
	}
	return gobelin

}

// ///////////////////////////////////////////////////////////////////////////////LE MENU/////////////////////////////////////////////////////////////////////
func Menu() {
	ClearConsole()
	fmt.Println("Menu:")

	// Créez un nouveau tableau pour le menu principal
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Option", "Description"})

	// Ajoutez les options du menu principal avec leur description au tableau
	table.Append([]string{"1", "Afficher les informations du personnage"})
	table.Append([]string{"2", "Afficher mon inventaire"})
	table.Append([]string{"3", "Voir le marchand"})
	table.Append([]string{"4", "Boire ma potion"})
	table.Append([]string{"5", "Boire ma potion poison"})
	table.Append([]string{"6", "Prendre ma boule de feu"})
	table.Append([]string{"7", "Forgeron"})
	table.Append([]string{"8", "Lancer le combat"})
	table.Append([]string{"9", "Quitter"})

	// Rendez le tableau joli
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Affichez le tableau
	table.Render()

	// Reste du code de la fonction Menu
	var choix string
	fmt.Print("Choisissez une option : ")
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	switch choix {
	case "1":
		p1.displayInfo()
	case "2":
		p1.AffInventory()
	case "3":
		fmt.Print("\033[H\033[2J")
		p1.Mymarchand()
	case "4":
		p1.TakePot()
	case "5":
		p1.PoisonPot()
	case "6":
		p1.SpellBook()
		Menu()
	case "7":
		p1.ForgeMenu()
	case "8":
		ClearConsole()
		TrainingFight()
	case "9":
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide.")
	}
}

// ///////////////////////////////////////////////////////////AFFICHER LE PERSONNAGE//////////////////////////////////////////////////////////////////////////////////////
func (p Player) displayInfo() {
	ClearConsole()
	fmt.Println("Informations du personnage :")

	// Créez un nouveau tableau
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Attribut", "Valeur"})

	// Ajoutez les attributs et leurs valeurs dans le tableau
	table.Append([]string{"Nom", p.name})
	table.Append([]string{"Classe", p.class})
	table.Append([]string{"Niveau", fmt.Sprintf("%d", p.level)})
	table.Append([]string{"Vie maximale", fmt.Sprintf("%d", p.lifemax)})
	table.Append([]string{"Vie actuelle", fmt.Sprintf("%d", p.life)})

	// Affichez l'inventaire sous forme de liste
	inventoryList := ""
	for item, count := range p.inventory {
		inventoryList += fmt.Sprintf("%s: %d\n", item, count)
	}
	table.Append([]string{"Inventaire", inventoryList})

	table.Append([]string{"Money", fmt.Sprintf("%d", p.money)})

	// Affichez les compétences sous forme de liste
	skillList := ""
	for _, skill := range p.skill {
		skillList += skill + "\n"
	}
	table.Append([]string{"Compétences", skillList})

	table.Append([]string{"Equipement", fmt.Sprintf("%+v", p.equipment)})

	// Rendez la table jolie
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Affichez le tableau
	table.Render()

	ReturnMenu()
}

// ///////////////////////////////////////////////////////////AFFICHER L'inventaire'//////////////////////////////////////////////////////////////////////////////////////
func (p *Player) AffInventory() {
	ClearConsole()

	// Créez un nouveau tableau pour l'inventaire du joueur
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Objet", "Quantité"})

	// Parcourez l'inventaire du joueur et ajoutez les objets et leurs quantités au tableau
	for item, count := range p.inventory {
		table.Append([]string{item, fmt.Sprintf("%d", count)})
	}

	// Rendez le tableau joli
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Affichez le tableau
	table.Render()
	fmt.Println("Appuyer sur (4) si vous n'utilisez rien du tout :")
	// Reste du code de la fonction AffInventory
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	switch choice {
	case 1:
		p1.TakePot()
	case 2:
		p1.Mypoison()
	case 3:
		p1.Myboule()
	case 4:
		ReturnMenu()

	default:
		fmt.Println("Option invalide.")
	}

	fmt.Println("je te vends ca")
	ReturnMenu()
}

// ////////////////////////////////////////////////////////////////////FONCTION POUR GERER LA CAPACITE DE L'INVENTAIRE///////////////////////////////////////////////////////////
func (p *Player) upgradeInventorySlot() {
	ClearConsole()

	// Vérifie si le joueur peut utiliser une augmentation de l'inventaire
	if p.inventoryUpgrades >= 3 {
		fmt.Println("Vous avez atteint le nombre maximum d'augmentations de l'inventaire (3 fois).")
		ReturnMenu()
		return
	}

	// Vérifie si le joueur a assez d'argent pour acheter l'augmentation de l'inventaire (coût fixe à 30 pièces d'or)
	if p.money < 30 {
		fmt.Println("Vous n'avez pas assez d'argent pour acheter une augmentation de l'inventaire.")
		ReturnMenu()
		return
	}

	// Augmente la capacité maximale de l'inventaire de +10
	p.inventoryUpgrades++
	// Soustrait le coût de l'augmentation de l'inventaire de la bourse d'argent du joueur
	p.Managemoney(-30)

	fmt.Println("Vous avez augmenté la capacité maximale de votre inventaire de +10 emplacements.")
	ReturnMenu()
}

// ///////////////////////////////////////////////////////////AFFICHER Le marchand//////////////////////////////////////////////////////////////////////////////////////
func (p *Player) Mymarchand() {

	ClearConsole()
	fmt.Println("Choisir un produit:")
	color.Green("Shop:")

	// Créez un nouveau tableau pour les produits du marchand
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Produit", "Prix"})

	// Ajoutez les produits et leurs prix dans le tableau
	table.Append([]string{"Potions", "5 elixir"})
	table.Append([]string{"Potions de Poison", "5 elixir"})
	table.Append([]string{"Sort", "5 elixir"})
	table.Append([]string{"Fourrure de Loup", "4 elixir"})
	table.Append([]string{"Peau de Troll", "7 elixir"})
	table.Append([]string{"Cuir de Sanglier", "3 elixir"})
	table.Append([]string{"Plume de Corbeau", "1 elixir"})
	table.Append([]string{"Augmentation d'inventaire", "30 elixir"}) // Ajout de l'objet "Augmentation d'inventaire"

	// Rendez le tableau joli
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Affichez le tableau
	table.Render()

	// Montant d'argent restant après l'achat
	fmt.Printf("Argent restant : %d elixir\n", p.money)

	// Reste du code de la fonction Mymarchand
	fmt.Println()
	fmt.Print("Choisissez un produit : ")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		ReturnMenu()
	}

	switch choice {
	case 1:
		p1.Mypotion()
	case 2:
		p1.Mypoison()
	case 3:
		p1.Myboule()
	case 4:
		p1.BuyItem("Fourrure de Loup", 4)
	case 5:
		p1.BuyItem("Peau de Troll", 7)
	case 6:
		p1.BuyItem("Cuir de Sanglier", 3)
	case 7:
		p1.BuyItem("Plume de Corbeau", 1)
	case 8:
		p1.upgradeInventorySlot()
	case 9:
		fmt.Println("Au revoir !")
		ReturnMenu()
	default:
		fmt.Println("Option invalide.")
	}

}

///////////////////////////////////////////////////////////////////////////////////////buy////////////////////////////////////////////////

func (p *Player) BuyItem(item string, cost int) {
	ClearConsole()
	if p.money >= cost {
		fmt.Printf("Tu as acheté %s à %d pièces d'or\n", item, cost)
		p.Managemoney(-cost)

		// Ajoutez l'item à l'inventaire du joueur
		if _, exists := p.inventory[item]; exists {
			p.inventory[item]++
		} else {
			p.inventory[item] = 1
		}

		fmt.Print("\033[H\033[2J")
		p.Mymarchand()
	} else {
		fmt.Println("Tu n'as pas assez d'oseilles")
	}
}

///////////////////////////////////////////////////////////////////////////////////FONCTION POUR ACHETER LA POTION///////////////////////////////////////////////////

func (p *Player) Mypotion() {
	ClearConsole()
	if p.money >= 5 {
		fmt.Println("Tu as acheté une potion de vie à 5 elixir")
		p.Managemoney(-5)

		p.inventory["potions"]++
		fmt.Print("\033[H\033[2J")
		p.Mymarchand()
	} else {
		println("Tu n'as pas assez d'oseilles")
	}
}

//////////////////////////////////////////////////////////////////////////////////// FONCTION POUR ACHETER LE POISON///////////////////////////////////////////////////

func (p *Player) Mypoison() {
	ClearConsole()
	if p.money >= 5 {
		fmt.Println("Tu as acheté une potion de poison à 5 elixir")
		p.Managemoney(-5)
		fmt.Println("Tu as acheté une potion de de poison")
		p.inventory["potion de poison"]++
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Println("Tu n'as pas assez d'oseilles")
	}
	p.Mymarchand()
	ClearConsole()
}

// ///////////////////////////////////////////////////////////////////////////////POUR ACHETER LA BOULE DE FEU QUE UNE FOIS //////////////////////////////////////////////////
func (p *Player) Myboules() {
	ClearConsole()
	if p.inventory["Sort"] > 0 {
		fmt.Println("Tu as acheté Un sort")
		p.inventory["boule de feu"]++
		fmt.Print("\033[H\033[2J")
	}
	p.Mymarchand()
	ClearConsole()
}

// ///////////////////////////////////////////////////////////////////////////////MON PETIT PORTEFEUILLE HAHAAA //////////////////////////////////////////
func (p *Player) Managemoney(nb int) {
	p.money += nb
}

// ////////////////////////////////////////////////////BOIRE POTION PENDANT LE COMBAT////////////////////////
func (p *Player) TakePot2() {
	ClearConsole()
	if p.inventory["potions"] > 0 {
		fmt.Println("Vous avez :", p.life)
		p.life += 50
		if p.life > p.lifemax {
			p.life = p.lifemax
		}
		fmt.Println("Vous avez bu une potion")
		color.Green("Vous avez maintenant : %d", p.life)
		p.inventory["potions"]--
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Println("Vous n'avez plus de potion")
	}

}

///////////////////////////////////FONCTION BOIRE LA POTION DE VIE///////////////////////////////////////////////////

func (p *Player) TakePot() {
	ClearConsole()
	if p.inventory["potions"] > 0 {
		fmt.Println("Vous avez :", p.life)
		p.life += 50
		if p.life > p.lifemax {
			p.life = p.lifemax
		}
		fmt.Println("Vous avez bu une potion")
		color.Green("Vous avez maintenant : %d", p.life)
		p.inventory["potions"]--
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Println("Vous n'avez plus de potion")
	}
	ReturnMenu()

}

///////////////////////////////////FONCTION BOIRE LE POISON///////////////////////////////////////////////////

func (p *Player) PoisonPot() {
	ClearConsole()
	fmt.Println("Vous avez pris une potion poison")
	p1.Dead()
	fmt.Println("Vous avez :", p.life)
	for i := 0; i < 3; i++ {
		p.life -= 10
		fmt.Println("Vous perdez 10 PV")
		fmt.Println("Vous avez :", p.life, " / ", p.lifemax)
		time.Sleep(3 * time.Second)
		p.inventory["potion de poison"]--
	}

	color.Green("Vous avez maintenant : %d", p.life)
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
	ReturnMenu()
}

///////////////////////////////////FONCTION BOULE DE FEU//////////////////////////////////////////////////////////////////////////////////////

func (p *Player) SpellBook() {
	newSpell := "Boule de feu"

	// Vérifiez si le sort n'est pas déjà dans la liste de compétences
	for _, spell := range p.skill {
		if spell == newSpell {
			color.Green("Vous avez déjà appris ce sort.%d")
			return
		}
	}

	/////////////////////////////////////////////////////// Ajoutez le sort à la liste de compétences///////////////////////////////////////////////
	p.skill = append(p.skill, newSpell)
	fmt.Printf("Vous avez appris le sort : %s\n", newSpell)
}

func (p *Player) Myboule() {
	ClearConsole()
	fmt.Println("Tu as acheté une boule de feu")
	p.inventory["boule de feu"]++
	fmt.Print("\033[H\033[2J")

	// Appel de SpellBook pour ajouter le sort "Boule de feu" à la liste de compétences
	p.SpellBook()

	ReturnMenu()
}

// /////////////////////////////////////////////////////FONCTION POUR RESSUCITER///////////////////////////////////////////////////
func (p *Player) Dead() {
	ClearConsole()
	if p.life <= 0 {
		fmt.Println("le joueur est mort")
		p.life += 50
		fmt.Println("Le joueur a été ressuscité avec 50% de points de vie maximum")
		ReturnMenu()
		ClearConsole()
	}
}

/////////////////////////////////////////////////////////////////////////////AFFICHER LE FORGERON/////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////FONCTION POUR GERER LE SOUS MENU DU FORGERON/////////////////////////////////////////////////////////////

func (p *Player) ForgeMenu() {
	ClearConsole()
	fmt.Println("Forgeron - Que voulez-vous fabriquer ?")
	fmt.Println("Appuyer sur (4) si vous n'utilisez rien du tout :")

	// Créez un nouveau tableau pour les options du forgeron
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Équipement", "Ressources Requises"})

	// Ajoutez les options du forgeron avec les ressources requises au tableau
	table.Append([]string{"1 ° Chapeau de l'aventurier", "Plume de Corbeau (1), Cuir de Sanglier (1)"})
	table.Append([]string{"2 ° Tunique de l'aventurier", "Fourrure de Loup (2), Peau de Troll (1)"})
	table.Append([]string{"3 ° Bottes de l'aventurier", "Fourrure de Loup (1), Cuir de Sanglier (1)"})

	// Rendez le tableau joli
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Affichez le tableau
	table.Render()

	// Reste du code de la fonction ForgeMenu
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		ReturnMenu()
	}

	switch choice {
	case 1:
		p1.CraftEquipment("Chapeau de l'aventurier", map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1})
	case 2:
		p1.CraftEquipment("Tunique de l'aventurier", map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1})
	case 3:
		p1.CraftEquipment("Bottes de l'aventurier", map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1})
	case 4:
		ReturnMenu()
	default:
		fmt.Println("Option invalide.")
	}
}

////////////////////////////////////////////////////////////////////////CONCERNANT L EQUIPEMENT DU FORGERON //////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////FONCTION POUR GERER LA FABRICATION DES EQUIPEMENTS////////////////////////////////////////////////////////
func (p *Player) CraftEquipment(equipment string, requiredResources map[string]int) {
	ClearConsole()

	// Retournez l'équipement précédent dans l'inventaire s'il est équipé
	switch equipment {
	case "Chapeau de l'aventurier":
		if p.equipment.Head != "" {
			// Ajoutez l'équipement précédent à l'inventaire du joueur
			if _, exists := p.inventory[p.equipment.Head]; exists {
				p.inventory[p.equipment.Head]++
			} else {
				p.inventory[p.equipment.Head] = 1
			}
			// Réduisez les points de vie maximum en conséquence
			p.lifemax -= 10
		}
	case "Tunique de l'aventurier":
		// Faites de même pour le torse
	case "Bottes de l'aventurier":
		// Faites de même pour les pieds
	}

	// Vérifie si le joueur a les ressources nécessaires pour fabriquer l'équipement
	for resource, quantity := range requiredResources {
		if p.inventory[resource] < quantity {
			fmt.Printf("Vous n'avez pas assez de %s pour fabriquer %s.\n", resource, equipment)
			ReturnMenu()
			return
		}
	}

	// Vérifie si le joueur a assez d'argent pour fabriquer l'équipement (coût de fabrication fixe à 5 pièces d'or)
	if p.money < 5 {
		fmt.Printf("Vous n'avez pas assez d'argent pour fabriquer %s.\n", equipment)
		ReturnMenu()
		return
	}

	// Vérifie si le joueur a de la place dans son inventaire pour ajouter l'équipement
	if len(p.inventory) >= 10 {
		fmt.Println("Votre inventaire est plein. Vous ne pouvez pas fabriquer de nouveaux équipements.")
		ReturnMenu()
		return
	}

	// Soustrait les ressources nécessaires de l'inventaire du joueur
	for resource, quantity := range requiredResources {
		p.inventory[resource] -= quantity
	}

	// Soustrait le coût de fabrication de la bourse d'argent du joueur
	p.Managemoney(-5)

	// Équipez l'équipement fabriqué dans la structure Equipment
	switch equipment {
	case "Chapeau de l'aventurier":
		p.equipment.Head = equipment
		p.lifemax += 10
	case "Tunique de l'aventurier":
		p.equipment.Torso = equipment
		p.lifemax += 25
	case "Bottes de l'aventurier":
		p.equipment.Feet = equipment
		p.lifemax += 15
	}

	// Ajoute l'équipement fabriqué à l'inventaire du joueur
	if _, exists := p.inventory[equipment]; exists {
		p.inventory[equipment]++
	} else {
		p.inventory[equipment] = 1
	}

	fmt.Printf("Vous avez fabriqué %s et il a été ajouté à votre équipement.\n", equipment)
	ReturnMenu()
}

// ///////////////////////////////////////////////////////////partir s'entrainer//////////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////Lancer combat d'entrainement////////////////////////////
func TrainingFight() {
	joueurPv := 50 // Points de vie du joueur
	tourDeCombat := 1

	gobelinNom := "Gobelin"
	gobelinPvMax := 30
	gobelinPvActuels := gobelinPvMax
	gobelinAttaque := 5

	fmt.Printf("Un %s apparaît !\n", gobelinNom)
	time.Sleep(1000000000)

	for {
		fmt.Printf("\nTour %d de combat\n", tourDeCombat)

		// Tour du joueur
		fmt.Printf("Points de vie du joueur : %d/%d\n", joueurPv, 50)
		fmt.Printf("Points de vie du %s : %d/%d\n", gobelinNom, gobelinPvActuels, gobelinPvMax)

		// Le joueur choisit son attaque
		var choixAttaque int
		fmt.Println("Choisissez votre attaque :")
		fmt.Println("1. Boule de feu")
		fmt.Println("2. Poison")
		fmt.Println("3. Boit ta potion de vie")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choixAttaque)

		switch choixAttaque {
		case 1:
			fmt.Println("Le joueur lance une boule de feu !")
			gobelinPvActuels -= 15 // Supposons que la boule de feu inflige 15 points de dégâts
		case 2:
			fmt.Println("Le joueur utilise du poison !")
			gobelinPvActuels -= 5 // Supposons que le poison inflige 5 points de dégâts par tour
		case 3:
			fmt.Println("Le joueur boit sa potion de vie !")
			p1.TakePot2()
		default:
			fmt.Println("Attaque non reconnue. Le joueur perd son tour.")
		}

		if gobelinPvActuels <= 0 {
			fmt.Printf("Le joueur a vaincu le %s !\n", gobelinNom)
			break
		}

		// Tour du monstre
		fmt.Println("Le monstre attaque le joueur !")
		joueurPv -= gobelinAttaque // Le monstre inflige des dégâts égaux à son attaque

		if joueurPv <= 0 {
			fmt.Printf("%s a vaincu le joueur. Le combat est terminé.\n", gobelinNom)
			break
		}
		// Code pour le combat avec le gobelin...

		victoryAnimation() // Affichez l'animation de victoire

		// Continuez avec le reste du combat ou retournez au menu principal, selon votre logique de jeu.

		tourDeCombat++
	}
}
func victoryAnimation() {
	// Vous pouvez utiliser des chaînes ASCII art pour créer votre animation.
	animation := `
    __   __            _    _ _       
    \ \ / /__ __ _ __| |__| | | ___  
     \ V / -_| '_/ _| '_ \ | |(_-<  
      \_/\___|_| \__|_.__/_|_|/__/
    `
	fmt.Println(animation)
}

func GoblinPattern(joueur *Player, gobelin Monster) {
	// Initialisez les variables pour suivre le nombre de tours et le modificateur de dégâts
	tourCompte := 0
	modificateurDegats := 1.0

	// Continuez le combat tant que le joueur ou le gobelin n'est pas vaincu
	for joueur.life > 0 && gobelin.pvActuels > 0 {
		// Incrémentez le nombre de tours
		tourCompte++

		// Vérifiez s'il est temps pour le gobelin d'utiliser l'attaque à 200% de dégâts
		if tourCompte%3 == 0 {
			modificateurDegats = 2.0
		} else {
			modificateurDegats = 1.0
		}

		// Calculez les dégâts infligés par le gobelin
		degats := int(float64(gobelin.attaque) * modificateurDegats)

		// Infligez des dégâts au joueur
		joueur.life -= degats

		// Assurez-vous que la vie du joueur ne devienne pas négative
		if joueur.life < 0 {
			joueur.life = 0
		}

		// Affichez les informations sur l'attaque
		fmt.Printf("%s inflige à %s %d de dégâts\n", gobelin.nom, joueur.name, degats)
		fmt.Printf("%s : %d/%d PV\n\n", joueur.name, joueur.life, joueur.lifemax)

		// Vérifiez si le joueur est vaincu
		if joueur.life == 0 {
			fmt.Printf("%s a été vaincu par %s !\n", joueur.name, gobelin.nom)
		} else {
			// Attendez un moment pour simuler le délai entre les tours
			time.Sleep(1 * time.Second)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////MISSION 3 INITIATIVE //////////////////////////////////////////////////////////
