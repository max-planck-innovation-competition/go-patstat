package models

/*
CREATE TABLE tls222_appln_jp_class (
    appln_id integer DEFAULT 0 NOT NULL,
    jp_class_scheme varchar(5) DEFAULT '' NOT NULL,
    jp_class_symbol varchar(50) DEFAULT '' NOT NULL
);
*/

// Tls222ApplnJpClass is a structure representing a row of the tls222_appln_jp_class table in the database
// FI (File Index) and F-Terms are used by the Japanese patent office for classification
// (TLS222_APPLN_JP_CLASS)
type Tls222ApplnJpClass struct {
	ApplnID       int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	JpClassScheme string `json:"jpClassScheme" gorm:"primaryKey;column:jp_class_scheme;type:varchar(5);default:'';not null"`
	JpClassSymbol string `json:"jpClassSymbol" gorm:"primaryKey;column:jp_class_symbol;type:varchar(50);default:'';not null"`
}

func (m *Tls222ApplnJpClass) TableName() string {
	return "tls222_appln_jp_class"
}
