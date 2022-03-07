package models

/*
CREATE TABLE tls202_appln_title (
	appln_id integer DEFAULT 0 NOT NULL,
	appln_title_lg char(2) DEFAULT '' NOT NULL,
	appln_title text NOT NULL
);
*/

// Tls202ApplnTitle is the structure for the application title
type Tls202ApplnTitle struct {
	ApplnId      int    `json:"applnId" gorm:"column:appln_id;type:integer;default:0;not null"`
	ApplnTitleLg string `json:"applnTitleLg" gorm:"column:appln_title_lg;type:char(2);default:'';not null"`
	ApplnTitle   string `json:"applnTitle" gorm:"column:appln_title;type:text;not null"`
}

func (obj *Tls202ApplnTitle) TableName() string {
	return "tls202_appln_title"
}
