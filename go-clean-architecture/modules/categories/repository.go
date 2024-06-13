package categories

import "gorm.io/gorm"

type CategoryRepository interface {
	GetAGetAll() ([]Category, error)
	GetById(id string) (Category, error)
	Create(category Category) error
	Update(id string, category Category) error
	Delete(id string) error
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepositoryImpl{
		db: db,
	}
}

// Create implements CategoryRepository.
func (c CategoryRepositoryImpl) Create(category Category) error {
	panic("unimplemented")
}

// Delete implements CategoryRepository.
func (c CategoryRepositoryImpl) Delete(id string) error {
	panic("unimplemented")
}

// GetAGetAll implements CategoryRepository.
func (c CategoryRepositoryImpl) GetAGetAll() ([]Category, error) {
	panic("unimplemented")
}

// GetById implements CategoryRepository.
func (c CategoryRepositoryImpl) GetById(id string) (Category, error) {
	panic("unimplemented")
}

// Update implements CategoryRepository.
func (c CategoryRepositoryImpl) Update(id string, category Category) error {
	panic("unimplemented")
}
