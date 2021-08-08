package menu

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

	// File Picker
	var fileLocation string

	uri := binding.NewString()
	_ = uri.Set("please choose file ...")

	showFilePicker := func() {
		onChosen := func(f fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println(err)
				return
			}
			if f == nil {
				return
			}
			fileLocation = f.URI().String()
			fileLocation = fileLocation[7:]
			fmt.Println("chosen: ", fileLocation)
			_ = uri.Set(f.URI().String())
		}
		dialog.ShowFileOpen(onChosen, myWindow)
	}

	entry := widget.NewEntryWithData(uri)
	entry.Disable()

	button := widget.NewButtonWithIcon("OPEN", theme.FileIcon(), showFilePicker)
	card := widget.NewCard("File Picker", "",
		container.New(layout.NewVBoxLayout(), entry, button))

	form := &widget.Form{
		Items: []*widget.FormItem{ // use entrys in form
			{Text: "Subject", Widget: subject, HintText: "Subject of Email"},
			{Text: "Email Address", Widget: fromEmail, HintText: "Email Address, sending from"},
			{Text: "Password", Widget: passEmail, HintText: "Password for Email"},
			{Text: "SMTP Host", Widget: smtpHost, HintText: "Host for Email Service"},
			{Text: "SMTP Port", Widget: smtpPort, HintText: "Port Number for Host"},
			{Text: "HTML File", Widget: card},
			{Widget: button},
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
