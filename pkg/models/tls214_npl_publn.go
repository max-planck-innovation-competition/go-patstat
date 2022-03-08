package models

/*
CREATE TABLE tls214_npl_publn (
    npl_publn_id varchar(32) DEFAULT '0' NOT NULL,
    xp_nr integer DEFAULT 0 NOT NULL,
    npl_type char(1) DEFAULT '' NOT NULL,
    npl_biblio text DEFAULT '' NOT NULL,
    npl_author varchar(1000) DEFAULT '' NOT NULL,
    npl_title1 varchar(1000) DEFAULT '' NOT NULL,
    npl_title2 varchar(1000) DEFAULT '' NOT NULL,
    npl_editor varchar(500) DEFAULT '' NOT NULL,
    npl_volume varchar(50) DEFAULT '' NOT NULL,
    npl_issue varchar(50) DEFAULT '' NOT NULL,
    npl_publn_date varchar(8) DEFAULT '' NOT NULL,
    npl_publn_end_date varchar(8) DEFAULT '' NOT NULL,
    npl_publisher varchar(500) DEFAULT '' NOT NULL,
    npl_page_first varchar(200) DEFAULT '' NOT NULL,
    npl_page_last varchar(200) DEFAULT '' NOT NULL,
    npl_abstract_nr varchar(50) DEFAULT '' NOT NULL,
    npl_doi varchar(500) DEFAULT '' NOT NULL,
    npl_isbn varchar(30) DEFAULT '' NOT NULL,
    npl_issn varchar(30) DEFAULT '' NOT NULL,
    online_availability varchar(500) DEFAULT '' NOT NULL,
    online_classification varchar(35) DEFAULT '' NOT NULL,
    online_search_date varchar(8) DEFAULT '' NOT NULL
);
*/

type Tls214NplPubln struct {
	NplPublnID           string `json:"nplPublnId" gorm:"primaryKey;column:npl_publn_id;type:varchar(32);'default:'0';not null"`
	XpNr                 int    `json:"xpNr" gorm:"column:xp_nr;type:integer;default:0;not null"`
	NplType              string `json:"nplType" gorm:"column:npl_type;type:char(1);default:'';not null"`
	NplBiblio            string `json:"nplBiblio" gorm:"column:npl_biblio;type:text;default:'';not null"`
	NplAuthor            string `json:"nplAuthor" gorm:"column:npl_author;type:varchar(1000);default:'';not null"`
	NplTitle1            string `json:"nplTitle1" gorm:"column:npl_title1;type:varchar(1000);default:'';not null"`
	NplTitle2            string `json:"nplTitle2" gorm:"column:npl_title2;type:varchar(1000);default:'';not null"`
	NplEditor            string `json:"nplEditor" gorm:"column:npl_editor;type:varchar(500);default:'';not null"`
	NplVolume            string `json:"nplVolume" gorm:"column:npl_volume;type:varchar(50);default:'';not null"`
	NplIssue             string `json:"nplIssue" gorm:"column:npl_issue;type:varchar(50);default:'';not null"`
	NplPublnDate         string `json:"nplPublnDate" gorm:"column:npl_publn_date;type:varchar(8);default:'';not null"`
	NplPublnEndDate      string `json:"nplPublnEndDate" gorm:"column:npl_publn_end_date;type:varchar(8);default:'';not null"`
	NplPublisher         string `json:"nplPublisher" gorm:"column:npl_publisher;type:varchar(500);default:'';not null"`
	NplPageFirst         string `json:"nplPageFirst" gorm:"column:npl_page_first;type:varchar(200);default:'';not null"`
	NplPageLast          string `json:"nplPageLast" gorm:"column:npl_page_last;type:varchar(200);default:'';not null"`
	NplAbstractNr        string `json:"nplAbstractNr" gorm:"column:npl_abstract_nr;type:varchar(50);default:'';not null"`
	NplDoi               string `json:"nplDoi" gorm:"column:npl_doi;type:varchar(500);default:'';not null"`
	NplIsbn              string `json:"nplIsbn" gorm:"column:npl_isbn;type:varchar(30);default:'';not null"`
	NplIssn              string `json:"nplIssn" gorm:"column:npl_issn;type:varchar(30);default:'';not null"`
	OnlineAvailability   string `json:"onlineAvailability" gorm:"column:online_availability;type:varchar(500);default:'';not null"`
	OnlineClassification string `json:"onlineClassification" gorm:"column:online_classification;type:varchar(35);default:'';not null"`
	OnlineSearchDate     string `json:"onlineSearchDate" gorm:"column:online_search_date;type:varchar(8);default:'';not null"`
	// relations
	Citations []Tls212Citation `json:"citations" gorm:"foreignKey:cited_npl_publn_id"`
}

func (m *Tls214NplPubln) TableName() string {
	return "tls214_npl_publn"
}
