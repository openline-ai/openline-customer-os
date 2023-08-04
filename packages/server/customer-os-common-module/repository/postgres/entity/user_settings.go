package entity

import "github.com/google/uuid"

type UserSettingsEntity struct {
	ID                          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TenantName                  string    `gorm:"size:255;not null"`
	UserName                    string    `gorm:"size:255;not null;uniqueIndex:idx_user_settings"`
	GoogleOAuthAllScopesEnabled bool      `gorm:"type:boolean;not null;default:false"`
	GoogleOAuthUserAccessToken  string    `gorm:"type:text"`
}

func (UserSettingsEntity) TableName() string {
	return "user_settings"
}
