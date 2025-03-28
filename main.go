// main.go
package main

import (
	"os"
	"NetTools-CLI/app" 
	)
	
func main() {
	resolver := app.NewNetResolver()
	app := app.Gerar(resolver)
	app.Run(os.Args)
}
