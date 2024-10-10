package response

const (
	StatusOk              = 200
	CodeSuccess           = 201
	DeleteSuccess         = 204
	ErrBadRequest         = 40000
	ErrBadRequestPage     = 40001
	ErrBadRequestPageSize = 40002
	ErrBadRequestMime     = 40003
	ErrBadRequestPostId   = 40004
	ErrUnauthorize        = 40101
	ErrInvalid            = 40102
	ErrYourSelf           = 40103
	ErrUserExist          = 40900
	ErrLogin              = 40400
	ErrNotFoundUser       = 40401
	ErrFindPost           = 40402
	ErrFileTooLarge       = 41300
)

var msg = map[int]string{
	StatusOk:              "Ok",
	CodeSuccess:           "Success",
	DeleteSuccess:         "Delete no error",
	ErrBadRequest:         "Bad request",
	ErrBadRequestPage:     "Page should be number",
	ErrBadRequestPageSize: "Page size should be number",
	ErrBadRequestMime:     "Can only use file .png, .jpg, .jpeg, .gif",
	ErrBadRequestPostId:   "Id must be number, can't convert from this request",
	ErrUnauthorize:        "Unauthorize",
	ErrInvalid:            "Invalid Token",
	ErrYourSelf:           "Not your self, cant create/update/delete anything for another",
	ErrUserExist:          "User exist",
	ErrLogin:              "Wrong username or password",
	ErrNotFoundUser:       "User not found",
	ErrFindPost:           "Can not found post or post was deleted",
	ErrFileTooLarge:       "File too large, only allowed 6MB",
}

// ErrOutOfDate:   "Token is out of date",
