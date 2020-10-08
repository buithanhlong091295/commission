package mongo

import (
	"xtek/exchange/commission/internal/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewCommissionRepo ...
func NewUsersMemberRepo(dbName string, session *mgo.Session) (*UsersMemberRepo, error) {
	collectionName := "users-members"
	// init user repository
	commissionRepo := &UsersMemberRepo{
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

// UsersMemberRepo ...
type UsersMemberRepo struct {
	dbName         string
	collectionName string
	session        *mgo.Session
}

// Create ...
func (r *UsersMemberRepo) Create(ct *models.UsersMember) error {
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

// GetByUserID ...
func (r *CommissionRepo) GetByUserID(userID string) (*models.UsersMember, error) {
	sc := r.session.Copy()
	defer sc.Close()

	var response *models.UsersMember
	query := bson.M{
		"userID": userID,
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
