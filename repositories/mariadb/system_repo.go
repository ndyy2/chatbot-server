// repositories/mariadb/system_repo.go
package mariadb

import (
	"ai-assistant/models/mariadb"

	"gorm.io/gorm"
)

type SystemRepository struct {
	db *gorm.DB
}

func NewSystemRepository(db *gorm.DB) *SystemRepository {
	return &SystemRepository{db: db}
}

func (r *SystemRepository) GetSetting(key string) (string, error) {
	var setting mariadb.SystemSetting
	result := r.db.Where("key = ?", key).First(&setting)
	return setting.Value, result.Error
}

func (r *SystemRepository) UpdateSetting(key, value string) error {
	return r.db.Model(&mariadb.SystemSetting{}).
		Where("key = ?", key).
		Update("value", value).Error
}