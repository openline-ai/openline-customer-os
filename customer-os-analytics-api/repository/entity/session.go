package entity

type SessionEntity struct {
	ID          string `gorm:"column:domain_sessionid;type:varchar(128);NOT NULL" json:"sessionId" binding:"required"`
	AppId       string `gorm:"column:app_id;type:varchar(255);NOT NULL" json:"appName" binding:"required"`
	TrackerName string `gorm:"column:name_tracker;type:varchar(128);NOT NULL" json:"trackerName" binding:"required"`
	Tenant      string `gorm:"column:tenant;type:varchar(64);NOT NULL" json:"tenant" binding:"required"`
	Country     string `gorm:"column:geo_country"`
	Region      string `gorm:"column:geo_region_name"`
	City        string `gorm:"column:geo_city"`
}

type SessionEntities []SessionEntity

func (SessionEntity) TableName() string {
	return "derived.sessions"
}
