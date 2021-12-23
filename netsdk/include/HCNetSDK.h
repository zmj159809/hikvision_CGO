#ifndef _HC_NET_SDK_H_
#define _HC_NET_SDK_H_

#ifndef _WINDOWS_
    #if (defined(_WIN32) || defined(_WIN64))
        #include <winsock2.h>
        #include <windows.h>
    #endif
#endif

#ifndef __PLAYRECT_defined
    #define __PLAYRECT_defined
    typedef struct __PLAYRECT
    {
        int x;
        int y;
        int uWidth;
        int uHeight;
    }PLAYRECT;
#endif

#if (defined(_WIN32)) //windows
    typedef  unsigned __int64   UINT64;
    typedef  signed   __int64   INT64;
#elif defined(__linux__) || defined(__APPLE__) //linux
    #define  BOOL  int
      #include <stdint.h>
      typedef uint32_t    DWORD;
      typedef uint16_t    WORD;
      typedef uint16_t    SHORT;
      typedef uint16_t    USHORT;
      typedef int32_t     LONG;
      typedef uint8_t     BYTE;
      typedef uint32_t    UINT;
      typedef void*       LPVOID;
      typedef void*       HANDLE;
      typedef uint32_t *  LPDWORD;
      typedef uint64_t    UINT64;

    #ifndef TRUE
        #define TRUE  1
    #endif
    #ifndef FALSE
        #define FALSE 0
    #endif
    #ifndef NULL
        #define NULL 0
    #endif

    #define __stdcall
    #define CALLBACK

    #define NET_DVR_API extern "C"
    typedef unsigned int   COLORKEY;
    typedef unsigned int   COLORREF;

    #ifndef __HWND_defined
        #define __HWND_defined
        #if defined(__linux__)
            typedef unsigned int HWND;
        #else
            typedef void* HWND;
        #endif
    #endif

    #ifndef __HDC_defined
        #define __HDC_defined
        #if defined(__linux__)
            typedef struct __DC
            {
                void*   surface;        //SDL Surface
                HWND    hWnd;           //HDC window handle
            }DC;
            typedef DC* HDC;
        #else
            typedef void* HDC;
        #endif
    #endif

    typedef struct tagInitInfo
    {
        int uWidth;
        int uHeight;
    }INITINFO;
#endif

#define NAME_LEN                32      //用户名长度
#define SERIALNO_LEN            48      //序列号长度
#define MACADDR_LEN                6       //mac地址长度
#define NET_DVR_DEV_ADDRESS_MAX_LEN 129
#define NET_DVR_LOGIN_USERNAME_MAX_LEN 64
#define NET_DVR_LOGIN_PASSWD_MAX_LEN 64
#define MAX_ALARMHOST_ALARMIN_NUM            512//网络报警主机最大报警输入口数
#define MAX_ALARMHOST_ALARMOUT_NUM            512//网络报警主机最大报警输出口数
#define MAX_ALARMHOST_SUBSYSTEM             32//报警主机最大子系统数

#define HOLIDAY_GROUP_NAME_LEN          32  //假日组名称长度
#define MAX_HOLIDAY_PLAN_NUM            16  //假日组最大假日计划数
#define TEMPLATE_NAME_LEN               32  //计划模板名称长度
#define MAX_HOLIDAY_GROUP_NUM           16   //计划模板最大假日组数
#define DOOR_NAME_LEN                   32  //门名称
#define STRESS_PASSWORD_LEN             8   //胁迫密码长度
#define SUPER_PASSWORD_LEN              8   //胁迫密码长度
#define GROUP_NAME_LEN                  32  //群组名称长度
#define GROUP_COMBINATION_NUM           8   //群组组合数
#define MULTI_CARD_GROUP_NUM            4   //单门最大多重卡组数
#define ACS_CARD_NO_LEN                 32  //门禁卡号长度
#define NET_SDK_EMPLOYEE_NO_LEN         32  //工号长度
#define NET_SDK_UUID_LEN                36  //UUID长度
#define NET_SDK_EHOME_KEY_LEN           32  //EHome Key长度
#define CARD_PASSWORD_LEN               8   //卡密码长度
#define MAX_DOOR_NUM                    32  //最大门数
#define MAX_CARD_RIGHT_PLAN_NUM         4   //卡权限最大计划个数
#define MAX_GROUP_NUM_128               128 //最大群组数
#define MAX_CARD_READER_NUM             64  //最大读卡器数
#define MAX_SNEAK_PATH_NODE             8   //最大后续读卡器数
#define MAX_MULTI_DOOR_INTERLOCK_GROUP  8   //最大多门互锁组数
#define MAX_INTER_LOCK_DOOR_NUM         8   //一个多门互锁组中最大互锁门数
#define MAX_CASE_SENSOR_NUM             8   //最大case sensor触发器数
#define MAX_DOOR_NUM_256                256 //最大门数
#define MAX_READER_ROUTE_NUM            16  //最大刷卡循序路径
#define MAX_FINGER_PRINT_NUM            10  //最大指纹个数
#define MAX_CARD_READER_NUM_512            512 //最大读卡器数
#define NET_SDK_MULTI_CARD_GROUP_NUM_20     20   //单门最大多重卡组数


