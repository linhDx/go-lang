package model

type (
	NoteModel struct {
	BaseNote
	Content string `json:"content"`
}
	TransformedNote struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Content string   `json:"content"`
	}
)
