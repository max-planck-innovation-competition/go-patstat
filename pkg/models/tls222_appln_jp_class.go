package models

/*
CREATE TABLE tls222_appln_jp_class (
    appln_id integer DEFAULT 0 NOT NULL,
    jp_class_scheme varchar(5) DEFAULT '' NOT NULL,
    jp_class_symbol varchar(50) DEFAULT '' NOT NULL
);
*/

type Tls222ApplnJpClass struct {
	ApplnID       int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	JpClassScheme string `json:"jpClassScheme" gorm:"primaryKey;column:jp_class_scheme;type:varchar(5);default:'';not null"`
	JpClassSymbol string `json:"jpClassSymbol" gorm:"primaryKey;column:jp_class_symbol;type:varchar(50);default:'';not null"`
}

func (m *Tls222ApplnJpClass) TableName() string {
	return "tls222_appln_jp_class"
}
