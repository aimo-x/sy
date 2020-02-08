package app

import (
	"encoding/csv"
	"os"
)

// Export ...
type Export struct {
}

// CSV 导出csv
func (ex *Export) CSV(filePath string, data [][]string) (err error) {
	fp, err := os.Create("./" + filePath) // 创建文件句柄
	if err != nil {
		return
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(fp)         //创建一个新的写入文件流
	err = w.WriteAll(data)
	if err != nil {
		return err
	}
	w.Flush()
	return
}