#define NET_DVR_GET_ALARMHOST_MAIN_STATUS_V40    2072   // 获取主要状态V40
#define NET_DVR_GET_ALARMHOST_MAIN_STATUS_V51    2083   // 获取主要状态V51
// 获取设备参数
#define NET_DVR_GET_WEEK_PLAN_CFG               2100    //获取门状态周计划参数
#define NET_DVR_SET_WEEK_PLAN_CFG               2101    //设置门状态周计划参数
#define NET_DVR_GET_DOOR_STATUS_HOLIDAY_PLAN    2102    //获取门状态假日计划参数
#define NET_DVR_SET_DOOR_STATUS_HOLIDAY_PLAN    2103    //设置门状态假日计划参数
#define NET_DVR_GET_DOOR_STATUS_HOLIDAY_GROUP   2104    //获取门状态假日组参数
#define NET_DVR_SET_DOOR_STATUS_HOLIDAY_GROUP   2105    //设置门状态假日组参数
#define NET_DVR_GET_DOOR_STATUS_PLAN_TEMPLATE   2106    //获取门状态计划模板参数
#define NET_DVR_SET_DOOR_STATUS_PLAN_TEMPLATE   2107    //设置门状态计划模板参数
#define NET_DVR_GET_DOOR_CFG                    2108    //获取门参数
#define NET_DVR_SET_DOOR_CFG                    2109    //设置门参数
#define NET_DVR_GET_DOOR_STATUS_PLAN            2110    //获取门状态计划参数
#define NET_DVR_SET_DOOR_STATUS_PLAN            2111    //设置门状态计划参数
#define NET_DVR_GET_GROUP_CFG                   2112    //获取群组参数
#define NET_DVR_SET_GROUP_CFG                   2113    //设置群组参数
#define NET_DVR_GET_MULTI_CARD_CFG              2114    //获取多重卡参数
#define NET_DVR_SET_MULTI_CARD_CFG              2115    //设置多重卡参数
#define NET_DVR_GET_CARD_CFG                    2116    //获取卡参数
#define NET_DVR_SET_CARD_CFG                    2117    //设置卡参数
#define NET_DVR_CLEAR_ACS_PARAM                    2118    //清空门禁主机参数
#define NET_DVR_GET_SNEAK_CFG                    2119    //获取反潜回参数
#define NET_DVR_SET_SNEAK_CFG                   2120    //设置反潜回参数
#define NET_DVR_GET_MULTI_DOOR_INTERLOCK_CFG    2121    //获取多门互锁参数
#define NET_DVR_SET_MULTI_DOOR_INTERLOCK_CFG    2122    //设置多门互锁参数
#define NET_DVR_GET_ACS_WORK_STATUS                2123    //获取门禁主机工作状态
#define NET_DVR_GET_VERIFY_WEEK_PLAN            2124    //获取读卡器验证方式周计划参数
#define NET_DVR_SET_VERIFY_WEEK_PLAN            2125    //设置读卡器验证方式周计划参数
#define NET_DVR_GET_CARD_RIGHT_WEEK_PLAN        2126    //获取卡权限周计划参数
#define NET_DVR_SET_CARD_RIGHT_WEEK_PLAN        2127    //设置卡权限周计划参数
#define NET_DVR_GET_VERIFY_HOLIDAY_PLAN         2128    //获取读卡器验证方式假日计划参数
#define NET_DVR_SET_VERIFY_HOLIDAY_PLAN         2129    //设置读卡器验证方式假日计划参数
#define NET_DVR_GET_CARD_RIGHT_HOLIDAY_PLAN     2130    //获取卡权限假日计划参数
#define NET_DVR_SET_CARD_RIGHT_HOLIDAY_PLAN     2131    //设置卡权限假日计划参数
#define NET_DVR_GET_VERIFY_HOLIDAY_GROUP        2132    //获取读卡器验证方式假日组参数
#define NET_DVR_SET_VERIFY_HOLIDAY_GROUP        2133    //设置读卡器验证方式假日组参数
#define NET_DVR_GET_CARD_RIGHT_HOLIDAY_GROUP    2134    //获取卡权限假日组参数
#define NET_DVR_SET_CARD_RIGHT_HOLIDAY_GROUP    2135    //设置卡权限假日组参数
#define NET_DVR_GET_VERIFY_PLAN_TEMPLATE        2136    //获取读卡器验证方式计划模板参数
#define NET_DVR_SET_VERIFY_PLAN_TEMPLATE        2137    //设置读卡器验证方式计划模板参数
#define NET_DVR_GET_CARD_RIGHT_PLAN_TEMPLATE    2138    //获取卡权限计划模板参数
#define NET_DVR_SET_CARD_RIGHT_PLAN_TEMPLATE    2139    //设置卡权限计划模板参数
#define NET_DVR_GET_CARD_READER_CFG                2140    //获取读卡器参数
#define NET_DVR_SET_CARD_READER_CFG             2141    //设置读卡器参数
#define NET_DVR_GET_CARD_READER_PLAN            2142    //获取读卡器验证计划参数
#define NET_DVR_SET_CARD_READER_PLAN            2143    //设置读卡器验证计划参数
#define NET_DVR_GET_CASE_SENSOR_CFG                2144    //获取事件触发器参数
#define NET_DVR_SET_CASE_SENSOR_CFG             2145    //设置事件触发器参数
#define NET_DVR_GET_CARD_READER_ANTI_SNEAK_CFG  2146    //获取读卡器反潜回参数
#define NET_DVR_SET_CARD_READER_ANTI_SNEAK_CFG  2147    //设置读卡器反潜回参数
#define NET_DVR_GET_PHONE_DOOR_RIGHT_CFG        2148    //获取手机关联门权限参数
#define NET_DVR_SET_PHONE_DOOR_RIGHT_CFG        2149    //获取手机关联门权限参数
#define NET_DVR_GET_FINGERPRINT_CFG             2150    //获取指纹参数
#define NET_DVR_SET_FINGERPRINT_CFG             2151    //设置指纹参数
#define NET_DVR_DEL_FINGERPRINT_CFG             2152    //删除指纹参数
#define NET_DVR_GET_EVENT_CARD_LINKAGE_CFG      2153    //获取事件卡号联动配置参数
#define NET_DVR_SET_EVENT_CARD_LINKAGE_CFG      2154    //设置事件卡号联动配置参数
#define NET_DVR_GET_ANTI_SNEAK_HOST_CFG            2155    //获取主机组反潜回参数
#define NET_DVR_SET_ANTI_SNEAK_HOST_CFG         2156    //设置主机组反潜回参数
#define NET_DVR_GET_READER_ANTI_SNEAK_HOST_CFG  2157    //获取主机组读卡器反潜回参数
#define NET_DVR_SET_READER_ANTI_SNEAK_HOST_CFG  2158    //设置主机组读卡器反潜回参数
#define NET_DVR_GET_ACS_CFG                     2159    //获取门禁主机参数
#define NET_DVR_SET_ACS_CFG                     2160    //设置门禁主机参数
#define NET_DVR_GET_CARD_PASSWD_CFG                2161    //获取卡密码开门使能配置
#define NET_DVR_SET_CARD_PASSWD_CFG             2162    //设置卡密码开门使能配置
#define NET_DVR_GET_CARD_USERINFO_CFG           2163    //获取卡号关联用户信息参数
#define NET_DVR_SET_CARD_USERINFO_CFG           2164    //设置卡号关联用户信息参数

