package models

/*
CREATE TABLE tls215_citn_categ (
    pat_publn_id integer DEFAULT 0 NOT NULL,
    citn_replenished integer DEFAULT 0 NOT NULL,
    citn_id smallint DEFAULT 0 NOT NULL,
    citn_categ varchar(10) DEFAULT '' NOT NULL,
    relevant_claim smallint DEFAULT 0 NOT NULL
);
*/

// Tls215CitnCateg is a structure representing a row of the tls215_citn_categ table
// Each citation has one or more categories (TLS215_CITN_CATEG), which indicate the
// relevance of the citations. E. g. citation category “X” indicates that the claimed invention
// cannot be considered as novel due to the existence of the cited document.
type Tls215CitnCateg struct {
	PatPublnID      int    `json:"patPublnId" gorm:"primaryKey;column:pat_publn_id;type:integer;default:0;not null"`
	CitnReplenished int    `json:"citnReplenished" gorm:"primaryKey;column:citn_replenished;type:integer;default:0;not null"`
	CitnID          int16  `json:"citnId" gorm:"primaryKey;column:citn_id;type:smallint;default:0;not null"`
	CitnCateg       string `json:"citnCateg" gorm:"primaryKey;column:citn_categ;type:varchar(10);default:'';not null"`
	RelevantClaim   int16  `json:"relevantClaim" gorm:"primaryKey;column:relevant_claim;type:smallint;default:0;not null"`
}

func (m *Tls215CitnCateg) TableName() string {
	return "tls215_citn_categ"
}
