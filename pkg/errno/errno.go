// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errno

import (
	"errors"
	"fmt"
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
	Success                = NewErrNo(int32(SuccessCode), "Success")
	ServiceErr             = NewErrNo(int32(ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int32(ParamErrCode), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int32(UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int32(AuthorizationFailedErrCode), "Authorization failed")
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
