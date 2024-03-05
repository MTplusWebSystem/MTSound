// main.go
package main

import (
	"mtsound/modules"
	"os"
)

func main() {
	rules := os.Args[1:]
	modules.Rules(rules)
}
