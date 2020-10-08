package stores

import "xtek/exchange/commission/internal/models"

// CommissionRepo ...
type CommissionRepo interface {
	Create(ct *models.CommissionHistory) error
	GetByID(ID string) (*models.CommissionHistory, error)
	Update(c *models.CommissionHistory) error
	GetAllByStatus(status models.CommissionStatus) ([]*models.CommissionHistory, error)
	GetAllWithPaginateNFilter(query map[string]interface{}, limit, page int) ([]*models.CommissionHistory, error)
	CountByQuery(query map[string]interface{}) (int, error)
	UpdateByQuery(ID string, query map[string]interface{}) error
}
