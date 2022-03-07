package models

/*
CREATE TABLE tls230_appln_techn_field (
    appln_id integer DEFAULT 0 NOT NULL,
    techn_field_nr smallint DEFAULT 0 NOT NULL,
    weight integer DEFAULT 1 NOT NULL
);
*/

type Tls230ApplnTechnField struct {
	ApplnID      int   `json:"applnId" gorm:"column:appln_id;type:integer;default:0;not null"`
	TechnFieldNr int16 `json:"technFieldNr" gorm:"column:techn_field_nr;type:smallint;default:0;not null"`
	Weight       int   `json:"weight" gorm:"column:weight;type:integer;default:1;not null"`
}

func (m *Tls230ApplnTechnField) TableName() string {
	return "tls230_appln_techn_field"
}
