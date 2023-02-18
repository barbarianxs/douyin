package errno

const (
	//service error
	ParamParseErrCode = 10002
	//General incoming parameter error

	TokenExpiredErrCode       = 10205
	TokenValidationErrCode    = 10206
	TokenInvalidErrCode       = 10207
	UserNameValidationErrCode = 10208
	PasswordValidationErrCode = 10209
	//Video-related incoming parameter error
	VideoDataGetErrCode  = 10301
	VideoDataCopyErrCode = 10302
	//Comment-related incoming parameter error
	CommentTextErrCode = 10401
	//Relation-related incoming parameter error
)

var (
	//service error
	ParamParseErr = NewErrNo(ParamParseErrCode, "Could not parse the param")
	//General incoming parameter error
	//User-related incoming parameter error
	LoginErr              = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr       = NewErrNo(UserNotExistErrCode, "User does not exists")
	TokenExpiredErr       = NewErrNo(TokenExpiredErrCode, "Token has been expired")
	TokenValidationErr    = NewErrNo(TokenInvalidErrCode, "Token is not active yet")
	TokenInvalidErr       = NewErrNo(TokenInvalidErrCode, "Token Invalid")
	UserNameValidationErr = NewErrNo(UserNameValidationErrCode, "Username is invalid")
	PasswordValidationErr = NewErrNo(PasswordValidationErrCode, "Password is invalid")
	//Video-related incoming parameter error
	VideoDataGetErr  = NewErrNo(VideoDataGetErrCode, "Could not get video data")
	VideoDataCopyErr = NewErrNo(VideoDataCopyErrCode, "Could not copy video data")
	//Comment-related incoming parameter error
	CommentTextErr = NewErrNo(CommentTextErrCode, "Comment text too long")
	//Relation-related incoming parameter error
)