#define NET_DVR_GET_ACS_EXTERNAL_DEV_CFG        2165    //获取门禁主机串口外设参数
#define NET_DVR_SET_ACS_EXTERNAL_DEV_CFG        2166    //设置门禁主机串口外设参数
#define NET_DVR_GET_PERSONNEL_CHANNEL_CFG       2167    //获取人员通道参数
#define NET_DVR_SET_PERSONNEL_CHANNEL_CFG       2168    //设置人员通道参数
#define NET_DVR_SET_PLATFORM_VERIFY_CFG         2169    //下发平台认证结果
#define NET_DVR_GET_PERSON_STATISTICS_CFG        2170   //获取人数统计参数
#define NET_DVR_SET_PERSON_STATISTICS_CFG        2171   //设置人数统计参数
#define NET_DVR_GET_ACS_SCREEN_DISPLAY_CFG        2172   //获取屏幕字符串显示参数
#define NET_DVR_SET_ACS_SCREEN_DISPLAY_CFG        2173   //设置屏幕字符串显示参数
#define NET_DVR_GET_GATE_TIME_CFG               2174  //获取人员通道闸门时间参数
#define NET_DVR_SET_GATE_TIME_CFG               2175  //设置人员通道闸门时间参数
#define NET_DVR_GET_LOCAL_CONTROLLER_STATUS     2176    //获取就地控制器状态
#define NET_DVR_GET_ONLINE_LOCAL_CONTROLLER     2177    //搜索在线就地控制器
#define NET_DVR_GET_CARD_CFG_V50                2178    //获取新卡参数(V50)
#define NET_DVR_SET_CARD_CFG_V50                2179    //设置新卡参数(V50)
#define NET_DVR_GET_ACS_WORK_STATUS_V50         2180    //获取门禁主机工作状态(V50)
#define NET_DVR_GET_EVENT_CARD_LINKAGE_CFG_V50  2181    //获取事件卡号联动配置参数(V50)
#define NET_DVR_SET_EVENT_CARD_LINKAGE_CFG_V50  2182    //设置事件卡号联动配置参数(V50)
#define NET_DVR_GET_FINGERPRINT_CFG_V50         2183    //获取指纹参数V50
#define NET_DVR_SET_FINGERPRINT_CFG_V50         2184    //设置指纹参数V50

#define NET_DVR_GET_SAFETYCABIN_STATE            2197    //获取防护舱状态
#define NET_DVR_GET_RS485_CASCADE_CFG            2198   //获取Rs485级联设备配置
#define NET_DVR_SET_RS485_CASCADE_CFG            2199   //设置Rs485级联设备配置

#define NET_DVR_GET_CARD                 2560
#define NET_DVR_SET_CARD                 2561
#define NET_DVR_DEL_CARD                 2562
#define NET_DVR_GET_FINGERPRINT          2563
#define NET_DVR_SET_FINGERPRINT          2564
#define NET_DVR_DEL_FINGERPRINT          2565
#define NET_DVR_GET_FACE                 2566
#define NET_DVR_SET_FACE                 2567



typedef enum
{
    NET_SDK_GET_NEXT_STATUS_SUCCESS = 1000,    // �ɹ���ȡ�����ݣ��ͻ��˴����걾�����ݺ���Ҫ�ٴε���NET_DVR_RemoteConfigGetNext��ȡ��һ������
        NET_SDK_GET_NETX_STATUS_NEED_WAIT,        // ��ȴ��豸�������ݣ���������NET_DVR_RemoteConfigGetNext����
        NET_SDK_GET_NEXT_STATUS_FINISH,            // ����ȫ��ȡ�꣬��ʱ�ͻ��˿ɵ���NET_DVR_StopRemoteConfig����������
        NET_SDK_GET_NEXT_STATUS_FAILED,            // �����쳣���ͻ��˿ɵ���NET_DVR_StopRemoteConfig����������
}NET_SDK_GET_NEXT_STATUS;


#define LIGHT_PWRON        2    /* 接通灯光电源 */
#define WIPER_PWRON        3    /* 接通雨刷开关 */
#define FAN_PWRON        4    /* 接通风扇开关 */
#define HEATER_PWRON    5    /* 接通加热器开关 */
#define AUX_PWRON1        6    /* 接通辅助设备开关 */
#define AUX_PWRON2        7    /* 接通辅助设备开关 */
#define SET_PRESET        8    /* 设置预置点 */
#define CLE_PRESET        9    /* 清除预置点 */

#define ZOOM_IN            11    /* 焦距以速度SS变大(倍率变大) */
#define ZOOM_OUT        12    /* 焦距以速度SS变小(倍率变小) */
#define FOCUS_NEAR      13  /* 焦点以速度SS前调 */
#define FOCUS_FAR       14  /* 焦点以速度SS后调 */
#define IRIS_OPEN       15  /* 光圈以速度SS扩大 */
#define IRIS_CLOSE      16  /* 光圈以速度SS缩小 */

#define TILT_UP            21    /* 云台以SS的速度上仰 */
#define TILT_DOWN        22    /* 云台以SS的速度下俯 */
#define PAN_LEFT        23    /* 云台以SS的速度左转 */
#define PAN_RIGHT        24    /* 云台以SS的速度右转 */
#define UP_LEFT            25    /* 云台以SS的速度上仰和左转 */
#define UP_RIGHT        26    /* 云台以SS的速度上仰和右转 */
#define DOWN_LEFT        27    /* 云台以SS的速度下俯和左转 */
#define DOWN_RIGHT        28    /* 云台以SS的速度下俯和右转 */
#define PAN_AUTO        29    /* 云台以SS的速度左右自动扫描 */

