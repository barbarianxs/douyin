
package errno

import (
	"errors"
	"fmt"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
)
const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	LoginErrCode               = 10005
	UserNotExistErrCode        = 10006

	IdNotEqualErrCode      = 20001
	ActionUnSupportErrCode = 20002

	FollowSelfErrcode     = 60001
	UserIDErrCode         = 60002
	ActionTypeErrCode     = 60003
	RelationActionErrCode = 60004
	RelationExistErrCode  = 60005

	CommentNotFound = 80001
	CommentError    = 88848

	Token2UserIdErrCode        = 90001 //从token中获取userid的时候出错
	BrokenAccessControlErrCode = 90002 //越权错误
)
type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(int32(user.ErrCode_SuccessCode), "Success")
	ServiceErr             = NewErrNo(int32(user.ErrCode_ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int32(user.ErrCode_ParamErrCode), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int32(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int32(user.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
	MessageIsNullExistErr  = NewErrNo(int32(relation.ErrCode_MessageIsNullErrCode ), "Message is Null")
	
	FollowSelfErr          = NewErrNo(int32(FollowSelfErrcode), "Follow self err")
	UserIDErr              = NewErrNo(int32(UserIDErrCode), "UserID  is wrong")
	ActionTypeErr          = NewErrNo(int32(ActionTypeErrCode), "ActionType is  unlegal")
	RelationActionErr      = NewErrNo(int32(RelationActionErrCode), "In Dal, RelationAction RowsAffected > 1")
	Token2UserIdErr        = NewErrNo(int32(Token2UserIdErrCode), "UseID is unable to get from toke")
	BrokenAccessControlErr = NewErrNo(int32(BrokenAccessControlErrCode), "Broken Access Control!")
	RelationExistErr       = NewErrNo(int32(RelationExistErrCode), "The follow relation has existed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
