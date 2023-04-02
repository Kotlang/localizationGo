package models

type LanguageListModel struct {
	IsoCode  string `json:"iso_code"`
	Language string `json:"language"`
}

func (m *LanguageListModel) Id() string {
	return m.IsoCode
}
