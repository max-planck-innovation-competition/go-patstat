package models

/*
CREATE TABLE tls227_pers_publn (
    person_id integer DEFAULT 0 NOT NULL,
    pat_publn_id integer DEFAULT 0 NOT NULL,
    applt_seq_nr integer DEFAULT 0 NOT NULL,
    invt_seq_nr integer DEFAULT 0 NOT NULL
);
*/

type Tls227PersPubln struct {
	PersonID   int `json:"personId" gorm:"primaryKey;column:person_id;type:integer;default:0;not null"`
	PatPublnID int `json:"patPublnId" gorm:"primaryKey;column:pat_publn_id;type:integer;default:0;not null"`
	AppltSeqNr int `json:"appltSeqNr" gorm:"primaryKey;column:applt_seq_nr;type:integer;default:0;not null"`
	InvtSeqNr  int `json:"invtSeqNr" gorm:"primaryKey;column:invt_seq_nr;type:integer;default:0;not null"`
	// relations
	Person Tls206Person `json:"person" gorm:"foreignKey:person_id;"`
}

func (m *Tls227PersPubln) TableName() string {
	return "tls227_pers_publn"
}
