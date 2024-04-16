package file

import (
	"bufio"
	"os"
	"strings"

	"github.com/halokid/ColorfulRabbit/utils"
)

func readLogFile(logfile string) ([][]string, error) {
  file, err := os.Open(logfile)
  if err != nil {
    return nil, err

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var text []string
  for scanner.Scan() {
    text = append(text, scanner.Text())
  }
  file.Close()

  var lines [][]string
  for i, each_ln := range text {
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
  return lines, nil
}
