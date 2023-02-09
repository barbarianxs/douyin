
package consts

const (
	// NoteTableName   = "note"
	UserTableName   = "user"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	ApiServiceName  = "api"
	// NoteServiceName = "note"
	UserServiceName = "user"
	MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	NoteServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
)
