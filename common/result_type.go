package common

var (
	SqlCountType = NewSqlType("SqlCountType")
	SqlPageType  = NewSqlType("SqlPageType")
	SqlListType  = NewSqlType("SqlListType")
	SqlMapType   = NewSqlType("SqlMapType")
)

type SqlType struct {
	name string
}

func NewSqlType(name string) SqlType {
	return SqlType{
		name: name,
	}
}

func (s SqlType) String() string {
	return s.name
}