#define FILL_PRE_SEQ    30    /* 将预置点加入巡航序列 */
#define SET_SEQ_DWELL    31    /* 设置巡航点停顿时间 */
#define SET_SEQ_SPEED    32    /* 设置巡航速度 */
#define CLE_PRE_SEQ        33    /* 将预置点从巡航序列中删除 */
#define STA_MEM_CRUISE    34    /* 开始记录轨迹 */
#define STO_MEM_CRUISE    35    /* 停止记录轨迹 */
#define RUN_CRUISE        36    /* 开始轨迹 */
#define RUN_SEQ            37    /* 开始巡航 */
#define STOP_SEQ        38    /* 停止巡航 */
#define GOTO_PRESET        39    /* 快球转到预置点 */

#define DEL_SEQ         43  /* 删除巡航路径 */
#define STOP_CRUISE        44    /* 停止轨迹 */
#define DELETE_CRUISE    45    /* 删除单条轨迹 */
#define DELETE_ALL_CRUISE 46/* 删除所有轨迹 */

#define PAN_CIRCLE      50   /* 云台以SS的速度自动圆周扫描 */
#define DRAG_PTZ        51   /* 拖动PTZ */
#define LINEAR_SCAN     52   /* 区域扫描 */ //2014-03-15
#define CLE_ALL_PRESET  53   /* 预置点全部清除 */
#define CLE_ALL_SEQ     54   /* 巡航全部清除 */
#define CLE_ALL_CRUISE  55   /* 轨迹全部清除 */

#define POPUP_MENU      56   /* 显示操作菜单 */

#define TILT_DOWN_ZOOM_IN    58    /* 云台以SS的速度下俯&&焦距以速度SS变大(倍率变大) */
#define TILT_DOWN_ZOOM_OUT  59  /* 云台以SS的速度下俯&&焦距以速度SS变小(倍率变小) */
#define PAN_LEFT_ZOOM_IN    60  /* 云台以SS的速度左转&&焦距以速度SS变大(倍率变大)*/
#define PAN_LEFT_ZOOM_OUT   61  /* 云台以SS的速度左转&&焦距以速度SS变小(倍率变小)*/
#define PAN_RIGHT_ZOOM_IN    62  /* 云台以SS的速度右转&&焦距以速度SS变大(倍率变大) */
#define PAN_RIGHT_ZOOM_OUT  63  /* 云台以SS的速度右转&&焦距以速度SS变小(倍率变小) */
#define UP_LEFT_ZOOM_IN     64  /* 云台以SS的速度上仰和左转&&焦距以速度SS变大(倍率变大)*/
#define UP_LEFT_ZOOM_OUT    65  /* 云台以SS的速度上仰和左转&&焦距以速度SS变小(倍率变小)*/
#define UP_RIGHT_ZOOM_IN    66  /* 云台以SS的速度上仰和右转&&焦距以速度SS变大(倍率变大)*/
#define UP_RIGHT_ZOOM_OUT   67  /* 云台以SS的速度上仰和右转&&焦距以速度SS变小(倍率变小)*/
#define DOWN_LEFT_ZOOM_IN   68  /* 云台以SS的速度下俯和左转&&焦距以速度SS变大(倍率变大) */
#define DOWN_LEFT_ZOOM_OUT  69  /* 云台以SS的速度下俯和左转&&焦距以速度SS变小(倍率变小) */
#define DOWN_RIGHT_ZOOM_IN    70  /* 云台以SS的速度下俯和右转&&焦距以速度SS变大(倍率变大) */
#define DOWN_RIGHT_ZOOM_OUT    71  /* 云台以SS的速度下俯和右转&&焦距以速度SS变小(倍率变小) */
#define TILT_UP_ZOOM_IN        72    /* 云台以SS的速度上仰&&焦距以速度SS变大(倍率变大) */
#define TILT_UP_ZOOM_OUT    73


//NET_DVR_Login_V30()参数结构
typedef struct tagNET_DVR_DEVICEINFO_V30
{
    BYTE sSerialNumber[SERIALNO_LEN];  //序列号
    BYTE byAlarmInPortNum;                //报警输入个数
    BYTE byAlarmOutPortNum;                //报警输出个数
    BYTE byDiskNum;                    //硬盘个数
    BYTE byDVRType;                    //设备类型, 1:DVR 2:ATM DVR 3:DVS ......
    BYTE byChanNum;                    //模拟通道个数
    BYTE byStartChan;                    //起始通道号,例如DVS-1,DVR - 1
    BYTE byAudioChanNum;                //语音通道数
    BYTE byIPChanNum;                    //最大数字通道个数，低位
    BYTE byZeroChanNum;            //零通道编码个数 //2010-01-16
    BYTE byMainProto;            //主码流传输协议类型 0-private, 1-rtsp,2-同时支持private和rtsp
    BYTE bySubProto;                //子码流传输协议类型0-private, 1-rtsp,2-同时支持private和rtsp
    BYTE bySupport;
    BYTE bySupport1;
    BYTE bySupport2;
    WORD wDevType;
    BYTE bySupport3;
    BYTE byMultiStreamProto;//是否支持多码流,按位表示,0-不支持,1-支持,bit1-码流3,bit2-码流4,bit7-主码流，bit-8子码流
    BYTE byStartDChan;        //起始数字通道号,0表示无效
    BYTE byStartDTalkChan;    //起始数字对讲通道号，区别于模拟对讲通道号，0表示无效
    BYTE byHighDChanNum;        //数字通道个数，高位
    BYTE bySupport4;
    BYTE byLanguageType;
    BYTE byVoiceInChanNum;   //音频输入通道数
    BYTE byStartVoiceInChanNo; //音频输入起始通道号 0表示无效
    BYTE  bySupport5;
    BYTE  bySupport6;
    BYTE  byMirrorChanNum;    //镜像通道个数，<录播主机中用于表示导播通道>
    WORD wStartMirrorChanNo;  //起始镜像通道号
    BYTE bySupport7;
    BYTE  byRes2;        //保留
}NET_DVR_DEVICEINFO_V30, *LPNET_DVR_DEVICEINFO_V30;

