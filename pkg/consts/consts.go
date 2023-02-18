
package consts

const (
	MessageTableName   = "message"
	UserTableName   = "user"
	VideoTableName   = "video"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Messages        = "messages"
	Videos			= "videos"
	ApiServiceName  = "api"
	// NoteServiceName = "note"
	UserServiceName = "user"
	MessageServiceName = "message"
	MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	MessageServiceAddr = ":9001"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
)
