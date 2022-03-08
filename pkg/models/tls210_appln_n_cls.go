package models

/*
CREATE TABLE tls210_appln_n_cls (
    appln_id integer DEFAULT 0 NOT NULL,
    nat_class_symbol varchar(15) DEFAULT '' NOT NULL
);
*/

// Tls210ApplnNCls is a structure representing a record of the tls210_appln_n_cls table in the database.
// In the past, some offices have used their own national classification system
// (TLS210_APPLN_N_CLS)
type Tls210ApplnNCls struct {
	ApplnID        int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	NatClassSymbol string `json:"natClassSymbol" gorm:"primaryKey;column:nat_class_symbol;type:varchar(15);default:'';not null"`
}

func (m *Tls210ApplnNCls) TableName() string {
	return "tls210_appln_n_cls"
}
