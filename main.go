package main

import (
	"fmt"
)

type User struct {
	Username string `json:"username"`
}

func main() {
	r := NewRouter().Prefix("/app-name")

	//r.GET("/user", UserController)
	//r.POST("/user", CreateUserController)

	r.Routes("/user", func(route *GroupRoutes) {
		route.OnGET(UserController)
		route.OnPOST(CreateUserController)
	})
	r.GET("/user/{id}/{name}", UserControllerWithId)

	r.Use(LoggerMiddleware, CORSMiddleware)
	r.Listen(8000)
}

func UserController(ctx *Context) {
	_, _ = ctx.Write("This is users")
	fmt.Println("API JALAN")
}

func UserControllerWithId(ctx *Context) {
	_, _ = ctx.Write("ID : " + ctx.Param("id") + " NAME : " + ctx.Param("name"))
}

func CreateUserController(ctx *Context) {
	var user User
	_ = ctx.Body(&user)

	_= ctx.JSON(user)
}

func AuthMiddleware(handler Handler) Handler {
	return func(ctx *Context) {
		fmt.Println("AUTH MIDDLEWARE MULAI")
		handler(ctx)
		fmt.Println("AUTH MIDDLEWARE SELESAI")
	}
}

func LoggerMiddleware(handler Handler) Handler {
	return func(ctx *Context) {
		fmt.Println("LOGGER MIDDLEWARE MULAI")
		handler(ctx)
		fmt.Println("LOGGER MIDDLEWARE SELESAI")
	}
}

func CORSMiddleware(handler Handler) Handler {
	return func(ctx *Context) {
		fmt.Println("CORS MIDDLEWARE MULAI")
		handler(ctx)
		fmt.Println("CORS MIDDLEWARE SELESAI")
	}
}















//func binarySearch(target int64, nums []int64) int {
//	left := 0
//	right := len(nums) - 1
//
//	for left <= right {
//		middle := (left + right) / 2
//
//		if nums[middle] == target {
//			return middle
//		}
//
//		if target < nums[middle] {
//			right = middle - 1
//		}
//
//		if target > nums[middle] {
//			left = middle + 1
//		}
//	}
//
//	return -1
//}
//
//func basicSearch(target int64, nums []int64) int64 {
//	for i, num := range nums {
//		if target == num {
//			return int64(i)
//		}
//	}
//	return -1
//}

// ~~~~~ RouteEntry ~~~~~ //

//type RouteEntry struct {
//	Path        *regexp.Regexp
//	Method      string
//	HandlerFunc http.HandlerFunc
//}
//
//func (ent *RouteEntry) Match(r *http.Request) map[string]string {
//	match := ent.Path.FindStringSubmatch(r.URL.Path)
//	if match == nil {
//		return nil // No match found
//	}
//
//	// Create a map to store URL parameters in
//	params := make(map[string]string)
//	groupNames := ent.Path.SubexpNames()
//	for i, group := range match {
//		params[groupNames[i]] = group
//	}
//
//	return params
//}
//
//// ~~~~~ Router ~~~~~ //
//
//type Router struct {
//	routes []RouteEntry
//}
//
//func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
//	// NOTE: ^ means start of string and $ means end. Without these,
//	//   we'll still match if the path has content before or after
//	//   the expression (/foo/bar/baz would match the "/bar" route).
//	exactPath := regexp.MustCompile("^" + path + "$")
//
//	e := RouteEntry{
//		Method:      method,
//		Path:        exactPath,
//		HandlerFunc: handlerFunc,
//	}
//	rtr.routes = append(rtr.routes, e)
//}
//
//func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	for _, e := range rtr.routes {
//		params := e.Match(r)
//		if params == nil {
//			continue // No match found
//		}
//
//		// Create new request with params stored in context
//		ctx := context.WithValue(r.Context(), "params", params)
//		e.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))
//		return
//	}
//
//	http.NotFound(w, r)
//}
//
//// ~~~~~ Helpers ~~~~~ //
//
//func URLParam(r *http.Request, name string) string {
//	ctx := r.Context()
//	params := ctx.Value("params").(map[string]string)
//	return params[name]
//}