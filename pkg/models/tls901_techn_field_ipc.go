package models

/*
CREATE TABLE tls901_techn_field_ipc (
    ipc_maingroup_symbol varchar(8) DEFAULT '' NOT NULL,
    techn_field_nr smallint DEFAULT 0 NOT NULL,
    techn_sector varchar(50) DEFAULT '' NOT NULL,
    techn_field varchar(50) DEFAULT '' NOT NULL
);
*/

// Tls901TechnFieldIpc is a structure representing a row of the tls901_techn_field_ipc table in the database
// This is the reference table which contains the mapping between 35 technology fields and
// the much more detailed IPC classification. These technology fields allow for the easy
// grouping of applications based on technology. The same technology fields are used by EPO
// and WIPO for their statistics.
type Tls901TechnFieldIpc struct {
	IpcMaingroupSymbol string `json:"ipcMaingroupSymbol" gorm:"primaryKey;column:ipc_maingroup_symbol;type:varchar(8);default:'';not null"`
	TechnFieldNr       int16  `json:"technFieldNr" gorm:"column:techn_field_nr;type:smallint;default:0;not null"`
	TechnSector        string `json:"technSector" gorm:"column:techn_sector;type:varchar(50);default:'';not null"`
	TechnField         string `json:"technField" gorm:"column:techn_field;type:varchar(50);default:'';not null"`
}

func (m *Tls901TechnFieldIpc) TableName() string {
	return "tls901_techn_field_ipc"
}
