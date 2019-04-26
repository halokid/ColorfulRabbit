package ColorfulRabbit

// 获取slice某元素的index
func IndexOf(element string, data []string) (int) {
  for k, v := range data {
    if element == v {
      return k
    }
  }
  return -1 //not found.
}




