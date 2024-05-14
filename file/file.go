package file

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filepath string) ([]string, error) {
  file, err := os.Open(filepath)
  if err != nil {
    return nil, err
  }

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var text []string
  for scanner.Scan() {
    text = append(text, scanner.Text())
  }
  file.Close()
  log.Println("1:", len(text))

  /*
  var lines [][]string
  for i, each_ln := range text {
    log.Println("each line:", each_ln)
    if i > 0 {
      sl := strings.Split(each_ln, "|")
      sl = utils.RemoveEmptyElements(sl)

      for i, slItem := range sl {
        sl[i] = strings.Replace(slItem, " ", "", -1)
        sl[i] = strings.Replace(sl[i], "\t", "", -1)
      }

      if len(sl) > 1 {
        lines = append(lines, sl)
      }
    }
  }
  */
  return text, nil
}


