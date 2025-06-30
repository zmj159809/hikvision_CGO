package main

import (
	"flag"
	"fmt"
	"github.com/zmj159809/hikvision_CGO/netsdk"
	"unsafe"
)

var (
	ip       = flag.String("ip", "", "ip")
	username = flag.String("name", "admin", "username")
	password = flag.String("psw", "ly123456", "password")
)

type MessCallBack struct {
}

func main() {
	//初始化设备和日志
	flag.Parse()
	err := netsdk.NetInit("./", true) //日志路径
	defer netsdk.NetCleanup()

	//登录
	userID, err := netsdk.NetLoginV40(*ip, *username, *password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("登陆成功，userID ： ", userID)
	defer func() {
		err = netsdk.NetLogout(userID)
		if err != nil {
			panic(fmt.Sprintf("logout err ：%v", err))
		}
	}()

	//注册回调函数
	eventCB := &MessCallBack{}
	eventCBId := netsdk.NewObjectId(eventCB)
	fmt.Println(eventCBId)
	err = netsdk.SetDVRMessCallBack(eventCBId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("注册回调函数成功")

	var signal string
	var DefenceIds []int32
	for {

		fmt.Println("********请选择功能********")
		fmt.Println("*      1、查询门状态     *")
		fmt.Println("*      2、控制门          *")
		fmt.Println("*      3、查询卡          *")
		fmt.Println("*      4、查询防区状态    *")
		fmt.Println("*      5、布防           *")
		fmt.Println("*      q、退出           *")
		fmt.Println("**************************")
		_, _ = fmt.Scanln(&signal)
		switch signal {
		case "1":
			var status netsdk.NET_DVR_ACS_WORK_STATUS
			err := netsdk.GetDoorStatus(userID, &status)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(status.ST_byMagneticStatus)
			}
		case "2":
			var doorIndex int32
			var ctrl uint32
			fmt.Println("请输入控制门的编号")
			_, _ = fmt.Scanln(&doorIndex)
			fmt.Println("请选择操(0- 关闭，1- 打开，2 -常开，3- 常关)")
			_, _ = fmt.Scanln(&ctrl)
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
			for k, v := range cardInfo {
				fmt.Printf("CardNo: %s | UserName : %s \n", k, v)
			}

		case "4":
			var AStatus netsdk.NET_DVR_ALARMHOST_MAIN_STATUS_V51
			err := netsdk.GetAlarmHostMainStatus(userID, &AStatus)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("512个防区状态(0xff-无效，0-对应防区当前无报警，1-对应防区当前有报警)：")
				fmt.Println(AStatus.ST_byAlarmInMemoryStatus)
				//fmt.Println("状态全部内容")
				//fmt.Println(AStatus)
			}
		case "5":
			//布防
			DefenceId, err := netsdk.DoDefence(userID)
			if err != nil {
				fmt.Println(err)
				return
			}
			DefenceIds = append(DefenceIds, DefenceId)
			fmt.Println("布防成功  防区id:", DefenceId)

		case "q":
			//撤防
			for _, v := range DefenceIds {
				netsdk.CloseDefence(v)
			}
			DefenceIds = make([]int32, 0)
			fmt.Println("退出成功")
			return
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}

// Invoke 回调函数返回信息处理
func (p *MessCallBack) Invoke(lCommand int, ip string, pAlarmInfo unsafe.Pointer, dwBufLen int) bool {

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

		fmt.Println("主类型：", AlarmInfo.ST_dwMajor.GetMajorString())
		fmt.Println("次类型：", AlarmInfo.ST_dwMinor.GetMinorString(AlarmInfo.ST_dwMajor))
		fmt.Println("报警IP：", ip)
		var cardNo string
		for k, v := range AlarmInfo.ST_struAcsEventInfo.ST_byCardNo {
			if v == 0 {
				cardNo = string(AlarmInfo.ST_struAcsEventInfo.ST_byCardNo[:k])
			}
		}
		fmt.Println("卡号为：", cardNo)
		fmt.Println("工号为：", fmt.Sprint(AlarmInfo.ST_struAcsEventInfo.ST_dwEmployeeNo))
		//if AlarmInfo.ST_byAcsEventInfoExtend == 1 {
		//	ExtendInfo := *(*netsdk.NET_DVR_ACS_EVENT_INFO_EXTEND)(AlarmInfo.ST_pAcsEventInfoExtend)
		//	fmt.Println("门禁主机扩展事件信息 ： ", ExtendInfo)
		//}
	}
	return true
}
