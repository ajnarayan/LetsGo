# LetsGo
```
$ cd LetsGo

$ go build 

$ ./letsgo
```


#### Go Workspace Basics
Every Go code will reside in your workspace. A workspace is a directory in your file system to which the path is stored in environment variable `GOPATH`. The most basic structure should contain a `src` directory.

Each Go package is created as a subdirectory within `src`. The following is the example from [Go Article](https://golang.org/doc/code.html)

```
src/
    github.com/LetsGo/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```  

#### Useful links
* [Golang Code documentation](https://golang.org/doc/code.html)
* [Common Golang Early Syntax Errors](https://golangtutorials.blogspot.com/2011/05/early-syntax-errors-and-other-minor.html)
* [GoLang Basics](https://github.com/alco/gostart)
* [Go Specs](https://tip.golang.org/ref/spec)

