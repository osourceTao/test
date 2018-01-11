package aha


import "log"

func getHcNum(x int) int {
	fn := [10]int{6,2,5,5,4,5,6,3,7,6}
	num := 0
	for {
		if x / 10 != 0 {
			num = num + fn[x % 10]
			x = x / 10
		} else {
			break
		}
	}
	num = num + fn[x % 10]
	return num
	
}


func QuanPaiLie(hcNum int) {
	sum := 0
	for i := 0 ; i <= 1111;i++ {
		for j := 0; j <= 1111; j++ {
			k := i + j
			total := getHcNum(i) + getHcNum(j) + getHcNum(k)
			if total == hcNum {
				log.Println(i,j,k)
				sum ++
			}
		}
	}
	log.Println(sum)
}


