package menu

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MainMenu(app fyne.App) {
	myWindow := app.NewWindow("Email Sender")
	myWindow.Resize(fyne.NewSize(800, 800))
	myWindow.CenterOnScreen()

	// Declare Entry Lines
	subject := widget.NewEntry()
	fromEmail := widget.NewEntry()
	passEmail := widget.NewPasswordEntry()
	smtpHost := widget.NewEntry()
	smtpPort := widget.NewEntry()

	// Set PlaceHolder Values
	subject.SetPlaceHolder("Subject")
	fromEmail.SetPlaceHolder("example@gmail.com")
	passEmail.SetPlaceHolder("")
	smtpHost.SetPlaceHolder("smtp.gmail.com")
	smtpPort.SetPlaceHolder("587")

	form := &widget.Form{
		Items: []*widget.FormItem{ // use entrys in form
			{Text: "Subject", Widget: subject, HintText: "Subject of Email"},
			{Text: "Email Address", Widget: fromEmail, HintText: "Email Address, sending from"},
			{Text: "Password", Widget: passEmail, HintText: "Password for Email"},
			{Text: "SMTP Host", Widget: smtpHost, HintText: "Host for Email Service"},
			{Text: "SMTP Port", Widget: smtpPort, HintText: "Port Number for Host"},
		},
		OnSubmit: func() { // Print for now to console
			log.Println("Form submitted:", subject.Text)
			log.Println("fromEmail:", fromEmail.Text)
			log.Println("passEmail:", passEmail.Text)
			log.Println("smtpHost:", smtpHost.Text)
			log.Println("smtpPort:", smtpPort.Text)
		},
	}

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
