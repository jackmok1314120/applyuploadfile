package utils

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type Coins struct {
	Sort      string `json:"sort"`
	ChainName string `json:"chain_name"`
	CoinName  string `json:"coin_name"`
	FullName  string `json:"full_name"`
}

func ExcelToDB(inFile string) []Coins {
	var (
		ct []Coins
	)

	// 打开文件
	xlFile, err := xlsx.OpenFile(inFile)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		//遍历行读取
		r := 0
		for _, row := range sheet.Rows {
			c := Coins{}
			count := 0
			// 遍历每行的列读取
			for _, cell := range row.Cells {
				text := cell.String()
				if count == 0 {
					c.Sort = text
				}
				if count == 1 {
					c.ChainName = text
				}
				if count == 2 {
					c.CoinName = text
				}
				if count == 3 {
					c.FullName = text
				}
				count++
				fmt.Printf("%20s", text)
			}
			if r != 0 {
				ct = append(ct, c)
			}
			r = 1
			fmt.Print("\n")
		}
	}
	fmt.Println("\n\nimport success")
	return ct
}
