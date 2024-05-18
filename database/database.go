package database

var connection string

func init() {
	connection = "MySQL"
}

func Getdatabase() string {
	return connection
}
