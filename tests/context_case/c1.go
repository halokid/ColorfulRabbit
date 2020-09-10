package main

import (
  "context"
  "github.com/pkg/errors"
  "sync"
)

/**
context的使用场景之 RPC调用
 */

func Rpc(ctx context.Context, url string) error {
  result := make(chan int)
  err := make(chan error)

  go func() {
    // call rpc, if success set to result, if fail set to err
    isSuccess := true
    if isSuccess {
      result <-1
    } else {
      err <-errors.New("some error happen")
    }
  }()

  // todo: 一旦哪个rpc产生错误， 所有的rpc都退出
  select {
  case <-ctx.Done():
    // other rpc call fail
    return ctx.Err()
  case e := <-err:
    // local rpc call fail, return err
    return e
  case <-result:
    // local rpc success, dont return err
    return nil
  }
}

func main() {
  ctx, cancel := context.WithCancel(context.Background())

  // rpc1 call
  err := Rpc(ctx, "http://rpc1")
  if err != nil {
    return
  }

  wg := sync.WaitGroup{}

  // rpc2 call
  wg.Add(1)
  go func() {
    defer wg.Done()
    err := Rpc(ctx, "http:/rpc2")
    if err != nil {
      cancel()
    }
  }()

  // rpc3 call
  wg.Add(1)
  go func() {
    defer wg.Done()
    err := Rpc(ctx, "http://rpc3")
    if err != nil {
      cancel()
    }
  }()

  // rpc4 call
  wg.Add(1)
  go func() {
    defer wg.Done()
    err := Rpc(ctx, "http://rpc4")
    if err != nil {
      cancel()
    }
  }()

  wg.Wait()
}










