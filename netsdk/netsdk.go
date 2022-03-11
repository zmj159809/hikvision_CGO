package netsdk

/*
#cgo CFLAGS:  -I./include
#cgo LDFLAGS: -L./lib/Linux -lhcnetsdk
#cgo LDFLAGS: -Wl,-rpath=./lib/Linux:./lib/Linux/HCNetSDK

#include "HCNetSDK.h"
#include <unistd.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
static int ST_Defence(int uid);
extern int MessageCallback(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *pAlarmInfo, DWORD dwBufLen, void* pUser);

int ST_Defence(int uid1){

	NET_DVR_SETUPALARM_PARAM struSetupParam={0};
	struSetupParam.dwSize=sizeof(NET_DVR_SETUPALARM_PARAM);
    struSetupParam.byDeployType = 1;//设置布防类型为实时布防

	LONG mHandle = NET_DVR_SetupAlarmChan_V41(uid1,&struSetupParam);


    return mHandle;
}
*/
import "C"
import (
	"errors"
	"fmt"
	"log"
	"unsafe"
)

// IF_fMEssCallBack 回调函数结构体
type IF_fMEssCallBack interface {
	Invoke(lCommand int, ip string, pBuf unsafe.Pointer, dwBufLen int) bool
}

// NetInit 初始化设备和日志
func NetInit(sdkLog string) error {

	//初始化资源
	ret := C.NET_DVR_Init()
	if int(ret) != 1 {
		fmt.Printf("NET_DVR_Init failed,error code = %v\n", C.NET_DVR_GetLastError())
		return errors.New(fmt.Sprintf("NET_DVR_Init failed,error code = %v\n", C.NET_DVR_GetLastError()))
	}
	C.NET_DVR_SetLogToFile(3, C.CString(sdkLog), 1)

	return nil
}

// NetLoginV40 登录
func NetLoginV40(deviceIp, username, password string) (int32, error) {
	var userLoginInfo C.NET_DVR_USER_LOGIN_INFO
	var deviceInfo C.NET_DVR_DEVICEINFO_V40

	userLoginInfo.wPort = 8000 // your device port,default 8000

	pUsername := C.CBytes([]byte(username))
	defer C.free(pUsername)
	// 使用memcpy函数进行拷贝
	C.memcpy(unsafe.Pointer(&userLoginInfo.sUserName), pUsername, C.ulong(len(username)))

	pPassword := C.CBytes([]byte(password))
	defer C.free(pPassword)
	C.memcpy(unsafe.Pointer(&userLoginInfo.sPassword), pPassword, C.ulong(len(password)))

	pDeviceIp := C.CBytes([]byte(deviceIp))
	defer C.free(pDeviceIp)
	C.memcpy(unsafe.Pointer(&userLoginInfo.sDeviceAddress), pDeviceIp, C.ulong(len(deviceIp)))

	// 调用登录接口
	uid := C.NET_DVR_Login_V40((C.LPNET_DVR_USER_LOGIN_INFO)(&userLoginInfo), (C.LPNET_DVR_DEVICEINFO_V40)(&deviceInfo))

	if int32(uid) < 0 {
		if err := isErr("Login"); err != nil {
			return -1, errors.New(fmt.Sprintf("ip: %s 登录失败,原因%v", deviceIp, err.Error()))
		}
		return -1, errors.New(fmt.Sprintf("ip: %s 登录失败", deviceIp))
	}
	return int32(uid), nil
}

// isErr  获取上一个发生的错误
func isErr(operation string) error {
	errno := int64(C.NET_DVR_GetLastError())
	if errno > 0 {
		reMsg := fmt.Sprintf("%s失败,失败代码号：%d", operation, errno)
		return errors.New(reMsg)
	}
	return nil
}

// NetLogout 退出登录
func NetLogout(uid int32) error {
	C.NET_DVR_Logout_V30(C.LONG(uid))
	if err := isErr("Logout"); err != nil {
		return err
	}
	return nil
}

// NetCleanup 释放SDK资源，在程序结束之前调用。
func NetCleanup() {
	C.NET_DVR_Cleanup()
}

//GetDoorStatus  获取门状态
func GetDoorStatus(uid int32, pBuf unsafe.Pointer) (err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("GetDoorStatus panic : ", e)
		}
	}()
	var lpBytesReturned C.DWORD
	var PBuf = (C.LPVOID)(pBuf)
	ret := C.NET_DVR_GetDVRConfig(C.LONG(uid), C.NET_DVR_GET_ACS_WORK_STATUS, C.LONG(0xFFFFFFF), PBuf, (C.DWORD)(unsafe.Sizeof(NET_DVR_ACS_WORK_STATUS{})), (C.LPDWORD)(&lpBytesReturned))

	if int32(ret) == 0 {
		if err = isErr("Get door status"); err != nil {
			return errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
		}
		return errors.New(fmt.Sprintf("get door status error，uid :%d", uid))
	}

	return nil
}

