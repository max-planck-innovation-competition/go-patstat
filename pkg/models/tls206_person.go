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

type Tls206Person struct {
	PersonID         int    `json:"personId" gorm:"primaryKey;column:person_id;type:integer;default:0;not null"`
	PersonName       string `json:"personName" gorm:"column:person_name;type:text;default:'';not null"`
	PersonNameOrigLg string `json:"personNameOrigLg" gorm:"column:person_name_orig_lg;type:varchar(500);default:'';not null"`
	PersonAddress    string `json:"personAddress" gorm:"column:person_address;type:text;default:'';not null"`
	PersonCtryCode   string `json:"personCtryCode" gorm:"column:person_ctry_code;type:char(2);default:'';not null"`
	Nuts             string `json:"nuts" gorm:"column:nuts;type:varchar(5);default:'';not null"`
	NutsLevel        int16  `json:"nutsLevel" gorm:"column:nuts_level;type:smallint;default:9;not null"`
	DocStdNameID     int    `json:"docStdNameId" gorm:"column:doc_std_name_id;type:integer;default:0;not null"`
	DocStdName       string `json:"docStdName" gorm:"column:doc_std_name;type:varchar(500);default:'';not null"`
	PsnID            int    `json:"psnId" gorm:"column:psn_id;type:integer;default:0;not null"`
	PsnName          string `json:"psnName" gorm:"column:psn_name;type:varchar(500);default:'';not null"`
	PsnLevel         int16  `json:"psnLevel" gorm:"column:psn_level;type:smallint;default:0;not null"`
	PsnSector        string `json:"psnSector" gorm:"column:psn_sector;type:varchar(50);default:'';not null"`
	HanID            int    `json:"hanId" gorm:"column:han_id;type:integer;default:0;not null"`
	HanName          string `json:"hanName" gorm:"column:han_name;type:text;default:'';not null"`
	HanHarmonized    int    `json:"hanHarmonized" gorm:"column:han_harmonized;type:integer;default:0;not null"`
	// relations
	Unmodified Tls226PersonOrig `json:"unmodified" gorm:"foreignKey:person_id"`
}

func (m *Tls206Person) TableName() string {
	return "tls206_person"
}
