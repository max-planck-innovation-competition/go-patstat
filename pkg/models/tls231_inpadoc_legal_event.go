package models

import "time"

/*
CREATE TABLE tls231_inpadoc_legal_event (
    event_id integer DEFAULT 0 NOT NULL,
    appln_id integer DEFAULT 0 NOT NULL,
    event_seq_nr smallint DEFAULT 0,
    event_type char(3) DEFAULT '',
    event_auth char(2) DEFAULT '',
    event_code varchar(4)  DEFAULT '',
    event_filing_date date DEFAULT '9999-12-31' NOT NULL,
    event_publn_date date DEFAULT '9999-12-31' NOT NULL,
    event_effective_date date DEFAULT '9999-12-31' NOT NULL,
    event_text text DEFAULT '',
    ref_doc_auth char(2) DEFAULT '',
    ref_doc_nr varchar(20) DEFAULT '',
    ref_doc_kind char(2) DEFAULT '',
    ref_doc_date date DEFAULT '9999-12-31' NOT NULL,
    ref_doc_text text DEFAULT '',
    party_type varchar(3) DEFAULT '',
    party_seq_nr smallint default 0,
    party_new text DEFAULT '',
    party_old text DEFAULT '',
    spc_nr varchar(40) DEFAULT '',
    spc_filing_date date DEFAULT '9999-12-31' NOT NULL,
    spc_patent_expiry_date date DEFAULT '9999-12-31' NOT NULL,
    spc_extension_date date DEFAULT '9999-12-31' NOT NULL,
    spc_text text DEFAULT '',
    designated_states text DEFAULT '',
    extension_states varchar(30) DEFAULT '',
    fee_country char(2) DEFAULT '',
    fee_payment_date date DEFAULT '9999-12-31' NOT NULL,
    fee_renewal_year smallint DEFAULT 9999 NOT NULL,
    fee_text text DEFAULT '',
    lapse_country char(2) DEFAULT '',
    lapse_date date  DEFAULT '9999-12-31' NOT NULL,
    lapse_text text DEFAULT '',
    reinstate_country char(2) DEFAULT '',
    reinstate_date date DEFAULT '9999-12-31' NOT NULL,
    reinstate_text text DEFAULT '',
    class_scheme varchar(4) DEFAULT '',
    class_symbol varchar(50) DEFAULT ''
);
*/

// Tls231InpadocLegalEvent is a structure representing a row in the tls231_inpadoc_legal_event table
// This table holds the INPADOC data, which contains information on legal events that
// occurred during the life of a patent, either before or after grant. Typical events are request
// for examination, payment of renewal fees, lapse of the patent, change of ownership,
// withdrawal of the application, patent applications entering the national phase, patents which
// have been opposed or revoked, etc.
type Tls231InpadocLegalEvent struct {
	EventID             int       `json:"eventId" gorm:"primaryKey;column:event_id;type:integer;default:0;not null"`
	ApplnID             int       `json:"applnId" gorm:"column:appln_id;type:integer;default:0;not null"`
	EventSeqNr          int16     `json:"eventSeqNr" gorm:"column:event_seq_nr;type:smallint;default:0;"`
	EventType           string    `json:"eventType" gorm:"column:event_type;type:char(3);default:'';"`
	EventAuth           string    `json:"eventAuth" gorm:"column:event_auth;type:char(2);default:'';"`
	EventCode           string    `json:"eventCode" gorm:"column:event_code;type:varchar(4);"`
	EventFilingDate     time.Time `json:"eventFilingDate" gorm:"column:event_filing_date;type:date;default:'9999-12-31';not null"`
	EventPublnDate      time.Time `json:"eventPublnDate" gorm:"column:event_publn_date;type:date;default:'9999-12-31';not null"`
	EventEffectiveDate  time.Time `json:"eventEffectiveDate" gorm:"column:event_effective_date;type:date;default:'9999-12-31';not null"`
	EventText           string    `json:"eventText" gorm:"column:event_text;type:text;default:'';"`
	RefDocAuth          string    `json:"refDocAuth" gorm:"column:ref_doc_auth;type:char(2);default:'';"`
	RefDocNr            string    `json:"refDocNr" gorm:"column:ref_doc_nr;type:varchar(20);default:'';"`
	RefDocKind          string    `json:"refDocKind" gorm:"column:ref_doc_kind;type:char(2);default:'';"`
	RefDocDate          time.Time `json:"refDocDate" gorm:"column:ref_doc_date;type:date;default:'9999-12-31';not null"`
	RefDocText          string    `json:"refDocText" gorm:"column:ref_doc_text;type:text;default:'';"`
	PartyType           string    `json:"partyType" gorm:"column:party_type;type:varchar(3);default:'';"`
	PartySeqNr          int16     `json:"partySeqNr" gorm:"column:party_seq_nr;type:smallint;"`
	PartyNew            string    `json:"partyNew" gorm:"column:party_new;type:text;default:'';"`
	PartyOld            string    `json:"partyOld" gorm:"column:party_old;type:text;default:'';"`
	SpcNr               string    `json:"spcNr" gorm:"column:spc_nr;type:varchar(40);default:'';"`
	SpcFilingDate       time.Time `json:"spcFilingDate" gorm:"column:spc_filing_date;type:date;default:'9999-12-31';not null"`
	SpcPatentExpiryDate time.Time `json:"spcPatentExpiryDate" gorm:"column:spc_patent_expiry_date;type:date;default:'9999-12-31';not null"`
	SpcExtensionDate    time.Time `json:"spcExtensionDate" gorm:"column:spc_extension_date;type:date;default:'9999-12-31';not null"`
	SpcText             string    `json:"spcText" gorm:"column:spc_text;type:text;default:'';"`
	DesignatedStates    string    `json:"designatedStates" gorm:"column:designated_states;type:text;default:'';"`
	ExtensionStates     string    `json:"extensionStates" gorm:"column:extension_states;type:varchar(30);default:'';"`
	FeeCountry          string    `json:"feeCountry" gorm:"column:fee_country;type:char(2);default:'';"`
	FeePaymentDate      time.Time `json:"feePaymentDate" gorm:"column:fee_payment_date;type:date;default:'9999-12-31';not null"`
	FeeRenewalYear      int16     `json:"feeRenewalYear" gorm:"column:fee_renewal_year;type:smallint;default:9999;not null"`
	FeeText             string    `json:"feeText" gorm:"column:fee_text;type:text;default:'';"`
	LapseCountry        string    `json:"lapseCountry" gorm:"column:lapse_country;type:char(2);default:'';"`
	LapseDate           time.Time `json:"lapseDate" gorm:"column:lapse_date;type:date;"`
	LapseText           string    `json:"lapseText" gorm:"column:lapse_text;type:text;default:'';"`
	ReinstateCountry    string    `json:"reinstateCountry" gorm:"column:reinstate_country;type:char(2);default:'';"`
	ReinstateDate       time.Time `json:"reinstateDate" gorm:"column:reinstate_date;type:date;default:'9999-12-31';not null"`
	ReinstateText       string    `json:"reinstateText" gorm:"column:reinstate_text;type:text;default:'';"`
	ClassScheme         string    `json:"classScheme" gorm:"column:class_scheme;type:varchar(4);default:'';"`
	ClassSymbol         string    `json:"classSymbol" gorm:"column:class_symbol;type:varchar(50);default:'';"`
	// relations
	LegalEventCode *Tls803LegalEventCode `json:"legalEventCode" gorm:"foreignKey:event_code,event_auth"`
}

func (m *Tls231InpadocLegalEvent) TableName() string {
	return "tls231_inpadoc_legal_event"
}
