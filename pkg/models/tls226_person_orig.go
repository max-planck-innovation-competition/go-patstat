package models

/*

CREATE TABLE tls226_person_orig (
    person_orig_id integer DEFAULT 0 NOT NULL,
    person_id integer DEFAULT 0 NOT NULL,
    source char(5) DEFAULT '' NOT NULL,
    source_version varchar(10) DEFAULT '' NOT NULL,
    name_freeform varchar(500) DEFAULT '' NOT NULL,
    person_name_orig_lg varchar(500) DEFAULT '' NOT NULL,
    last_name varchar(500) DEFAULT '' NOT NULL,
    first_name varchar(500) DEFAULT '' NOT NULL,
    middle_name varchar(500) DEFAULT '' NOT NULL,
    address_freeform varchar(1000) DEFAULT '' NOT NULL,
    address_1 varchar(500) DEFAULT '' NOT NULL,
    address_2 varchar(500) DEFAULT '' NOT NULL,
    address_3 varchar(500) DEFAULT '' NOT NULL,
    address_4 varchar(500) DEFAULT '' NOT NULL,
    address_5 varchar(500) DEFAULT '' NOT NULL,
    street varchar(500) DEFAULT '' NOT NULL,
    city varchar(200) DEFAULT '' NOT NULL,
    zip_code varchar(30) DEFAULT '' NOT NULL,
    state char(2) DEFAULT '' NOT NULL,
    person_ctry_code char(2) DEFAULT '' NOT NULL,
    residence_ctry_code char(2) DEFAULT '' NOT NULL,
    role varchar(2) DEFAULT '' NOT NULL
);
*/

// Tls226PersonOrig is a structure representing a row in the tls226_person_orig table
// TLS226_PERSON_ORIG: Unmodified person data
// This table is best suited for detailed analysis of person data.
// A row contains the name and address of a person (applicant and/or inventors; physical
// person or legal person). The data is taken from various data sources. It is kept in the
// "original" form, i.e. the data has not been cleaned, aggregated or otherwise modified.
// Depending on the data structure of each data source, not all attributes of this table are
// populated for every person.
// Each row has one corresponding row in TLS206_PERSON. In these tables the data has
// been cleaned and unified and its table structure has been simplified and normalised.
type Tls226PersonOrig struct {
	PersonOrigID      int    `csv:"column" json:"personOrigId" gorm:"primaryKey;column:person_orig_id;type:integer;default:0;not null"`
	PersonID          int    `csv:"person_id" json:"personId" gorm:"column:person_id;type:integer;default:0;not null"`
	Source            string `csv:"source" json:"source" gorm:"column:source;type:char(5);default:'';not null"`
	SourceVersion     string `csv:"source_version" json:"sourceVersion" gorm:"column:source_version;type:varchar(10);default:'';not null"`
	NameFreeform      string `csv:"name_freeform" json:"nameFreeform" gorm:"column:name_freeform;type:varchar(500);default:'';not null"`
	PersonNameOrigLg  string `csv:"person_name_orig_lg" json:"personNameOrigLg" gorm:"column:person_name_orig_lg;type:varchar(500);default:'';not null"`
	LastName          string `csv:"last_name" json:"lastName" gorm:"column:last_name;type:varchar(500);default:'';not null"`
	FirstName         string `csv:"first_name" json:"firstName" gorm:"column:first_name;type:varchar(500);default:'';not null"`
	MiddleName        string `csv:"middle_name" json:"middleName" gorm:"column:middle_name;type:varchar(500);default:'';not null"`
	AddressFreeform   string `csv:"address_freeform" json:"addressFreeform" gorm:"column:address_freeform;type:varchar(1000);default:'';not null"`
	Address1          string `csv:"address_1" json:"address1" gorm:"column:address_1;type:varchar(500);default:'';not null"`
	Address2          string `csv:"address_2" json:"address2" gorm:"column:address_2;type:varchar(500);default:'';not null"`
	Address3          string `csv:"address_3" json:"address3" gorm:"column:address_3;type:varchar(500);default:'';not null"`
	Address4          string `csv:"address_4" json:"address4" gorm:"column:address_4;type:varchar(500);default:'';not null"`
	Address5          string `csv:"address_5" json:"address5" gorm:"column:address_5;type:varchar(500);default:'';not null"`
	Street            string `csv:"street" json:"street" gorm:"column:street;type:varchar(500);default:'';not null"`
	City              string `csv:"city" json:"city" gorm:"column:city;type:varchar(200);default:'';not null"`
	ZipCode           string `csv:"zip_code" json:"zipCode" gorm:"column:zip_code;type:varchar(30);default:'';not null"`
	State             string `csv:"state" json:"state" gorm:"column:state;type:char(2);default:'';not null"`
	PersonCtryCode    string `csv:"person_ctry_code" json:"personCtryCode" gorm:"column:person_ctry_code;type:char(2);default:'';not null"`
	ResidenceCtryCode string `csv:"residence_ctry_code" json:"residenceCtryCode" gorm:"column:residence_ctry_code;type:char(2);default:'';not null"`
	Role              string `csv:"role" json:"role" gorm:"column:role;type:varchar(2);default:'';not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls226PersonOrig) TableName() string {
	return "tls226_person_orig"
}
