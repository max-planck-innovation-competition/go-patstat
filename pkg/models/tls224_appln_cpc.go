package models

/*
CREATE TABLE tls224_appln_cpc (
    appln_id integer DEFAULT 0 NOT NULL,
    cpc_class_symbol varchar(19) DEFAULT '' NOT NULL
);
*/

// Tls224ApplnCpc is the structure of the tls224_appln_cpc table
// CPC – Cooperative Patent Classification in an extension of IPC.
// It has been created in 2013 and is maintained by EPO and the US patent office.
// Many major offices are nowadays using CPC, in addition to IPC.
// CPC symbols are assigned on a family level (TLS225_DOCDB_FAM_CPC).
//
// The table contains for each application its assigned cooperative patent classifications
// (CPC symbols). All applications of the same DOCDB family have the same CPC symbols
// assigned.
type Tls224ApplnCpc struct {
	ApplnID        int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	CpcClassSymbol string `json:"cpcClassSymbol" gorm:"primaryKey;column:cpc_class_symbol;type:varchar(19);default:'';not null"`
}

func (m *Tls224ApplnCpc) TableName() string {
	return "tls224_appln_cpc"
}
