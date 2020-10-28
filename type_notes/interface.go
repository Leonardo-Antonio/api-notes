package type_notes

type iTypeNote interface {
	GetTypes() (TypeNotes []model, err error)
	CreateActivity(name string) error
	DeleteType(id int) error
}
