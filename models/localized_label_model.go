package models

type LocalizedLabelModel struct {
	Key         string `json:"key"`
	Translation string `json:"translation"`
}

func (m *LocalizedLabelModel) Id() string {
	return m.Key
}
