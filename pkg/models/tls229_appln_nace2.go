package models

/*
CREATE TABLE tls229_appln_nace2 (
    appln_id integer DEFAULT 0 NOT NULL,
    nace2_code varchar(5) DEFAULT '' NOT NULL,
    weight numeric DEFAULT 1 NOT NULL
);

*/

type Tls229ApplnNace2 struct {
	ApplnID   int    `json:"applnId" gorm:"column:appln_id;type:integer;default:0;not null"`
	Nace2Code string `json:"nace2Code" gorm:"column:nace2_code;type:varchar(5);default:'';not null"`
	Weight    int    `json:"weight" gorm:"column:weight;type:integer;default:1;not null"`
}

func (m *Tls229ApplnNace2) TableName() string {
	return "tls229_appln_nace2"
}
