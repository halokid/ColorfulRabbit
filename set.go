package ColorfulRabbit
/**
数据结构之set类型
 */

var Exists = struct{}{}

type Set struct {
  m    map[interface{}]struct{}
}

func NewSet() *Set {
  m := make(map[interface{}]struct{})
  return &Set{
    m:      m,
  }
}

func (s *Set) Add(items ...interface{}) error {
  for _, item := range items {
    s.m[item] = Exists
  }
  return nil
}

func (s *Set) Contains(item interface{}) bool {
  _, ok := s.m[item]
  return ok
}

func (s *Set) Size() int {
  return len(s.m)
}

func (s *Set) Clear() {
  s.m = make(map[interface{}]struct{})
}

func (s *Set) Equal(other *Set) bool {
  // 判断两个set是否相等
  if s.Size() != other.Size() {
    return false
  }

  for key := range s.m {
    if !other.Contains(key) {
      return false
    }
  }
  return true
}

func (s *Set) IsSubset(other *Set) bool {
  // 判断s是否是other的子集
  if s.Size() > other.Size() {
    return false
  }

  for key := range s.m {
    if !other.Contains(key) {
      return false
    }
  }
  return true
}












