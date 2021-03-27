package main

import (
	"fmt"
	"notesclient/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type listFetcher func(int) ([]api.Note, error)
type itemFetcher func(int) (*api.Note, error)

func (n *NotesClient) generateHeader(note *api.Note) string {
	title := fmt.Sprintf("%s | by %s at %v", note.Title, n.fetchAuthorName(note.UserID),
		note.CreatedAt.Local().Format("Jan 2 15:04 2006"))
	return title
}

//NewNotesListWidget ...
func (n *NotesClient) NewNotesListWidget(fetchList listFetcher) fyne.CanvasObject {
	page := 1
	newNoteHeader := func(note api.Note) fyne.CanvasObject {
		but := widget.NewButton(n.generateHeader(&note), func() {
			win := n.NewViewerWindow(note.ID)
			win.Show()
		})
		but.Alignment = widget.ButtonAlignLeading
		return but
	}
	list := container.NewVBox()
	updateHandler := func() {
		page = 1
		notes, err := fetchList(page)
		if err != nil {
			dialog.ShowError(err, n.Window)
			return
		}
		list.Objects = nil
		for _, note := range notes {
			list.Add(newNoteHeader(note))
		}
		list.Refresh()
	}
	loadNextHandler := func() {
		page++
		notes, err := fetchList(page)
		if err != nil {
			dialog.ShowError(err, n.Window)
			return
		}
		for _, note := range notes {
			list.Add(newNoteHeader(note))
		}

	}
	createHandler := func() {
		win := n.NewCreateWindow()
		win.SetOnClosed(func() {
			updateHandler()
		})
		win.Show()
	}
	topButtons := container.NewHBox(widget.NewButton("update", updateHandler), widget.NewButton("create", createHandler))
	updateHandler()
	return container.NewBorder(topButtons, widget.NewButton("load next", loadNextHandler), nil, nil, container.NewVScroll(list))
}
