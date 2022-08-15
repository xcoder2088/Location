package models

// TemplateData hold data sent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FlotMap   map[string]float32

	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
