package models

import "time"

/*
CREATE TABLE tls209_appln_ipc (
    appln_id integer DEFAULT 0 NOT NULL,
    ipc_class_symbol varchar(15) DEFAULT '' NOT NULL,
    ipc_class_level char(1) DEFAULT '' NOT NULL,
    ipc_version date DEFAULT '9999-12-31' NOT NULL,
    ipc_value char(1) DEFAULT '' NOT NULL,
    ipc_position char(1) DEFAULT '' NOT NULL,
    ipc_gener_auth char(2) DEFAULT '' NOT NULL
);
*/

type Tls209ApplnIpc struct {
	ApplnID        int       `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	IpcClassSymbol string    `json:"ipcClassSymbol" gorm:"column:ipc_class_symbol;type:varchar(15);default:'';not null"`
	IpcClassLevel  string    `json:"ipcClassLevel" gorm:"column:ipc_class_level;type:char(1);default:'';not null"`
	IpcVersion     time.Time `json:"ipcVersion" gorm:"column:ipc_version;type:date;default:'9999-12-31';not null"`
	IpcValue       string    `json:"ipcValue" gorm:"column:ipc_value;type:char(1);default:'';not null"`
	IpcPosition    string    `json:"ipcPosition" gorm:"column:ipc_position;type:char(1);default:'';not null"`
	IpcGenerAuth   string    `json:"ipcGenerAuth" gorm:"column:ipc_gener_auth;type:char(2);default:'';not null"`
}

func (m *Tls209ApplnIpc) TableName() string {
	return "tls209_appln_ipc"
}
