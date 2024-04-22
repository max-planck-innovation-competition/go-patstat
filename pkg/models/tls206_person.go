package models

/*
CREATE TABLE tls206_person (
	person_id integer DEFAULT 0 NOT NULL,
	person_name text DEFAULT '' NOT NULL,
	person_name_orig_lg varchar(500) DEFAULT '' NOT NULL,
	person_address text DEFAULT '' NOT NULL,
	person_ctry_code char(2) DEFAULT '' NOT NULL,
	nuts varchar(5) DEFAULT '' NOT NULL,
	nuts_level smallint DEFAULT 9 NOT NULL,
	doc_std_name_id integer DEFAULT 0 NOT NULL,
	doc_std_name varchar(500) DEFAULT '' NOT NULL,
	psn_id integer DEFAULT 0 NOT NULL,
	psn_name varchar(500) DEFAULT '' NOT NULL,
	psn_level smallint DEFAULT 0 NOT NULL,
	psn_sector varchar(50) DEFAULT '' NOT NULL,
	han_id integer DEFAULT 0 NOT NULL,
	han_name text DEFAULT '' NOT NULL,
	han_harmonized integer DEFAULT 0 NOT NULL
);

*/

// Tls206Person is a structure representing a person
// Table that contains the key data on applicants and inventors such as: the person name,
// the address and the country/territory of residence (which is not necessarily the nationality).
// Several types of names are available:
// - The name as delivered by the offices
// - The name in original language, possibly in a non-Latin character set
// - The name as standardised by the EPO (DOCDB standardised name)
// - The PATSTAT standardised name
// - The name as standardised by the OECD (OECD Harmonised Applicant Name)
type Tls206Person struct {
	PersonID         int    `csv:"person_id" json:"personId" gorm:"primaryKey;column:person_id;type:integer;not null"`
	PersonName       string `csv:"person_name" json:"personName" gorm:"column:person_name;type:text;default:'';not null"`
	PersonNameOrigLg string `csv:"person_name_orig_lg" json:"personNameOrigLg" gorm:"column:person_name_orig_lg;type:varchar(500);default:'';not null"`
	PersonAddress    string `csv:"person_address" json:"personAddress" gorm:"column:person_address;type:text;default:'';not null"`
	PersonCtryCode   string `csv:"person_ctry_code" json:"personCtryCode" gorm:"column:person_ctry_code;type:char(2);default:'';not null"`
	Nuts             string `csv:"nuts" json:"nuts" gorm:"column:nuts;type:varchar(5);default:'';not null"`
	NutsLevel        int16  `csv:"nuts_level" json:"nutsLevel" gorm:"column:nuts_level;type:smallint;default:9;not null"`
	DocStdNameID     int    `csv:"doc_std_name_id" json:"docStdNameId" gorm:"column:doc_std_name_id;type:integer;default:0;not null"`
	DocStdName       string `csv:"doc_std_name" json:"docStdName" gorm:"column:doc_std_name;type:varchar(500);default:'';not null"`
	PsnID            int    `csv:"psn_id" json:"psnId" gorm:"column:psn_id;type:integer;default:0;not null"`
	PsnName          string `csv:"psn_name" json:"psnName" gorm:"column:psn_name;type:varchar(500);default:'';not null"`
	PsnLevel         int16  `csv:"psn_level" json:"psnLevel" gorm:"column:psn_level;type:smallint;default:0;not null"`
	PsnSector        string `csv:"psn_sector" json:"psnSector" gorm:"column:psn_sector;type:varchar(50);default:'';not null"`
	HanID            int    `csv:"han_id" json:"hanId" gorm:"column:han_id;type:integer;default:0;not null"`
	HanName          string `csv:"han_name" json:"hanName" gorm:"column:han_name;type:text;default:'';not null"`
	HanHarmonized    int    `csv:"han_harmonized" json:"hanHarmonized" gorm:"column:han_harmonized;type:integer;default:0;not null"`
	// relations
	// Unmodified *Tls226PersonOrig `csv:"-" json:"unmodified" gorm:"foreignKey:person_id;references:person_id;default:null"`
	// Relations
	//PersonApplications []*Tls207PersAppln `json:"personsApplications" gorm:"foreignKey:person_id;"`
}

// TableName sets the sql table name for this struct type
func (m *Tls206Person) TableName() string {
	return "tls206_person"
}
