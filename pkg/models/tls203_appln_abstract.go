package models

// Tls203ApplnAbstr is the structure for the application abstract
// This is the 1-paragraph summary of the invention which is shown on the first page of a publication.
// By design, in PATSTAT abstracts are related not to the individual publication, but to the application of
// the publication.
// Abstracts can be in any language. PATSTAT contains only 1 abstract per application.
// Abstracts in English language are preferred. (TLS203_APPLN_ABSTRACT)
//
// This table contains the English language abstract, if available. If there is no abstract in
// English, then it contains the most recent abstract in another language.
type Tls203ApplnAbstr struct {
	ApplnId         int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	ApplnAbstractLg string `json:"applnAbstractLg" gorm:"column:appln_abstract_lg;type:char(2);default:'';not null"`
	ApplnAbstract   string `json:"applnTitle" gorm:"column:appln_abstract;type:text;not null"`
}

// TableName sets the sql table name for this struct type
func (obj *Tls203ApplnAbstr) TableName() string {
	return "tls203_appln_abstr"
}
