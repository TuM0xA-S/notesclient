package api

import "fmt"

func createUserURL() string {
	return "/api/users"
}

func loginURL() string {
	return "/api/me"
}

func notesListURL(page int) string {
	return fmt.Sprintf("/api/me/notes?page=%d", page)
}

func createNoteURL() string {
	return "/api/me/notes"
}

func noteURL(noteID int) string {
	return fmt.Sprintf("/api/me/notes/%d", noteID)
}

func userURL() string {
	return "/api/me"
}

func publishedNotesListURL(page int) string {
	return fmt.Sprintf("/api/notes?page=%d", page)
}

func publishedNoteURL(noteID int) string {
	return fmt.Sprintf("/api/notes/%d", noteID)
}

func anotherUserURL(userID int) string {
	return fmt.Sprintf("/api/users/%d", userID)
}
