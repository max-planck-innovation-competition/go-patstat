package models

/*
CREATE TABLE tls216_appln_contn (
    appln_id integer DEFAULT 0 NOT NULL,
    parent_appln_id integer DEFAULT 0 NOT NULL,
    contn_type char(3) DEFAULT '' NOT NULL
);
*/

type Tls216ApplnContn struct {
	ApplnID       int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	ParentApplnID int    `json:"parentApplnId" gorm:"primaryKey;column:parent_appln_id;type:integer;default:0;not null"`
	ContnType     string `json:"contnType" gorm:"column:contn_type;type:char(3);default:'';not null"`
}

func (m *Tls216ApplnContn) TableName() string {
	return "tls216_appln_contn"
}
