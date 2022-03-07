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

type Tls226PersonOrig struct {
	PersonOrigID      int    `json:"personOrigId" gorm:"column:person_orig_id;type:integer;default:0;not null"`
	PersonID          int    `json:"personId" gorm:"column:person_id;type:integer;default:0;not null"`
	Source            string `json:"source" gorm:"column:source;type:char(5);default:'';not null"`
	SourceVersion     string `json:"sourceVersion" gorm:"column:source_version;type:varchar(10);default:'';not null"`
	NameFreeform      string `json:"nameFreeform" gorm:"column:name_freeform;type:varchar(500);default:'';not null"`
	PersonNameOrigLg  string `json:"personNameOrigLg" gorm:"column:person_name_orig_lg;type:varchar(500);default:'';not null"`
	LastName          string `json:"lastName" gorm:"column:last_name;type:varchar(500);default:'';not null"`
	FirstName         string `json:"firstName" gorm:"column:first_name;type:varchar(500);default:'';not null"`
	MiddleName        string `json:"middleName" gorm:"column:middle_name;type:varchar(500);default:'';not null"`
	AddressFreeform   string `json:"addressFreeform" gorm:"column:address_freeform;type:varchar(1000);default:'';not null"`
	Address1          string `json:"address1" gorm:"column:address_1;type:varchar(500);default:'';not null"`
	Address2          string `json:"address2" gorm:"column:address_2;type:varchar(500);default:'';not null"`
	Address3          string `json:"address3" gorm:"column:address_3;type:varchar(500);default:'';not null"`
	Address4          string `json:"address4" gorm:"column:address_4;type:varchar(500);default:'';not null"`
	Address5          string `json:"address5" gorm:"column:address_5;type:varchar(500);default:'';not null"`
	Street            string `json:"street" gorm:"column:street;type:varchar(500);default:'';not null"`
	City              string `json:"city" gorm:"column:city;type:varchar(200);default:'';not null"`
	ZipCode           string `json:"zipCode" gorm:"column:zip_code;type:varchar(30);default:'';not null"`
	State             string `json:"state" gorm:"column:state;type:char(2);default:'';not null"`
	PersonCtryCode    string `json:"personCtryCode" gorm:"column:person_ctry_code;type:char(2);default:'';not null"`
	ResidenceCtryCode string `json:"residenceCtryCode" gorm:"column:residence_ctry_code;type:char(2);default:'';not null"`
	Role              string `json:"role" gorm:"column:role;type:varchar(2);default:'';not null"`
}

func (m *Tls226PersonOrig) TableName() string {
	return "tls226_person_orig"
}
