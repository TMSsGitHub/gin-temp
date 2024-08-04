package model

import (
	"gin-temp/internal/global/utils"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt uint64 `json:"createdAt,omitempty"`
	UpdatedAt uint64 `json:"updatedAt,omitempty"`
	DeletedAt uint64 `json:"deletedAt,omitempty"`
}

// BeforeCreate
// 创建前自动填充创建时间和更新时间
func (base *BaseModel) BeforeCreate(_ *gorm.DB) error {
	now := utils.GetCurrentMs()
	base.CreatedAt = now
	base.UpdatedAt = now
	return nil
}

// BeforeUpdate
// 更新前自动填充更新时间
func (base *BaseModel) BeforeUpdate(_ *gorm.DB) error {
	base.UpdatedAt = utils.GetCurrentMs()
	return nil
}
