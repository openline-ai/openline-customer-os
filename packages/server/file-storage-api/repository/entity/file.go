package entity

type File struct {
	ID        string `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	TenantId  string `gorm:"column:tenant_id;type:varchar(255);NOT NULL" json:"tenantId" binding:"required"`
	Name      string `gorm:"column:name;type:varchar(255);NOT NULL;" json:"name" binding:"required"`
	Extension string `gorm:"column:extension;type:varchar(255);NOT NULL;" json:"extension" binding:"required"`
	MIME      string `gorm:"column:mime;type:varchar(255);NOT NULL;" json:"mime" binding:"required"`
	Length    int64  `gorm:"column:length;type:bigint;NOT NULL;" json:"length" binding:"required"`
}

func (File) TableName() string {
	return "files"
}
