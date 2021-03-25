package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// NotesClient represents application
type NotesClient struct {
	Token binding.String
	Host  binding.String
	App   fyne.App
}

// NewNotesClient constructs notes client
func NewNotesClient() *NotesClient {
	n := &NotesClient{}
	n.App = app.New()
	n.Host = binding.NewString()
	n.Token = binding.NewString()
	n.InitializeMainWindow()
	return n
}

//Run notes client
func (n *NotesClient) Run() {
	n.App.Run()
}

// InitializeMainWindow creates new master window
func (n *NotesClient) InitializeMainWindow() {
	w := n.App.NewWindow("notes gui client")
	tabs := container.NewAppTabs(
		container.NewTabItem("publications", widget.NewLabel("publications here")),
		container.NewTabItem("my notes", widget.NewLabel("my notes here")),
		container.NewTabItem("account", widget.NewLabel("account here")),
		container.NewTabItem("server", NewServerWidget(n.Host)),
	)

	w.SetContent(tabs)
	w.Show()
}

// NewServerWidget ...
func NewServerWidget(host binding.String) fyne.CanvasObject {
	c := widget.NewForm()
	e := widget.NewEntryWithData(host)
	e.SetPlaceHolder("host[:port]")
	c.Append("Host:", e)
	c.SubmitText = "Connect"
	c.OnSubmit = func() {
		http
	}
	return c
}
