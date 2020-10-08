package models

import (
	"math/big"
	"time"

	"github.com/richard-xtek/go-grpc-micro-kit/utils/math"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// CommissionType ...
type CommissionType int

var (
	// CommissionLendingType ...
	CommissionLendingType CommissionType = 0
	// CommissionTradeType ...
	CommissionTradeType CommissionType = 1
)

// CommissionStatus ...
type CommissionStatus int

var (
	// CommissionCompletedStatus ...
	CommissionCompletedStatus CommissionStatus = 0
	// CommssionFailStatus ...
	CommssionFailStatus CommissionStatus = 1
	// CommissionFreezedStatus ...
	CommissionFreezedStatus CommissionStatus = 2
)

// CommissionHistory ...
type CommissionHistory struct {
	ID         string           `json:"id,omitempty" bson:"_id,omitempty"`
	SenderID   string           `json:"senderID" bson:"senderID"`
	ReceiverID string           `json:"receiverID" bson:"receiverID"`
	Amount     *big.Int         `json:"amount" bson:"amount"`
	Type       CommissionType   `json:"type" bson:"type"`
	Status     CommissionStatus `json:"status" bson:"status"`
	CreatedAt  time.Time        `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt" bson:"updatedAt"`
	Coin       string           `json:"coin" bson:"coin"`
}

// BeforeCreate ...
func (b *CommissionHistory) BeforeCreate() error {
	b.ID = uuid.NewV4().String()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

// CommissionHistoryRecord ...
type CommissionHistoryRecord struct {
	ID         string           `json:"id,omitempty" bson:"_id,omitempty"`
	SenderID   string           `json:"senderID" bson:"senderID"`
	ReceiverID string           `json:"receiverID" bson:"receiverID"`
	Amount     string           `json:"amount" bson:"amount"`
	Type       CommissionType   `json:"type" bson:"type"`
	Status     CommissionStatus `json:"status" bson:"status"`
	CreatedAt  time.Time        `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt" bson:"updatedAt"`
	Coin       string           `json:"coin" bson:"coin"`
}

// SetBSON ...
func (b *CommissionHistory) SetBSON(raw bson.Raw) error {
	var decode *CommissionHistoryRecord
	err := raw.Unmarshal(&decode)
	if err != nil {
		// ll.Error("Set BSON Error: ", zap.Error(err))
		return err
	}
	b.ID = decode.ID
	b.SenderID = decode.SenderID
	b.ReceiverID = decode.ReceiverID
	b.CreatedAt = decode.CreatedAt
	b.UpdatedAt = decode.UpdatedAt
	b.Coin = decode.Coin
	if decode.Amount != "" {
		b.Amount = math.DBStringToBigInt(decode.Amount)
	} else {
		b.Amount = big.NewInt(0)
	}
	b.Type = decode.Type
	b.Status = decode.Status
	return nil
}

// GetBSON ...
func (b *CommissionHistory) GetBSON() (interface{}, error) {
	or := CommissionHistoryRecord{
		SenderID:   b.SenderID,
		ReceiverID: b.ReceiverID,
		CreatedAt:  b.CreatedAt,
		UpdatedAt:  b.UpdatedAt,
		Type:       b.Type,
		Status:     b.Status,
		Coin:       b.Coin,
	}
	if b.ID == "" {
		or.ID = uuid.NewV4().String()
	} else {
		or.ID = b.ID
	}
	if b.Amount != nil {
		or.Amount = math.BigIntToDBString(b.Amount)
	}
	return or, nil
}

// ParseAmountToFloat ..
func (b *CommissionHistory) ParseAmountToFloat() float64 {
	amount := math.GeneralFormatToFloat(b.Amount)
	return amount
}

// GetCompletedStatus ...
func (b *CommissionHistory) GetCompletedStatus() CommissionStatus {
	return CommissionCompletedStatus
}

// GetFailStatus ...
func (b *CommissionHistory) GetFailStatus() CommissionStatus {
	return CommssionFailStatus
}
