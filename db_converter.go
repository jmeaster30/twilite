package twilite

type DbConverter interface {
	IntoDatabaseType() (string, error)

	IntoDatabaseData() error
	FromDatabaseData()
}
