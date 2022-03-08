package models

/*
CREATE TABLE tls222_appln_jp_class (
    appln_id integer DEFAULT 0 NOT NULL,
    jp_class_scheme varchar(5) DEFAULT '' NOT NULL,
    jp_class_symbol varchar(50) DEFAULT '' NOT NULL
);
*/

// Tls222ApplnJpClass is a structure representing a row of the tls222_appln_jp_class table in the database
// FI (File Index) and F-Terms are used by the Japanese patent office for classification
// (TLS222_APPLN_JP_CLASS)
// FI and F-terms linked to JP application (only):
// FI (File Index) has been developed to expand IPC in some technical fields. FI consists of
// an IPC symbol and an IPC-subdivision symbol and/or file discrimination symbol added to
// the IPC symbol.
// F-TERMS (File Forming Terms) re-classify or further segment each specific technical field
// of IPC from a variety of viewpoints (i.e., objective, application, structure, material,
// manufacturing process, processing, etc.).
// Japan’s Patent Map Guidance System (PMGS) provides useful information about JP
// national classification of FI and F-terms in English. You may retrieve the classification list
// and an explanation for each classification:
// https://www5.j-platpat.inpit.go.jp/pms/tokujitsu/pmgs_en/PMGS_EN_GM101_Top.action
type Tls222ApplnJpClass struct {
	ApplnID       int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	JpClassScheme string `json:"jpClassScheme" gorm:"primaryKey;column:jp_class_scheme;type:varchar(5);default:'';not null"`
	JpClassSymbol string `json:"jpClassSymbol" gorm:"primaryKey;column:jp_class_symbol;type:varchar(50);default:'';not null"`
}

func (m *Tls222ApplnJpClass) TableName() string {
	return "tls222_appln_jp_class"
}
