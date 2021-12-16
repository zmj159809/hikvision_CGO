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
	int ret =-1;
	NET_DVR_SETUPALARM_PARAM struSetupParam={0};
	struSetupParam.dwSize=sizeof(NET_DVR_SETUPALARM_PARAM);
    struSetupParam.byDeployType = 1;//设置布防类型为实时布防

	LONG mHandle = NET_DVR_SetupAlarmChan_V41(uid1,&struSetupParam);
	if (mHandle<0){
	  return ret;
	}
	ret = 0;

    return ret;
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
func GetDoorStatus(uid int32) (NET_DVR_ACS_WORK_STATUS, error) {
	var DoorStatus C.LPNET_DVR_ACS_WORK_STATUS
	var lpBytesReturned C.DWORD
	ret := C.NET_DVR_GetDVRConfig(C.LONG(uid), C.NET_DVR_GET_ACS_WORK_STATUS, C.LONG(0xFFFFFFF), (C.LPVOID)(unsafe.Pointer(&DoorStatus)), (C.DWORD)(unsafe.Sizeof(NET_DVR_ACS_WORK_STATUS{})), (C.LPDWORD)(&lpBytesReturned))

	var Dstatus NET_DVR_ACS_WORK_STATUS
	Dstatus = *(*NET_DVR_ACS_WORK_STATUS)(unsafe.Pointer(&DoorStatus))

	if int32(ret) == 0 {
		if err := isErr("Get door status"); err != nil {
			return Dstatus, errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
		}
		return Dstatus, errors.New(fmt.Sprintf("get door status error，uid :%d", uid))
	}

	return Dstatus, nil
}

//export MessageCallback
func MessageCallback(lCommand C.int, pAlarmer *C.struct_tagNET_DVR_ALARMER, pAlarmInfo *C.char, dwBufLen C.DWORD, pUser unsafe.Pointer) C.int {

	id := *(*ObjectId)(pUser)
	interf := id.Get()
	obj, ok := interf.(IF_fMEssCallBack)
	Alarmer := *(*NET_DVR_ALARMER)(unsafe.Pointer(pAlarmer))
	ip := Alarmer.ST_sDeviceIP
	if ok {
		ret := obj.Invoke(int(lCommand), string(ip[:]), unsafe.Pointer(pAlarmInfo), int(dwBufLen))
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
	fmt.Println("注册回调函数成功")
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

	ret := C.NET_DVR_ControlGateway(C.LONG(uid), C.LONG(DoorIndex), C.DWORD(ctrl))
	if int32(ret) == 0 {
		if err := isErr("Control door"); err != nil {
			return errors.New(fmt.Sprintf("uid:[%d] 失败,原因%v", uid, err.Error()))
		}
		return errors.New(fmt.Sprintf("Control door error，uid :%d", uid))
	}
	return nil
}