//登录设备信息
typedef struct tagNET_DVR_DEVICEINFO_V40
{
    NET_DVR_DEVICEINFO_V30 struDeviceV30;
    BYTE  bySupportLock;        //设备支持锁定功能，该字段由SDK根据设备返回值来赋值的。bySupportLock为1时，dwSurplusLockTime和byRetryLoginTime有效
    BYTE  byRetryLoginTime;        //剩余可尝试登陆的次数，用户名，密码错误时，此参数有效
    BYTE  byPasswordLevel;      //admin密码安全等级0-无效，1-默认密码，2-有效密码,3-风险较高的密码。当用户的密码为出厂默认密码（12345）或者风险较高的密码时，上层客户端需要提示用户更改密码。
    BYTE  byProxyType;  //代理类型，0-不使用代理, 1-使用socks5代理, 2-使用EHome代理
    DWORD dwSurplusLockTime;    //剩余时间，单位秒，用户锁定时，此参数有效
    BYTE  byCharEncodeType;     //字符编码类型
    BYTE  bySupportDev5;//支持v50版本的设备参数获取，设备名称和设备类型名称长度扩展为64字节
    BYTE  bySupport;  //能力集扩展，位与结果：0- 不支持，1- 支持
    BYTE  byLoginMode; //登录模式 0-Private登录 1-ISAPI登录
    DWORD dwOEMCode;
    int iResidualValidity;   //该用户密码剩余有效天数，单位：天，返回负值，表示密码已经超期使用，例如“-3表示密码已经超期使用3天”
    BYTE  byResidualValidity; // iResidualValidity字段是否有效，0-无效，1-有效
    BYTE  byRes2[243];
}NET_DVR_DEVICEINFO_V40, *LPNET_DVR_DEVICEINFO_V40;

typedef void (*fLoginResultCallBack) (LONG lUserID, DWORD dwResult, LPNET_DVR_DEVICEINFO_V30 lpDeviceInfo , void* pUser);
typedef void (*REALDATACALLBACK) (LONG lPlayHandle, DWORD dwDataType, BYTE *pBuffer, DWORD dwBufSize, void* pUser);

typedef struct tagNET_DVR_USER_LOGIN_INFO
{
    char sDeviceAddress[NET_DVR_DEV_ADDRESS_MAX_LEN];
    BYTE byUseTransport;    //是否启用能力集透传，0--不启用透传，默认，1--启用透传
    WORD wPort;
    char sUserName[NET_DVR_LOGIN_USERNAME_MAX_LEN];
    char sPassword[NET_DVR_LOGIN_PASSWD_MAX_LEN];
    fLoginResultCallBack cbLoginResult;
    void *pUser;
    BOOL bUseAsynLogin;
    BYTE byProxyType; //0:不使用代理，1：使用标准代理，2：使用EHome代理
    BYTE byUseUTCTime;    //0-不进行转换，默认,1-接口上输入输出全部使用UTC时间,SDK完成UTC时间与设备时区的转换,2-接口上输入输出全部使用平台本地时间，SDK完成平台本地时间与设备时区的转换
    BYTE byLoginMode; //0-Private 1-ISAPI 2-自适应
    BYTE byHttps;    //0-不适用tls，1-使用tls 2-自适应
    LONG iProxyID;    //代理服务器序号，添加代理服务器信息时，相对应的服务器数组下表值
    BYTE byVerifyMode;  //认证方式，0-不认证，1-双向认证，2-单向认证；认证仅在使用TLS的时候生效;
    BYTE byRes3[119];
}NET_DVR_USER_LOGIN_INFO,*LPNET_DVR_USER_LOGIN_INFO;

//图片质量
typedef struct tagNET_DVR_JPEGPARA
{
    WORD    wPicSize;
    WORD    wPicQuality;            /* 图片质量系数 0-最好 1-较好 2-一般 */
}NET_DVR_JPEGPARA, *LPNET_DVR_JPEGPARA;

//软解码预览参数
typedef struct tagNET_DVR_CLIENTINFO
{
    LONG lChannel;
    LONG lLinkMode;
    HWND hPlayWnd;
    char* sMultiCastIP;
    BYTE byProtoType;
    BYTE byRes[3];
}NET_DVR_CLIENTINFO, *LPNET_DVR_CLIENTINFO;

#define STREAM_ID_LEN   32

//预览V40接口
typedef struct tagNET_DVR_PREVIEWINFO
{
    LONG lChannel;
    DWORD dwStreamType;
    DWORD dwLinkMode;
    HWND hPlayWnd;
    DWORD bBlocked;
    DWORD bPassbackRecord;
    BYTE byPreviewMode;
    BYTE byStreamID[STREAM_ID_LEN];
    BYTE byProtoType;
    BYTE byRes1;
    BYTE byVideoCodingType;
    DWORD dwDisplayBufNum;
    BYTE byNPQMode;
    BYTE byRes[215];
}NET_DVR_PREVIEWINFO, *LPNET_DVR_PREVIEWINFO;

//报警设备信息
typedef struct tagNET_DVR_ALARMER
{
    BYTE byUserIDValid;                 /* userid是否有效 0-无效，1-有效 */
    BYTE bySerialValid;                 /* 序列号是否有效 0-无效，1-有效 */
    BYTE byVersionValid;                /* 版本号是否有效 0-无效，1-有效 */
    BYTE byDeviceNameValid;             /* 设备名字是否有效 0-无效，1-有效 */
    BYTE byMacAddrValid;                /* MAC地址是否有效 0-无效，1-有效 */
    BYTE byLinkPortValid;               /* login端口是否有效 0-无效，1-有效 */
    BYTE byDeviceIPValid;               /* 设备IP是否有效 0-无效，1-有效 */
    BYTE bySocketIPValid;               /* socket ip是否有效 0-无效，1-有效 */
    LONG lUserID;                       /* NET_DVR_Login()返回值, 布防时有效 */
    BYTE sSerialNumber[SERIALNO_LEN];    /* 序列号 */
    DWORD dwDeviceVersion;                /* 版本信息 高16位表示主版本，低16位表示次版本*/
    char sDeviceName[NAME_LEN];            /* 设备名字 */
    BYTE byMacAddr[MACADDR_LEN];        /* MAC地址 */
    WORD wLinkPort;                     /* link port */
    char sDeviceIP[128];                /* IP地址 */
    char sSocketIP[128];                /* 报警主动上传时的socket IP地址 */
    BYTE byIpProtocol;                  /* Ip协议 0-IPV4, 1-IPV6 */
    BYTE byRes1[2];
    BYTE bJSONBroken;                   //JSON断网续传标志。0：不续传；1：续传
    WORD wSocketPort;
    BYTE byRes2[6];
}NET_DVR_ALARMER, *LPNET_DVR_ALARMER;

