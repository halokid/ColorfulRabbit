package ColorfulRabbit
/**
配置类使用的函数
 */
import "github.com/Unknwon/goconfig"

type Conf struct {
  //attrsCache  []map[string]interface{}
  EnvFile     string
  ConfFile    string
}

var logx = &Logx{
  DebugFlag:  true,
}

func (cf *Conf) GetEnv() string {
  // 获取开发环境
  cFile, err := goconfig.LoadConfigFile(cf.EnvFile)
  CheckFatal(err, "load env file error")

  env, err := cFile.GetValue("default", "env")
  CheckFatal(err, "not env setting in conf file")
  return env
}



func (cf *Conf) GetVal(key string) string {
  // 根据开发环境配置文件的key获取值
  env := cf.GetEnv()
  //key := env + "_" + keyName

  c, err := goconfig.LoadConfigFile(cf.ConfFile)
  CheckFatal(err, "no conf file")
  val, err := c.GetValue(env, key)
  CheckFatal(err, "no key " + key + " in conf file")
  return val
}


func GetValx(filePath, sec, key string) string {
  // 获取指定的配置文件的配置值
  c, err := goconfig.LoadConfigFile(filePath)
  CheckFatal(err, "no conf file")
  val, err := c.GetValue(sec, key)
  CheckFatal(err, "no key " + key + " in conf file at section " + sec)
  return val
}






