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

type Tls215CitnCateg struct {
	PatPublnID      int    `json:"patPublnId" gorm:"column:pat_publn_id;type:integer;default:0;not null"`
	CitnReplenished int    `json:"citnReplenished" gorm:"column:citn_replenished;type:integer;default:0;not null"`
	CitnID          int16  `json:"citnId" gorm:"column:citn_id;type:smallint;default:0;not null"`
	CitnCateg       string `json:"citnCateg" gorm:"column:citn_categ;type:varchar(10);default:'';not null"`
	RelevantClaim   int16  `json:"relevantClaim" gorm:"column:relevant_claim;type:smallint;default:0;not null"`
}

func (m *Tls215CitnCateg) TableName() string {
	return "tls215_citn_categ"
}
