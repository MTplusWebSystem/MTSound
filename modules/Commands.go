package modules

import "fmt"

func Rules(rules []string) {
	var tam_rules int = len(rules)

	rule := rules[0]
	switch rule {
	case "next":
		if tam_rules == 2{
			fmt.Printf("Indo para a música %s\n", rules[1])
		}else{
			fmt.Println("Indo para próxima música")
		}
	case "back":
		if tam_rules == 2{
			fmt.Printf("Voltando para música %s\n", rules[1])
		}else{
			fmt.Println("Voltando para próxima música")
		}
	case "create":
		if tam_rules == 3 && rules[1] == "playlist"{
			fmt.Printf("Criando a playlist %s\n", rules[2])
		} else {
			fmt.Println("Comando create inválido. Uso: create playlist name")
		}			
	case "play":
		fmt.Println("Música iniciada")
	case "pause":
		fmt.Println("Música pausada")
	default:
		fmt.Println("Uso: mtsound <comando> <argumento>")
	}
}
