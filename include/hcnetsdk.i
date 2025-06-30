%module HKIPcamera
%{
// 必须包含的头文件
#include "HCNetSDK.h"
#include "plaympeg4.h"  // 若涉及解码回调需包含
#include <opencv2/opencv.hpp>  // 若涉及图像处理需包含

// 其他必要的宏定义和全局变量声明
#define USECOLOR 1
extern LONG lUserID;
extern NET_DVR_DEVICEINFO_V30 struDeviceInfo;
%}

// // 启用Numpy支持（若需要传递图像数据）
// %include "numpy.i"
// %init %{
// import_array();
// %}

// 类型映射：处理C++与Python之间的数据类型转换
%include "typemaps.i"
%apply char * { BYTE * };  // 处理字节流参数
%apply int *OUTPUT { int *pnPort };  // 处理输出型参数

// 封装海康SDK的核心函数和结构体
%include "HCNetSDK.h"

// 封装登录相关函数
%rename(login_v30) NET_DVR_Login_V30;
extern LONG NET_DVR_Login_V30(
    const char *sDVRIP, 
    WORD wDVRPort, 
    const char *sUserName, 
    const char *sPassword, 
    LPNET_DVR_DEVICEINFO_V30 lpDeviceInfo
);

// 封装实时预览回调函数（需自定义回调处理）
%typemap(in) (LONG lRealHandle, void* pUser) {
    $1 = ($1_type)PyLong_AsLong($input);
    $2 = NULL;  // 可根据需要传递Python对象
}

// 封装图像回调处理（示例）
%inline %{
void DecodeFrameCallback(FRAME_INFO *pFrameInfo, BYTE *pBuf) {
    // 此处添加解码和图像处理的C++代码
    // 参考搜索结果中的yv12toYUV和DecCBFun实现[2,5](@ref)
}
%}

// 异常回调封装
%callback("%s_cb");
extern void NET_DVR_SetExceptionCallBack_V30(
    DWORD dwType, 
    LONG lUserID, 
    void (*fExceptionCallBack)(DWORD dwType, LONG lUserID, LONG lHandle, void *pUser), 
    void *pUser
);
%nocallback;

// 其他常用API封装（按需添加）
%rename(ptz_control) NET_DVR_PTZControl_Other;
extern BOOL NET_DVR_PTZControl_Other(
    LONG lUserID, 
    LONG lChannel, 
    DWORD dwPTZCommand, 
    DWORD dwStop, 
    DWORD dwSpeed
);