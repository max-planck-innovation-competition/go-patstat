package models

/*
CREATE TABLE tls205_tech_rel (
    appln_id integer DEFAULT 0 NOT NULL,
    tech_rel_appln_id integer DEFAULT 0 NOT NULL
);
*/

// Tls205TechRel is a structure representing a row of the tls205_tech_rel table in the database
// Technical relations (TLS205_TECH_REL)
// Technical relations are "priority-like" relations between applications which have been
// detected by EPO examiners but which have not been published by a patent office. From a
// statistical point of view you should consider them equal to the priority and continuation
// relations established in TLS204_APPLN_PRIOR and in TLS216_APPLN_CONTN.
type Tls205TechRel struct {
	ApplnID        int `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;not null"`
	TechRelApplnID int `json:"techRelApplnId" gorm:"primaryKey;column:tech_rel_appln_id;type:integer;default:0;not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls205TechRel) TableName() string {
	return "tls205_tech_rel"
}
