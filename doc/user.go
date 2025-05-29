package doc

import "go-swagger-demo/api"

// swagger:route POST /user user createUserRequest
// 创建用户.
// responses:
//   200: createUserResponse
//   default: errResponse

// swagger:route GET /user/{name} user getUserRequest
// 查询用户.
// responses:
//   200: getUserResponse
//   default: errResponse

// swagger:parameters createUserRequest
type createUserParamsWrapper struct { //对请求体的描述
	// in:body
	Body api.User
}

// swagger:parameters getUserRequest
type getUserParamsWrapper struct { //对请求路径参数的描述
	// in:path
	Name string `json:"name"`
}

// swagger:response createUserResponse
type createUserResponseWrapper struct { //对创建用户响应体的描述
	// in:body
	Body api.User
}

// swagger:response getUserResponse
type getUserResponseWrapper struct { //对查询用户响应体的描述
	// in:body
	Body api.User
}

// swagger:response errResponse
type errResponseWrapper struct { //对错误响应体的描述
	Code    int    `json:"code"`
	Message string `json:"message"`
}
