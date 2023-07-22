package models

/*
CREATE TABLE tls230_appln_techn_field (
    appln_id integer DEFAULT 0 NOT NULL,
    techn_field_nr smallint DEFAULT 0 NOT NULL,
    weight integer DEFAULT 1 NOT NULL
);
*/

// Tls230ApplnTechnField is a structure representing a row of the tls230_appln_techn_field table in the database
// This table tells to which degree an application belongs to one or more technical fields.
type Tls230ApplnTechnField struct {
	ApplnID      int   `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;not null"`
	TechnFieldNr int16 `json:"technFieldNr" gorm:"primaryKey;column:techn_field_nr;type:smallint;default:0;not null"`
	Weight       int   `json:"weight" gorm:"column:weight;type:integer;default:1;not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls230ApplnTechnField) TableName() string {
	return "tls230_appln_techn_field"
}
