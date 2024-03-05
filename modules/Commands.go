package modules

import "fmt"

func Rules(rules []string) {
	var tam_rules int = len(rules)
	if tam_rules < 2 {
		fmt.Println("Uso: mtsound <comando> <argumento>")
		return
	}

	rule := rules[0]
	switch rule {
	case "next":
		if tam_rules == 2{
			fmt.Printf("Indo para a música %s\n", rules[1])
		}
	case "back":
		fmt.Printf("Voltando para música %s\n", rules[1])
	case "play":
		fmt.Println("Música iniciada")
	case "pause":
		fmt.Println("Música pausada")
	default:
		fmt.Println("Comando não reconhecido.")
	}
}
