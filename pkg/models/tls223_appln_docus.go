package models

/*
CREATE TABLE tls223_appln_docus (
    appln_id integer DEFAULT 0 NOT NULL,
    docus_class_symbol varchar(50) DEFAULT '' NOT NULL
);
*/

// Tls223ApplnDocus is a structure representing a tls223_appln_docus table row
// USPC codes have been used by the US office for classification until recently
// (TLS223_APPLN_DOCUS)
// USPC codes linked to US applications (only):
// US patent classification (USPC) codes are typically expressed in the following format
// "482/1" on US publications. The first number, 482, represents the class of invention. The
// number following the slash is the subclass of invention within the class. There are about
// 450 classes of invention and about 150 000 subclasses of invention in the USPC.
type Tls223ApplnDocus struct {
	ApplnID          int    `json:"applnId" gorm:"primaryKey;column:appln_id;type:integer;default:0;not null"`
	DocusClassSymbol string `json:"docusClassSymbol" gorm:"primaryKey;column:docus_class_symbol;type:varchar(50);default:'';not null"`
}

func (m *Tls223ApplnDocus) TableName() string {
	return "tls223_appln_docus"
}
