package api

type RestClient struct {
	Host  string
	Token string
}

// CreateUser action
func (rc *RestClient) CreateUser(u *User) error {
	resp, err := RequestWithJSON("POST", rc.Host+createUserURL(), u, "")
	if err != nil {
		return err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return err
	}

	return ErrFrom(rd)
}

// Login action
func (rc *RestClient) Login(u *User) error {
	resp, err := RequestWithJSON("POST", rc.Host+loginURL(), u, "")
	if err != nil {
		return err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return err
	}

	if rd.Success {
		rc.Token = rd.AccessToken
	}

	return ErrFrom(rd)
}

// UserDetails action
func (rc *RestClient) UserDetails() (*User, error) {
	resp, err := RequestWithJSON("GET", rc.Host+userURL(), nil, rc.Token)
	if err != nil {
		return nil, err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	return &rd.User, ErrFrom(rd)
}

// NotesList action
func (rc *RestClient) NotesList() ([]Note, error) {
	resp, err := RequestWithJSON("GET", rc.Host+notesListURL(), nil, rc.Token)
	if err != nil {
		return nil, err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	return rd.Notes, ErrFrom(rd)
}

// CreateNote action
func (rc *RestClient) CreateNote(n *Note) error {
	resp, err := RequestWithJSON("POST", rc.Host+createNoteURL(), n, rc.Token)
	if err != nil {
		return err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return err
	}

	return ErrFrom(rd)
}

// NoteDetails action
func (rc *RestClient) NoteDetails(noteID int) (*Note, error) {
	resp, err := RequestWithJSON("GET", rc.Host+noteURL(noteID), nil, rc.Token)
	if err != nil {
		return nil, err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	return &rd.Note, ErrFrom(rd)
}

// NoteRemove action
func (rc *RestClient) NoteRemove(noteID int) error {
	resp, err := RequestWithJSON("DELETE", rc.Host+noteURL(noteID), nil, rc.Token)
	if err != nil {
		return err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return err
	}

	return ErrFrom(rd)
}

// NotePatch action
func (rc *RestClient) NotePatch(noteID int) error {
	resp, err := RequestWithJSON("PUT", rc.Host+noteURL(noteID), nil, rc.Token)
	if err != nil {
		return err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return err
	}

	return ErrFrom(rd)
}

// PublishedNotesList action
func (rc *RestClient) PublishedNotesList() ([]Note, error) {
	resp, err := RequestWithJSON("GET", rc.Host+publishedNotesListURL(), nil, "")
	if err != nil {
		return nil, err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	return rd.Notes, ErrFrom(rd)
}

// PublishedNoteDetails action
func (rc *RestClient) PublishedNoteDetails(noteID int) (*Note, error) {
	resp, err := RequestWithJSON("GET", rc.Host+publishedNoteURL(noteID), nil, "")
	if err != nil {
		return nil, err
	}
	rd, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	return &rd.Note, ErrFrom(rd)
}
