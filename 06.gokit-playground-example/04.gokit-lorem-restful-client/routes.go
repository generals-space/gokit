package lorem_restful

import "github.com/gorilla/mux"

// Routes 接口路由表
type Routes struct {
	Lorem *mux.Route
}

// Initialize 初始化所有路由接口, 由transport部分调用, 然后将handler逐一挂载到不同路由.
func Initialize(router *mux.Router) Routes {
	return Routes{
		Lorem: router.Methods("POST").Path("/lorem/{type}/{min}/{max}").Name("lorem"),
	}
}
