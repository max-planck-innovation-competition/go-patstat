package models

/*
CREATE TABLE tls228_docdb_fam_citn (
    docdb_family_id integer DEFAULT 0 NOT NULL,
    cited_docdb_family_id integer DEFAULT 0 NOT NULL
);
*/

// Tls228DocdbFamCitn is a structure representing a row in the tls228_docdb_fam_citn table in the database
// For analysis, citations on the level of families (TLS228_DOCDB_FAM_CITN) are often
// regarded as more interesting than on the level of publications
// (The 3 levels: Family – Application – Publication).
type Tls228DocdbFamCitn struct {
	DocdbFamilyID      int `json:"docdbFamilyId" gorm:"primaryKey;column:docdb_family_id;type:integer;default:0;not null"`
	CitedDocdbFamilyID int `json:"citedDocdbFamilyId" gorm:"primaryKey;column:cited_docdb_family_id;type:integer;default:0;not null"`
}

func (m *Tls228DocdbFamCitn) TableName() string {
	return "tls228_docdb_fam_citn"
}
