package models

import "time"

/*
CREATE TABLE tls211_pat_publn (
    pat_publn_id integer DEFAULT 0 NOT NULL,
    publn_auth char(2) DEFAULT '' NOT NULL,
    publn_nr varchar(15) DEFAULT '' NOT NULL,
    publn_nr_original varchar(100) DEFAULT '' NOT NULL,
    publn_kind char(2) DEFAULT '' NOT NULL,
    appln_id integer,
    publn_date date DEFAULT '9999-12-31' NOT NULL,
    publn_lg char(2) DEFAULT '' NOT NULL,
    publn_first_grant char(1) DEFAULT '' NOT NULL,
    publn_claims integer DEFAULT 0 NOT NULL
);
*/

// Tls211PatPubln is a structure representing a row from the 'tls211_pat_publn' table
// This table contains the key bibliographical data elements relevant to identify patent
// publications. These elements can be found on the first page of printed patent documents.
// For example: publication authority, publication number, publication kind and publication
// date. This table is directly linked to the TLS201_APPLN table via the appln_id, a surrogate
// key that groups all the data elements from a single patent application. 2 important extra
// elements that cannot be found on a patent publication have been added:
// a) The PUBLN_FIRST_GRANT:
// indication that this publication was the first indication of a patent grant
// b) PUBLN_CLAIMS:
// number of claims (only available for a number of publishing authorities).
type Tls211PatPubln struct {
	PatPublnID      int       `json:"patPublnId" gorm:"primaryKey;column:pat_publn_id;type:integer;default:0;not null"`
	PublnAuth       string    `json:"publnAuth" gorm:"column:publn_auth;type:char(2);default:'';not null"`
	PublnNr         string    `json:"publnNr" gorm:"column:publn_nr;type:varchar(15);default:'';not null"`
	PublnNrOriginal string    `json:"publnNrOriginal" gorm:"column:publn_nr_original;type:varchar(100);default:'';not null"`
	PublnKind       string    `json:"publnKind" gorm:"column:publn_kind;type:char(2);default:'';not null"`
	ApplnID         int       `json:"applnId"  gorm:"column:appln_id;type:integer"`
	PublnDate       time.Time `json:"publnDate" gorm:"column:publn_date;type:date;default:'9999-12-31';not null"`
	PublnLg         string    `json:"publnLg" gorm:"column:publn_lg;type:char(2);default:'';not null"`
	PublnFirstGrant string    `json:"publnFirstGrant" gorm:"column:publn_first_grant;type:char(1);default:'';not null"`
	PublnClaims     int       `json:"publnClaims" gorm:"column:publn_claims;type:integer;default:0;not null"`
	// relations
	Application *Tls201Appln      `json:"application" gorm:"foreignKey:appln_id"`
	Persons     []*Tls206Person   `json:"persons" gorm:"many2many:tls227_pers_publn;foreignKey:pat_publn_id;joinForeignKey:person_id;"`
	Citations   []*Tls212Citation `json:"citations" gorm:"foreignKey:pat_publn_id"`
	Cited       []*Tls212Citation `json:"cited" gorm:"foreignKey:cited_pat_publn_id"`
}

func (m *Tls211PatPubln) TableName() string {
	return "tls211_pat_publn"
}
