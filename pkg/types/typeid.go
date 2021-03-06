package types

type TypeId int

const (
    INVALID = iota
    PARAMETER_OFFSET
    BOOLEAN
    TINYINT
    SMALLINT
    INTEGER
    BIGINT
    DECIMAL
    TIMESTAMP
    DATE
    VARCHAR
    VARBINARY
    ARRAY
    UDT
)
