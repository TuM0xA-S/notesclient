package main

import (
	"notesclient/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// NotesClient represents application
type NotesClient struct {
	api.RestClient
	App    fyne.App
	Window fyne.Window
}

// NewNotesClient constructs notes client
func NewNotesClient() *NotesClient {
	n := &NotesClient{}
	n.App = app.New()
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
	n.Window = w
	tabs := container.NewAppTabs(
		container.NewTabItem("publications", widget.NewLabel("publications here")),
		container.NewTabItem("my notes", widget.NewLabel("my notes here")),
		container.NewTabItem("account", widget.NewLabel("account here")),
		container.NewTabItem("server", n.NewServerWidget()),
	)

	w.SetContent(tabs)
	w.Show()
}

// NewServerWidget ...
func (n *NotesClient) NewServerWidget() fyne.CanvasObject {
	c := widget.NewForm()
	e := widget.NewEntry()
	e.SetPlaceHolder("http[s]://host[:port]")
	c.Append("Host:", e)
	c.SubmitText = "Connect"
	c.OnSubmit = func() {
		n.Host = e.Text
		_, err := n.PublishedNotesList()
		dialog.ShowError(err, n.Window)
	}
	return c
}
