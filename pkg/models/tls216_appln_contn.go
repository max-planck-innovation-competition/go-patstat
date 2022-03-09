package models

/*
CREATE TABLE tls216_appln_contn (
    appln_id integer DEFAULT 0 NOT NULL,
    parent_appln_id integer DEFAULT 0 NOT NULL,
    contn_type char(3) DEFAULT '' NOT NULL
);
*/

// Tls216ApplnContn is a structure representing a tls216_appln_contn table row
// Continuations (TLS216_APPLN_CONTN)
// In a similar way as the TLS204_APPLN_PRIOR establishes the priority links between
// applications , the links between parent and child applications for various types relations
// such as: continuation (in part), divisional applications, internal priorities are defined via the
// TLS216_APPLN_CONTN table. Continuation (in part) is generally only applicable to US
// patent applications. This table should be considered as a priority-like relationship similar to
// the TLS204_APPLN_PRIOR table.
type Tls216ApplnContn struct {
	ApplnID       int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	ParentApplnID int    `json:"parentApplnId" gorm:"primaryKey;column:parent_appln_id;type:integer;default:0;not null"`
	ContnType     string `json:"contnType" gorm:"column:contn_type;type:char(3);default:'';not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls216ApplnContn) TableName() string {
	return "tls216_appln_contn"
}
