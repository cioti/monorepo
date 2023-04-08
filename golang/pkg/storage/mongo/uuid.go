package mongo

import (
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type MUUID struct {
	uuid.UUID
}

func NewMUUID(id uuid.UUID) MUUID {
	return MUUID{
		UUID: id,
	}
}

func (mu MUUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.Binary, bsoncore.AppendBinary(nil, 4, mu.UUID[:]), nil
}

func (mu *MUUID) UnmarshalBSONValue(t bsontype.Type, raw []byte) error {
	if t != bsontype.Binary {
		return fmt.Errorf("invalid format on unmarshal bson value")
	}

	_, data, _, ok := bsoncore.ReadBinary(raw)
	if !ok {
		return fmt.Errorf("not enough bytes to unmarshal bson value")
	}

	copy(mu.UUID[:], data)

	return nil
}
