package api

import "fmt"

func createUserURL() string {
	return "/api/user/create"
}

func loginURL() string {
	return "/api/user/login"
}

func notesListURL() string {
	return "/api/me/notes"
}

func createNoteURL() string {
	return "/api/me/notes/create"
}

func noteURL(noteID int) string {
	return fmt.Sprintf("/api/me/notes/%d", noteID)
}

func userURL() string {
	return "/api/me"
}

func publishedNotesListURL() string {
	return "/api/notes"
}

func publishedNoteURL(noteID int) string {
	return fmt.Sprintf("/api/note/%d", noteID)
}