//报警布防参数结构体。
typedef struct tagNET_DVR_SETUPALARM_PARAM
{
    DWORD dwSize;
    BYTE  byLevel; //布防优先级，0-一等级（高），1-二等级（中），2-三等级（低）
    BYTE  byAlarmInfoType; //上传报警信息类型（抓拍机支持），0-老报警信息（NET_DVR_PLATE_RESULT），1-新报警信息(NET_ITS_PLATE_RESULT)2012-9-28
    BYTE  byRetAlarmTypeV40; //0--返回NET_DVR_ALARMINFO_V30或NET_DVR_ALARMINFO, 1--设备支持NET_DVR_ALARMINFO_V40则返回NET_DVR_ALARMINFO_V40，不支持则返回NET_DVR_ALARMINFO_V30或NET_DVR_ALARMINFO
    BYTE  byRetDevInfoVersion; //CVR上传报警信息回调结构体版本号 0-COMM_ALARM_DEVICE， 1-COMM_ALARM_DEVICE_V40
    BYTE  byRetVQDAlarmType; //VQD报警上传类型，0-上传报报警NET_DVR_VQD_DIAGNOSE_INFO，1-上传报警NET_DVR_VQD_ALARM
    //1-表示人脸侦测报警扩展(INTER_FACE_DETECTION),0-表示原先支持结构(INTER_FACESNAP_RESULT)
    BYTE  byFaceAlarmDetection;
    //Bit0- 表示二级布防是否上传图片: 0-上传，1-不上传
    //Bit1- 表示开启数据上传确认机制；0-不开启，1-开启
    BYTE  bySupport;
    //断网续传类型
    //bit0-车牌检测（IPC） （0-不续传，1-续传）
    //bit1-客流统计（IPC）  （0-不续传，1-续传）
    //bit2-热度图统计（IPC） （0-不续传，1-续传）
    //bit3-人脸抓拍（IPC） （0-不续传，1-续传）
    //bit4-人脸对比（IPC） （0-不续传，1-续传）
    BYTE  byBrokenNetHttp;
    WORD  wTaskNo;    //任务处理号 和 (上传数据NET_DVR_VEHICLE_RECOG_RESULT中的字段dwTaskNo对应 同时 下发任务结构 NET_DVR_VEHICLE_RECOG_COND中的字段dwTaskNo对应)
    BYTE  byDeployType;    //布防类型：0-客户端布防，1-实时布防
    BYTE  byRes1[3];
    BYTE  byAlarmTypeURL;//bit0-表示人脸抓拍报警上传（INTER_FACESNAP_RESULT）；0-表示二进制传输，1-表示URL传输（设备支持的情况下，设备支持能力根据具体报警能力集判断,同时设备需要支持URL的相关服务，当前是”云存储“）
                         //bit1-表示EVENT_JSON中图片数据长传类型；0-表示二进制传输，1-表示URL传输（设备支持的情况下，设备支持能力根据具体报警能力集判断）
    BYTE  byCustomCtrl;//Bit0- 表示支持副驾驶人脸子图上传: 0-不上传,1-上传,(注：只在公司内部8600/8200等平台开放)
}NET_DVR_SETUPALARM_PARAM, *LPNET_DVR_SETUPALARM_PARAM;

//门禁主机工作状态结构体。
typedef struct tagNET_DVR_ACS_WORK_STATUS
{
    DWORD dwSize;
    BYTE byDoorLockStatus[MAX_DOOR_NUM]; //门锁状态，0-关，1-开
    BYTE byDoorStatus[MAX_DOOR_NUM]; //门状态，1-休眠，2-常开状态，3-常闭状态，4-普通状态
    BYTE byMagneticStatus[MAX_DOOR_NUM]; //门磁状态，0-闭合，1-开启
    BYTE byCaseStatus[MAX_CASE_SENSOR_NUM]; ////事件触发器状态，0-无输入，1-有输入
    WORD wBatteryVoltage; //蓄电池电压值，实际值乘10，单位：伏特
    BYTE byBatteryLowVoltage; //蓄电池是否处于低压状态，0-否，1-是
    BYTE byPowerSupplyStatus; //设备供电状态，1-交流电供电，2-蓄电池供电
    BYTE byMultiDoorInterlockStatus; //多门互锁状态，0-关闭，1-开启
    BYTE byAntiSneakStatus; //反潜回状态，0-关闭，1-开启
    BYTE byHostAntiDismantleStatus; //主机防拆状态，0-关闭，1-开启
    BYTE byIndicatorLightStatus; //指示灯状态，0-掉线，1-在线
    BYTE byCardReaderOnlineStatus[MAX_CARD_READER_NUM]; //读卡器在线状态，0-不在线，1-在线
    BYTE byCardReaderAntiDismantleStatus[MAX_CARD_READER_NUM]; //读卡器防拆状态，0-关闭，1-开启
    BYTE byCardReaderVerifyMode[MAX_CARD_READER_NUM]; //读卡器当前验证方式，1-刷卡，2-刷卡+密码，3-刷卡，4-刷卡或密码
    BYTE bySetupAlarmStatus[MAX_ALARMHOST_ALARMIN_NUM];//报警输入口布防状态，0-对应报警输入口处于撤防状态，1-对应报警输入口处于布防状态
    BYTE byAlarmInStatus[MAX_ALARMHOST_ALARMIN_NUM]; //报警输入口报警状态，0-对应报警输入口当前无报警，1-对应报警输入口当前有报警
    BYTE byAlarmOutStatus[MAX_ALARMHOST_ALARMOUT_NUM]; //报警输出口状态，0-对应报警输出口无报警，1-对应报警输出口有报警
    DWORD dwCardNum; //已添加的卡数量
    BYTE byRes2[32];
}NET_DVR_ACS_WORK_STATUS, *LPNET_DVR_ACS_WORK_STATUS;

