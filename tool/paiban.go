package main

import (
	"fmt"

	"github.com/golang-module/carbon"
)

func main() {
	carbon.SetLocale("Asia/Shanghai")
	startTime := carbon.NewCarbon().SetDateTime(2024, 4, 17, 18, 0, 0).Timestamp()
	endTime := carbon.NewCarbon().SetDateTime(2024, 5, 6, 18, 0, 0).Timestamp()

	isWork := true
	i := 0
	addHour := 12 // 默认工作12小时
	currentTime := startTime
	var action string

	for currentTime < endTime {
		//下一阶段时间
		addHour = 24
		if i == 0 || i%2 == 0 {
			addHour = 12
		}

		currentCarbon := carbon.NewCarbon().CreateFromTimestamp(currentTime)
		nextCarbon := currentCarbon.AddHours(addHour)

		//工作休息标识
		action = "休息"
		if isWork {
			action = "上班"
		}

		//if !isWork {
		fmt.Printf("%s 至 %s %s(%d小时)\n", currentCarbon.Format("m-d H:00"), nextCarbon.Format("m-d H:00"), action, addHour)
		//}

		isWork = !isWork
		currentTime = nextCarbon.Timestamp()
		i++
	}
}
