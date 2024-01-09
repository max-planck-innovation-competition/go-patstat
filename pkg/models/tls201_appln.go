package models

import (
	"time"
)

/*
CREATE TABLE tls201_appln (
    appln_id integer DEFAULT 0 NOT NULL,
    appln_auth char(2) DEFAULT '' NOT NULL,
    appln_nr varchar(15) DEFAULT '' NOT NULL,
    appln_kind char(2) DEFAULT '' NOT NULL,
    appln_filing_date date DEFAULT '9999-12-31' NOT NULL,
    appln_filing_year smallint DEFAULT 9999 NOT NULL,
    appln_nr_epodoc varchar(20) DEFAULT '' NOT NULL,
    appln_nr_original varchar(100) DEFAULT '' NOT NULL,
    ipr_type char(2) DEFAULT '' NOT NULL,
    receiving_office char(2) DEFAULT '' NOT NULL,
    internat_appln_id integer DEFAULT 0 NOT NULL,
    int_phase char(1) DEFAULT 'N' NOT NULL,
    reg_phase char(1) DEFAULT 'N' NOT NULL,
    nat_phase char(1) DEFAULT 'N' NOT NULL,
    earliest_filing_date date DEFAULT '9999-12-31' NOT NULL,
    earliest_filing_year smallint DEFAULT 9999 NOT NULL,
    earliest_filing_id integer DEFAULT 0 NOT NULL,
    earliest_publn_date date DEFAULT '9999-12-31' NOT NULL,
    earliest_publn_year smallint DEFAULT 9999 NOT NULL,
    earliest_pat_publn_id integer DEFAULT 0 NOT NULL,
    granted char(1) DEFAULT 'N' NOT NULL,
    docdb_family_id integer DEFAULT 0 NOT NULL,
    inpadoc_family_id integer DEFAULT 0 NOT NULL,
    docdb_family_size smallint DEFAULT 0 NOT NULL,
    nb_citing_docdb_fam smallint DEFAULT 0 NOT NULL,
    nb_applicants smallint DEFAULT 0 NOT NULL,
    nb_inventors smallint DEFAULT 0 NOT NULL
);
*/

