package main

import (
	"flag"
	"fmt"
	"hkdoorSDK/netsdk"
	"unsafe"
)

var (
	ip       = flag.String("ip", "", "ip")
	username = flag.String("name", "admin", "username")
	password = flag.String("psw", "ly123456", "password")
)

type ST_fMessCallBack struct {
}

func main() {
	//初始化设备和日志
	flag.Parse()
	err := netsdk.NetInit("./") //日志路径
	defer netsdk.NetCleanup()

	//登录
	userID, err := netsdk.NetLoginV40(*ip, *username, *password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("登陆成功，userID ： ", userID)
	defer netsdk.NetLogout(userID)

	//注册回调函数
	eventCB := &ST_fMessCallBack{}
	eventCBId := netsdk.NewObjectId(eventCB)
	fmt.Println(eventCBId)
	err = netsdk.SetDVRMessCallBack(eventCBId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("注册回调函数成功")

	//布防
	DefenceId, err := netsdk.DoDefence(userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(DefenceId)
	defer netsdk.CloseDefence(DefenceId)

	var signal string

	for {

		fmt.Println("********请选择功能********")
		fmt.Println("*      1、查询门状态     *")
		fmt.Println("*      2、控制门          *")
		fmt.Println("*      3、查询卡          *")
		fmt.Println("*      q、退出           *")
		fmt.Println("**************************")
		fmt.Scanln(&signal)
		switch signal {
		case "1":
			status, err := netsdk.GetDoorStatus(userID)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(status.ST_byMagneticStatus)
				fmt.Println("状态全部内容")
				fmt.Println(status)
			}
		case "2":
			var doorIndex int32
			var ctrl uint32
			fmt.Println("请输入控制门的编号")
			fmt.Scanln(&doorIndex)
			fmt.Println("请选择操(0- 关闭，1- 打开，2 -常开，3- 常关)")
			fmt.Scanln(&ctrl)
			err := netsdk.ControlDoor(userID, doorIndex, ctrl)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("执行操作成功")
			}
		case "3":
			cardInfo, err := netsdk.GetCardInfo(userID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cardInfo)
		case "q":
			fmt.Println("退出登录")
			return
		default:
			fmt.Println("输入有误，请重新输入")

		}

	}

}

// Invoke 回调函数返回信息处理
func (p *ST_fMessCallBack) Invoke(lCommand int, ip string, pAlarmInfo unsafe.Pointer, dwBufLen int) bool {

	fmt.Println(lCommand, ip, pAlarmInfo, dwBufLen)
	if lCommand == netsdk.COMM_ALARM_ACS {
		AlarmInfo := *(*netsdk.NET_DVR_ACS_ALARM_INFO)(pAlarmInfo)
		fmt.Println(AlarmInfo)
		fmt.Println("门禁主机事件")
		fmt.Println("发生时间:",
			AlarmInfo.ST_struTime.ST_dwYear,
			AlarmInfo.ST_struTime.ST_dwMonth,
			AlarmInfo.ST_struTime.ST_dwDay,
			AlarmInfo.ST_struTime.ST_dwHour,
			AlarmInfo.ST_struTime.ST_dwMinute,
			AlarmInfo.ST_struTime.ST_dwSecond,
		)

		fmt.Println("主类型：", AlarmInfo.ST_dwMajor)
		fmt.Println("次类型：", AlarmInfo.ST_dwMinor)
		fmt.Println("报警IP：", ip)
		fmt.Println("卡号为：", AlarmInfo.ST_struAcsEventInfo.ST_byCardNo)
		fmt.Println("工号为：", AlarmInfo.ST_struAcsEventInfo.ST_dwEmployeeNo)
		if AlarmInfo.ST_byAcsEventInfoExtend == 1 {
			ExtendInfo := *(*netsdk.NET_DVR_ACS_EVENT_INFO_EXTEND)(AlarmInfo.ST_pAcsEventInfoExtend)
			fmt.Println("门禁主机扩展事件信息 ： ", ExtendInfo)
		}
	}
	return true
}
