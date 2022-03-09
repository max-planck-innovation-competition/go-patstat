package models

import "time"

/*
CREATE TABLE tls225_docdb_fam_cpc (
    docdb_family_id integer DEFAULT 0 NOT NULL,
    cpc_class_symbol varchar(19) DEFAULT '' NOT NULL,
    cpc_gener_auth varchar(2) DEFAULT '' NOT NULL,
    cpc_version date DEFAULT '9999-12-31' NOT NULL,
    cpc_position char(1) DEFAULT '' NOT NULL,
    cpc_value char(1) DEFAULT '' NOT NULL,
    cpc_action_date date DEFAULT '9999-12-31' NOT NULL,
    cpc_status char(1) DEFAULT '' NOT NULL,
    cpc_data_source char(1) DEFAULT '' NOT NULL
);
*/

// Tls225DocdbFamCpc is the structure for tls225_docdb_fam_cpc table.
// CPC – Cooperative Patent Classification in an extension of IPC.
// It has been created in 2013 and is maintained by EPO and the US patent office.
// Many major offices are nowadays using CPC, in addition to IPC.
// CPC symbols are assigned on a family level (TLS225_DOCDB_FAM_CPC).
//
// TLS225_DOCDB_FAM_CPC: Cooperative Patent Classification by DOCDB family
// All applications of the same DOCDB simple family have the same cooperative patent
// classifications (CPC symbols) assigned. The same CPC symbol can be assigned to the
// same DOCDB family by one or more patent offices.
// This table contains detailed information for each assignment.
//
// All applications of the same DOCDB simple family have the same cooperative patent
// classifications (CPC symbols) assigned. The same CPC symbol can be assigned to the
// same DOCDB family by one or more patent offices.
// This table contains detailed information for each assignment.

type Tls225DocdbFamCpc struct {
	DocdbFamilyID  int       `json:"docdbFamilyId" gorm:"primaryKey;column:docdb_family_id;type:integer;default:0;not null"`
	CpcClassSymbol string    `json:"cpcClassSymbol" gorm:"primaryKey;column:cpc_class_symbol;type:varchar(19);default:'';not null"`
	CpcGenerAuth   string    `json:"cpcGenerAuth" gorm:"primaryKey;column:cpc_gener_auth;type:varchar(2);default:'';not null"`
	CpcVersion     time.Time `json:"cpcVersion" gorm:"column:cpc_version;type:date;default:'9999-12-31';not null"`
	CpcPosition    string    `json:"cpcPosition" gorm:"column:cpc_position;type:char(1);default:'';not null"`
	CpcValue       string    `json:"cpcValue" gorm:"column:cpc_value;type:char(1);default:'';not null"`
	CpcActionDate  time.Time `json:"cpcActionDate" gorm:"column:cpc_action_date;type:date;default:'9999-12-31';not null"`
	CpcStatus      string    `json:"cpcStatus" gorm:"column:cpc_status;type:char(1);default:'';not null"`
	CpcDataSource  string    `json:"cpcDataSource" gorm:"column:cpc_data_source;type:char(1);default:'';not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls225DocdbFamCpc) TableName() string {
	return "tls225_docdb_fam_cpc"
}
