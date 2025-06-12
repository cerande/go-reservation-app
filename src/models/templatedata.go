package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int32
	FloatMap  map[string]float32
	DataMap   map[string]any
	Flash     string
	Warning   string
	Error     string
}
