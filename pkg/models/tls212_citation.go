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

// Tls212Citation Citations (TLS212_CITATION) are references from patent publications to documents which
// are regarded as relevant for the patent procedure. They are identified in various stages in that
// procedure by various roles: by the applicant before application, during search and examination
// by the patent office, during an opposition procedure, by a third party etc.
// Patent publications typically cite other patent publications or non-patent literature; in less
// frequent cases applications are also cited.
// Each citation has one or more categories (TLS215_CITN_CATEG), which indicate the
// relevance of the citations. E. g. citation category “X” indicates that the claimed invention
// cannot be considered as novel due to the existence of the cited document.
//
// This table establishes the links between publications, applications and non-patent
// literature documents with regards to citations. Forward and backward citations are defined
// as well as the citation generating authority (e.g. search authority) and the procedural step
// in which the citation was created (e.g. search report or opposition procedure).
type Tls212Citation struct {
	PatPublnID      int    `json:"patPublnId" gorm:"primaryKey;column:pat_publn_id;type:integer;not null"`
	CitnReplenished int    `json:"citnReplenished" gorm:"primaryKey;column:citn_replenished;type:integer;default:0;not null"`
	CitnID          int16  `json:"citnId" gorm:"primaryKey;column:citn_id;type:smallint;default:0;not null"`
	CitnOrigin      string `json:"citnOrigin" gorm:"column:citn_origin;type:char(3);default:'';not null"`
	CitedPatPublnID int    `json:"citedPatPublnId" gorm:"column:cited_pat_publn_id;type:integer;default:0;not null"`
	CitedApplnID    int    `json:"citedApplnId" gorm:"column:cited_appln_id;type:integer;default:0;not null"`
	PatCitnSeqNr    int16  `json:"patCitnSeqNr" gorm:"column:pat_citn_seq_nr;type:smallint;default:0;not null"`
	CitedNplPublnID string `json:"citedNplPublnId" gorm:"column:cited_npl_publn_id;type:varchar(32);default:0;not null"`
	NplCitnSeqNr    int16  `json:"nplCitnSeqNr" gorm:"column:npl_citn_seq_nr;type:smallint;default:0;not null"`
	CitnGenerAuth   string `json:"citnGenerAuth" gorm:"column:citn_gener_auth;type:char(2);default:'';not null"`
	// relations
	Category *Tls215CitnCateg `json:"category" gorm:"foreignKey:pat_publn_id,citn_replenished,citn_id;references:pat_publn_id,citn_replenished,citn_id"`
}

// TableName sets the sql table name for this struct type
func (m *Tls212Citation) TableName() string {
	return "tls212_citation"
}