//卡参数配置条件结构体。old
typedef struct tagNET_DVR_CARD_COND
{
    DWORD dwSize;
    DWORD dwCardNum; //设置或获取卡数量，获取时置为0xffffffff表示获取所有卡信息
    BYTE  byRes[64];
}NET_DVR_CARD_COND, *LPNET_DVR_CARD_COND;

//卡参数配置条件结构体。
typedef struct _NET_DVR_CARD_CFG_COND
{
    DWORD dwSize;
    DWORD dwCardNum; //设置或获取卡数量，获取时置为0xffffffff表示获取所有卡信息
    BYTE  byCheckCardNo; //设备是否进行卡号校验，0-不校验，1-校验
    BYTE           byRes1[3];
    WORD wLocalControllerID; //就地控制器序号，表示往就地控制器下发离线卡参数，0代表是门禁主机
    BYTE  byRes2[2];
    DWORD dwLockID;  //锁ID
    BYTE  byRes3[20];
}NET_DVR_CARD_CFG_COND, *LPNET_DVR_CARD_CFG_COND;


//时间EX
typedef struct tagNET_DVR_TIME_EX
{
    WORD wYear;
    BYTE byMonth;
    BYTE byDay;
    BYTE byHour;
    BYTE byMinute;
    BYTE bySecond;
    BYTE byRes;
}NET_DVR_TIME_EX,*LPNET_DVR_TIME_EX;

//有效期参数
typedef struct tagNET_DVR_VALID_PERIOD_CFG
{
    BYTE byEnable; //使能有效期，0-不使能，1使能
    BYTE byBeginTimeFlag;      //是否限制起始时间的标志，0-不限制，1-限制
    BYTE byEnableTimeFlag;     //是否限制终止时间的标志，0-不限制，1-限制
    BYTE byTimeDurationNo;     //有效期索引,从0开始（时间段通过SDK设置给锁，后续在制卡时，只需要传递有效期索引即可，以减少数据量）
    NET_DVR_TIME_EX struBeginTime; //有效期起始时间
    NET_DVR_TIME_EX struEndTime; //有效期结束时间
    BYTE byTimeType; //时间类型：0-设备本地时间（默认），1-UTC时间（对于struBeginTime，struEndTime字段有效）
    BYTE byRes2[31];
}NET_DVR_VALID_PERIOD_CFG, *LPNET_DVR_VALID_PERIOD_CFG;

//返回卡信息结构体
typedef struct _NET_DVR_CARD_RECORD
{
    DWORD                      dwSize;
    BYTE                        byCardNo[ACS_CARD_NO_LEN];
    BYTE                        byCardType;
    BYTE                        byLeaderCard;
    BYTE                        byUserType;
    BYTE                        byRes1;
    BYTE                        byDoorRight[MAX_DOOR_NUM_256];
    NET_DVR_VALID_PERIOD_CFG    struValid;
    BYTE                        byBelongGroup[MAX_GROUP_NUM_128];
    BYTE                        byCardPassword[CARD_PASSWORD_LEN];
    WORD                        wCardRightPlan[MAX_DOOR_NUM_256];
    DWORD                       dwMaxSwipeTimes;
    DWORD                       dwSwipeTimes;
    DWORD                       dwEmployeeNo;
    BYTE                        byName[NAME_LEN];
    //按位表示，0-无权限，1-有权限
    //第0位表示：弱电报警
    //第1位表示：开门提示音
    //第2位表示：限制客卡
    //第3位表示：通道
    //第4位表示：反锁开门
    //第5位表示：巡更功能
    DWORD                      dwCardRight;
    BYTE                       byRes[256];
}NET_DVR_CARD_RECORD, *LPNET_DVR_CARD_RECORD;

typedef struct tagNET_DVR_ALARMHOST_MAIN_STATUS_V51
{
    DWORD  dwSize;
    BYTE   bySetupAlarmStatus[MAX_ALARMHOST_ALARMIN_NUM]; //防区布防状态，(最大支持512个防区查询)，0xff-无效，0-对应防区处于撤防状态，1-对应防区处于布防状态，2-对应防区处于布防中
    BYTE   byAlarmInStatus[MAX_ALARMHOST_ALARMIN_NUM]; //防区报警状态（触发状态），(最大支持512个防区查询)，0xff-无效，0-对应防区当前无报警，1-对应防区当前有报警
    BYTE   byAlarmOutStatus[MAX_ALARMHOST_ALARMOUT_NUM]; //触发器状态，(最大支持512个触发器查询)，0xff-无效，0-对应触发器无报警，1-对应触发器有报警，2-未关联，3-离线，4-心跳异常
    BYTE   byBypassStatus[MAX_ALARMHOST_ALARMIN_NUM]; //防区旁路状态，数组下标表示0对应防区1，0xff-无效，0-表示防区没有旁路 1-表示防区旁路
    BYTE   bySubSystemGuardStatus[MAX_ALARMHOST_SUBSYSTEM/*32*/]; //子系统布防状态，0xff-无效，0-对应子系统处于撤防状态，1-对应子系统处于布防状态，2-对应子系统处于布防中
    BYTE   byAlarmInFaultStatus[MAX_ALARMHOST_ALARMIN_NUM]; //防区故障状态，0xff-无效，0-对应防区处于正常状态，1-对应防区处于故障状态
    BYTE   byAlarmInMemoryStatus[MAX_ALARMHOST_ALARMIN_NUM]; //防区报警记忆状态（报警状态）， 0xff-无效，0-对应防区当前无报警，1-对应防区当前有报警
    BYTE   byAlarmInTamperStatus[MAX_ALARMHOST_ALARMIN_NUM]; //防区防拆状态，0xff-无效，0-对应防区无报警，1-对应防区有报警
    BYTE   byEnableSubSystem[MAX_ALARMHOST_SUBSYSTEM/*32*/]; //子系统启用状态，0-无效，1-对应子系统未启用，2-对应子系统启用
    BYTE   bySubSystemGuardType[MAX_ALARMHOST_SUBSYSTEM]; //子系统布防类型，0-无效，1-外出布防，2-即时布防，3-在家布防
    BYTE   bySubSystemAlarm[MAX_ALARMHOST_SUBSYSTEM]; //子系统报警状态，0-无效，1-正常，2-报警
    BYTE   byAlarmOutCharge[MAX_ALARMHOST_ALARMOUT_NUM]; //触发器电量状态，(最大支持512个触发器查询)，0-无效，1-正常，2-电量低
    BYTE   byAlarmOutTamperStatus[MAX_ALARMHOST_ALARMOUT_NUM]; //触发器防拆状态，(最大支持512个触发器查询)，0-无效，1-防拆，2-无防拆
    BYTE   byAlarmInShieldedStatus[MAX_ALARMHOST_ALARMIN_NUM]; //防区屏蔽状态，0-无效，1-屏蔽，2-非屏蔽
    BYTE   byAlarmOutLinkage[MAX_ALARMHOST_ALARMOUT_NUM]; //触发器联动事件类型，(最大支持512个触发器查询)，0-无效，1-报警，2-布防，3-撤防，4-手动控制
    BYTE   byRes[512]; //保留字节
}NET_DVR_ALARMHOST_MAIN_STATUS_V51, *LPNET_DVR_ALARMHOST_MAIN_STATUS_V51;


