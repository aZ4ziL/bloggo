package models

type Category struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"size:100" json:"name"`
	Slug     string    `gorm:"size:100;unique;index" json:"slug"`
	Logo     string    `gorm:"size:255;null" json:"logo"`
	Articles []Article `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"articles"`
}

// CreateNewCategory create new category
func CreateNewCategory(c *Category) error {
	err := db.Create(c).Error
	return err
}

// GetCategoryByID get category by passing the ID
func GetCategoryByID(id uint) (Category, error) {
	var c Category
	err := db.Model(&Category{}).Where("id = ?", id).Preload("Articles").First(&c).Error
	return c, err
}

// GetCategoryBySlug get category by passing the SLUG
func GetCategoryBySlug(slug string) (Category, error) {
	var c Category
	err := db.Model(&Category{}).Where("slug = ?", slug).Preload("Articles").First(&c).Error
	return c, err
}

// GetAllCategories return all categories
func GetAllCategories() []Category {
	var categories []Category
	db.Model(&Category{}).Find(&categories).Preload("Articles")
	return categories
}
