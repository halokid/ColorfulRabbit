package ColorfulRabbit
/**
生成密码
 */
import (
  "fmt"
  "math/rand"
)

//var (
//  length  int
//  charset string
//)

const (
  NUmStr  = "0123456789"
  CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
  SpecStr = "+=-@#~,.[]()!%^*$"
)

/**
//解析参数
func parseArgs() {
  //需要接受指针，就传递地址,&
  flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
  flag.StringVar(&charset, "t", "num",
    //反引号以原样输出
    `-t 制定密码生成的字符集,
        num:只使用数字[0-9],
        char:只使用英文字母[a-zA-Z],
        mix:使用数字和字母，
        advance:使用数字、字母以及特殊字符`)
  flag.Parse()
}
*/

//检测字符串中的空格
func test1() {
  for i := 0; i < len(CharStr); i++ {
    if CharStr[i] != ' ' {
      fmt.Printf("%c", CharStr[i])
    }
  }
}

func GenerPwd(charset string, length int) string {
  //初始化密码切片
  var passwd []byte = make([]byte, length, length)
  //源字符串
  var sourceStr string
  //判断字符类型,如果是数字
  if charset == "num" {
    sourceStr = NUmStr
    //如果选的是字符
  } else if charset == "char" {
    sourceStr = charset
    //如果选的是混合模式
  } else if charset == "mix" {
    sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)
    //如果选的是高级模式
  } else if charset == "advance" {
    sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
  } else {
    sourceStr = NUmStr
  }
  fmt.Println("source:", sourceStr)

  //遍历，生成一个随机index索引,
  for i := 0; i < length; i++ {
    index := rand.Intn(len(sourceStr))
    passwd[i] = sourceStr[index]
  }
  return string(passwd)
}






