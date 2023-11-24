package database

var connection string

func init() {
	connection = "MYSQL"
}

func GetDatabaseName() string {
	return connection
}
