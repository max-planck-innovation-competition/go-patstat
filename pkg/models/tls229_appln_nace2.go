package models

/*
CREATE TABLE tls229_appln_nace2 (
    appln_id integer DEFAULT 0 NOT NULL,
    nace2_code varchar(5) DEFAULT '' NOT NULL,
    weight numeric DEFAULT 1 NOT NULL
);

*/

// Tls229ApplnNace2 is a structure representing a row of the tls229_appln_nace2 table in the database
// This table tells to which degree an application belongs to one or more industries.
type Tls229ApplnNace2 struct {
	ApplnID   int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;not null"`
	Nace2Code string `json:"nace2Code" gorm:"primaryKey;column:nace2_code;type:varchar(5);default:'';not null"`
	Weight    int    `json:"weight" gorm:"column:weight;type:integer;default:1;not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls229ApplnNace2) TableName() string {
	return "tls229_appln_nace2"
}
