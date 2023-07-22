package models

/*
CREATE TABLE tls207_pers_appln (
    person_id integer DEFAULT 0 NOT NULL,
    appln_id integer DEFAULT 0 NOT NULL,
    applt_seq_nr smallint DEFAULT 0 NOT NULL,
    invt_seq_nr smallint DEFAULT 0 NOT NULL
);
*/

// Tls207PersAppln is a structure representing a row of data from the 'tls207_pers_appln' table
// This table links the applicants and inventors of the most recent publication to an
// application. Publications which contain only persons with non-Latin names
// (e.g. with Chinese characters) are not considered here.
type Tls207PersAppln struct {
	PersonID   int   `json:"personId" gorm:"primaryKey;column:person_id;type:integer;not null"`
	ApplnID    int   `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	AppltSeqNr int16 `json:"appltSeqNr" gorm:"primaryKey;column:applt_seq_nr;type:smallint;default:0;not null"`
	InvtSeqNr  int16 `json:"invtSeqNr" gorm:"primaryKey;column:invt_seq_nr;type:smallint;default:0;not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls207PersAppln) TableName() string {
	return "tls207_pers_appln"
}
