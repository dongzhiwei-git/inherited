package main

import "fmt"

func main() {
	//f, err := excelize.OpenFile("txt.xlsx")
	//if err != nil {
	//	println(err.Error())
	//	fmt.Println("er")
	//	return
	//}
	// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("txt001", "B2")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//println(cell)

	//获取 Sheet1 上所有单元格
	//rows, err := f.GetRows("txt001")
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		print(colCell, "\t")
	//	}
	//	println()
	//}
	//计时器
	//t1 := time.Now()
	//rows, _ := f.GetRows("txt001")
	//for err, _ := range rows {
	//	//fmt.Println(err)
	//	if err == 0 {
	//		fmt.Printf("2")
	//	}
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//	//println(row[0], row[1], row[2], row[3], row[4])
	//
	//}
	//elapsed := time.Since(t1)
	//fmt.Println("App elapsed: ", elapsed)

	//并发读
	//t1 := time.Now()
	//wg := sync.WaitGroup{}
	//wg.Add(16010)
	//count := 0
	//rows, _ := f.GetRows("txt001")
	//for _, row := range rows {
	//	go func(i []string) {
	//		if i == nil {
	//			fmt.Printf("2")
	//		}
	//		wg.Done()
	//	}(row)
	//
	//}
	//wg.Wait()
	//elapsed := time.Since(t1)
	//fmt.Println("App elapsed: ", elapsed, "count=", count)

	ch := make(chan int, 10)

	//ch <- 1
	<-ch
	fmt.Println(1)

}