BOOL NET_DVR_Init();
BOOL NET_DVR_Cleanup();

LONG NET_DVR_Login_V30(char *sDVRIP, WORD wDVRPort, char *sUserName, char *sPassword, LPNET_DVR_DEVICEINFO_V30 lpDeviceInfo);
LONG NET_DVR_Login_V40(LPNET_DVR_USER_LOGIN_INFO pLoginInfo,LPNET_DVR_DEVICEINFO_V40 lpDeviceInfo);
BOOL NET_DVR_Logout(LONG lUserID);
BOOL NET_DVR_Logout_V30(LONG lUserID);

BOOL NET_DVR_SetConnectTime(DWORD dwWaitTime, DWORD dwTryTimes);
BOOL NET_DVR_SetReconnect(DWORD dwInterval, BOOL bEnableRecon);
BOOL NET_DVR_CaptureJPEGPicture(LONG lUserID, LONG lChannel, LPNET_DVR_JPEGPARA lpJpegPara, char *sPicFileName);

BOOL  NET_DVR_SetLogToFile(DWORD nLogLevel , char * strLogDir, BOOL bAutoDel);

BOOL NET_DVR_CaptureJPEGPicture(LONG lUserID, LONG lChannel, LPNET_DVR_JPEGPARA lpJpegPara, char *sPicFileName);
LONG NET_DVR_RealPlay_V30(LONG lUserID, LPNET_DVR_CLIENTINFO lpClientInfo, void(*fRealDataCallBack_V30) (LONG lRealHandle, DWORD dwDataType, BYTE *pBuffer, DWORD dwBufSize, void* pUser), void* pUser, BOOL bBlocked);
BOOL NET_DVR_ClosePreview(LONG lUserID, DWORD nSessionID);
BOOL NET_DVR_ClosePlayBack(LONG lUserID, DWORD nSessionID);
LONG NET_DVR_RealPlay_V40(LONG lUserID, LPNET_DVR_PREVIEWINFO lpPreviewInfo, REALDATACALLBACK fRealDataCallBack_V30, void* pUser);

BOOL NET_DVR_SaveRealData(LONG lRealHandle,char *sFileName);
BOOL NET_DVR_StopSaveRealData(LONG lRealHandle);

BOOL NET_DVR_PTZControlWithSpeed(LONG lRealHandle, DWORD dwPTZCommand, DWORD dwStop, DWORD dwSpeed);
BOOL NET_DVR_StopRealPlay(LONG lRealHandle);

//启用写日志文件
BOOL NET_DVR_SetLogToFile(DWORD nLogLevel,char * strLogDir, BOOL bAutoDel);

//布防撤防
LONG NET_DVR_SetupAlarmChan_V41(LONG lUserID, LPNET_DVR_SETUPALARM_PARAM lpSetupParam);
BOOL NET_DVR_CloseAlarmChan_V30(LONG lAlarmHandle);

//设置报警回调函数
typedef BOOL (CALLBACK *MSGCallBack_V31)(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *pAlarmInfo, DWORD dwBufLen, void* pUser);
BOOL NET_DVR_SetDVRMessageCallBack_V31(MSGCallBack_V31 fMessageCallBack, void* pUser);

//获取设备状态
BOOL NET_DVR_GetDVRConfig(LONG lUserID, DWORD dwCommand,LONG lChannel, LPVOID lpOutBuffer, DWORD dwOutBufferSize, LPDWORD lpBytesReturned);

//控制门状态
BOOL NET_DVR_ControlGateway(LONG lUserID, LONG lGatewayIndex, DWORD dwStaic);

//建立长连接
typedef void(CALLBACK *fRemoteConfigCallback)(DWORD dwType, void* lpBuffer, DWORD dwBufLen, void* pUserData);
 LONG  NET_DVR_StartRemoteConfig(LONG lUserID, DWORD dwCommand, LPVOID lpInBuffer, DWORD dwInBufferLen, fRemoteConfigCallback cbStateCallback, LPVOID pUserData);
 BOOL  NET_DVR_StopRemoteConfig(LONG lHandle);
 LONG  NET_DVR_GetNextRemoteConfig(LONG lHandle, void* lpOutBuff, DWORD dwOutBuffSize);
 BOOL  NET_DVR_GetRemoteConfigState(LONG lHandle, void *pState);
 BOOL  NET_DVR_SendRemoteConfig(LONG lHandle, DWORD dwDataType, char *pSendBuf, DWORD dwBufSize);
 LONG  NET_DVR_SendWithRecvRemoteConfig(LONG lHandle, void* lpInBuff, DWORD dwInBuffSize, void* lpOutBuff, DWORD dwOutBuffSize, DWORD *dwOutDataLen);

//获取错误码
DWORD NET_DVR_GetLastError();

#endif