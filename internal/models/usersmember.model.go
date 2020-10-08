package models

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// UsersMember ...
type UsersMember struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	UserID       string `json:"senderID" bson:"senderID"`
	SponsorID    string `json:"sponsorID" bson:"sponsorID"`
	TotalMembers int32  `json:"totalMembers" bson:"totalMembers"`
}

// BeforeCreate ...
func (b *UsersMember) BeforeCreate() error {
	b.ID = uuid.NewV4().String()
	return nil
}

// UsersMemberRecord ...
type UsersMemberRecord struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	UserID       string `json:"senderID" bson:"senderID"`
	SponsorID    string `json:"sponsorID" bson:"sponsorID"`
	TotalMembers int32  `json:"totalMembers" bson:"totalMembers"`
}

// SetBSON ...
func (b *UsersMember) SetBSON(raw bson.Raw) error {
	var decode *UsersMemberRecord
	err := raw.Unmarshal(&decode)
	if err != nil {
		// ll.Error("Set BSON Error: ", zap.Error(err))
		return err
	}
	b.ID = decode.ID
	b.SponsorID = decode.SponsorID
	b.TotalMembers = decode.TotalMembers
	return nil
}

// GetBSON ...
func (b *UsersMember) GetBSON() (interface{}, error) {
	or := UsersMemberRecord{
		UserID:       b.UserID,
		SponsorID:    b.SponsorID,
		TotalMembers: b.TotalMembers,
	}
	if b.ID == "" {
		or.ID = uuid.NewV4().String()
	} else {
		or.ID = b.ID
	}
	return or, nil
}
