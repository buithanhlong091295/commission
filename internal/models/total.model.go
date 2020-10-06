package models

import (
	"math/big"

	"github.com/richard-xtek/go-grpc-micro-kit/utils/math"

	"gopkg.in/mgo.v2/bson"
)

// Totals ..
type Totals struct {
	ID    string   `bson:"_id"`
	Total *big.Int `bson:"total"`
}

// TotalsRecord ..
type TotalsRecord struct {
	ID    string          `bson:"_id"`
	Total bson.Decimal128 `bson:"total"`
}

// SetBSON ...
func (t *Totals) SetBSON(raw bson.Raw) error {
	var decode *TotalsRecord
	err := raw.Unmarshal(&decode)
	if err != nil {
		//ll.Error("Set BSON Error: ", zap.Error(err))
		return err
	}

	t.ID = decode.ID
	t.Total = math.DBStringToBigInt(decode.Total.String())
	return nil
}

// GetBSON ...
func (t *Totals) GetBSON() (interface{}, error) {
	or := TotalsRecord{
		ID: t.ID,
	}

	v, err := bson.ParseDecimal128(t.Total.String())
	if err != nil {
		return nil, err
	}
	or.Total = v
	return or, nil
}