//export MessageCallback
func MessageCallback(lCommand C.int, pAlarmer *C.struct_tagNET_DVR_ALARMER, pAlarmInfo *C.char, dwBufLen C.DWORD, pUser unsafe.Pointer) C.int {
	defer func() {
		if e := recover(); e != nil {
			log.Println("MessageCallback panic : ", e)
		}
	}()
	id := *(*ObjectId)(pUser)
	interf := id.Get()
	obj, ok := interf.(IF_fMEssCallBack)
	Alarm := *(*NET_DVR_ALARMER)(unsafe.Pointer(pAlarmer))
	ip := Alarm.ST_sDeviceIP
	var ipStr string
	for k, v := range ip {
		if v == 0 {
			ipStr = string(ip[:k])
			break
		}
	}
	if ok {
		ret := obj.Invoke(int(lCommand), ipStr, unsafe.Pointer(pAlarmInfo), int(dwBufLen))
		if ret {
			return 1
		}
	}

	return 0
}

// SetDVRMessCallBack 设置回调函数
func SetDVRMessCallBack(dwUser ObjectId) error {
	if dwUser.IsNil() {
		log.Println("dwUser is nil")
		return errors.New("dwUser is nil")
	}
	ret := C.NET_DVR_SetDVRMessageCallBack_V31(C.MSGCallBack_V31(C.MessageCallback), unsafe.Pointer(&dwUser))
	if int32(ret) == 0 {
		//注册失败
		return errors.New("注册回调函数失败")
	}

	return nil
}

// DoDefence 布防
func DoDefence(uid int32) (int32, error) {
	ret := C.ST_Defence(C.int(uid))
	if int32(ret) == -1 {
		err := isErr("布防")
		return int32(ret), err
	}
	fmt.Println("布防成功")
	return int32(ret), nil
}

// CloseDefence 撤防
func CloseDefence(dId int32) {
	C.NET_DVR_CloseAlarmChan_V30(C.int(dId))
}

// ControlDoor 控制门状态 0- 关闭，1- 打开），2- 常开），3- 常关，
func ControlDoor(uid int32, DoorIndex int32, ctrl uint32) error {
	defer func() {
		if e := recover(); e != nil {
			log.Println("ControlDoor panic : ", e)
		}
	}()
	ret := C.NET_DVR_ControlGateway(C.LONG(uid), C.LONG(DoorIndex), C.DWORD(ctrl))
	if int32(ret) == 0 {
		if err := isErr("Control door"); err != nil {
			return errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
		}
		return errors.New(fmt.Sprintf("Control door error，uid :%d", uid))
	}
	return nil
}

// GetCardInfo 获取门禁设备下的卡信息
func GetCardInfo(uid int32) (cardInfo map[string]string, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("GetCardInfo panic : ", e)
		}
	}()
	var getcard NET_DVR_CARD_COND
	getcard.ST_dwSize = DWORD(unsafe.Sizeof(NET_DVR_CARD_COND{}))
	getcard.ST_dwCardNum = 0xffffffff
	cardInfo = make(map[string]string)

	var CardStruct C.NET_DVR_CARD_RECORD
	var cardStruct NET_DVR_CARD_RECORD
	//var userDate byte
	//建立长连接
	ret1 := C.NET_DVR_StartRemoteConfig(C.LONG(uid), C.NET_DVR_GET_CARD, (C.LPVOID)(unsafe.Pointer(&getcard)), C.DWORD(unsafe.Sizeof(NET_DVR_CARD_COND{})), nil, nil)
	if int32(ret1) < 0 {
		if err := isErr("StartRemoteConfig"); err != nil {
			return cardInfo, errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
		}
		return cardInfo, errors.New(fmt.Sprintf("GetCard Info error，uid :%d", uid))
	}

	//循环获取卡数据
	for {
		ret2 := int32(C.NET_DVR_GetNextRemoteConfig(ret1, unsafe.Pointer(&CardStruct), C.DWORD(1300)))
		//fmt.Println(*(*NET_DVR_CARD_RECORD)(unsafe.Pointer(&CardStruct)))
		//fmt.Println(unsafe.Sizeof(NET_DVR_CARD_RECORD{}))
		//fmt.Println("NET_DVR_GetNextRemoteConfig 返回结果",ret2)
		if int64(ret2) < 0 {
			if err := isErr("GetNextRemoteConfig1"); err != nil {
				return cardInfo, errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
			}
			return cardInfo, errors.New(fmt.Sprintf("GetCard Info error，uid :%d", uid))
		}
		if ret2 == 1000 {
			cardStruct = *(*NET_DVR_CARD_RECORD)(unsafe.Pointer(&CardStruct))
			for k, v := range cardStruct.byCardNo {
				if v == 0 {
					cardInfo[string(cardStruct.byCardNo[:k])] = string(cardStruct.byName[:])
					break
				}
			}
			fmt.Println("get card info :cardNo:", string(cardStruct.byCardNo[:]), "|name : ", string(cardStruct.byName[:]))
			continue
		}
		if ret2 == 1001 {
			continue
		}
		if ret2 == 1002 {
			ret3 := C.NET_DVR_StopRemoteConfig(ret1)
			if int32(ret3) == 0 {
				if err := isErr("GetNextRemoteConfig"); err != nil {
					return cardInfo, errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
				}
				return cardInfo, errors.New(fmt.Sprintf("Control door error，uid :%d", uid))
			}
			return cardInfo, nil
		}
		if ret2 == 1003 {
			ret3 := C.NET_DVR_StopRemoteConfig(ret1)
			if int32(ret3) == 0 {
				if err := isErr("GetNextRemoteConfig"); err != nil {
					return cardInfo, errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
				}
				return cardInfo, errors.New(fmt.Sprintf("Control door error，uid :%d", uid))
			}
			return cardInfo, nil
		}

	}

}

