package common

type SqlResultI interface {
	GetSql() string

	SetSql(sql string)

	GetSqlType() SqlType

	String() string
}

type SqlResult struct {
	sql     string
	sqlType SqlType
}

func NewSqlResult(sql string, sqlType SqlType) *SqlResult {
	return &SqlResult{
		sql:     sql,
		sqlType: sqlType,
	}
}

func (s *SqlResult) GetSql() string {
	return s.sql
}

func (s *SqlResult) SetSql(sql string) {
	s.sql = sql
}

func (s *SqlResult) GetSqlType() SqlType {
	return s.sqlType
}

func (s *SqlResult) String() string {
	return s.GetSql()
}

type SqlCountResult struct {
	*SqlResult
}

func NewSqlCountResult(sql string) *SqlCountResult {
	return &SqlCountResult{
		SqlResult: NewSqlResult(sql, SqlCountType),
	}
}

type SqlPageResult struct {
	*SqlResult

	page   int
	size   int
	limit  int
	offset int
}

func NewSqlPageResult(sql string) *SqlPageResult {
	return &SqlPageResult{
		SqlResult: NewSqlResult(sql, SqlPageType),
	}
}

func (s *SqlPageResult) SetPage(page int) *SqlPageResult {
	s.page = page
	return s
}

func (s *SqlPageResult) GetPage() int {
	return s.page
}

func (s *SqlPageResult) SetSize(size int) *SqlPageResult {
	s.size = size
	s.limit = size
	return s
}

func (s *SqlPageResult) GetSize() int {
	return s.size
}

func (s *SqlPageResult) GetLimit() int {
	return s.limit
}

func (s *SqlPageResult) SetOffset(offset int) *SqlPageResult {
	s.offset = offset
	return s
}

func (s *SqlPageResult) GetOffset() int {
	return s.offset
}

type SqlListResult struct {
	*SqlResult
}

func NewSqlListResult(sql string) *SqlListResult {
	return &SqlListResult{
		SqlResult: NewSqlResult(sql, SqlListType),
	}
}

type SqlMapResult struct {
	*SqlResult
}

func NewSqlMapResult(sql string) *SqlMapResult {
	return &SqlMapResult{
		SqlResult: NewSqlResult(sql, SqlMapType),
	}
}
