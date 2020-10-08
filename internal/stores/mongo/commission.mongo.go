package mongo

import (
	"math/big"
	"reflect"
	"time"
	"xtek/exchange/commission/internal/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewCommissionRepo ...
func NewCommissionRepo(dbName string, session *mgo.Session) (*CommissionRepo, error) {
	collectionName := "commissions"
	// init user repository
	commissionRepo := &CommissionRepo{
		dbName:         dbName,
		collectionName: collectionName,
		session:        session,
	}
	// copy session connect mongo
	sc := session.Copy()
	defer sc.Close()
	// set index for collection
	index := mgo.Index{
		Key: []string{"id"},
	}

	err := sc.DB(dbName).C(collectionName).EnsureIndex(index)
	if err != nil {
		return nil, err
	}

	return commissionRepo, nil
}

// CommissionRepo ...
type CommissionRepo struct {
	dbName         string
	collectionName string
	session        *mgo.Session
}

// Create ...
func (r *CommissionRepo) Create(ct *models.CommissionHistory) error {
	sc := r.session.Copy()
	defer sc.Close()

	err := ct.BeforeCreate()
	if err != nil {
		return err
	}
	err = sc.DB(r.dbName).C(r.collectionName).Insert(ct)
	if err != nil {
		return err
	}

	return nil
}

// GetByID ...
func (r *CommissionRepo) GetByID(ID string) (*models.CommissionHistory, error) {
	sc := r.session.Copy()
	defer sc.Close()

	var response *models.CommissionHistory
	query := bson.M{
		"_id": ID,
	}
	err := sc.DB(r.dbName).C(r.collectionName).Find(query).One(&response)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *CommissionRepo) Update(c *models.CommissionHistory) error {
	sc := r.session.Copy()
	defer sc.Close()

	c.UpdatedAt = time.Now()

	err := sc.DB(r.dbName).C(r.collectionName).UpdateId(c.ID, c)
	if err != nil {
		return err
	}

	return nil
}

// GetAllByStatus ...
func (r *CommissionRepo) GetAllByStatus(status models.CommissionStatus) ([]*models.CommissionHistory, error) {
	sc := r.session.Copy()
	defer sc.Close()

	var response []*models.CommissionHistory
	query := bson.M{
		"status": status,
	}

	err := sc.DB(r.dbName).C(r.collectionName).Find(query).All(&response)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetAllWithPaginateNFilter ...
func (r *CommissionRepo) GetAllWithPaginateNFilter(query map[string]interface{}, limit, page int) ([]*models.CommissionHistory, error) {
	sc := r.session.Copy()
	defer sc.Close()
	var data []*models.CommissionHistory
	skip := (page - 1) * limit
	time := 10 * time.Second
	err := sc.DB(r.dbName).C(r.collectionName).Find(query).Sort("-createdAt").SetMaxTime(time).Skip(skip).Limit(limit).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CountByQuery ...
func (r *CommissionRepo) CountByQuery(query map[string]interface{}) (int, error) {
	sc := r.session.Copy()
	defer sc.Close()
	total, err := sc.DB(r.dbName).C(r.collectionName).Find(query).Count()
	if err != nil {
		return 0, err
	}
	return total, nil
}

// UpdateByQuery ...
func (r *CommissionRepo) UpdateByQuery(ID string, query map[string]interface{}) error {
	sc := r.session.Copy()
	defer sc.Close()

	err := sc.DB(r.dbName).C(r.collectionName).UpdateId(ID, query)
	if err != nil {
		return err
	}

	return nil
}

// GetTotalAmountByUserIDNCoinNStatus ..
func (r *CommissionRepo) GetTotalAmountByUserIDNCoinNStatus(coin, userID string, status models.CommissionStatus) (*big.Int, error) {
	sc := r.session.Copy()
	defer sc.Close()

	var res []*models.Totals
	result := reflect.ValueOf(&res).Interface()

	match := bson.M{
		"receiverID": userID,
		"status":     status,
		"coin":       coin,
	}
	group := bson.M{
		"_id": "$coin",
		"total": bson.M{
			"$sum": bson.M{
				"$convert": bson.M{
					"input": "$amount",
					"to":    "decimal",
				},
			},
		},
	}
	query := []bson.M{
		bson.M{"$match": match},
		bson.M{"$group": group},
	}
	err := sc.DB(r.dbName).C(r.collectionName).Pipe(query).All(result)
	if err != nil {
		return nil, err
	}
	if len(res) <= 0 {
		return big.NewInt(0), nil
	}
	return res[0].Total, nil
}
