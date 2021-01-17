**gin源码阅读笔记**

>```go
>//Demo
>package main
>
>import (
>	"github.com/gin-gonic/gin"
>	"net/http"
>	"github.com/mattn/go-colorable"
>)
>
>func main() {
>    //解决windows下日志输出乱码及颜色问题
>	gin.ForceConsoleColor()
>	gin.DefaultWriter = colorable.NewColorableStdout()
>	
>    r := gin.Default()
>	r.GET("/home", func(ctx *gin.Context) {
>		ctx.JSON(http.StatusOK, gin.H{
>			"message": "OK!",
>		})
>	})
>	r.Run(":9090")
>}
>
>//Default
>//Default方法构成如下
>func Default() *Engine {
>	debugPrintWARNINGDefault()
>    //初始化并返回一个Engine结构体
>	engine := New()
>    //调用Use函数， 添加中间件Logger() 与 Recovery()
>	engine.Use(Logger(), Recovery())
>	return engine
>}
>
>//engine := New()
>//对engine进行初始化
>func New() *Engine {
>	debugPrintWARNINGNew()
>	engine := &Engine{
>		//内嵌路由组结构
>		RouterGroup: RouterGroup{
>			Handlers: nil,
>			basePath: "/",
>			root:     true,
>		},
>		FuncMap:                template.FuncMap{},
>		RedirectTrailingSlash:  true,
>		RedirectFixedPath:      false,
>		HandleMethodNotAllowed: false,
>		ForwardedByClientIP:    true,
>		AppEngine:              defaultAppEngine,
>		UseRawPath:             false,
>		RemoveExtraSlash:       false,
>		UnescapePathValues:     true,
>		MaxMultipartMemory:     defaultMultipartMemory,
>		trees:                  make(methodTrees, 0, 9),
>		delims:                 render.Delims{Left: "{{", Right: "}}"},
>		secureJsonPrefix:       "while(1);",
>	}
>    //将 engine 重新赋值给路由组对象中的 engine
>    //互相转换
>	engine.RouterGroup.engine = engine
>	engine.pool.New = func() interface{} {
>		return engine.allocateContext()
>	}
>	return engine
>}
>
>//engine结构体
>// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
>// Create an instance of Engine, by using New() or Default()
>type Engine struct {
>    //路由组
>	RouterGroup
>
>	// Enables automatic redirection if the current route can't be matched but a
>	// handler for the path with (without) the trailing slash exists.
>	// For example if /foo/ is requested but a route only exists for /foo, the
>	// client is redirected to /foo with http status code 301 for GET requests
>	// and 307 for all other request methods.
>    //自动进行重定向
>	RedirectTrailingSlash bool
>
>	// If enabled, the router tries to fix the current request path, if no
>	// handle is registered for it.
>	// First superfluous path elements like ../ or // are removed.
>	// Afterwards the router does a case-insensitive lookup of the cleaned path.
>	// If a handle can be found for this route, the router makes a redirection
>	// to the corrected path with status code 301 for GET requests and 307 for
>	// all other request methods.
>	// For example /FOO and /..//Foo could be redirected to /foo.
>	// RedirectTrailingSlash is independent of this option.
>    //修复路径
>	RedirectFixedPath bool
>
>	// If enabled, the router checks if another method is allowed for the
>	// current route, if the current request can not be routed.
>	// If this is the case, the request is answered with 'Method Not Allowed'
>	// and HTTP status code 405.
>	// If no other Method is allowed, the request is delegated to the NotFound
>	// handler.
>    //判断是否有其他方法满足请求
>	HandleMethodNotAllowed bool
>	ForwardedByClientIP    bool
>
>	// #726 #755 If enabled, it will thrust some headers starting with
>	// 'X-AppEngine...' for better integration with that PaaS.
>	AppEngine bool
>
>	// If enabled, the url.RawPath will be used to find parameters.
>	UseRawPath bool
>
>	// If true, the path value will be unescaped.
>	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,
>	// as url.Path gonna be used, which is already unescaped.
>    //不转义路径
>	UnescapePathValues bool
>
>	// Value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
>	// method call.
>	MaxMultipartMemory int64
>
>	// RemoveExtraSlash a parameter can be parsed from the URL even with extra slashes.
>	// See the PR #1817 and issue #1644
>	RemoveExtraSlash bool
>
>	delims           render.Delims
>	secureJsonPrefix string
>	HTMLRender       render.HTMLRender
>	FuncMap          template.FuncMap
>	allNoRoute       HandlersChain
>	allNoMethod      HandlersChain
>	noRoute          HandlersChain
>	noMethod         HandlersChain
>	pool             sync.Pool
>	trees            methodTrees
>}
>
>//engine.Use(Logger(), Recovery())
>//Use 实现
>func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
>    //路由组handle添加中间件（HandlerFunc）
>	engine.RouterGroup.Use(middleware...)
>    //将404Handler、405Handler添加进入HandlerChain
>	engine.rebuild404Handlers()
>	engine.rebuild405Handlers()
>	return engine
>}
>
>// Use adds middleware to the group, see example code in GitHub.
>func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
>    //顺序添加middleware(本质上是handlers)
>	group.Handlers = append(group.Handlers, middleware...)
>	return group.returnObj()
>}
>
>func (engine *Engine) rebuild404Handlers() {
>	engine.allNoRoute = engine.combineHandlers(engine.noRoute)
>}
>func (engine *Engine) rebuild405Handlers() {
>	engine.allNoMethod = engine.combineHandlers(engine.noMethod)
>}
>
>func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
>    //确定handle数量， 超出上限即panic
>	finalSize := len(group.Handlers) + len(handlers)
>	if finalSize >= int(abortIndex) {
>		panic("too many handlers")
>	}
>    //合并handle， 先放入group中的handle， 而后handlersChain
>	mergedHandlers := make(HandlersChain, finalSize)
>	copy(mergedHandlers, group.Handlers)
>	copy(mergedHandlers[len(group.Handlers):], handlers)
>	return mergedHandlers
>}
>
>//r.GET("/home", func(ctx *gin.Context)
>func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
>	return group.handle(http.MethodGet, relativePath, handlers)
>}
>
>func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
>    //计算绝对路径
>	absolutePath := group.calculateAbsolutePath(relativePath)
>	//合并handler
>    	handlers = group.combineHandlers(handlers)
>    
>	group.engine.addRoute(httpMethod, absolutePath, handlers)
>	return group.returnObj()
>}
>
>func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
>	assert1(path[0] == '/', "path must begin with '/'")
>	assert1(method != "", "HTTP method can not be empty")
>	assert1(len(handlers) > 0, "there must be at least one handler")
>
>	debugPrintRoute(method, path, handlers)
>	root := engine.trees.get(method)
>	if root == nil {
>		root = new(node)
>		root.fullPath = "/"
>		engine.trees = append(engine.trees, methodTree{method: method, root: root})
>	}
>	root.addRoute(path, handlers)
>}
>
>type H map[string]interface{}
>
>// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
>// It is a shortcut for http.ListenAndServe(addr, router)
>// Note: this method will block the calling goroutine indefinitely unless an error happens.
>func (engine *Engine) Run(addr ...string) (err error) {
>	defer func() { debugPrintError(err) }()
>
>	address := resolveAddress(addr)
>	debugPrint("Listening and serving HTTP on %s\n", address)
>	err = http.ListenAndServe(address, engine)
>	return
>}
>
>func ListenAndServe(addr string, handler Handler) error {
>	server := &Server{Addr: addr, Handler: handler}
>	return server.ListenAndServe()
>}
>
>type Handler interface {
>	ServeHTTP(ResponseWriter, *Request)
>}
>
>// ServeHTTP conforms to the http.Handler interface.
>//请求来了，从 engine.pool 里拿一个空的context，丢到 engine.handleHTTPRequest 处理，然后回收。
>func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
>	c := engine.pool.Get().(*Context)
>	c.writermem.reset(w)
>	c.Request = req
>	c.reset()
>
>	engine.handleHTTPRequest(c)
>
>	engine.pool.Put(c)
>}
>
>func (engine *Engine) handleHTTPRequest(c *Context) {
>	httpMethod := c.Request.Method
>	rPath := c.Request.URL.Path
>	unescape := false
>	if engine.UseRawPath && len(c.Request.URL.RawPath) > 0 {
>		rPath = c.Request.URL.RawPath
>		unescape = engine.UnescapePathValues
>	}
>
>	if engine.RemoveExtraSlash {
>		rPath = cleanPath(rPath)
>	}
>
>	// Find root of the tree for the given HTTP method
>	t := engine.trees
>	for i, tl := 0, len(t); i < tl; i++ {
>		if t[i].method != httpMethod {
>			continue
>		}
>		root := t[i].root
>		// Find route in tree
>		value := root.getValue(rPath, c.Params, unescape)
>		if value.handlers != nil {
>			c.handlers = value.handlers
>			c.Params = value.params
>			c.fullPath = value.fullPath
>			c.Next()
>			c.writermem.WriteHeaderNow()
>			return
>		}
>		if httpMethod != "CONNECT" && rPath != "/" {
>			if value.tsr && engine.RedirectTrailingSlash {
>				redirectTrailingSlash(c)
>				return
>			}
>			if engine.RedirectFixedPath && redirectFixedPath(c, root, engine.RedirectFixedPath) {
>				return
>			}
>		}
>		break
>	}
>
>	if engine.HandleMethodNotAllowed {
>		for _, tree := range engine.trees {
>			if tree.method == httpMethod {
>				continue
>			}
>			if value := tree.root.getValue(rPath, nil, unescape); value.handlers != nil {
>				c.handlers = engine.allNoMethod
>				serveError(c, http.StatusMethodNotAllowed, default405Body)
>				return
>			}
>		}
>	}
>	c.handlers = engine.allNoRoute
>	serveError(c, http.StatusNotFound, default404Body)
>}
>
>//调用下一个handle
>func (c *Context) Next() {
>	c.index++
>	for c.index < int8(len(c.handlers)) {
>		c.handlers[c.index](c)
>		c.index++
>	}
>}
>```
>
>