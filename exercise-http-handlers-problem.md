[https://tour.go-zh.org/methods/14](https://tour.go-zh.org/methods/14)

#**练习：HTTP 处理**
实现下面的类型，并在其上定义 `ServeHTTP` 方法。在 `web` 服务器中注册它们来处理指定的路径。

`type String string`

	type Struct struct {`
    	Greeting string
    	Punct    string
    	Who      string
	}
例如，可以使用如下方式注册处理方法：

	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
在启动你的 `http` 服务器后，你将能够访问： `http://localhost:4000/string` 和 `http://localhost:4000/struct`.

*注意：* 这个例子无法在基于 `web` 的用户界面下运行。 为了尝试编写 `web` 服务，你可能需要 安装 `Go`。

#**遇到的问题**
当我运行服务端程序时，出现以下错误提示:
	
    2016/01/04 16:51:09 listen tcp 127.0.0.1:4000: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
对于服务端程序出现的这种问题，百度/google尝试了很多方法，都没能成功解决。
我将程序中的`4000`端口改为`8080`端口就可以正常访问了。