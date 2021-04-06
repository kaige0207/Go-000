package main

func main() {
	
}

type temporary interface {
	Temporary() bool    // IsTemporary returns true if err is temporary.
}

//自定义错误类型
type MyError struct {
}

func (e *MyError) Error() string {
	return "xxx"
}

func (e *MyError) Temporary() bool {
	//xxx
	return true
}
//对外暴露的调用方法
func (e *MyError) IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}
