package ColorfulRabbit

import "github.com/Unknwon/goconfig"

type Conf struct {
  //attrsCache  []map[string]interface{}
  EnvFile     string
  ConfFile    string
}

var logx = &Logx{
  DebugFlag:  true,
}

// 获取开发环境
func (cf *Conf) GetEnv() string {
  cFile, err := goconfig.LoadConfigFile(cf.EnvFile)
  CheckFatal(err, "load env file error")

  env, err := cFile.GetValue("default", "env")
  CheckFatal(err, "not env setting in conf file")
  return env
}



// 根据配置文件key获取值
func (cf *Conf) GetVal(key string) string {
  env := cf.GetEnv()
  //key := env + "_" + keyName

  c, err := goconfig.LoadConfigFile(cf.ConfFile)
  CheckFatal(err, "no conf file")
  val, err := c.GetValue(env, key)
  CheckFatal(err, "no key " + key + " in conf file")
  return val
}







