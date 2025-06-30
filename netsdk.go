// Package hikvision_CGO 提供用于在linux环境 golang直接调用海康的sdk。
// 详细用法参见示例代码。
package hikvision_CGO

/*
#cgo CFLAGS:  -I./include
//运行时链接
#cgo LDFLAGS: -L./lib/Linux -lhcnetsdk
//编译时链接
#cgo LDFLAGS: -Wl,-rpath=./lib/Linux:./lib/Linux/HCNetSDKCom

#include "HCNetSDK.h"
#include <unistd.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
static int ST_Defence(int uid);
extern int MessageCallback(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *pAlarmInfo, DWORD dwBufLen, void* pUser); //回调函数声明 golang实现

//ST_Defence golang要转换数据很麻烦,直接用C写了
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

//--------------------------------------注册回调函数------------------------------------

// MessageCallback  处理从C库接收的报警信息，返回状态码。
// 参数说明：
//   - lCommand: 报警命令类型
//   - pAlarmInfo: 报警信息字符串指针
//export MessageCallback
func MessageCallback(lCommand C.int, pAlarmer *C.struct_tagNET_DVR_ALARMER, pAlarmInfo *C.char, dwBufLen C.DWORD, pUser unsafe.Pointer) C.int {
	defer func() {
		if e := recover(); e != nil {
			log.Println("MessageCallback panic : ", e)
		}
	}()
	if pAlarmer == nil || pUser == nil {
		log.Println("pAlarmer or pUser is nil")
		return 0
	}

	id := *(*ObjectId)(pUser)

	obj := id.Get()
	Alarm := *(*NET_DVR_ALARMER)(unsafe.Pointer(pAlarmer))
	ip := Alarm.ST_sDeviceIP

	ipStr := C.GoString((*C.char)(unsafe.Pointer(&ip[0])))

	ret := obj.Invoke(int(lCommand), ipStr, unsafe.Pointer(pAlarmInfo), int(dwBufLen))
	if ret {
		return 1
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

//--------------------------------------功能性调用------------------------------------

// NetInit 初始化设备和日志
// sdkLog 日志文件路径 , ifSdkLog 是否开启日志
// 示例： netsdk.NetInit("./",true)  日志文件名为sdk内部定义
func NetInit(sdkLog string, ifSdkLog bool) error {

	//初始化资源
	ret := C.NET_DVR_Init()
	if int(ret) != 1 {
		fmt.Printf("NET_DVR_Init failed,error code = %v\n", C.NET_DVR_GetLastError())
		return errors.New(fmt.Sprintf("NET_DVR_Init failed,error code = %v\n", C.NET_DVR_GetLastError()))
	}
	if ifSdkLog {
		cStr := C.CString(sdkLog)
		defer C.free(unsafe.Pointer(cStr)) // 确保释放
		// 日志等级， 日志目录 ，自动删除
		C.NET_DVR_SetLogToFile(3, cStr, 1)
	}
	return nil
}

// NetLoginV40 登录
// deviceIp 设备IP地址 username  用户名 password 密码
// 示例： NetLoginV40("192.168.1.64", "admin", "12345")
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
// operation  操作名称
// 示例：isErr("Login")
func isErr(operation string) error {
	errno := int64(C.NET_DVR_GetLastError())
	if errno > 0 {
		reMsg := fmt.Sprintf("%s失败,失败代码号：%d", operation, errno)
		return errors.New(reMsg)
	}
	return nil
}

// NetLogout 退出登录
// uid  用户ID 由NetLoginV40函数返回
// 示例：NetLogout(uid)
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
// uid 用户ID 由NetLoginV40函数返回 pBuf 门禁主机工作状态结构体（NET_DVR_ACS_WORK_STATUS ）用于接收调用返回的信息
//示例：var status netsdk.NET_DVR_ACS_WORK_STATUS
//      GetDoorStatus(uid, &status)
func GetDoorStatus(uid int32, pBuf *NET_DVR_ACS_WORK_STATUS) (err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("GetDoorStatus panic : ", e)
		}
	}()
	var lpBytesReturned C.DWORD
	var PBuf = (C.LPVOID)(unsafe.Pointer(pBuf))
	ret := C.NET_DVR_GetDVRConfig(C.LONG(uid), C.NET_DVR_GET_ACS_WORK_STATUS, C.LONG(0xFFFFFFF), PBuf, (C.DWORD)(unsafe.Sizeof(NET_DVR_ACS_WORK_STATUS{})), (C.LPDWORD)(&lpBytesReturned))

	if int32(ret) == 0 {
		if err = isErr("Get door status"); err != nil {
			return errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
		}
		return errors.New(fmt.Sprintf("get door status error，uid :%d", uid))
	}

	return nil
}

// DoDefence 布防
// uid 用户ID 由NetLoginV40函数返回
func DoDefence(uid int32) (int32, error) {
	ret := C.ST_Defence(C.int(uid))
	if int32(ret) == -1 {
		err := isErr("布防")
		return int32(ret), err
	}
	log.Println("布防成功")
	return int32(ret), nil
}

// CloseDefence 撤防
// dId 布防ID 由DoDefence函数返回
func CloseDefence(dId int32) {
	C.NET_DVR_CloseAlarmChan_V30(C.int(dId))
}

// ControlDoor 控制门状态
// uid 用户ID 由NetLoginV40函数返回
// ctrl 控制命令，0- 关闭，1- 打开），2- 常开），3- 常关，
// DoorIndex 门编号，从0开始  对应的是 NET_DVR_ACS_WORK_STATUS.ST_byMagneticStatus 数组中的第几位门
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
// uid 用户ID 由NetLoginV40函数返回
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
// uid 用户ID 由NetLoginV40函数返回
// status 报警主机状态信息结构体 用于接收函数调用返回的数据
func GetAlarmHostMainStatus(uid int32, status *NET_DVR_ALARMHOST_MAIN_STATUS_V51) (err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("GetAlarmHostMainStatus panic : ", e)
		}
	}()
	var AlarmHostStatus = (C.LPVOID)(unsafe.Pointer(status))
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
