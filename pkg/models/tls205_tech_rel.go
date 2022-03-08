package models

/*
CREATE TABLE tls205_tech_rel (
    appln_id integer DEFAULT 0 NOT NULL,
    tech_rel_appln_id integer DEFAULT 0 NOT NULL
);
*/

// Tls205TechRel is a structure representing a row of the tls205_tech_rel table in the database
// Technical relations (TLS205_TECH_REL)
type Tls205TechRel struct {
	ApplnID        int `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	TechRelApplnID int `json:"techRelApplnId" gorm:"primaryKey;column:tech_rel_appln_id;type:integer;default:0;not null"`
}

func (m *Tls205TechRel) TableName() string {
	return "tls205_tech_rel"
}
