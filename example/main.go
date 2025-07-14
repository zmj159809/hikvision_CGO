package main

import (
	"flag"
	"fmt"
	"unsafe"

	"github.com/zmj159809/hikvision_CGO"
	"github.com/zmj159809/hikvision_CGO/logger"
)

var (
	ip       = flag.String("ip", "", "设备IP地址")
	username = flag.String("name", "admin", "用户名")
	password = flag.String("psw", "ly123456", "密码")
)

// MessCallBack 回调函数接口实现
type MessCallBack struct {
	// 此处添加正常业务逻辑处理返回事件的函数,然后再invoke的时候调用
}

func main() {
	// 初始化设备和日志
	flag.Parse()

	if *ip == "" {
		logger.ErrorArgs("请提供设备IP地址 -ip 192.168.1.100")
		return
	}

	if err := hikvision_CGO.NetInit("./", true); err != nil {
		logger.Errorf("初始化SDK失败: %v", err)
		return
	}
	defer hikvision_CGO.NetCleanup()

	// 登录
	userID, err := hikvision_CGO.NetLoginV40(*ip, *username, *password)
	if err != nil {
		logger.Errorf("登录失败: %v", err)
		return
	}
	logger.Infof("登录成功，userID: %v", userID)

	defer func() {
		if err := hikvision_CGO.NetLogout(userID); err != nil {
			logger.Errorf("登出失败: %v", err)
		}
	}()

	// 注册回调函数
	eventCB := &MessCallBack{}
	eventCBId := hikvision_CGO.NewObjectId(eventCB)
	logger.Infof("创建回调对象ID: %v", eventCBId)

	if err := hikvision_CGO.SetDVRMessCallBack(eventCBId); err != nil {
		logger.Errorf("注册回调函数失败: %v", err)
		return
	}
	logger.InfoArgs("注册回调函数成功")

	var signal string
	var DefenceIds []int32

	for {
		logger.InfoArgs("********请选择功能********")
		logger.InfoArgs("*      1、查询门状态     *")
		logger.InfoArgs("*      2、控制门          *")
		logger.InfoArgs("*      3、查询卡          *")
		logger.InfoArgs("*      4、查询防区状态    *")
		logger.InfoArgs("*      5、布防           *")
		logger.InfoArgs("*      q、退出           *")
		logger.InfoArgs("**************************")

		if _, err := fmt.Scanln(&signal); err != nil {
			logger.Warnf("输入读取失败: %v", err)
			continue
		}

		switch signal {
		case "1":
			handleDoorStatus(userID)
		case "2":
			handleDoorControl(userID)
		case "3":
			handleCardInfo(userID)
		case "4":
			handleAlarmStatus(userID)
		case "5":
			defenceID, err := handleDefence(userID)
			if err != nil {
				logger.Errorf("布防失败: %v", err)
				continue
			}
			DefenceIds = append(DefenceIds, defenceID)
		case "q":
			// 撤防
			for _, v := range DefenceIds {
				hikvision_CGO.CloseDefence(v)
			}
			DefenceIds = make([]int32, 0)
			logger.InfoArgs("退出成功")
			return
		default:
			logger.WarnArgs("输入有误，请重新输入")
		}
	}
}

// handleDoorStatus 处理门状态查询
func handleDoorStatus(userID int32) {
	var status hikvision_CGO.NET_DVR_ACS_WORK_STATUS
	if err := hikvision_CGO.GetDoorStatus(userID, &status); err != nil {
		logger.Errorf("获取门状态失败: %v", err)
		return
	}
	logger.Infof("门状态: %v", status.ST_byMagneticStatus)
}

