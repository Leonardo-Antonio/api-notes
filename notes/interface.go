package notes

type iNotes interface {
	DeleteNote(IDS model) error
	UpdateNote(note model) error
	GetAllNotes(ID int) (notes []model, err error)
	NewNote(note model) error
}
