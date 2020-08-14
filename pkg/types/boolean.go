package types

type Boolean struct {
    Type
}

func (b *Boolean) Size() int {
    return 1
}

func (b *Boolean) String() string {
    return "BOOLEAN"
}

type BooleanValue struct {
}

func (b *BooleanValue) GetTypeId() TypeId {
    return BOOLEAN
}
