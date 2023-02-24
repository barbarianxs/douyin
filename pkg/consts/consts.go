
package consts

const (
	MessageTableName   = "message"
	RelationTableName   = "relation"
	CommentTableName = "comment"
	UserTableName   = "user"
	FavoriteTableName = "favorite"
	VideoTableName   = "video"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Messages        = "messages"
	Videos			= "videos"
	ApiServiceName  = "api"
	// NoteServiceName = "note"
	UserServiceName = "user"
	InteractServiceName = "interact"
	RelationServiceName = "message"
	MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	RelationServiceAddr = ":9001"
	InteractServiceAddr = ":9002"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
	VideoSavePath   = "https://douyin-test-guo.oss-cn-hangzhou.aliyuncs.com/video/"
	CoverPath := "https://douyin-test-main.oss-cn-hangzhou.aliyuncs.com/img/"
	StatusOK = 200

	//favorite actiontype,1是点赞，2是取消点赞
	Like   = 1
	Unlike = 2
)