// Tls201Appln is the structure for application table
// The core domain object is the Application, which is a request for patent protection for an
// invention filed with the EPO or another patent office.
// During the life of a patent, various publications are issued. An application has at least one publication,
// otherwise it would still be treated as confidential and would not be accessible in any database.
// Exceptions are applications that, for example, have been used as a priority or have
// been cited, but then revoked before publication.
// Every application (TLS201_APPLN) has at least 1 publication (TLS211_PAT_PUBLN). Every
// application belongs to exactly 1 simple family (also called DOCDB family) and to exactly 1
// extended family (also called INPADOC family) (TLS201_APPLN).
// Strictly speaking, title, abstract, persons and classifications are part of the publication.
// However, by design, in PATSTAT these domain objects are related not to the individual
// publication, but to the application of the publication.
//
// This table contains the key bibliographical data elements relevant to identify the patent
// application. Most of the elements in this table can be found on the first page of a printed
// patent document. E.g.: application authority, application number and application filing date.
// From a database structure point of view, this table is very important because it links to
// many other database tables via the attribute APPLN_ID.
// 2022 Autumn:
// Deprecated attributes removed
// (APPLN_NR_EPODOC from table TLS201_APPLN, and EVENT_IMPACT from table TLS803_LEGAL_EVENT_CODE).
type Tls201Appln struct {
	ApplnId            int       `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;"`
	ApplnAuth          string    `json:"applnAuth" gorm:"column:appln_auth;type:char(2);default:'';not null"`
	ApplnNr            string    `json:"applnNr" gorm:"column:appln_nr;type:varchar(15);default:'';not null"`
	ApplnKind          string    `json:"applnKind" gorm:"column:appln_kind;type:char(2);default:'';not null"`
	ApplnFilingDate    time.Time `json:"applnFilingDate" gorm:"column:appln_filing_date;type:date;default:'9999-12-31';not null"`
	ApplnFilingYear    int       `json:"applnFilingYear" gorm:"column:appln_filing_year;type:smallint;default:9999;not null"`
	ApplnNrOriginal    string    `json:"applnNrOriginal" gorm:"column:appln_nr_original;type:varchar(100);default:'';not null"`
	IprType            string    `json:"iprType" gorm:"column:ipr_type;type:char(2);default:'';not null"`
	ReceivingOffice    string    `json:"receivingOffice" gorm:"column:receiving_office;type:char(2);default:'';not null"`
	InternatApplnId    int       `json:"internatApplnId" gorm:"column:internat_appln_id;type:integer;default:0;not null"`
	IntPhase           string    `json:"intPhase" gorm:"column:int_phase;type:char(1);default:'N';not null"`
	RegPhase           string    `json:"regPhase" gorm:"column:reg_phase;type:char(1);default:'N';not null"`
	NatPhase           string    `json:"natPhase" gorm:"column:nat_phase;type:char(1);default:'N';not null"`
	EarliestFilingDate time.Time `json:"earliestFilingDate" gorm:"column:earliest_filing_date;type:date;default:'9999-12-31';not null"`
	EarliestFilingYear int       `json:"earliestFilingYear" gorm:"column:earliest_filing_year;type:smallint;default:9999;not null"`
	EarliestFilingId   int       `json:"earliestFilingId" gorm:"column:earliest_filing_id;type:integer;default:0;not null"`
	EarliestPublnDate  time.Time `json:"earliestPublnDate" gorm:"column:earliest_publn_date;type:date;default:'9999-12-31';not null"`
	EarliestPublnYear  int       `json:"earliestPublnYear" gorm:"column:earliest_publn_year;type:smallint;default:9999;not null"`
	EarliestPatPublnId int       `json:"earliestPatPublnId" gorm:"column:earliest_pat_publn_id;type:integer;default:0;not null"`
	Granted            string    `json:"granted" gorm:"column:granted;type:char(1);default:'N';not null"`
	DocDbFamilyId      int       `json:"docDbFamilyId" gorm:"column:docdb_family_id;type:integer;default:0;not null"`
	InpadocFamilyId    int       `json:"inpadocFamilyId" gorm:"column:inpadoc_family_id;type:integer;default:0;not null"`
	DocdbFamilySize    int       `json:"docdbFamilySize" gorm:"column:docdb_family_size;type:smallint;default:0;not null"`
	NbCitingDocdbFam   int       `json:"nbCitingDocdbFam" gorm:"column:nb_citing_docdb_fam;type:smallint;default:0;not null"`
	NbApplicants       int       `json:"nbApplicants" gorm:"column:nb_applicants;type:smallint;default:0;not null"`
	NbInventors        int       `json:"nbInventors" gorm:"column:nb_inventors;type:smallint;default:0;not null"`
	// relations
	Title                     *Tls202ApplnTitle          `json:"title" gorm:"foreignKey:appln_id"`
	Abstract                  *Tls203ApplnAbstr          `json:"abstract" gorm:"foreignKey:appln_id"`
	Priorities                []*Tls204ApplnPrior        `json:"priorities" gorm:"foreignKey:appln_id"`
	PrioritiesApplications    []*Tls204ApplnPrior        `json:"prioritiesApplications" gorm:"foreignKey:prior_appln_id"`
	TechRelations             []*Tls205TechRel           `json:"techRelations" gorm:"foreignKey:appln_id"`
	TechRelationsApplications []*Tls205TechRel           `json:"techRelationsApplications" gorm:"foreignKey:tech_rel_appln_id"`
	PersonApplications        []*Tls207PersAppln         `json:"personsApplications" gorm:"foreignKey:appln_id;"`
	IpcClasses                []*Tls209ApplnIpc          `json:"ipcClasses" gorm:"foreignKey:appln_id"`
	NationalClasses           []*Tls210ApplnNCls         `json:"nationalClasses" gorm:"foreignKey:appln_id"`
	Publications              []*Tls211PatPubln          `json:"publications" gorm:"foreignKey:appln_id"`
	Continuations             []*Tls216ApplnContn        `json:"continuations" gorm:"foreignKey:appln_id"`
	ParentContinuations       []*Tls216ApplnContn        `json:"parentContinuations" gorm:"foreignKey:parent_appln_id"`
	Citations                 []*Tls212Citation          `json:"citations" gorm:"foreignKey:cited_appln_id"`
	FamilyCitations           []*Tls228DocdbFamCitn      `json:"familyCitations" gorm:"foreignKey:docdb_family_id"`
	FamilyCited               []*Tls228DocdbFamCitn      `json:"familyCited" gorm:"foreignKey:cited_docdb_family_id"`
	JpClasses                 []*Tls222ApplnJpClass      `json:"jpClasses" gorm:"foreignKey:appln_id"`
	CpcClasses                []*Tls224ApplnCpc          `json:"cpcClasses" gorm:"foreignKey:appln_id"`
	FamilyCpcClasses          []*Tls225DocdbFamCpc       `json:"familyCpcClasses" gorm:"foreignKey:docdb_family_id"`
	NaceCodes                 []*Tls229ApplnNace2        `json:"naceCodes" gorm:"foreignKey:appln_id"`
	TechnicalFields           []*Tls230ApplnTechnField   `json:"technicalFields" gorm:"foreignKey:appln_id"`
	LegalEvents               []*Tls231InpadocLegalEvent `json:"legalEvents" gorm:"foreignKey:appln_id"`
	// Deprecated
	// UsClasses                 []*Tls223ApplnDocus        `json:"usClasses" gorm:"foreignKey:appln_id"` // Deprecated
	// ApplnNrEpodoc      string    `json:"applnNrEpodoc" gorm:"column:appln_nr_epodoc;type:varchar(20);default:'';not null"` // Deprecated
}

// TableName sets the sql table name for this struct type
func (obj *Tls201Appln) TableName() string {
	return "tls201_appln"
}
