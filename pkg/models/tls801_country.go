package models

/*
CREATE TABLE tls801_country (
    ctry_code char(2) DEFAULT '' NOT NULL,
    iso_alpha3 varchar(3) DEFAULT '' NOT NULL,
    st3_name varchar(100) DEFAULT '' NOT NULL,
    organisation_flag char(1) DEFAULT '' NOT NULL,
    continent varchar(25) DEFAULT '' NOT NULL,
    eu_member varchar(1) DEFAULT '' NOT NULL,
    epo_member varchar(1) DEFAULT '' NOT NULL,
    oecd_member varchar(1) DEFAULT '' NOT NULL,
    discontinued varchar(1) DEFAULT '' NOT NULL
);
*/

// Tls801Country is a structure representing a row from the 'tls801_country' table
// Contains information about states/countries/territories and IP organisations, e.g. their twoand three-letter code, their (short) name and whether they are member of the EU, the EPO
// or the OECD. This table is based on WIPO standard ST.3.
type Tls801Country struct {
	CtryCode         string `json:"ctryCode" gorm:"primaryKey;column:ctry_code;type:char(2);default:'';not null"`
	IsoAlpha3        string `json:"isoAlpha3" gorm:"column:iso_alpha3;type:varchar(3);default:'';not null"`
	St3Name          string `json:"st3Name" gorm:"column:st3_name;type:varchar(100);default:'';not null"`
	OrganisationFlag string `json:"organisationFlag" gorm:"column:organisation_flag;type:char(1);default:'';not null"`
	Continent        string `json:"continent" gorm:"column:continent;type:varchar(25);default:'';not null"`
	EuMember         string `json:"euMember" gorm:"column:eu_member;type:varchar(1);default:'';not null"`
	EpoMember        string `json:"epoMember" gorm:"column:epo_member;type:varchar(1);default:'';not null"`
	OecdMember       string `json:"oecdMember" gorm:"column:oecd_member;type:varchar(1);default:'';not null"`
	Discontinued     string `json:"discontinued" gorm:"column:discontinued;type:varchar(1);default:'';not null"`
}

// TableName sets the sql table name for this struct type
func (m *Tls801Country) TableName() string {
	return "tls801_country"
}
