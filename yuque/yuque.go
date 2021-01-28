package yuque

import (
  "github.com/halokid/ColorfulRabbit"
  "github.com/mozillazg/request"
  //"log"
  "net/http"
)

type Yuque struct {
  Enpoint     string
  XAuth       string
  NameSpace   string
}

func NewYuque() *Yuque {
  return &Yuque{
    Enpoint: EndPoint,
    XAuth:  Xautrh,
    NameSpace: NameSpace,
  }
}

func (y *Yuque) GetDoc(docId, slug string) string {
  // 获取文档
  api := y.GenApi("getDoc", docId, slug)
  //log.Printf("api url------------------ %+v", api)
  bodyHtml := y.DoGet(api)
  return bodyHtml
}

func (y *Yuque) GenApi(act string, docId, slug string) string {
  // 生产API地址
  switch act {
  case "getDoc":
    return y.Enpoint + "repos/" + y.NameSpace + docId + "/docs/" + slug
  default:
    return ""
  }
}

func (y *Yuque) DoGet(api string) string {
  // 执行请求动作
  c := new(http.Client)
  req := request.NewRequest(c)
  y.Auth(req)
  rsp, err := req.Get(api)
  ColorfulRabbit.CheckError(err, "DoGet get api Error")
  //log.Printf("rsp status ----------- %+v", rsp.StatusCode)
  js, err := rsp.Json()
  ColorfulRabbit.CheckError(err, "DoGet body2json Error")
  bodyHtml := js.Get("data").Get("body_html").MustString()
  return bodyHtml
}

func (y *Yuque) Auth(req *request.Request) error {
  req.Headers = map[string]string{
    "X-Auth-Token":     y.XAuth,
  }
  return nil
}