// handleDoorControl 处理门控制
func handleDoorControl(userID int32) {
	var doorIndex int32
	var ctrl uint32

	fmt.Print("请输入控制门的编号: ")
	if _, err := fmt.Scanln(&doorIndex); err != nil {
		logger.Errorf("输入门编号失败: %v", err)
		return
	}

	fmt.Print("请选择操作(0-关闭，1-打开，2-常开，3-常关): ")
	if _, err := fmt.Scanln(&ctrl); err != nil {
		logger.Errorf("输入控制命令失败: %v", err)
		return
	}

	if err := hikvision_CGO.ControlDoor(userID, doorIndex, ctrl); err != nil {
		logger.Errorf("控制门失败: %v", err)
		return
	}
	logger.InfoArgs("执行操作成功")
}

// handleCardInfo 处理卡信息查询
func handleCardInfo(userID int32) {
	cardInfo, err := hikvision_CGO.GetCardInfo(userID)
	if err != nil {
		logger.Errorf("获取卡信息失败: %v", err)
		return
	}

	if len(cardInfo) == 0 {
		logger.InfoArgs("未找到卡信息")
		return
	}

	logger.Infof("共找到 %d 张卡:", len(cardInfo))
	for cardNo, userName := range cardInfo {
		logger.Infof("卡号: %s | 用户名: %s", cardNo, userName)
	}
}

// handleAlarmStatus 处理报警状态查询
func handleAlarmStatus(userID int32) {
	var aStatus hikvision_CGO.NET_DVR_ALARMHOST_MAIN_STATUS_V51
	if err := hikvision_CGO.GetAlarmHostMainStatus(userID, &aStatus); err != nil {
		logger.Errorf("获取报警状态失败: %v", err)
		return
	}

	logger.InfoArgs("512个防区状态(0xff-无效，0-对应防区当前无报警，1-对应防区当前有报警):")
	logger.Infof("防区状态: %v", aStatus.ST_byAlarmInMemoryStatus)
}

// handleDefence 处理布防
func handleDefence(userID int32) (int32, error) {
	defenceID, err := hikvision_CGO.DoDefence(userID)
	if err != nil {
		return -1, err
	}
	logger.Infof("布防成功，防区ID: %v", defenceID)
	return defenceID, nil
}

// Invoke 回调函数返回信息处理
func (p *MessCallBack) Invoke(lCommand int, ip string, pAlarmInfo unsafe.Pointer, dwBufLen int) bool {
	logger.Debugf("收到回调: Command=%v, IP=%v, BufLen=%v", lCommand, ip, dwBufLen)

	if lCommand == hikvision_CGO.COMM_ALARM_ACS {
		alarmInfo := *(*hikvision_CGO.NET_DVR_ACS_ALARM_INFO)(pAlarmInfo)
		logger.Debugf("报警信息: %v", alarmInfo)

		logger.InfoArgs("门禁主机事件")
		logger.Infof("发生时间: %d-%02d-%02d %02d:%02d:%02d",
			alarmInfo.ST_struTime.ST_dwYear,
			alarmInfo.ST_struTime.ST_dwMonth,
			alarmInfo.ST_struTime.ST_dwDay,
			alarmInfo.ST_struTime.ST_dwHour,
			alarmInfo.ST_struTime.ST_dwMinute,
			alarmInfo.ST_struTime.ST_dwSecond,
		)

		logger.Infof("主类型: %v", alarmInfo.ST_dwMajor.GetMajorString())
		logger.Infof("次类型: %v", alarmInfo.ST_dwMinor.GetMinorString(alarmInfo.ST_dwMajor))
		logger.Infof("报警IP: %v", ip)

		// 提取卡号
		var cardNo string
		for k, v := range alarmInfo.ST_struAcsEventInfo.ST_byCardNo {
			if v == 0 {
				cardNo = string(alarmInfo.ST_struAcsEventInfo.ST_byCardNo[:k])
				break
			}
		}
		logger.Infof("卡号: %v", cardNo)
		logger.Infof("工号: %v", alarmInfo.ST_struAcsEventInfo.ST_dwEmployeeNo)
	}

	return true // 如果return false 会导致认为处理失败重回队列，所以返回true
}