// GetAlarmHostMainStatus 获取报警主机的状态信息结构体
func GetAlarmHostMainStatus(uid int32, status unsafe.Pointer) (err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("GetAlarmHostMainStatus panic : ", e)
		}
	}()
	var AlarmHostStatus = (C.LPVOID)(status)
	var lpBytesReturned C.DWORD
	ret := C.NET_DVR_GetDVRConfig(C.LONG(uid), C.NET_DVR_GET_ACS_WORK_STATUS, C.LONG(0xFFFFFFF), AlarmHostStatus, (C.DWORD)(unsafe.Sizeof(NET_DVR_ALARMHOST_MAIN_STATUS_V51{})), (C.LPDWORD)(&lpBytesReturned))

	if int32(ret) == 0 {
		if err = isErr("Get door status"); err != nil {
			return errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
		}
		return errors.New(fmt.Sprintf("get door status error，uid :%d", uid))
	}
	return nil
}

// GetMajorString 返回主类型字符串
func (d DWORD) GetMajorString() string {
	switch d {
	case MAJOR_ALARM:
		return "报警"
	case MAJOR_EXCEPTION:
		return "异常"
	case MAJOR_OPERATION:
		return "操作"
	case MAJOR_EVENT:
		return "事件"
	default:
		return "unknown"
	}
}

