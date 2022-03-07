package models

/*
CREATE TABLE tls212_citation (
    pat_publn_id integer DEFAULT 0 NOT NULL,
    citn_replenished integer DEFAULT 0 NOT NULL,
    citn_id smallint DEFAULT 0 NOT NULL,
    citn_origin char(3) DEFAULT '' NOT NULL,
    cited_pat_publn_id integer DEFAULT 0 NOT NULL,
    cited_appln_id integer DEFAULT 0 NOT NULL,
    pat_citn_seq_nr smallint DEFAULT 0 NOT NULL,
    cited_npl_publn_id varchar(32) DEFAULT 0 NOT NULL,
    npl_citn_seq_nr smallint DEFAULT 0 NOT NULL,
    citn_gener_auth char(2) DEFAULT '' NOT NULL
);
*/

type Tls212Citation struct {
	PatPublnID      int    `json:"patPublnId" gorm:"primaryKey;column:pat_publn_id;type:integer;default:0;not null"`
	CitnReplenished int    `json:"citnReplenished" gorm:"column:citn_replenished;type:integer;default:0;not null"`
	CitnID          int16  `json:"citnId" gorm:"column:citn_id;type:smallint;default:0;not null"`
	CitnOrigin      string `json:"citnOrigin" gorm:"column:citn_origin;type:char(3);default:'';not null"`
	CitedPatPublnID int    `json:"citedPatPublnId" gorm:"column:cited_pat_publn_id;type:integer;default:0;not null"`
	CitedApplnID    int    `json:"citedApplnId" gorm:"column:cited_appln_id;type:integer;default:0;not null"`
	PatCitnSeqNr    int16  `json:"patCitnSeqNr" gorm:"column:pat_citn_seq_nr;type:smallint;default:0;not null"`
	CitedNplPublnID string `json:"citedNplPublnId" gorm:"column:cited_npl_publn_id;type:varchar(32);default:0;not null"`
	NplCitnSeqNr    int16  `json:"nplCitnSeqNr" gorm:"column:npl_citn_seq_nr;type:smallint;default:0;not null"`
	CitnGenerAuth   string `json:"citnGenerAuth" gorm:"column:citn_gener_auth;type:char(2);default:'';not null"`
}

func (m *Tls212Citation) TableName() string {
	return "tls212_citation"
}
