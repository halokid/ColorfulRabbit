package main

import (
	"context"
	"net/http"
)

type Fookey string

var Username = Fookey("user-name")
var Userid = Fookey("user-id")

func foo(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), Userid, "1")
    ctx2 := context.WithValue(ctx, Username, "Spoo")
    next(w, r.WithContext(ctx2))
  }
}

func GetUserName (ctx context.Context) string {
  if ret, ok := ctx.Value(Username).(string); ok {
    return ret
  }
  return ""
}

func GetUserId (ctx context.Context) string {
  if ret, ok := ctx.Value(Userid).(string); ok {
    return ret
  }
  return ""
}

func test(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("welcome: "))
  w.Write([]byte(GetUserId(r.Context())))
  w.Write([]byte(" "))
  w.Write([]byte(GetUserName(r.Context())))
}

func c2() {
  http.Handle("/", foo(test))
  http.ListenAndServe(":8080", nil)
}



