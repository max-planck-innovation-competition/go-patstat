package models

/*
CREATE TABLE tls204_appln_prior (
    appln_id integer DEFAULT 0 NOT NULL,
    prior_appln_id integer DEFAULT 0 NOT NULL,
    prior_appln_seq_nr smallint DEFAULT 0 NOT NULL
);
*/

// Tls204ApplnPrior is a table generated from tls204_appln_prior.
// Priorities (TLS204_APPLN_PRIOR)
type Tls204ApplnPrior struct {
	ApplnID         int   `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	PriorApplnID    int   `json:"priorApplnId" gorm:"primaryKey;column:prior_appln_id;type:integer;default:0;not null"`
	PriorApplnSeqNr int16 `json:"priorApplnSeqNr" gorm:"column:prior_appln_seq_nr;type:smallint;default:0;not null"`
}

func (m *Tls204ApplnPrior) TableName() string {
	return "tls204_appln_prior"
}
