package models

/*
CREATE TABLE tls207_pers_appln (
    person_id integer DEFAULT 0 NOT NULL,
    appln_id integer DEFAULT 0 NOT NULL,
    applt_seq_nr smallint DEFAULT 0 NOT NULL,
    invt_seq_nr smallint DEFAULT 0 NOT NULL
);
*/

type Tls207PersAppln struct {
	PersonID   int   `json:"personId" gorm:"primaryKey;column:person_id;type:integer;default:0;not null"`
	ApplnID    int   `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	AppltSeqNr int16 `json:"appltSeqNr" gorm:"primaryKey;column:applt_seq_nr;type:smallint;default:0;not null"`
	InvtSeqNr  int16 `json:"invtSeqNr" gorm:"primaryKey;column:invt_seq_nr;type:smallint;default:0;not null"`
}

func (m *Tls207PersAppln) TableName() string {
	return "tls207_pers_appln"
}
