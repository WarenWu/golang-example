package channel

import(
	"strconv"
	"fmt"
) 


type Empty interface{}

var empty Empty

const N int = 64

func SemMoe() {
	src := make([]int, N)
	dst := make([]string, N)
	sem := make(chan Empty)
	for i, v := range src {
		go func(i, v int) {
			dst[i] = doSomthig(i, v)
			sem <- empty
		}(i, v)
	}
	for i := 0; i < N; i++ {
		<-sem
	}
}

func doSomthig(i, v int) string {
	return ""
}

func FactoyMode() {
	wait := func() chan Empty {
		out := make(chan Empty)
		go func() {
			for i := 0; i < N; i++ {
				out <- i
			}
		}()
		return out
	}()

	for v := range wait{
		fmt.Println(v)
	}
}

func FilterMode()  {

	in:= func () chan int {
		ret := make(chan int)
		go func ()  {
			for i := 0; i < 1000; i++ {
				ret <- i
			}
		}()
		return ret
	}()

	for{
		prime := <- in
        in = func (in1 chan int, prime int)chan int {
			out := make(chan int)
			go func ()  {
				i:=<-in1
				if i%prime != 0{
					out <- i
				}
			}()
			return out
		}(in, prime)
	}
}



func ServerMode()  {
	type req struct{
		a, b int64
		reqCh chan string
	}

	type handleFunc func(a, b int64) string
	in := func (op handleFunc) chan *req {
		out := make(chan *req)
		go func() {
           for v := range out{
			// goroutine 和闭包使用要注意不要用上下文抓取，会出错
			//    go func(){
			// 	   req.reqCh <- op(req.a, req.b)
			//    }()
			   go func(req *req){
				   req.reqCh <- op(req.a, req.b)
			   }(v)
		   }
		}()
		return out
	}(func(a, b int64) string {
		return strconv.FormatInt(a+b,10)
	})
	
	var reqs [N]*req
	for i := 0; i < N; i++ {
        reqs[i] = &req{
			a:int64(i),
			b:int64(i + N),
			reqCh:make(chan string), 
		}
		in <- reqs[i]
	}

	for i := 0; i < N; i++ {
        if <-reqs[i].reqCh != strconv.FormatInt(int64(N + 2*i),10) {
            fmt.Println(i,":error")
		}else{
			fmt.Println(i,":success")
		}
	}
}

func ChainMode()  {
	prime := make(chan int)
	var left, right chan int
	left, right = prime, nil
	go func ()  {
		left <- 1
	}()

	for i := 1; i < 10; i++ {
		right = func (left chan int)  chan int{
			out := make(chan int)
			go func(){
                out <- (<-left + i)
			}()
			return out
		}(left)
		left = <- <-right
	}

	fmt.Println(<-right)
}