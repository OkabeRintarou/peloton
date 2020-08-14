package types

import "fmt"

type abstractType interface {
    fmt.Stringer
    GetTypeId() TypeId
    Size() int
}

type Type struct {
    typeId TypeId
}

func (t *Type) GetTypeId() TypeId {
    return t.typeId
}

func (t *Type) Size() int {
    return 0
}

func (t *Type) String() string {
    return "INVALID"
}

var globalTypes [14]abstractType

func init() {
    globalTypes[0] = &Type{INVALID}
    globalTypes[2] = &Boolean{Type{typeId: BOOLEAN}}
}
