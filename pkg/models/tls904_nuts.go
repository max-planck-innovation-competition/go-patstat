package models

/*
CREATE TABLE tls904_nuts (
    nuts varchar(5) DEFAULT ('') NOT NULL,
    nuts_level smallint DEFAULT '0',
    nuts_label varchar(250) DEFAULT ''
);
*/

type Tls904Nuts struct {
	Nuts      string `json:"nuts" gorm:"column:nuts;type:varchar(5);"`
	NutsLevel int16  `json:"nutsLevel" gorm:"column:nuts_level;type:smallint;default:'0';"`
	NutsLabel string `json:"nutsLabel" gorm:"column:nuts_label;type:varchar(250);default:'';"`
}

func (m *Tls904Nuts) TableName() string {
	return "tls904_nuts"
}
