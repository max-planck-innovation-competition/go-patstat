package models

// Tls203ApplnAbstr is the structure for the application abstract
type Tls203ApplnAbstr struct {
	ApplnId         int    `json:"applnId" gorm:"column:appln_id;type:integer;default:0;not null"`
	ApplnAbstractLg string `json:"applnTitleLg" gorm:"column:appln_title_lg;type:char(2);default:'';not null"`
	ApplnAbstract   string `json:"applnTitle" gorm:"column:appln_title;type:text;not null"`
}

func (obj *Tls203ApplnAbstr) TableName() string {
	return "tls203_appln_abstr"
}
