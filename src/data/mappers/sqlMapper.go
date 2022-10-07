package mappers

type SqlMapper struct {
	Sql string
	Args []interface{}
	
}

func NewSqlMapper(sql string, args []interface{}) *SqlMapper {
	return &SqlMapper{
		Sql: sql,
		Args: args,
	}
}

func Mapper(sql string, args []interface{}, err error) *SqlMapper {

	if err != nil {
		panic(err.Error())
	}

	return NewSqlMapper(sql, args)
}
