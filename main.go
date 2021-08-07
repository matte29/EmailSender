package main

import (
	"fyne.io/fyne/v2/app"
	menu "github.com/matte29/EmailSender/src/layout"
	resources "github.com/matte29/EmailSender/src/resources"
)

func main() {
	// Set App and Window
	myApp := app.New()
	myApp.SetIcon(resources.IconPng)
	menu.MainMenu(myApp)
}