// GetMinorString 获取告警次类型
func (d DWORD) GetMinorString(major DWORD) string {
	switch major {
	case 1: //报警
		switch d {
		case MINOR_ALARMIN_SHORT_CIRCUIT:
			return "防区短路报警"
		case MINOR_ALARMIN_BROKEN_CIRCUIT:
			return "防区断路报警"
		case MINOR_ALARMIN_EXCEPTION:
			return "防区异常报警"
		case MINOR_ALARMIN_RESUME:
			return "防区报警恢复"
		case MINOR_HOST_DESMANTLE_ALARM:
			return "设备防拆报警"
		case MINOR_HOST_DESMANTLE_RESUME:
			return "设备防拆恢复"
		case MINOR_CARD_READER_DESMANTLE_ALARM:
			return "读卡器防拆报警"
		case MINOR_CARD_READER_DESMANTLE_RESUME:
			return "读卡器防拆恢复"
		case MINOR_CASE_SENSOR_ALARM:
			return "事件输入报警"
		case MINOR_CASE_SENSOR_RESUME:
			return "事件输入恢复"
		case MINOR_STRESS_ALARM:
			return "胁迫报警"
		case MINOR_OFFLINE_ECENT_NEARLY_FULL:
			return "离线事件满90%报警"
		case MINOR_CARD_MAX_AUTHENTICATE_FAIL:
			return "卡号认证失败超次报警"
		case MINOR_SD_CARD_FULL:
			return "SD卡存储满报警"
		case MINOR_LINKAGE_CAPTURE_PIC:
			return "联动抓拍事件报警"
		case MINOR_SECURITY_MODULE_DESMANTLE_ALARM:
			return "门控安全模块防拆报警"
		case MINOR_SECURITY_MODULE_DESMANTLE_RESUME:
			return "门控安全模块防拆恢复"
		case MINOR_POS_START_ALARM:
			return "POS开启"
		case MINOR_POS_END_ALARM:
			return "POS结束"
		case MINOR_FACE_IMAGE_QUALITY_LOW:
			return "人脸图像画质低"
		case MINOR_FINGE_RPRINT_QUALITY_LOW:
			return "指纹图像画质低"
		case MINOR_FIRE_IMPORT_SHORT_CIRCUIT:
			return "消防输入短路报警"
		case MINOR_FIRE_IMPORT_BROKEN_CIRCUIT:
			return "消防输入断路报警"
		case MINOR_FIRE_IMPORT_RESUME:
			return "消防输入恢复"
		case MINOR_FIRE_BUTTON_TRIGGER:
			return "消防按钮触发"
		case MINOR_FIRE_BUTTON_RESUME:
			return "消防按钮恢复"
		case MINOR_MAINTENANCE_BUTTON_TRIGGER:
			return "维护按钮触发"
		case MINOR_MAINTENANCE_BUTTON_RESUME:
			return "维护按钮恢复"
		case MINOR_EMERGENCY_BUTTON_TRIGGER:
			return "紧急按钮触发"
		case MINOR_EMERGENCY_BUTTON_RESUME:
			return "紧急按钮恢复"
		case MINOR_DISTRACT_CONTROLLER_ALARM:
			return "分控器防拆报警"
		case MINOR_DISTRACT_CONTROLLER_RESUME:
			return "分控器防拆报警恢复"
		case MINOR_CHANNEL_CONTROLLER_DESMANTLE_ALARM:
			return "通道控制器防拆报警"
		case MINOR_CHANNEL_CONTROLLER_DESMANTLE_RESUME:
			return "通道控制器防拆报警恢复"
		case MINOR_CHANNEL_CONTROLLER_FIRE_IMPORT_ALARM:
			return "通道控制器消防输入报警"
		case MINOR_CHANNEL_CONTROLLER_FIRE_IMPORT_RESUME:
			return "通道控制器消防输入报警恢复"
		case MINOR_LEGAL_EVENT_NEARLY_FULL:
			return "合法事件满90%报警"
		case MINOR_LOCK_HIJIACK_ALARM:
			return "智能锁防劫持报警"
		default:
			return "unknown"
		}
	case 2:
		switch d {
		case MINOR_NET_BROKEN:
			return "网络断开"
		case MINOR_RS485_DEVICE_ABNORMAL:
			return "RS485连接状态异常"
		case MINOR_RS485_DEVICE_REVERT:
			return "RS485连接状态异常恢复"
		case MINOR_DEV_POWER_ON:
			return "设备上电启动"
		case MINOR_DEV_POWER_OFF:
			return "设备掉电关闭"
		case MINOR_WATCH_DOG_RESET:
			return "看门狗复位"
		case MINOR_LOW_BATTERY:
			return "蓄电池电压低"
		case MINOR_BATTERY_RESUME:
			return "蓄电池电压恢复正常"
		case MINOR_AC_OFF:
			return "交流电断电"
		case MINOR_AC_RESUME:
			return "交流电恢复"
		case MINOR_NET_RESUME:
			return "网络恢复"
		case MINOR_FLASH_ABNORMAL:
			return "FLASH读写异常"
		case MINOR_CARD_READER_OFFLINE:
			return "读卡器掉线"
		case MINOR_CARD_READER_RESUME:
			return "读卡器掉线恢复"
		case MINOR_INDICATOR_LIGHT_OFF:
			return "指示灯关闭"
		case MINOR_INDICATOR_LIGHT_RESUME:
			return "指示灯恢复"
		case MINOR_CHANNEL_CONTROLLER_OFF:
			return "通道控制器掉线"
		case MINOR_CHANNEL_CONTROLLER_RESUME:
			return "通道控制器恢复"
		case MINOR_SECURITY_MODULE_OFF:
			return "门控安全模块掉线"
		case MINOR_SECURITY_MODULE_RESUME:
			return "门控安全模块掉线恢复"
		case MINOR_LOCAL_CONTROL_NET_BROKEN:
			return "就地控制器网络断开"
		case MINOR_LOCAL_CONTROL_NET_RSUME:
			return "就地控制器网络恢复"
		case MINOR_MASTER_RS485_LOOPNODE_BROKEN:
			return "主控RS485环路节点断开"
		case MINOR_MASTER_RS485_LOOPNODE_RESUME:
			return "主控RS485环路节点恢复"
		case MINOR_LOCAL_CONTROL_OFFLINE:
			return "就地控制器掉线"
		case MINOR_LOCAL_CONTROL_RESUME:
			return "就地控制器掉线恢复"
		case MINOR_LOCAL_DOWNSIDE_RS485_LOOPNODE_BROKEN:
			return "就地下行RS485环路断开"
		case MINOR_LOCAL_DOWNSIDE_RS485_LOOPNODE_RESUME:
			return "就地下行RS485环路恢复"
		case MINOR_DISTRACT_CONTROLLER_ONLINE:
			return "分控器在线"
		case MINOR_DISTRACT_CONTROLLER_OFFLINE:
			return "分控器离线"
		case MINOR_ID_CARD_READER_NOT_CONNECT:
			return "身份证阅读器未连接（智能专用）"
		case MINOR_ID_CARD_READER_RESUME:
			return "身份证阅读器连接恢复（智能专用）"
		case MINOR_FINGER_PRINT_MODULE_NOT_CONNECT:
			return "指纹模组未连接（智能专用）"
		case MINOR_FINGER_PRINT_MODULE_RESUME:
			return "指纹模组连接恢复（智能专用）"
		case MINOR_CAMERA_NOT_CONNECT:
			return "摄像头未连接"
		case MINOR_CAMERA_RESUME:
			return "摄像头连接恢复"
		case MINOR_COM_NOT_CONNECT:
			return "COM口未连接"
		case MINOR_COM_RESUME:
			return "COM口连接恢复"
		case MINOR_DEVICE_NOT_AUTHORIZE:
			return "设备未授权"
		case MINOR_PEOPLE_AND_ID_CARD_DEVICE_ONLINE:
			return "人证设备在线"
		case MINOR_PEOPLE_AND_ID_CARD_DEVICE_OFFLINE:
			return "人证设备离线"
		case MINOR_LOCAL_LOGIN_LOCK:
			return "本地登录锁定"
		case MINOR_LOCAL_LOGIN_UNLOCK:
			return "本地登录解锁"
		case MINOR_BATTERY_ELECTRIC_LOW:
			return "电池电压低（仅人脸设备使用）"
		case MINOR_BATTERY_ELECTRIC_RESUME:
			return "电池电压恢复正常（仅人脸设备使用）"
		case MINOR_SUBMARINEBACK_COMM_BREAK:
			return "与反潜回服务器通信断开"
		case MINOR_SUBMARINEBACK_COMM_RESUME:
			return "与反潜回服务器通信恢复"
		case MINOR_MOTOR_SENSOR_EXCEPTION:
			return "电机或传感器异常"
		case MINOR_CAN_BUS_EXCEPTION:
			return "CAN总线异常"
		case MINOR_CAN_BUS_RESUME:
			return "CAN总线恢复"
		case MINOR_GATE_TEMPERATURE_OVERRUN:
			return "闸机腔体温度超限"
		case MINOR_IR_EMITTER_EXCEPTION:
			return "红外对射异常"
		case MINOR_IR_EMITTER_RESUME:
			return "红外对射恢复"
		case MINOR_LAMP_BOARD_COMM_EXCEPTION:
			return "灯板通信异常"
		case MINOR_LAMP_BOARD_COMM_RESUME:
			return "灯板通信恢复"
		case MINOR_IR_ADAPTOR_COMM_EXCEPTION:
			return "红外转接板通信异常"
		case MINOR_IR_ADAPTOR_COMM_RESUME:
			return "红外转接板通信恢复"

		default:
			return "unknown"
		}
	case 3:
		switch d {
		case MINOR_LOCAL_LOGIN:
			return "本地登陆"
		case MINOR_LOCAL_LOGOUT:
			return "本地注销登陆"
		case MINOR_LOCAL_UPGRADE:
			return "本地升级"
		case MINOR_REMOTE_LOGIN:
			return "远程登录"
		case MINOR_REMOTE_LOGOUT:
			return "远程注销登陆"
		case MINOR_REMOTE_ARM:
			return "远程布防"
		case MINOR_REMOTE_DISARM:
			return "远程撤防"
		case MINOR_REMOTE_REBOOT:
			return "远程重启"
		case MINOR_REMOTE_UPGRADE:
			return "远程升级"
		case MINOR_REMOTE_CFGFILE_OUTPUT:
			return "远程导出配置文件"
		case MINOR_REMOTE_CFGFILE_INTPUT:
			return "远程导入配置文件"
		case MINOR_REMOTE_ALARMOUT_OPEN_MAN:
			return "远程手动开启报警输出"
		case MINOR_REMOTE_ALARMOUT_CLOSE_MAN:
			return "远程手动关闭报警输出"
		case MINOR_REMOTE_OPEN_DOOR:
			return "远程开门"
		case MINOR_REMOTE_CLOSE_DOOR:
			return "远程关门（对于梯控，表示受控）"
		case MINOR_REMOTE_ALWAYS_OPEN:
			return "远程常开（对于梯控，表示自由）"
		case MINOR_REMOTE_ALWAYS_CLOSE:
			return "远程常关（对于梯控，表示禁用）"
		case MINOR_REMOTE_CHECK_TIME:
			return "远程手动校时"
		case MINOR_NTP_CHECK_TIME:
			return "NTP自动校时"
		case MINOR_REMOTE_CLEAR_CARD:
			return "远程清空卡号"
		case MINOR_REMOTE_RESTORE_CFG:
			return "远程恢复默认参数"
		case MINOR_ALARMIN_ARM:
			return "防区布防"
		case MINOR_ALARMIN_DISARM:
			return "防区撤防"
		case MINOR_LOCAL_RESTORE_CFG:
			return "本地恢复默认参数"
		case MINOR_REMOTE_CAPTURE_PIC:
			return "远程抓拍"
		case MINOR_MOD_NET_REPORT_CFG:
			return "修改网络中心参数配置"
		case MINOR_MOD_GPRS_REPORT_PARAM:
			return "修改GPRS中心参数配置"
		case MINOR_MOD_REPORT_GROUP_PARAM:
			return "修改中心组参数配置"
		case MINOR_UNLOCK_PASSWORD_OPEN_DOOR:
			return "解除码输入"
		case MINOR_AUTO_RENUMBER:
			return "自动重新编号"
		case MINOR_AUTO_COMPLEMENT_NUMBER:
			return "自动补充编号"
		case MINOR_NORMAL_CFGFILE_INPUT:
			return "导入普通配置文件"
		case MINOR_NORMAL_CFGFILE_OUTTPUT:
			return "导出普通配置文件"
		case MINOR_CARD_RIGHT_INPUT:
			return "导入卡权限参数"
		case MINOR_CARD_RIGHT_OUTTPUT:
			return "导出卡权限参数"
		case MINOR_LOCAL_USB_UPGRADE:
			return "本地U盘升级"
		case MINOR_REMOTE_VISITOR_CALL_LADDER:
			return "访客呼梯"
		case MINOR_REMOTE_HOUSEHOLD_CALL_LADDER:
			return "住户呼梯"
		case MINOR_REMOTE_ACTUAL_GUARD:
			return "远程实时布防"
		case MINOR_REMOTE_ACTUAL_UNGUARD:
			return "远程实时撤防"
		case MINOR_REMOTE_CONTROL_NOT_CODE_OPER_FAILED:
			return "遥控器未对码操作失败"
		case MINOR_REMOTE_CONTROL_CLOSE_DOOR:
			return "遥控器关门"
		case MINOR_REMOTE_CONTROL_OPEN_DOOR:
			return "遥控器开门"
		case MINOR_REMOTE_CONTROL_ALWAYS_OPEN_DOOR:
			return "遥控器常开门"

		default:
			return "unknown"
		}
	case 5:
		switch d {
		case MINOR_LEGAL_CARD_PASS:
			return "合法卡认证通过"
		case MINOR_CARD_AND_PSW_PASS:
			return "刷卡加密码认证通过"
		case MINOR_CARD_AND_PSW_FAIL:
			return "刷卡加密码认证失败"
		case MINOR_CARD_AND_PSW_TIMEOUT:
			return "数卡加密码认证超时"
		case MINOR_CARD_AND_PSW_OVER_TIME:
			return "刷卡加密码超次"
		case MINOR_CARD_NO_RIGHT:
			return "未分配权限"
		case MINOR_CARD_INVALID_PERIOD:
			return "无效时段"
		case MINOR_CARD_OUT_OF_DATE:
			return "卡号过期"
		case MINOR_INVALID_CARD:
			return "无此卡号"
		case MINOR_ANTI_SNEAK_FAIL:
			return "反潜回认证失败"
		case MINOR_INTERLOCK_DOOR_NOT_CLOSE:
			return "互锁门未关闭"
		case MINOR_NOT_BELONG_MULTI_GROUP:
			return "卡不属于多重认证群组"
		case MINOR_INVALID_MULTI_VERIFY_PERIOD:
			return "卡不在多重认证时间段内"
		case MINOR_MULTI_VERIFY_SUPER_RIGHT_FAIL:
			return "多重认证模式超级权限认证失败"
		case MINOR_MULTI_VERIFY_REMOTE_RIGHT_FAIL:
			return "多重认证模式远程认证失败"
		case MINOR_MULTI_VERIFY_SUCCESS:
			return "多重认证成功"
		case MINOR_LEADER_CARD_OPEN_BEGIN:
			return "首卡开门开始"
		case MINOR_LEADER_CARD_OPEN_END:
			return "首卡开门结束"
		case MINOR_ALWAYS_OPEN_BEGIN:
			return "常开状态开始"
		case MINOR_ALWAYS_OPEN_END:
			return "常开状态结束"
		case MINOR_LOCK_OPEN:
			return "门锁打开"
		case MINOR_LOCK_CLOSE:
			return "门锁关闭"
		case MINOR_DOOR_BUTTON_PRESS:
			return "开门按钮打开"
		case MINOR_DOOR_BUTTON_RELEASE:
			return "开门按钮放开"
		case MINOR_DOOR_OPEN_NORMAL:
			return "正常开门（门磁）"
		case MINOR_DOOR_CLOSE_NORMAL:
			return "正常关门（门磁）"
		case MINOR_DOOR_OPEN_ABNORMAL:
			return "门异常打开（门磁）"
		case MINOR_DOOR_OPEN_TIMEOUT:
			return "门打开超时（门磁）"
		case MINOR_ALARMOUT_ON:
			return "报警输出打开"
		case MINOR_ALARMOUT_OFF:
			return "报警输出关闭"
		case MINOR_ALWAYS_CLOSE_BEGIN:
			return "常关状态开始"
		case MINOR_ALWAYS_CLOSE_END:
			return "常关状态结束"
		case MINOR_MULTI_VERIFY_NEED_REMOTE_OPEN:
			return "多重多重认证需要远程开门"
		case MINOR_MULTI_VERIFY_SUPERPASSWD_VERIFY_SUCCESS:
			return "多重认证超级密码认证成功事件"
		case MINOR_MULTI_VERIFY_REPEAT_VERIFY:
			return "多重认证重复认证事件"
		case MINOR_MULTI_VERIFY_TIMEOUT:
			return "多重认证重复认证事件"
		case MINOR_DOORBELL_RINGING:
			return "门铃响"
		case MINOR_FINGERPRINT_COMPARE_PASS:
			return "指纹比对通过"
		case MINOR_FINGERPRINT_COMPARE_FAIL:
			return "指纹比对失败"
		case MINOR_CARD_FINGERPRINT_VERIFY_PASS:
			return "刷卡加指纹认证通过"
		case MINOR_CARD_FINGERPRINT_VERIFY_FAIL:
			return "刷卡加指纹认证失败"
		case MINOR_CARD_FINGERPRINT_VERIFY_TIMEOUT:
			return "刷卡加指纹认证超时"
		case MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_PASS:
			return "刷卡加指纹加密码认证通过"
		case MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_FAIL:
			return "刷卡加指纹加密码认证失败"
		case MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_TIMEOUT:
			return "刷卡加指纹加密码认证超时"
		case MINOR_FINGERPRINT_PASSWD_VERIFY_PASS:
			return "指纹加密码认证通过"
		case MINOR_FINGERPRINT_PASSWD_VERIFY_FAIL:
			return "指纹加密码认证失败"
		case MINOR_FINGERPRINT_PASSWD_VERIFY_TIMEOUT:
			return "指纹加密码认证超时"
		case MINOR_FINGERPRINT_INEXISTENCE:
			return "指纹不存在"
		case MINOR_CARD_PLATFORM_VERIFY:
			return "刷卡平台认证"
		case MINOR_CALL_CENTER:
			return "呼叫中心事件"
		case MINOR_FIRE_RELAY_TURN_ON_DOOR_ALWAYS_OPEN:
			return "消防继电器导通触发门常开"
		case MINOR_FIRE_RELAY_RECOVER_DOOR_RECOVER_NORMAL:
			return "消防继电器恢复门恢复正常"
		case MINOR_FACE_AND_FP_VERIFY_PASS:
			return "人脸加指纹认证通过"
		case MINOR_FACE_AND_FP_VERIFY_FAIL:
			return "人脸加指纹认证失败"
		case MINOR_FACE_AND_FP_VERIFY_TIMEOUT:
			return "人脸加指纹认证超时"
		case MINOR_FACE_AND_PW_VERIFY_PASS:
			return "人脸加密码认证通过"
		case MINOR_FACE_AND_PW_VERIFY_FAIL:
			return "人脸加密码认证失败"
		case MINOR_FACE_AND_PW_VERIFY_TIMEOUT:
			return "人脸加密码认证超时"
		case MINOR_FACE_AND_CARD_VERIFY_PASS:
			return "人脸加刷卡认证通过"
		case MINOR_FACE_AND_CARD_VERIFY_FAIL:
			return "人脸加刷卡认证失败"
		case MINOR_FACE_AND_CARD_VERIFY_TIMEOUT:
			return "人脸加刷卡认证超时"
		case MINOR_FACE_AND_PW_AND_FP_VERIFY_PASS:
			return "人脸加密码加指纹认证通过"
		case MINOR_FACE_AND_PW_AND_FP_VERIFY_FAIL:
			return "人脸加密码加指纹认证失败"
		case MINOR_FACE_AND_PW_AND_FP_VERIFY_TIMEOUT:
			return "人脸加密码加指纹认证超时"
		case MINOR_FACE_CARD_AND_FP_VERIFY_PASS:
			return "人脸加刷卡加指纹认证通过"
		case MINOR_FACE_CARD_AND_FP_VERIFY_FAIL:
			return "人脸加刷卡加指纹认证失败"
		case MINOR_FACE_CARD_AND_FP_VERIFY_TIMEOUT:
			return "人脸加刷卡加指纹认证超时"
		case MINOR_EMPLOYEENO_AND_FP_VERIFY_PASS:
			return "工号加指纹认证通过"
		case MINOR_EMPLOYEENO_AND_FP_VERIFY_FAIL:
			return "工号加指纹认证失败"
		case MINOR_EMPLOYEENO_AND_FP_VERIFY_TIMEOUT:
			return "工号加指纹认证超时"
		case MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_PASS:
			return "工号加指纹加密码认证通过"
		case MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_FAIL:
			return "工号加指纹加密码认证失败"
		case MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_TIMEOUT:
			return "工号加指纹加密码认证超时"
		case MINOR_FACE_VERIFY_PASS:
			return "人脸认证通过"
		case MINOR_FACE_VERIFY_FAIL:
			return "人脸认证失败"
		case MINOR_EMPLOYEENO_AND_FACE_VERIFY_PASS:
			return "工号加人脸认证通过"
		case MINOR_EMPLOYEENO_AND_FACE_VERIFY_FAIL:
			return "工号加人脸认证失败"
		case MINOR_EMPLOYEENO_AND_FACE_VERIFY_TIMEOUT:
			return "工号加人脸认证超时"
		case MINOR_FACE_RECOGNIZE_FAIL:
			return "人脸识别失败"
		case MINOR_FIRSTCARD_AUTHORIZE_BEGIN:
			return "首卡授权开始"
		case MINOR_FIRSTCARD_AUTHORIZE_END:
			return "首卡授权结束"
		case MINOR_DOORLOCK_INPUT_SHORT_CIRCUIT:
			return "门锁输入短路报警"
		case MINOR_DOORLOCK_INPUT_BROKEN_CIRCUIT:
			return "门锁输入断路报警"
		case MINOR_DOORLOCK_INPUT_EXCEPTION:
			return "门锁输入异常报警"
		case MINOR_DOORCONTACT_INPUT_SHORT_CIRCUIT:
			return "门磁输入短路报警"
		case MINOR_DOORCONTACT_INPUT_BROKEN_CIRCUIT:
			return "门磁输入断路报警"
		case MINOR_DOORCONTACT_INPUT_EXCEPTION:
			return "门磁输入异常报警"
		case MINOR_OPENBUTTON_INPUT_SHORT_CIRCUIT:
			return "开门按钮输入短路报警"
		case MINOR_OPENBUTTON_INPUT_BROKEN_CIRCUIT:
			return "开门按钮输入断路报警"
		case MINOR_OPENBUTTON_INPUT_EXCEPTION:
			return "开门按钮输入异常报警"
		case MINOR_DOORLOCK_OPEN_EXCEPTION:
			return "门锁异常打开"
		case MINOR_DOORLOCK_OPEN_TIMEOUT:
			return "门锁打开超时"
		case MINOR_FIRSTCARD_OPEN_WITHOUT_AUTHORIZE:
			return "首卡未授权开门失败"
		case MINOR_CALL_LADDER_RELAY_BREAK:
			return "呼梯继电器断开"
		case MINOR_CALL_LADDER_RELAY_CLOSE:
			return "呼梯继电器闭合"
		case MINOR_AUTO_KEY_RELAY_BREAK:
			return "自动按键继电器断开"
		case MINOR_AUTO_KEY_RELAY_CLOSE:
			return "自动按键继电器闭合"
		case MINOR_KEY_CONTROL_RELAY_BREAK:
			return "按键梯控继电器断开"
		case MINOR_KEY_CONTROL_RELAY_CLOSE:
			return "按键梯控继电器闭合"
		case MINOR_EMPLOYEENO_AND_PW_PASS:
			return "工号加密码认证通过"
		case MINOR_EMPLOYEENO_AND_PW_FAIL:
			return "工号加密码认证失败"
		case MINOR_EMPLOYEENO_AND_PW_TIMEOUT:
			return "工号加密码认证超时"
		case MINOR_HUMAN_DETECT_FAIL:
			return "真人检测失败"
		case MINOR_PEOPLE_AND_ID_CARD_COMPARE_PASS:
			return "人证比对通过"
		case MINOR_PEOPLE_AND_ID_CARD_COMPARE_FAIL:
			return "人证比对失败"
		case MINOR_CERTIFICATE_BLOCKLIST:
			return "黑名单事件"
		case MINOR_LEGAL_MESSAGE:
			return "合法短信"
		case MINOR_ILLEGAL_MESSAGE:
			return "非法短信"
		case MINOR_MAC_DETECT:
			return "MAC侦测"
		case MINOR_DOOR_OPEN_OR_DORMANT_FAIL:
			return "门状态常闭或休眠状态认证失败"
		case MINOR_AUTH_PLAN_DORMANT_FAIL:
			return "认证计划休眠模式认证失败"
		case MINOR_CARD_ENCRYPT_VERIFY_FAIL:
			return "卡加密校验失败"
		case MINOR_SUBMARINEBACK_REPLY_FAIL:
			return "反潜回服务器应答失败"
		case MINOR_TRAILING:
			return "尾随通行"
		case MINOR_REVERSE_ACCESS:
			return "反向闯入"
		case MINOR_FORCE_ACCESS:
			return "外力冲撞"
		case MINOR_CLIMBING_OVER_GATE:
			return "翻越"
		case MINOR_PASSING_TIMEOUT:
			return "通行超时"
		case MINOR_INTRUSION_ALARM:
			return "误闯报警"
		case MINOR_FREE_GATE_PASS_NOT_AUTH:
			return "闸机自由通行时未认证通过"
		case MINOR_DROP_ARM_BLOCK:
			return "摆臂被阻挡"
		case MINOR_DROP_ARM_BLOCK_RESUME:
			return "摆臂阻挡消除"
		case MINOR_LOCAL_FACE_MODELING_FAIL:
			return "设备升级本地人脸建模失败"
		case MINOR_STAY_EVENT:
			return "逗留事件"
		case MINOR_PASSWORD_MISMATCH:
			return "密码不匹配"
		case MINOR_EMPLOYEE_NO_NOT_EXIST:
			return "工号不存在"
		case MINOR_COMBINED_VERIFY_PASS:
			return "组合认证通过"
		case MINOR_COMBINED_VERIFY_TIMEOUT:
			return "组合认证超时"
		case MINOR_VERIFY_MODE_MISMATCH:
			return "认证方式不匹配"

		default:
			return "unknown"
		}
	default:
		return "unknown"
	}
}
