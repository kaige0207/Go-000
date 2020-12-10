package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

/*
	作业：基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出
	代码：可以实现如下功能
			1.通过 errgroup 分别启动端口号为 8081 和 8082 的两个 http server
			2.注册了 linux signal 信号监听: Ctrl C 、终端退出、程序停止
			3.通过 errgroup 实现了一个服务退出，全部注销退出，
			  在收到 signal 中断信号或者客户端退出请求(addr/shutdown)后，两个服务正常关闭、signal 信号监听停止
*/

var (
	name1 = "Server1"
	name2 = "Server2"
	addr1 = ":8081"
	addr2 = ":8082"
)

type Handler struct {
	name string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello " + h.name))
}

func handleServer(ch chan string, name, addr string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", &Handler{name: name})
	//当客户端请求/shutdown时触发关闭服务器
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Bye Bye " + name))
		ch <- "Shutdown"
	})
	return &http.Server{Addr: addr, Handler: mux}
}

func listenServer(server *http.Server) error {
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func main() {
	group, _ := errgroup.WithContext(context.Background())
	chan1 := make(chan string)
	chan2 := make(chan string)
	ch := make(chan os.Signal, 1)
	server1 := handleServer(chan1, name1, addr1)
	server2 := handleServer(chan2, name2, addr2)
	log.Println("Register Signal...")
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	//启动并监听 Server1
	group.Go(func() error {
		log.Println("Server1 Start...")
		return listenServer(server1)
	})

	//启动并监听 Server2
	group.Go(func() error {
		log.Println("Server2 Start...")
		return listenServer(server2)
	})

	//当一个退出，全部注销退出
	group.Go(func() (err error) {
		// 收到任何一个信号就关闭服务
		select {
		case sigMsg := <-ch:
			fmt.Printf("Receive signal close: %v\n", sigMsg)
		case serMsg1 := <-chan1:
			fmt.Printf("Receive server1 request: %v\n", serMsg1)
		case serMsg2 := <-chan2:
			fmt.Printf("Receive server2 request: %v\n", serMsg2)
		}

		close(chan1)
		if err := server1.Shutdown(context.Background()); err != nil {
			return err
		}
		log.Println("Server1 closed!")

		close(chan2)
		if err := server2.Shutdown(context.Background()); err != nil {
			return err
		}
		log.Println("Server2 closed!")

		signal.Stop(ch)
		close(ch)
		log.Println("Signal closed!")

		return nil
	})

	if err := group.Wait(); err != nil {
		log.Printf("Shutdown errorr: %+v \n", err)
	}
}
