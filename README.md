## 可輸入數值陣列
在迴圈內加入兩個巢狀迴圈
並設定為可自行輸入數值狀態

用以練習迴圈實作
```golang
for i := 0; i < num; i++ {
		for e := 0; e < i; e++ {
			fmt.Printf("*")
		}
		for f := num - i; f > 0; f-- {
			fmt.Printf("-")
		}
		fmt.Println(" ")
	}
	for i := 1; i < num; i++ {
		for j := num - i; j > 0; j-- {
			fmt.Printf("*")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("-")
		}
		fmt.Println(" ")
```