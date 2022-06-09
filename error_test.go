package main

import (
    "context"
    "errors"
    "fmt"
    "golang.org/x/sync/errgroup"
    "log"
    "testing"
    "time"
    errs "github.com/pkg/errors"
)

func TestError(t *testing.T) {
    f := &FuImp{}
    if err := doWork(f); err != nil {
        log.Fatal(err)
    }
}

type Fu interface {
    Work1() error
    Work2() error
    Work3() error
    Errs() error
}

func doWork(v Fu) error {
    v.Work1()
    v.Work2()
    v.Work3()

    // 只在最后几种处理err
    if err := v.Errs(); err != nil {
        return err
    }
    return nil
}

type FuImp struct{}

func (f *FuImp) Work1() error { return nil }
func (f *FuImp) Work2() error { return nil }
func (f *FuImp) Work3() error { return nil }
func (f *FuImp) Errs() error  { return nil }

/**
  用errorgroup作为error的处理,wait 会返回子协程的error（子协程第一个出现的error），而且所有的子协程都会保证执行完
    所以使用errorgroup可以达到追踪子协程错误(只能知道第一个出现的err)
*/
func TestLsd(t *testing.T) {
    eg, ctx := errgroup.WithContext(context.Background())
    for i := 0; i < 100; i++ {
        i := i
        /**
          errorgroup的Go方法有个特点：
          就是在执行传入Go的方法时，如果传入的方法返回的error不是nil，也就是有error出现，
              就会把在调用WithContext时生成的和eg对应的context取消掉
        */
        eg.Go(func() error { // 会启动一个goroutine去执行的func
            time.Sleep(2 * time.Second)
            select {
            case <-ctx.Done():
                fmt.Println("Canceled:", i)
                return nil
            default:
                fmt.Println("End:", i)
                return nil
            }
        })
    }
    if err := eg.Wait(); err != nil { // 不论是否有error，wait会阻塞直到所有启动的goroutine都返回，只不过wait只返回第一个出现的err而已
        log.Fatal(err) // 返回的错误是一个出错的 err,其他的错误会被丢弃
    }
    <-ctx.Done()
    fmt.Println("over")
}


func TestLel(t *testing.T) {
    //controller
    if err := Controller(); err != nil {
        log.Printf("failed. err: %+v", err) // 打印err便可看到堆栈的信息
        log.Printf("failed. err: %v", errs.Cause(err)) // 根error
    } // 不在需要每层都打印了，而是层层追加error信息到堆栈里，在最后打印
}
func Controller() error {
    fmt.Println("执行controller层")
    if err := Service(); err != nil {
        return errs.WithMessage(err, "service failed \n")
    }
    return nil
}
func Service() error {
    fmt.Println("执行service层")
    if err := Dao(); err != nil { // 不需要每层都打印err，直接把本层的error信息追加到上层的err里就行了
        return errs.WithMessage(err, "Dao failed \n") // 在service层只要用withMessage追加信息就行了
    }
    return nil
}
func Dao() error {
    fmt.Println("执行dao层")
    // Wrapf会附加堆栈和信息到error上
    return errs.Wrapf(errors.New("sql failed") , "query sql : %s", "select * form t" )
}
