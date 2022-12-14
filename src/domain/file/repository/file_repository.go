package repository

import (
	"github.com/PowerReport/pfs/src/domain/file/model"
	"gorm.io/gorm"
)

type IFileRepository interface {
	UserViewed() (*gorm.DB, error)
	Get(uint64) (model.File, error)
	GetAll() ([]model.File, error)
	Save(model.File) (model.File, error)
	Update(model.File) (model.File, error)
	Delete(model.File) (model.File, error)
	SoftDelete(model.File) (model.File, error)
}
