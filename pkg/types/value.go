package types

type Value interface {
    GetTypeId() TypeId
}

type TinyIntValue struct {
}

func (t *TinyIntValue) GetTypeId() TypeId {
    return TINYINT
}

type SmallIntValue struct {
}

func (s *SmallIntValue) GetTypeId() TypeId {
    return SMALLINT
}

type IntegerValue struct {
}

func (i *IntegerValue) GetTypeId() TypeId {
    return INTEGER
}

type BigIntValue struct {
}

func (b *BigIntValue) GetTypeId() TypeId {
    return BIGINT
}

type DecimalValue struct {
}

func (d *DecimalValue) GetTypeId() TypeId {
    return DECIMAL
}

type DateValue struct {
}

func (d *DateValue) GetTypeId() TypeId {
    return DATE
}

type TimestampValue struct {
}

func (t *TimestampValue) GetTypeId() TypeId {
    return TIMESTAMP
}

type VarcharValue struct {
}

func (v *VarcharValue) GetTypeId() TypeId {
    return VARCHAR
}
