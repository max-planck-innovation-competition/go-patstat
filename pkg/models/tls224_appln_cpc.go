package models

/*
CREATE TABLE tls224_appln_cpc (
    appln_id integer DEFAULT 0 NOT NULL,
    cpc_class_symbol varchar(19) DEFAULT '' NOT NULL
);
*/

type Tls224ApplnCpc struct {
	ApplnID        int    `json:"applnId" gorm:"column:appln_id;type:integer;default:0;not null"`
	CpcClassSymbol string `json:"cpcClassSymbol" gorm:"column:cpc_class_symbol;type:varchar(19);default:'';not null"`
}

func (m *Tls224ApplnCpc) TableName() string {
	return "tls224_appln_cpc"
}
