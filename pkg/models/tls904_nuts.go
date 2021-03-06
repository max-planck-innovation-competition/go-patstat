package models

/*
CREATE TABLE tls904_nuts (
    nuts varchar(5) DEFAULT ('') NOT NULL,
    nuts_level smallint DEFAULT '0',
    nuts_label varchar(250) DEFAULT ''
);
*/

// Tls904Nuts is a table which contains the regions of the Nomenclature of Territorial Units for Statistics
// NUTS (Nomenclature of Territorial Units for Statistics) is a European Union standard for
// referencing the subdivisions of countries for statistical purposes. This reference table
// contains the regions of the NUTS levels 0 - 3
type Tls904Nuts struct {
	Nuts      string `json:"nuts" gorm:"primaryKey;column:nuts;type:varchar(5);"`
	NutsLevel int16  `json:"nutsLevel" gorm:"column:nuts_level;type:smallint;default:0;"`
	NutsLabel string `json:"nutsLabel" gorm:"column:nuts_label;type:varchar(250);default:'';"`
}

// TableName sets the sql table name for this struct type
func (m *Tls904Nuts) TableName() string {
	return "tls904_nuts"
}
