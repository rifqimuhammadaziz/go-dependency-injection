package simple

type Database struct {
	Name string
}

// create alias for avoid multiple type data error
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"}) // convert to DatabaseMongoDB
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"}) // convert to DatabasePostgreSQL
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(
	postgreSQL *DatabasePostgreSQL,
	mongoDB *DatabaseMongoDB,
) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongoDB}
}
