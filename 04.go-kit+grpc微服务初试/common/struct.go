package common

// User ...
type User struct {
	Name    string
	Company string
}

// Department ...
type Department struct{
	Name string
	Users []*User
}

var UserManagerHttpTransportAddr = ":8081"
var DepartmentHttpTransportAddr = ":8082"
var UserManagerGrpcTransportAddr = ":9001"
var DepartmentGrpcTransportAddr = ":9002"