package models

/*
CREATE TABLE tls902_ipc_nace2 (
    ipc varchar(8) DEFAULT '' NOT NULL,
    not_with_ipc varchar(8) DEFAULT '' NOT NULL,
    unless_with_ipc varchar(8) DEFAULT '' NOT NULL,
    nace2_code varchar(5) DEFAULT '' NOT NULL,
    nace2_weight smallint DEFAULT 1 NOT NULL,
    nace2_descr varchar(150) DEFAULT '' NOT NULL
);
*/

// Tls902IpcNace2 is a structure representing a row of data from the 'tls902_ipc_nace2' table.
// A reference table which contains the mapping between the IPC classification and the
// NACE2 codes for industrial sectors. The industrial sectors allow for the grouping of
// applications based on the industry.
type Tls902IpcNace2 struct {
	Ipc           string `json:"ipc" gorm:"primaryKey;column:ipc;type:varchar(8);default:'';not null"`
	NotWithIpc    string `json:"notWithIpc" gorm:"primaryKey;column:not_with_ipc;type:varchar(8);default:'';not null"`
	UnlessWithIpc string `json:"unlessWithIpc" gorm:"primaryKey;column:unless_with_ipc;type:varchar(8);default:'';not null"`
	Nace2Code     string `json:"nace2Code" gorm:"primaryKey;column:nace2_code;type:varchar(5);default:'';not null"`
	Nace2Weight   int16  `json:"nace2Weight" gorm:"column:nace2_weight;type:smallint;default:1;not null"`
	Nace2Descr    string `json:"nace2Descr" gorm:"column:nace2_descr;type:varchar(150);default:'';not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls902IpcNace2) TableName() string {
	return "tls902_ipc_nace2"
}
