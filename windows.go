package main

import (
	"notesclient/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// NewEditWindow ...
func (n *NotesClient) NewEditWindow(noteID int) (fyne.Window, error) {
	win := n.App.NewWindow("edit note")
	win.Resize(fyne.NewSize(500, 500))
	title := widget.NewEntry()
	title.SetPlaceHolder("title")
	body := widget.NewMultiLineEntry()
	body.SetPlaceHolder("body")
	public := widget.NewCheck("public", nil)
	createButton := widget.NewButton("submit", func() {
		notePatch := api.NotePatch{Title: title.Text, Body: body.Text, Published: public.Checked, ID: noteID}
		err := n.NoteUpdate(&notePatch)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			win.Close()
		}
	})
	content := container.NewBorder(title, container.NewVBox(public, createButton), nil, nil, body)
	note, err := n.NoteDetails(noteID)
	if err != nil {
		return nil, err
	}
	title.Text = note.Title
	body.Text = note.Body
	public.SetChecked(note.Published)
	content.Refresh()
	win.SetContent(content)

	return win, nil
}

func (n *NotesClient) fetchAuthorName(userID int) string {
	user, err := n.AnotherUserDetails(userID)
	if err != nil {
		return "<error>"
	}
	return user.Username
}

// NewCreateWindow ...
func (n *NotesClient) NewCreateWindow() fyne.Window {
	win := n.App.NewWindow("create note")
	win.Resize(fyne.NewSize(500, 500))
	title := widget.NewEntry()
	title.SetPlaceHolder("title")
	body := widget.NewMultiLineEntry()
	body.SetPlaceHolder("body")
	public := widget.NewCheck("public", nil)
	createButton := widget.NewButton("create", func() {
		note := api.Note{Title: title.Text, Body: body.Text, Published: public.Checked}
		err := n.CreateNote(&note)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			win.Close()
		}
	})
	content := container.NewBorder(title, container.NewVBox(public, createButton), nil, nil, body)
	win.SetContent(content)
	return win
}

// NewViewerWindow ...
func (n *NotesClient) NewViewerWindow(noteID int) fyne.Window {
	win := n.App.NewWindow("view note")
	win.Resize(fyne.NewSize(500, 500))
	title := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{
		Bold: true,
	})

	title.Wrapping = fyne.TextWrapWord
	body := widget.NewLabel("")
	body.Wrapping = fyne.TextWrapWord

	reload := widget.NewButton("reload", nil)
	bottomButtons := container.NewHBox()
	publicState := widget.NewLabel("")
	content := container.NewBorder(container.NewVBox(reload, title),
		container.NewVBox(publicState, bottomButtons), nil, nil, container.NewVScroll(body))
	win.SetContent(content)

	fetch := func(noteID int) (*api.Note, error) {
		note, err := n.NoteDetails(noteID)
		if err == nil {
			return note, err
		}
		return n.PublishedNoteDetails(noteID)
	}

	loadNote := func() {
		note, err := fetch(noteID)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		body.Text = note.Body
		title.Text = n.generateHeader(note)
		publicState.Text = "--published--"
		if !note.Published {
			publicState.SetText("--not published--")
		}
		content.Refresh()
	}
	reload.OnTapped = loadNote

	loadNote()

	_, err := n.NoteDetails(noteID)
	if err == nil {
		editHandler := func() {
			winEdit, err := n.NewEditWindow(noteID)
			if err != nil {
				dialog.ShowError(err, n.Window)
				return
			}
			winEdit.Show()
		}
		bottomButtons.Add(widget.NewButton("edit", editHandler))
		removeHandler := func() {
			err := n.NoteRemove(noteID)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			win.Close()
			n.Window.Content().Refresh()
		}
		bottomButtons.Add(widget.NewButton("delete", removeHandler))
	}
	return win
}
