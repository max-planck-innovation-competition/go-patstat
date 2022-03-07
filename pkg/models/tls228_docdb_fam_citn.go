package models

/*
CREATE TABLE tls228_docdb_fam_citn (
    docdb_family_id integer DEFAULT 0 NOT NULL,
    cited_docdb_family_id integer DEFAULT 0 NOT NULL
);
*/

type Tls228DocdbFamCitn struct {
	DocdbFamilyID      int `json:"docdbFamilyId" gorm:"column:docdb_family_id;type:integer;default:0;not null"`
	CitedDocdbFamilyID int `json:"citedDocdbFamilyId" gorm:"column:cited_docdb_family_id;type:integer;default:0;not null"`
}

func (m *Tls228DocdbFamCitn) TableName() string {
	return "tls228_docdb_fam_citn"
}
