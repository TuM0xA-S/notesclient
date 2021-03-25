package main

import (
	"encoding/json"
	"log"
	"notesclient/api"
	"os"

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

type config struct {
	Host     string
	Username string
	Password string
}

var cfg = &config{}

// NewNotesClient constructs notes client
func NewNotesClient() *NotesClient {
	n := &NotesClient{}
	n.App = app.New()
	f, err := os.Open("cfg.json")
	if err == nil {
		err := json.NewDecoder(f).Decode(cfg)
		if err != nil {
			log.Println(err)
		}
	}
	n.Host = cfg.Host

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
	w.Resize(fyne.NewSize(1000, 800))
	w.SetMaster()
	tabs := container.NewAppTabs(
		container.NewTabItem("publications", widget.NewLabel("publications here")),
		container.NewTabItem("my notes", widget.NewLabel("my notes here")),
		container.NewTabItem("account", n.NewAccountWidget()),
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
	e.Text = n.Host
	c.Append("Host:", e)
	c.SubmitText = "Connect"
	c.OnSubmit = func() {
		n.Host = e.Text
		_, err := n.PublishedNotesList()
		if err == nil {
			dialog.ShowInformation("connected", "get response from a server", n.Window)
			return
		}
		dialog.ShowError(err, n.Window)
	}
	return c
}

// NewAccountWidget ...
func (n *NotesClient) NewAccountWidget() fyne.CanvasObject {
	registerHandler := func() {
		passwordEntry := widget.NewEntry()
		usernameEntry := widget.NewEntry()
		fi := []*widget.FormItem{
			{Text: "username", Widget: usernameEntry},
			{Text: "password", Widget: passwordEntry},
		}
		callback := func(send bool) {
			if !send {
				return
			}
			u := &api.User{Username: usernameEntry.Text, Password: passwordEntry.Text}
			err := n.CreateUser(u)
			if err == nil {
				dialog.ShowInformation("user created", "now you can login", n.Window)
				return
			}
			dialog.ShowError(err, n.Window)
		}
		dialog.ShowForm("user data", "register", "cancel", fi, callback, n.Window)
	}
	var updateStatus func()
	loginHandler := func() {
		passwordEntry := widget.NewEntry()
		usernameEntry := widget.NewEntry()
		passwordEntry.Text = cfg.Password
		usernameEntry.Text = cfg.Username
		fi := []*widget.FormItem{
			{Text: "username", Widget: usernameEntry},
			{Text: "password", Widget: passwordEntry},
		}
		callback := func(send bool) {
			if !send {
				return
			}
			u := &api.User{Username: usernameEntry.Text, Password: passwordEntry.Text}
			err := n.Login(u)
			if err != nil {
				dialog.ShowError(err, n.Window)
				return
			}
			dialog.ShowInformation("logged", "now you can work with notes", n.Window)
			updateStatus()
		}
		dialog.ShowForm("user data", "login", "cancel", fi, callback, n.Window)
	}

	res := container.NewVBox()
	updateStatus = func() {
		user, err := n.UserDetails()
		if err != nil {
			res.Objects = []fyne.CanvasObject{widget.NewLabel("not logged"),
				container.NewHBox(widget.NewButton("register", registerHandler), widget.NewButton("login", loginHandler)),
			}
			return
		}
		res.Objects = []fyne.CanvasObject{widget.NewLabel("logged as " + user.Username)}
	}
	updateStatus()
	return res
}
