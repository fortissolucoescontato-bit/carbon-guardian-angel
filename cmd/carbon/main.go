package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("carbon-guardian-angel %s\n", version)
		return
	}

	fmt.Println("Carbon Guardian Angel: O Guardião da Identidade Nexo-Fortis.")
	fmt.Println("Uso: carbon-guardian-angel version")
}
