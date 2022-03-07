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

type Tls225DocdbFamCpc struct {
	DocdbFamilyID  int       `json:"docdbFamilyId" gorm:"column:docdb_family_id;type:integer;default:0;not null"`
	CpcClassSymbol string    `json:"cpcClassSymbol" gorm:"column:cpc_class_symbol;type:varchar(19);default:'';not null"`
	CpcGenerAuth   string    `json:"cpcGenerAuth" gorm:"column:cpc_gener_auth;type:varchar(2);default:'';not null"`
	CpcVersion     time.Time `json:"cpcVersion" gorm:"column:cpc_version;type:date;default:'9999-12-31';not null"`
	CpcPosition    string    `json:"cpcPosition" gorm:"column:cpc_position;type:char(1);default:'';not null"`
	CpcValue       string    `json:"cpcValue" gorm:"column:cpc_value;type:char(1);default:'';not null"`
	CpcActionDate  time.Time `json:"cpcActionDate" gorm:"column:cpc_action_date;type:date;default:'9999-12-31';not null"`
	CpcStatus      string    `json:"cpcStatus" gorm:"column:cpc_status;type:char(1);default:'';not null"`
	CpcDataSource  string    `json:"cpcDataSource" gorm:"column:cpc_data_source;type:char(1);default:'';not null"`
}

func (m *Tls225DocdbFamCpc) TableName() string {
	return "tls225_docdb_fam_cpc"
}
