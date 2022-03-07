package models

/*
CREATE TABLE tls205_tech_rel (
    appln_id integer DEFAULT 0 NOT NULL,
    tech_rel_appln_id integer DEFAULT 0 NOT NULL
);
*/

type Tls205TechRel struct {
	ApplnID        int `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	TechRelApplnID int `json:"techRelApplnId" gorm:"primaryKey;column:tech_rel_appln_id;type:integer;default:0;not null"`
}

func (m *Tls205TechRel) TableName() string {
	return "tls205_tech_rel"
}
