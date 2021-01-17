**go源码阅读 – http.router**

> http.router基于radix tree 基数树，关于基数树内容可见：https://en.wikipedia.org/wiki/Radix_tree
>
> ```go
> //Demo
> package main
> 
> import (
>     "fmt"
>     "net/http"
>     "log"
> 
>     "github.com/julienschmidt/httprouter"
> )
> 
> func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
>     fmt.Fprint(w, "Welcome!\n")
> }
> 
> func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
>     fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
> }
> 
> func main() {
>     router := httprouter.New()
>     router.GET("/", Index)
>     router.GET("/hello/:name", Hello)
> 
>     log.Fatal(http.ListenAndServe(":8080", router))
> }
> 
> //router := httprouter.New() 返回router实例
> func New() *Router {
> 	return &Router{
> 		RedirectTrailingSlash:  true,
> 		RedirectFixedPath:      true,
> 		HandleMethodNotAllowed: true,
> 		HandleOPTIONS:          true,
> 	}
> }
> 
> // Router is a http.Handler which can be used to dispatch requests to different
> // handler functions via configurable routes
> type Router struct {
>     //trees 路由森林，存储方法与路由树（节点中定义了handle）
> 	trees map[string]*node
> 
> 	paramsPool sync.Pool
> 	maxParams  uint16
> 	SaveMatchedRoutePath bool
> 	RedirectTrailingSlash bool
> 	RedirectFixedPath bool
> 	HandleMethodNotAllowed bool
> 	HandleOPTIONS bool
> 	GlobalOPTIONS http.Handler
> 	globalAllowed string
> 	NotFound http.Handler
> 	MethodNotAllowed http.Handler
> 	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
> }
> 
> //router实现了serveHttp方法，符合handle接口定义
> // ServeHTTP makes the router implement the http.Handler interface.
> func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
> 	if r.PanicHandler != nil {
> 		defer r.recv(w, req)
> 	}
> 
> 	path := req.URL.Path
> 
> 	if root := r.trees[req.Method]; root != nil {
> 		if handle, ps, tsr := root.getValue(path, r.getParams); handle != nil {
> 			if ps != nil {
> 				handle(w, req, *ps)
> 				r.putParams(ps)
> 			} else {
> 				handle(w, req, nil)
> 			}
> 			return
>     ...
> }
>         
> type node struct {
> 	path     	string
>     indices	 	string	//所有孩子的首字母构成的[]byte 转化而来
>     wildChild 	bool
>     nType 		nodeType
>     priority	uint32
>     children 	[]*node
>     handle		Handle
> }
>         
> 
> ```
>
> 

