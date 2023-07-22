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

// Tls209ApplnIpc is a structure representing ipc classification for an application
// IPC is the International Patent Classification, which is maintained by WIPO and used by all patent offices
//
// The table contains all international patent classifications linked to the applications. The set
// of classifications linked to a single application is a de-duplicated merge of all classifications
// of the various publication instances linked to the specific application. Additionally, only the
// latest version of the IPC classifications is used. This means that the user does not have to
// worry about reclassifications because older applications will always be classified according
// to the latest IPC version.
type Tls209ApplnIpc struct {
	ApplnID        int       `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;not null"`
	IpcClassSymbol string    `json:"ipcClassSymbol" gorm:"primaryKey;column:ipc_class_symbol;type:varchar(15);default:'';not null"`
	IpcClassLevel  string    `json:"ipcClassLevel" gorm:"column:ipc_class_level;type:char(1);default:'';not null"`
	IpcVersion     time.Time `json:"ipcVersion" gorm:"column:ipc_version;type:date;default:'9999-12-31';not null"`
	IpcValue       string    `json:"ipcValue" gorm:"column:ipc_value;type:char(1);default:'';not null"`
	IpcPosition    string    `json:"ipcPosition" gorm:"column:ipc_position;type:char(1);default:'';not null"`
	IpcGenerAuth   string    `json:"ipcGenerAuth" gorm:"column:ipc_gener_auth;type:char(2);default:'';not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls209ApplnIpc) TableName() string {
	return "tls209_appln_ipc"
}
