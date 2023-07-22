package models

/*
CREATE TABLE tls803_legal_event_code (
    event_auth char(2) DEFAULT '' NOT NULL,
    event_code varchar(4) DEFAULT '' NOT NULL,
    event_impact char(1) DEFAULT '',
    event_descr varchar(250) DEFAULT '',
    event_descr_orig varchar(250) DEFAULT '',
    event_category_code char(1) DEFAULT '',
    event_category_title varchar(100) DEFAULT ''
);
*/

// Tls803LegalEventCode is a structure representing a row of data from the 'tls803_legal_event_code' table
// This table contains all legal event codes which are used in EPO’s worldwide legal event
// database (also called INPADOC database). Similar legal event codes are grouped into legal
// event categories.
// 2022 Autumn:
// Deprecated attributes removed
// (APPLN_NR_EPODOC from table TLS201_APPLN, and EVENT_IMPACT from table TLS803_LEGAL_EVENT_CODE).
type Tls803LegalEventCode struct {
	EventAuth          string `json:"eventAuth" gorm:"primaryKey;column:event_auth;type:char(2);default:'';not null"`
	EventCode          string `json:"eventCode" gorm:"primaryKey;column:event_code;type:varchar(4);default:'';not null"`
	EventDescr         string `json:"eventDescr" gorm:"column:event_descr;type:varchar(250);default:'';"`
	EventDescrOrig     string `json:"eventDescrOrig" gorm:"column:event_descr_orig;type:varchar(250);default:'';"`
	EventCategoryCode  string `json:"eventCategoryCode" gorm:"column:event_category_code;type:char(1);default:'';"`
	EventCategoryTitle string `json:"eventCategoryTitle" gorm:"column:event_category_title;type:varchar(100);default:'';"`
	// EventImpact        string `json:"eventImpact" gorm:"column:event_impact;type:char(1);default:'';"` Deprecated
}

// TableName sets the sql table name for this struct type
func (m *Tls803LegalEventCode) TableName() string {
	return "tls803_legal_event_code"
}
