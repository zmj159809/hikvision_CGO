package hikvision_CGO

import "unsafe"

type LONG int32
type LLONG int64
type DWORD uint32
type WORD uint16
type BYTE byte
type BOOL int32

const (
	COMM_ALARM_ACS = 0x5002 //门禁主机报警信息
)

//事件类型
const (
	MAJOR_ALARM     DWORD = 0x1 //报警
	MAJOR_EXCEPTION DWORD = 0x2 //异常
	MAJOR_OPERATION DWORD = 0x3 //操作
	MAJOR_EVENT     DWORD = 0x5 //事件

	// 报警主类型 01
	MINOR_ALARMIN_SHORT_CIRCUIT                 DWORD = 0x400 //防区短路报警
	MINOR_ALARMIN_BROKEN_CIRCUIT                DWORD = 0x401 //防区断路报警
	MINOR_ALARMIN_EXCEPTION                     DWORD = 0x402 //防区异常报警
	MINOR_ALARMIN_RESUME                        DWORD = 0x403 //防区报警恢复
	MINOR_HOST_DESMANTLE_ALARM                  DWORD = 0x404 //设备防拆报警
	MINOR_HOST_DESMANTLE_RESUME                 DWORD = 0x405 //设备防拆恢复
	MINOR_CARD_READER_DESMANTLE_ALARM           DWORD = 0x406 //读卡器防拆报警
	MINOR_CARD_READER_DESMANTLE_RESUME          DWORD = 0x407 //读卡器防拆恢复
	MINOR_CASE_SENSOR_ALARM                     DWORD = 0x408 //事件输入报警
	MINOR_CASE_SENSOR_RESUME                    DWORD = 0x409 //事件输入恢复
	MINOR_STRESS_ALARM                          DWORD = 0x40a //胁迫报警
	MINOR_OFFLINE_ECENT_NEARLY_FULL             DWORD = 0x40b //离线事件满90%报警
	MINOR_CARD_MAX_AUTHENTICATE_FAIL            DWORD = 0x40c //卡号认证失败超次报警
	MINOR_SD_CARD_FULL                          DWORD = 0x40d //SD卡存储满报警
	MINOR_LINKAGE_CAPTURE_PIC                   DWORD = 0x40e //联动抓拍事件报警
	MINOR_SECURITY_MODULE_DESMANTLE_ALARM       DWORD = 0x40f //门控安全模块防拆报警
	MINOR_SECURITY_MODULE_DESMANTLE_RESUME      DWORD = 0x410 //门控安全模块防拆恢复
	MINOR_POS_START_ALARM                       DWORD = 0x411 //POS开启
	MINOR_POS_END_ALARM                         DWORD = 0x412 //POS结束
	MINOR_FACE_IMAGE_QUALITY_LOW                DWORD = 0x413 //人脸图像画质低
	MINOR_FINGE_RPRINT_QUALITY_LOW              DWORD = 0x414 //指纹图像画质低
	MINOR_FIRE_IMPORT_SHORT_CIRCUIT             DWORD = 0x415 //消防输入短路报警
	MINOR_FIRE_IMPORT_BROKEN_CIRCUIT            DWORD = 0x416 //消防输入断路报警
	MINOR_FIRE_IMPORT_RESUME                    DWORD = 0x417 //消防输入恢复
	MINOR_FIRE_BUTTON_TRIGGER                   DWORD = 0x418 //消防按钮触发
	MINOR_FIRE_BUTTON_RESUME                    DWORD = 0x419 //消防按钮恢复
	MINOR_MAINTENANCE_BUTTON_TRIGGER            DWORD = 0x41a //维护按钮触发
	MINOR_MAINTENANCE_BUTTON_RESUME             DWORD = 0x41b //维护按钮恢复
	MINOR_EMERGENCY_BUTTON_TRIGGER              DWORD = 0x41c //紧急按钮触发
	MINOR_EMERGENCY_BUTTON_RESUME               DWORD = 0x41d //紧急按钮恢复
	MINOR_DISTRACT_CONTROLLER_ALARM             DWORD = 0x41e //分控器防拆报警
	MINOR_DISTRACT_CONTROLLER_RESUME            DWORD = 0x41f //分控器防拆报警恢复
	MINOR_CHANNEL_CONTROLLER_DESMANTLE_ALARM    DWORD = 0x422 //通道控制器防拆报警
	MINOR_CHANNEL_CONTROLLER_DESMANTLE_RESUME   DWORD = 0x423 //通道控制器防拆报警恢复
	MINOR_CHANNEL_CONTROLLER_FIRE_IMPORT_ALARM  DWORD = 0x424 //通道控制器消防输入报警
	MINOR_CHANNEL_CONTROLLER_FIRE_IMPORT_RESUME DWORD = 0x425 //通道控制器消防输入报警恢复
	MINOR_LEGAL_EVENT_NEARLY_FULL               DWORD = 0x442 //合法事件满90%报警
	MINOR_LOCK_HIJIACK_ALARM                    DWORD = 0x95d //智能锁防劫持报警

	// 异常主类型 02
	MINOR_NET_BROKEN                           DWORD = 0x27  //网络断开
	MINOR_RS485_DEVICE_ABNORMAL                DWORD = 0x3a  //RS485连接状态异常
	MINOR_RS485_DEVICE_REVERT                  DWORD = 0x3b  //RS485连接状态异常恢复
	MINOR_DEV_POWER_ON                         DWORD = 0x400 //设备上电启动
	MINOR_DEV_POWER_OFF                        DWORD = 0x401 //设备掉电关闭
	MINOR_WATCH_DOG_RESET                      DWORD = 0x402 //看门狗复位
	MINOR_LOW_BATTERY                          DWORD = 0x403 // 蓄电池电压低
	MINOR_BATTERY_RESUME                       DWORD = 0x404 // 蓄电池电压恢复正常
	MINOR_AC_OFF                               DWORD = 0x405 // 交流电断电
	MINOR_AC_RESUME                            DWORD = 0x406 //交流电恢复
	MINOR_NET_RESUME                           DWORD = 0x407 //网络恢复
	MINOR_FLASH_ABNORMAL                       DWORD = 0x408 //FLASH读写异常
	MINOR_CARD_READER_OFFLINE                  DWORD = 0x409 // 读卡器掉线
	MINOR_CARD_READER_RESUME                   DWORD = 0x40a // 读卡器掉线恢复
	MINOR_INDICATOR_LIGHT_OFF                  DWORD = 0x40b // 指示灯关闭
	MINOR_INDICATOR_LIGHT_RESUME               DWORD = 0x40c // 指示灯恢复
	MINOR_CHANNEL_CONTROLLER_OFF               DWORD = 0x40d // 通道控制器掉线
	MINOR_CHANNEL_CONTROLLER_RESUME            DWORD = 0x40e // 通道控制器恢复
	MINOR_SECURITY_MODULE_OFF                  DWORD = 0x40f //门控安全模块掉线
	MINOR_SECURITY_MODULE_RESUME               DWORD = 0x410 // 门控安全模块掉线恢复
	MINOR_BATTERY_ELECTRIC_LOW                 DWORD = 0x411 // 电池电压低（仅人脸设备使用）
	MINOR_BATTERY_ELECTRIC_RESUME              DWORD = 0x412 //电池电压恢复正常（仅人脸设备使用）
	MINOR_LOCAL_CONTROL_NET_BROKEN             DWORD = 0x413 //就地控制器网络断开
	MINOR_LOCAL_CONTROL_NET_RSUME              DWORD = 0x414 // 就地控制器网络恢复
	MINOR_MASTER_RS485_LOOPNODE_BROKEN         DWORD = 0x415 // 主控RS485环路节点断开
	MINOR_MASTER_RS485_LOOPNODE_RESUME         DWORD = 0x416 //主控RS485环路节点恢复
	MINOR_LOCAL_CONTROL_OFFLINE                DWORD = 0x417 // 就地控制器掉线
	MINOR_LOCAL_CONTROL_RESUME                 DWORD = 0x418 // 就地控制器掉线恢复
	MINOR_LOCAL_DOWNSIDE_RS485_LOOPNODE_BROKEN DWORD = 0x419 // 就地下行RS485环路断开
	MINOR_LOCAL_DOWNSIDE_RS485_LOOPNODE_RESUME DWORD = 0x41a //就地下行RS485环路恢复
	MINOR_DISTRACT_CONTROLLER_ONLINE           DWORD = 0x41b //分控器在线
	MINOR_DISTRACT_CONTROLLER_OFFLINE          DWORD = 0x41c //分控器离线
	MINOR_ID_CARD_READER_NOT_CONNECT           DWORD = 0x41d //身份证阅读器未连接（智能专用）
	MINOR_ID_CARD_READER_RESUME                DWORD = 0x41e // 身份证阅读器连接恢复（智能专用）
	MINOR_FINGER_PRINT_MODULE_NOT_CONNECT      DWORD = 0x41f //指纹模组未连接（智能专用）
	MINOR_FINGER_PRINT_MODULE_RESUME           DWORD = 0x420 //指纹模组连接恢复（智能专用）
	MINOR_CAMERA_NOT_CONNECT                   DWORD = 0x421 //摄像头未连接
	MINOR_CAMERA_RESUME                        DWORD = 0x422 // 摄像头连接恢复
	MINOR_COM_NOT_CONNECT                      DWORD = 0x423 //COM口未连接
	MINOR_COM_RESUME                           DWORD = 0x424 //COM口连接恢复
	MINOR_DEVICE_NOT_AUTHORIZE                 DWORD = 0x425 //设备未授权
	MINOR_PEOPLE_AND_ID_CARD_DEVICE_ONLINE     DWORD = 0x426 // 人证设备在线
	MINOR_PEOPLE_AND_ID_CARD_DEVICE_OFFLINE    DWORD = 0x427 //人证设备离线
	MINOR_LOCAL_LOGIN_LOCK                     DWORD = 0x428 //本地登录锁定
	MINOR_LOCAL_LOGIN_UNLOCK                   DWORD = 0x429 //本地登录解锁
	MINOR_SUBMARINEBACK_COMM_BREAK             DWORD = 0x42a //与反潜回服务器通信断开
	MINOR_SUBMARINEBACK_COMM_RESUME            DWORD = 0x42b //与反潜回服务器通信恢复
	MINOR_MOTOR_SENSOR_EXCEPTION               DWORD = 0x42c //电机或传感器异常
	MINOR_CAN_BUS_EXCEPTION                    DWORD = 0x42d // CAN总线异常
	MINOR_CAN_BUS_RESUME                       DWORD = 0x42e // CAN总线恢复
	MINOR_GATE_TEMPERATURE_OVERRUN             DWORD = 0x42f // 闸机腔体温度超限
	MINOR_IR_EMITTER_EXCEPTION                 DWORD = 0x430 // 红外对射异常
	MINOR_IR_EMITTER_RESUME                    DWORD = 0x431 // 红外对射恢复
	MINOR_LAMP_BOARD_COMM_EXCEPTION            DWORD = 0x432 // 灯板通信异常
	MINOR_LAMP_BOARD_COMM_RESUME               DWORD = 0x433 //灯板通信恢复
	MINOR_IR_ADAPTOR_COMM_EXCEPTION            DWORD = 0x434 //红外转接板通信异常
	MINOR_IR_ADAPTOR_COMM_RESUME               DWORD = 0x435 //红外转接板通信恢复

	//操作类型 03
	MINOR_LOCAL_LOGIN                         DWORD = 0x50  //本地登陆
	MINOR_LOCAL_LOGOUT                        DWORD = 0x51  //本地注销登陆
	MINOR_LOCAL_UPGRADE                       DWORD = 0x5a  //本地升级
	MINOR_REMOTE_LOGIN                        DWORD = 0x70  //远程登录
	MINOR_REMOTE_LOGOUT                       DWORD = 0x71  //远程注销登陆
	MINOR_REMOTE_ARM                          DWORD = 0x79  //远程布防
	MINOR_REMOTE_DISARM                       DWORD = 0x7a  //远程撤防
	MINOR_REMOTE_REBOOT                       DWORD = 0x7b  //远程重启
	MINOR_REMOTE_UPGRADE                      DWORD = 0x7e  //远程升级
	MINOR_REMOTE_CFGFILE_OUTPUT               DWORD = 0x86  //远程导出配置文件
	MINOR_REMOTE_CFGFILE_INTPUT               DWORD = 0x87  //远程导入配置文件
	MINOR_REMOTE_ALARMOUT_OPEN_MAN            DWORD = 0xd6  //远程手动开启报警输出
	MINOR_REMOTE_ALARMOUT_CLOSE_MAN           DWORD = 0xd7  //远程手动关闭报警输出
	MINOR_REMOTE_OPEN_DOOR                    DWORD = 0x400 //远程开门
	MINOR_REMOTE_CLOSE_DOOR                   DWORD = 0x401 //远程关门（对于梯控，表示受控）
	MINOR_REMOTE_ALWAYS_OPEN                  DWORD = 0x402 //远程常开（对于梯控，表示自由）
	MINOR_REMOTE_ALWAYS_CLOSE                 DWORD = 0x403 //远程常关（对于梯控，表示禁用）
	MINOR_REMOTE_CHECK_TIME                   DWORD = 0x404 //远程手动校时
	MINOR_NTP_CHECK_TIME                      DWORD = 0x405 //NTP自动校时
	MINOR_REMOTE_CLEAR_CARD                   DWORD = 0x406 //远程清空卡号
	MINOR_REMOTE_RESTORE_CFG                  DWORD = 0x407 //远程恢复默认参数
	MINOR_ALARMIN_ARM                         DWORD = 0x408 //防区布防
	MINOR_ALARMIN_DISARM                      DWORD = 0x409 //防区撤防
	MINOR_LOCAL_RESTORE_CFG                   DWORD = 0x40a //本地恢复默认参数
	MINOR_REMOTE_CAPTURE_PIC                  DWORD = 0x40b //远程抓拍
	MINOR_MOD_NET_REPORT_CFG                  DWORD = 0x40c //修改网络中心参数配置
	MINOR_MOD_GPRS_REPORT_PARAM               DWORD = 0x40d //修改GPRS中心参数配置
	MINOR_MOD_REPORT_GROUP_PARAM              DWORD = 0x40e //修改中心组参数配置
	MINOR_UNLOCK_PASSWORD_OPEN_DOOR           DWORD = 0x40f //解除码输入
	MINOR_AUTO_RENUMBER                       DWORD = 0x410 //自动重新编号
	MINOR_AUTO_COMPLEMENT_NUMBER              DWORD = 0x411 //自动补充编号
	MINOR_NORMAL_CFGFILE_INPUT                DWORD = 0x412 //导入普通配置文件
	MINOR_NORMAL_CFGFILE_OUTTPUT              DWORD = 0x413 //导出普通配置文件
	MINOR_CARD_RIGHT_INPUT                    DWORD = 0x414 //导入卡权限参数
	MINOR_CARD_RIGHT_OUTTPUT                  DWORD = 0x415 //导出卡权限参数
	MINOR_LOCAL_USB_UPGRADE                   DWORD = 0x416 //本地U盘升级
	MINOR_REMOTE_VISITOR_CALL_LADDER          DWORD = 0x417 //访客呼梯
	MINOR_REMOTE_HOUSEHOLD_CALL_LADDER        DWORD = 0x418 //住户呼梯
	MINOR_REMOTE_ACTUAL_GUARD                 DWORD = 0x419 //远程实时布防
	MINOR_REMOTE_ACTUAL_UNGUARD               DWORD = 0x41a //远程实时撤防
	MINOR_REMOTE_CONTROL_NOT_CODE_OPER_FAILED DWORD = 0x41b //遥控器未对码操作失败
	MINOR_REMOTE_CONTROL_CLOSE_DOOR           DWORD = 0x41c //遥控器关门
	MINOR_REMOTE_CONTROL_OPEN_DOOR            DWORD = 0x41d //遥控器开门
	MINOR_REMOTE_CONTROL_ALWAYS_OPEN_DOOR     DWORD = 0x41e //遥控器常开门

	//事件类型 05
	MINOR_LEGAL_CARD_PASS                         DWORD = 0x01 //合法卡认证通过
	MINOR_CARD_AND_PSW_PASS                       DWORD = 0x02 //刷卡加密码认证通过
	MINOR_CARD_AND_PSW_FAIL                       DWORD = 0x03 //刷卡加密码认证失败
	MINOR_CARD_AND_PSW_TIMEOUT                    DWORD = 0x04 //数卡加密码认证超时
	MINOR_CARD_AND_PSW_OVER_TIME                  DWORD = 0x05 //刷卡加密码超次
	MINOR_CARD_NO_RIGHT                           DWORD = 0x06 //未分配权限
	MINOR_CARD_INVALID_PERIOD                     DWORD = 0x07 //无效时段
	MINOR_CARD_OUT_OF_DATE                        DWORD = 0x08 //卡号过期
	MINOR_INVALID_CARD                            DWORD = 0x09 //无此卡号
	MINOR_ANTI_SNEAK_FAIL                         DWORD = 0x0a //反潜回认证失败
	MINOR_INTERLOCK_DOOR_NOT_CLOSE                DWORD = 0x0b //互锁门未关闭
	MINOR_NOT_BELONG_MULTI_GROUP                  DWORD = 0x0c //卡不属于多重认证群组
	MINOR_INVALID_MULTI_VERIFY_PERIOD             DWORD = 0x0d //卡不在多重认证时间段内
	MINOR_MULTI_VERIFY_SUPER_RIGHT_FAIL           DWORD = 0x0e //多重认证模式超级权限认证失败
	MINOR_MULTI_VERIFY_REMOTE_RIGHT_FAIL          DWORD = 0x0f //多重认证模式远程认证失败
	MINOR_MULTI_VERIFY_SUCCESS                    DWORD = 0x10 //多重认证成功
	MINOR_LEADER_CARD_OPEN_BEGIN                  DWORD = 0x11 //首卡开门开始
	MINOR_LEADER_CARD_OPEN_END                    DWORD = 0x12 //首卡开门结束
	MINOR_ALWAYS_OPEN_BEGIN                       DWORD = 0x13 //常开状态开始
	MINOR_ALWAYS_OPEN_END                         DWORD = 0x14 //常开状态结束
	MINOR_LOCK_OPEN                               DWORD = 0x15 //门锁打开
	MINOR_LOCK_CLOSE                              DWORD = 0x16 //门锁关闭
	MINOR_DOOR_BUTTON_PRESS                       DWORD = 0x17 //开门按钮打开
	MINOR_DOOR_BUTTON_RELEASE                     DWORD = 0x18 //开门按钮放开
	MINOR_DOOR_OPEN_NORMAL                        DWORD = 0x19 //正常开门（门磁）
	MINOR_DOOR_CLOSE_NORMAL                       DWORD = 0x1a //正常关门（门磁）
	MINOR_DOOR_OPEN_ABNORMAL                      DWORD = 0x1b //门异常打开（门磁）
	MINOR_DOOR_OPEN_TIMEOUT                       DWORD = 0x1c //门打开超时（门磁）
	MINOR_ALARMOUT_ON                             DWORD = 0x1d //报警输出打开
	MINOR_ALARMOUT_OFF                            DWORD = 0x1e //报警输出关闭
	MINOR_ALWAYS_CLOSE_BEGIN                      DWORD = 0x1f //常关状态开始
	MINOR_ALWAYS_CLOSE_END                        DWORD = 0x20 //常关状态结束
	MINOR_MULTI_VERIFY_NEED_REMOTE_OPEN           DWORD = 0x21 //多重多重认证需要远程开门
	MINOR_MULTI_VERIFY_SUPERPASSWD_VERIFY_SUCCESS DWORD = 0x22 //多重认证超级密码认证成功事件
	MINOR_MULTI_VERIFY_REPEAT_VERIFY              DWORD = 0x23 //多重认证重复认证事件
	MINOR_MULTI_VERIFY_TIMEOUT                    DWORD = 0x24 //多重认证重复认证事件
	MINOR_DOORBELL_RINGING                        DWORD = 0x25 //门铃响
	MINOR_FINGERPRINT_COMPARE_PASS                DWORD = 0x26 //指纹比对通过
	MINOR_FINGERPRINT_COMPARE_FAIL                DWORD = 0x27 //指纹比对失败
	MINOR_CARD_FINGERPRINT_VERIFY_PASS            DWORD = 0x28 //刷卡加指纹认证通过
	MINOR_CARD_FINGERPRINT_VERIFY_FAIL            DWORD = 0x29 //刷卡加指纹认证失败
	MINOR_CARD_FINGERPRINT_VERIFY_TIMEOUT         DWORD = 0x2a //刷卡加指纹认证超时
	MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_PASS     DWORD = 0x2b //刷卡加指纹加密码认证通过
	MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_FAIL     DWORD = 0x2c //刷卡加指纹加密码认证失败
	MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_TIMEOUT  DWORD = 0x2d //刷卡加指纹加密码认证超时
	MINOR_FINGERPRINT_PASSWD_VERIFY_PASS          DWORD = 0x2e //指纹加密码认证通过
	MINOR_FINGERPRINT_PASSWD_VERIFY_FAIL          DWORD = 0x2f //指纹加密码认证失败
	MINOR_FINGERPRINT_PASSWD_VERIFY_TIMEOUT       DWORD = 0x30 //指纹加密码认证超时
	MINOR_FINGERPRINT_INEXISTENCE                 DWORD = 0x31 //指纹不存在
	MINOR_CARD_PLATFORM_VERIFY                    DWORD = 0x32 //刷卡平台认证
	MINOR_CALL_CENTER                             DWORD = 0x33 //呼叫中心事件
	MINOR_FIRE_RELAY_TURN_ON_DOOR_ALWAYS_OPEN     DWORD = 0x34 //消防继电器导通触发门常开
	MINOR_FIRE_RELAY_RECOVER_DOOR_RECOVER_NORMAL  DWORD = 0x35 //消防继电器恢复门恢复正常
	MINOR_FACE_AND_FP_VERIFY_PASS                 DWORD = 0x36 //人脸加指纹认证通过
	MINOR_FACE_AND_FP_VERIFY_FAIL                 DWORD = 0x37 //人脸加指纹认证失败
	MINOR_FACE_AND_FP_VERIFY_TIMEOUT              DWORD = 0x38 //人脸加指纹认证超时
	MINOR_FACE_AND_PW_VERIFY_PASS                 DWORD = 0x39 //人脸加密码认证通过
	MINOR_FACE_AND_PW_VERIFY_FAIL                 DWORD = 0x3a //人脸加密码认证失败
	MINOR_FACE_AND_PW_VERIFY_TIMEOUT              DWORD = 0x3b //人脸加密码认证超时
	MINOR_FACE_AND_CARD_VERIFY_PASS               DWORD = 0x3c //人脸加刷卡认证通过
	MINOR_FACE_AND_CARD_VERIFY_FAIL               DWORD = 0x3d //人脸加刷卡认证失败
	MINOR_FACE_AND_CARD_VERIFY_TIMEOUT            DWORD = 0x3e //人脸加刷卡认证超时
	MINOR_FACE_AND_PW_AND_FP_VERIFY_PASS          DWORD = 0x3f //人脸加密码加指纹认证通过
	MINOR_FACE_AND_PW_AND_FP_VERIFY_FAIL          DWORD = 0x40 //人脸加密码加指纹认证失败
	MINOR_FACE_AND_PW_AND_FP_VERIFY_TIMEOUT       DWORD = 0x41 //人脸加密码加指纹认证超时
	MINOR_FACE_CARD_AND_FP_VERIFY_PASS            DWORD = 0x42 //人脸加刷卡加指纹认证通过
	MINOR_FACE_CARD_AND_FP_VERIFY_FAIL            DWORD = 0x43 //人脸加刷卡加指纹认证失败
	MINOR_FACE_CARD_AND_FP_VERIFY_TIMEOUT         DWORD = 0x44 //人脸加刷卡加指纹认证超时
	MINOR_EMPLOYEENO_AND_FP_VERIFY_PASS           DWORD = 0x45 //工号加指纹认证通过
	MINOR_EMPLOYEENO_AND_FP_VERIFY_FAIL           DWORD = 0x46 //工号加指纹认证失败
	MINOR_EMPLOYEENO_AND_FP_VERIFY_TIMEOUT        DWORD = 0x47 //工号加指纹认证超时
	MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_PASS    DWORD = 0x48 //工号加指纹加密码认证通过
	MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_FAIL    DWORD = 0x49 //工号加指纹加密码认证失败
	MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_TIMEOUT DWORD = 0x4a //工号加指纹加密码认证超时
	MINOR_FACE_VERIFY_PASS                        DWORD = 0x4b //人脸认证通过
	MINOR_FACE_VERIFY_FAIL                        DWORD = 0x4c //人脸认证失败
	MINOR_EMPLOYEENO_AND_FACE_VERIFY_PASS         DWORD = 0x4d //工号加人脸认证通过
	MINOR_EMPLOYEENO_AND_FACE_VERIFY_FAIL         DWORD = 0x4e //工号加人脸认证失败
	MINOR_EMPLOYEENO_AND_FACE_VERIFY_TIMEOUT      DWORD = 0x4f //工号加人脸认证超时
	MINOR_FACE_RECOGNIZE_FAIL                     DWORD = 0x50 //人脸识别失败
	MINOR_FIRSTCARD_AUTHORIZE_BEGIN               DWORD = 0x51 //首卡授权开始
	MINOR_FIRSTCARD_AUTHORIZE_END                 DWORD = 0x52 //首卡授权结束
	MINOR_DOORLOCK_INPUT_SHORT_CIRCUIT            DWORD = 0x53 //门锁输入短路报警
	MINOR_DOORLOCK_INPUT_BROKEN_CIRCUIT           DWORD = 0x54 //门锁输入断路报警
	MINOR_DOORLOCK_INPUT_EXCEPTION                DWORD = 0x55 //门锁输入异常报警
	MINOR_DOORCONTACT_INPUT_SHORT_CIRCUIT         DWORD = 0x56 //门磁输入短路报警
	MINOR_DOORCONTACT_INPUT_BROKEN_CIRCUIT        DWORD = 0x57 //门磁输入断路报警
	MINOR_DOORCONTACT_INPUT_EXCEPTION             DWORD = 0x58 //门磁输入异常报警
	MINOR_OPENBUTTON_INPUT_SHORT_CIRCUIT          DWORD = 0x59 //开门按钮输入短路报警
	MINOR_OPENBUTTON_INPUT_BROKEN_CIRCUIT         DWORD = 0x5a //开门按钮输入断路报警
	MINOR_OPENBUTTON_INPUT_EXCEPTION              DWORD = 0x5b //开门按钮输入异常报警
	MINOR_DOORLOCK_OPEN_EXCEPTION                 DWORD = 0x5c //门锁异常打开
	MINOR_DOORLOCK_OPEN_TIMEOUT                   DWORD = 0x5d //门锁打开超时
	MINOR_FIRSTCARD_OPEN_WITHOUT_AUTHORIZE        DWORD = 0x5e //首卡未授权开门失败
	MINOR_CALL_LADDER_RELAY_BREAK                 DWORD = 0x5f //呼梯继电器断开
	MINOR_CALL_LADDER_RELAY_CLOSE                 DWORD = 0x60 //呼梯继电器闭合
	MINOR_AUTO_KEY_RELAY_BREAK                    DWORD = 0x61 //自动按键继电器断开
	MINOR_AUTO_KEY_RELAY_CLOSE                    DWORD = 0x62 //自动按键继电器闭合
	MINOR_KEY_CONTROL_RELAY_BREAK                 DWORD = 0x63 //按键梯控继电器断开
	MINOR_KEY_CONTROL_RELAY_CLOSE                 DWORD = 0x64 // 按键梯控继电器闭合
	MINOR_EMPLOYEENO_AND_PW_PASS                  DWORD = 0x65 //工号加密码认证通过
	MINOR_EMPLOYEENO_AND_PW_FAIL                  DWORD = 0x66 //工号加密码认证失败
	MINOR_EMPLOYEENO_AND_PW_TIMEOUT               DWORD = 0x67 //工号加密码认证超时
	MINOR_HUMAN_DETECT_FAIL                       DWORD = 0x68 //真人检测失败
	MINOR_PEOPLE_AND_ID_CARD_COMPARE_PASS         DWORD = 0x69 //人证比对通过
	MINOR_PEOPLE_AND_ID_CARD_COMPARE_FAIL         DWORD = 0x70 //人证比对失败
	MINOR_CERTIFICATE_BLOCKLIST                   DWORD = 0x71 //黑名单事件
	MINOR_LEGAL_MESSAGE                           DWORD = 0x72 //合法短信
	MINOR_ILLEGAL_MESSAGE                         DWORD = 0x73 //非法短信
	MINOR_MAC_DETECT                              DWORD = 0x74 //MAC侦测
	MINOR_DOOR_OPEN_OR_DORMANT_FAIL               DWORD = 0x75 //门状态常闭或休眠状态认证失败
	MINOR_AUTH_PLAN_DORMANT_FAIL                  DWORD = 0x76 //认证计划休眠模式认证失败
	MINOR_CARD_ENCRYPT_VERIFY_FAIL                DWORD = 0x77 //卡加密校验失败
	MINOR_SUBMARINEBACK_REPLY_FAIL                DWORD = 0x78 //反潜回服务器应答失败
	MINOR_TRAILING                                DWORD = 0x85 //尾随通行
	MINOR_REVERSE_ACCESS                          DWORD = 0x86 //反向闯入
	MINOR_FORCE_ACCESS                            DWORD = 0x87 //外力冲撞
	MINOR_CLIMBING_OVER_GATE                      DWORD = 0x88 //翻越
	MINOR_PASSING_TIMEOUT                         DWORD = 0x89 //通行超时
	MINOR_INTRUSION_ALARM                         DWORD = 0x8a //误闯报警
	MINOR_FREE_GATE_PASS_NOT_AUTH                 DWORD = 0x8b //闸机自由通行时未认证通过
	MINOR_DROP_ARM_BLOCK                          DWORD = 0x8c //摆臂被阻挡
	MINOR_DROP_ARM_BLOCK_RESUME                   DWORD = 0x8d //摆臂阻挡消除
	MINOR_LOCAL_FACE_MODELING_FAIL                DWORD = 0x8e //设备升级本地人脸建模失败
	MINOR_STAY_EVENT                              DWORD = 0x8f //逗留事件
	MINOR_PASSWORD_MISMATCH                       DWORD = 0x97 //密码不匹配
	MINOR_EMPLOYEE_NO_NOT_EXIST                   DWORD = 0x98 //工号不存在
	MINOR_COMBINED_VERIFY_PASS                    DWORD = 0x99 //组合认证通过
	MINOR_COMBINED_VERIFY_TIMEOUT                 DWORD = 0x9a //组合认证超时
	MINOR_VERIFY_MODE_MISMATCH                    DWORD = 0x9b //认证方式不匹配

)

// 使用cgo -godefs生成对齐正确的结构体
//--------------------登录返回信息--------------------------------//

// NET_DVR_ALARMER 报警设备信息结构体。
type NET_DVR_ALARMER struct {
	ST_byUserIDValid     BYTE      /* userid是否有效 0-无效，1-有效 */
	ST_bySerialValid     BYTE      /* 序列号是否有效 0-无效，1-有效 */
	ST_byVersionValid    BYTE      /* 序列号是否有效 0-无效，1-有效 */
	ST_byDeviceNameValid BYTE      /* 设备名字是否有效 0-无效，1-有效 */
	ST_byMacAddrValid    BYTE      /* MAC地址是否有效 0-无效，1-有效 */
	ST_byLinkPortValid   BYTE      /* login端口是否有效 0-无效，1-有效 */
	ST_byDeviceIPValid   BYTE      /* 设备IP是否有效 0-无效，1-有效 */
	ST_bySocketIPValid   BYTE      /* Socket ip是否有效 0-无效，1-有效 */
	ST_lUserID           LONG      /* NET_DVR_Login()返回值, 布防时有效 */
	ST_sSerialNumber     [48]BYTE  /* 序列号 */
	ST_dwDeviceVersion   DWORD     /* 版本信息 高16位表示主版本，低16位表示次版本*/
	ST_sDeviceName       [32]BYTE  /* 设备名字 */
	ST_byMacAddr         [6]BYTE   /* MAC地址 */
	ST_wLinkPort         WORD      /* link port */
	ST_sDeviceIP         [128]BYTE /* IP地址 */
	ST_byIpProtocol      BYTE      /* 报警主动上传时的socket IP地址 */
	ST_byRes1            [2]BYTE   /* Ip协议 0-IPV4, 1-IPV6 */
	ST_bJSONBroken       BYTE      //JSON断网续传标志。0：不续传；1：续传
	ST_wSocketPort       WORD
	ST_byRes2            [6]BYTE
}

//--------------------回调函数返回门禁设备告警信息-------------------//

// NET_DVR_ACS_ALARM_INFO 门禁主机报警信息结构体。
type NET_DVR_ACS_ALARM_INFO struct {
	ST_dwSize                  DWORD                  //结构体大小
	ST_dwMajor                 DWORD                  //报警主类型，具体定义见“Remarks”说明
	ST_dwMinor                 DWORD                  //报警次类型，次类型含义根据主类型不同而不同，具体定义见“Remarks”说明
	ST_struTime                NET_DVR_TIME           //报警时间
	ST_sNetUser                [16]BYTE               //网络操作的用户名
	ST_struRemoteHostAddr      NET_DVR_IPADDR         //远程主机地址
	ST_struAcsEventInfo        NET_DVR_ACS_EVENT_INFO //报警信息详细参数
	ST_dwPicDataLen            DWORD                  //图片数据大小，不为0是表示后面带数据
	ST_pPicData                unsafe.Pointer         //图片数据缓冲区
	ST_wInductiveEventType     WORD                   //归纳事件类型，0-无效，客户端判断该值为非0值后，报警类型通过归纳事件类型区分，否则通过原有报警主次类型（dwMajor、dwMinor）区分
	ST_byPicTransType          BYTE                   //图片数据传输方式: 0-二进制；1-url
	ST_byRes1                  BYTE                   //保留，置为0
	ST_dwIOTChannelNo          DWORD                  //IOT通道号
	ST_pAcsEventInfoExtend     unsafe.Pointer         //byAcsEventInfoExtend为1时，表示指向一个NET_DVR_ACS_EVENT_INFO_EXTEND结构体
	ST_byAcsEventInfoExtend    BYTE                   //pAcsEventInfoExtend是否有效：0-无效，1-有效
	ST_byTimeType              BYTE                   //时间类型：0-设备本地时间，1-UTC时间（struTime的时间）
	ST_byRes2                  BYTE                   //保留，置为0
	ST_byAcsEventInfoExtendV20 BYTE                   //pAcsEventInfoExtendV20是否有效：0-无效，1-有效
	ST_pAcsEventInfoExtendV20  unsafe.Pointer         //byAcsEventInfoExtendV20为1时，表示指向一个NET_DVR_ACS_EVENT_INFO_EXTEND_V20结构体
	ST_byRes                   [4]BYTE                //保留，置为0
}

// NET_DVR_TIME 时间参数结构体
type NET_DVR_TIME struct {
	ST_dwYear   DWORD //年
	ST_dwMonth  DWORD //月
	ST_dwDay    DWORD //日
	ST_dwHour   DWORD //时
	ST_dwMinute DWORD //分
	ST_dwSecond DWORD //秒
}

//NET_DVR_IPADDR IP地址结构体
type NET_DVR_IPADDR struct {
	ST_sIpV4 [16]BYTE  //IPv4地址
	ST_sIpV6 [128]BYTE //IPv6地址
}

//NET_DVR_ACS_EVENT_INFO 门禁主机事件信息
type NET_DVR_ACS_EVENT_INFO struct {
	ST_dwSize                         DWORD    //结构体大小
	ST_byCardNo                       [32]BYTE //卡号
	ST_byCardType                     BYTE     //卡类型：1- 普通卡，2- 残障人士卡，3- 黑名单卡，4- 巡更卡，5- 胁迫卡，6- 超级卡，7- 来宾卡，8- 解除卡，为0表示无效
	ST_byAllowListNo                  BYTE     //白名单单号，取值范围：1~8，0表示无效
	ST_byReportChannel                BYTE     //报告上传通道：1- 布防上传，2- 中心组1上传，3- 中心组2上传，0表示无效
	ST_byCardReaderKind               BYTE     //读卡器类型：0- 无效，1- IC读卡器，2- 身份证读卡器，3- 二维码读卡器，4- 指纹头
	ST_dwCardReaderNo                 DWORD    //读卡器编号，为0表示无效
	ST_dwDoorNo                       DWORD    // 门编号（或者梯控的楼层编号），为0表示无效（当接的设备为人员通道设备时，门1为进方向，门2为出方向）
	ST_dwVerifyNo                     DWORD    //多重卡认证序号，为0表示无效
	ST_dwAlarmInNo                    DWORD    //报警输入号，为0表示无效
	ST_dwAlarmOutNo                   DWORD    //报警输出号，为0表示无效
	ST_dwCaseSensorNo                 DWORD    //事件触发器编号
	ST_dwRs485No                      DWORD    //RS485通道号，为0表示无效
	ST_dwMultiCardGroupNo             DWORD    //群组编号
	ST_wAccessChanne                  WORD     //人员通道号
	ST_byDeviceNo                     BYTE     //设备编号，为0表示无效
	ST_byDistractControlNo            BYTE     //分控器编号，为0表示无效
	ST_dwEmployeeNo                   DWORD    //工号，为0无效
	ST_wLocalControllerID             WORD     //就地控制器编号，0-门禁主机，1-255代表就地控制器
	ST_byInternetAccess               BYTE     //网口ID：（1-上行网口1,2-上行网口2,3-下行网口1）
	ST_byType                         BYTE     //防区类型，0:即时防区,1-24小时防区,2-延时防区,3-内部防区,4-钥匙防区,5-火警防区,6-周界防区,7-24小时无声防区,8-24小时辅助防区,9-24小时震动防区,10-门禁紧急开门防区,11-门禁紧急关门防区，0xff-无
	ST_byMACAddr                      [6]BYTE  //物理地址，为0无效
	ST_bySwipeCardType                BYTE     //刷卡类型，0-无效，1-二维码
	ST_byMask                         BYTE     //是否带口罩：0-保留，1-未知，2-不戴口罩，3-戴口罩
	ST_dwSerialNo                     DWORD    //事件流水号，为0无效
	ST_byChannelControllerID          BYTE     //通道控制器ID，为0无效，1-主通道控制器，2-从通道控制器
	ST_byChannelControllerLampID      BYTE     //通道控制器灯板ID，为0无效（有效范围1-255）
	ST_byChannelControllerIRAdaptorID BYTE     //通道控制器红外转接板ID，为0无效（有效范围1-255）
	ST_byChannelControllerIREmitterID BYTE     //通道控制器红外对射ID，为0无效（有效范围1-255）
	ST_byRes                          [4]BYTE  //保留，置为0
}

// NET_DVR_ACS_EVENT_INFO_EXTEND 门禁主机扩展事件信息。
type NET_DVR_ACS_EVENT_INFO_EXTEND struct {
	ST_dwFrontSerialNo       DWORD    //事件流水号，为0无效（若该字段为0，平台根据dwSerialNo判断是否丢失事件；若该字段不为0，平台根据该字段和dwSerialNo字段共同判断是否丢失事件）（主要用于解决报警订阅后导致dwSerialNo不连续的情况
	ST_byUserType            BYTE     //人员类型：0-无效，1-普通人（主人），2-来宾（访客），3-黑名单人，4-管理员
	ST_byCurrentVerifyMode   BYTE     //读卡器当前验证方式：0-无效，1-休眠，2-刷卡+密码，3-刷卡，4-刷卡或密码，5-指纹，6-指纹+密码，7-指纹或刷卡，8-指纹+刷卡，9-指纹+刷卡+密码，10-人脸或指纹或刷卡或密码，11-人脸+指纹，12-人脸+密码，13-人脸+刷卡，14-人脸，15-工号+密码，16-指纹或密码，17-工号+指纹，18-工号+指纹+密码，19-人脸+指纹+刷卡，20-人脸+密码+指纹，21-工号+人脸，22-人脸或人脸+刷卡，23-指纹或人脸，24-刷卡或人脸或密码
	ST_byCurrentEvent        BYTE     //是否为实时事件：0-无效，1-是（实时事件），2-否（离线事件）
	ST_byPurePwdVerifyEnable BYTE     //设备是否支持纯密码认证：0-不支持，1-支持
	ST_byEmployeeNo          [32]BYTE //工号（人员ID）（对于设备来说，如果使用了工号（人员ID）字段，byEmployeeNo一定要传递，如果byEmployeeNo可转换为dwEmployeeNo，那么该字段也要传递；对于上层平台或客户端来说，优先解析byEmployeeNo字段，如该字段为空，再考虑解析dwEmployeeNo字段）
	ST_byAttendanceStatus    BYTE     //考勤状态：0-未定义,1-上班，2-下班，3-开始休息，4-结束休息，5-开始加班，6-结束加班
	ST_byStatusValue         BYTE     //考勤状态值
	ST_byRes2                [2]BYTE  //保留，置为0
	ST_byUUID                [36]BYTE //UUID（该字段仅在对接萤石平台过程中才会使用）
	ST_byDeviceName          [64]BYTE //设备序列号
	ST_byRes                 [24]BYTE //保留，置为0
}

//-----------返回门禁状态结构体-------------------------------//

// NET_DVR_ACS_WORK_STATUS  门禁主机工作状态结构体
type NET_DVR_ACS_WORK_STATUS struct {
	ST_dwSize                          DWORD     //结构体大小
	ST_byDoorLockStatus                [32]BYTE  //门锁状态：0- 关，1- 开
	ST_byDoorStatus                    [32]BYTE  //门状态：1- 休眠，2- 常开状态，3- 常闭状态，4- 普通状态
	ST_byMagneticStatus                [32]BYTE  //门磁状态：0- 闭合，1- 开启
	ST_byCaseStatus                    [8]BYTE   //事件报警输入状态：0- 无输入，1- 有输入
	ST_wBatteryVoltage                 WORD      //蓄电池电压值，实际值乘10，单位：伏特
	ST_byBatteryLowVoltage             BYTE      //蓄电池是否处于低压状态：0- 否，1- 是
	ST_byPowerSupplyStatus             BYTE      //设备供电状态：1- 交流电供电，2- 蓄电池供电
	ST_byMultiDoorInterlockStatus      BYTE      //多门互锁状态：0- 关闭，1- 开启
	ST_byAntiSneakStatus               BYTE      //反潜回状态：0-关闭，1-开启
	ST_byHostAntiDismantleStatus       BYTE      //主机防拆状态：0- 关闭，1- 开启
	ST_byIndicatorLightStatus          BYTE      //指示灯状态
	ST_byCardReaderOnlineStatus        [64]BYTE  //读卡器在线状态：0- 不在线，1- 在线
	ST_byCardReaderAntiDismantleStatus [64]BYTE  //读卡器防拆状态：0- 关闭，1- 开启
	ST_byCardReaderVerifyMode          [64]BYTE  //读卡器当前验证方式：0- 无效，1- 休眠，2- 刷卡+密码，3- 刷卡，4- 刷卡或密码，5- 指纹，6- 指纹加密码，7- 指纹或刷卡，8- 指纹加刷卡，9- 指纹加刷卡加密码
	ST_bySetupAlarmStatus              [512]BYTE //报警输入口布防状态：0- 对应报警输入口处于撤防状态，1- 对应报警输入口处于布防状态
	ST_byAlarmInStatus                 [512]BYTE //报警输入口报警状态：0- 对应报警输入口当前无报警，1- 对应报警输入口当前有报警
	ST_byAlarmOutStatus                [512]BYTE //报警输出口状态：0- 对应报警输出口无报警，1- 对应报警输出口有报警
	ST_dwCardNum                       DWORD     //已添加的卡数量
	ST_byRes2                          [32]BYTE  //保留，置为0
}

//-----------获取卡参数结构体-------------------------------//

// NET_DVR_CARD_COND 获取卡参数配置条件结构体。old
type NET_DVR_CARD_COND struct {
	ST_dwSize    DWORD
	ST_dwCardNum DWORD
	ST_byRes     [64]BYTE
}

// NET_DVR_CARD_RECORD 返回卡信息结构体
type NET_DVR_CARD_RECORD struct {
	dwSize     DWORD    //结构体大小
	byCardNo   [32]BYTE //卡号
	byCardType BYTE     /*卡类型：1- 普通卡（默认），2- 残障人士卡，3- 禁止名单卡，4- 巡
	更卡，5- 胁迫卡，6- 超级卡，7- 来宾卡，8- 解除卡，9- 员工卡，10- 应急卡，11- 应急管理卡（用
	于授权临时卡权限，本身不能开门），默认普通卡*/
	byLeaderCard BYTE /*是否为首卡：1- 是，0- 否*/
	byUserType   BYTE //用户类型：0 – 普通用户 1- 管理员用户
	byRes1       BYTE
	byDoorRight  [256]BYTE /*门权限（梯控的楼层权限、锁权限），按字节
	表示，1-为有权限，0-为无权限，从低位到高位依次表示对门（或者梯控楼层、锁）1-N 是否有权限	*/
	struValid     NET_DVR_VALID_PERIOD_CFG /*有效期参数（有效时间跨度为 1970 年 1 月 1 日 0 点 0 分 0 秒~2037 年 12 月 31 日 23 点 59 分 59 秒）*/
	byBelongGroup [128]BYTE                /*所属群组，按字节表示，1-属于，0-不
	属于，从低位到高位表示是否从属群组 1~N*/
	byCardPassword [8]BYTE   //卡密码
	wCardRightPlan [256]WORD /*卡权限计划，取值为计划模板编号，同个
	门（锁）不同计划模板采用权限或的方式处理 */
	dwMaxSwipeTimes DWORD    //最大刷卡次数，0 为无次数限制
	dwSwipeTimes    DWORD    //已刷卡次数
	dwEmployeeNo    DWORD    //工号（用户 ID），1~99999999，不能以 0 开头且不能重复
	byName          [32]BYTE //姓名
	dwCardRight     DWORD    //卡权限
	byRes           [256]BYTE
}

//NET_DVR_VALID_PERIOD_CFG 有效期参数结构体
type NET_DVR_VALID_PERIOD_CFG struct {
	byEnable         BYTE            //使能有效期，0-不使能，1使能
	byBeginTimeFlag  BYTE            //是否限制起始时间的标志，0-不限制，1-限制
	byEnableTimeFlag BYTE            //是否限制终止时间的标志，0-不限制，1-限制
	byTimeDurationNo BYTE            //有效期索引,从0开始（时间段通过SDK设置给锁，后续在制卡时，只需要传递有效期索引即可，以减少数据量）
	struBeginTime    NET_DVR_TIME_EX //有效期起始时间
	struEndTime      NET_DVR_TIME_EX //有效期结束时间
	byTimeType       BYTE            //时间类型：0-设备本地时间（默认），1-UTC时间（对于struBeginTime，struEndTime字段有效）
	byRes2           [31]BYTE
}

// NET_DVR_TIME_EX 时间参数结构体扩展
type NET_DVR_TIME_EX struct {
	ST_wYear    WORD //年
	ST_byMonth  BYTE //月
	ST_byDay    BYTE //日
	ST_byHour   BYTE //时
	ST_byMinute BYTE //分
	ST_bySecond BYTE //秒
	byRes       BYTE
}

// NET_DVR_ALARMHOST_MAIN_STATUS_V51 ----------------------获取报警主机防区状态参数------------------------//
type NET_DVR_ALARMHOST_MAIN_STATUS_V51 struct {
	ST_dwSize                  DWORD
	ST_bySetupAlarmStatus      [512]BYTE //防区布防状态，(最大支持512个防区查询)，0xff-无效，0-对应防区处于撤防状态，1-对应防区处于布防状态，2-对应防区处于布防中
	ST_byAlarmInStatus         [512]BYTE //防区报警状态（触发状态），(最大支持512个防区查询)，0xff-无效，0-对应防区当前无报警，1-对应防区当前有报警
	ST_byAlarmOutStatus        [512]BYTE //触发器状态，(最大支持512个触发器查询)，0xff-无效，0-对应触发器无报警，1-对应触发器有报警，2-未关联，3-离线，4-心跳异常
	ST_byBypassStatus          [512]BYTE //防区旁路状态，数组下标表示0对应防区1，0xff-无效，0-表示防区没有旁路 1-表示防区旁路
	ST_bySubSystemGuardStatus  [32]BYTE  //子系统布防状态，0xff-无效，0-对应子系统处于撤防状态，1-对应子系统处于布防状态，2-对应子系统处于布防中
	ST_byAlarmInFaultStatus    [512]BYTE //防区故障状态，0xff-无效，0-对应防区处于正常状态，1-对应防区处于故障状态
	ST_byAlarmInMemoryStatus   [512]BYTE //防区报警记忆状态（报警状态）， 0xff-无效，0-对应防区当前无报警，1-对应防区当前有报警
	ST_byAlarmInTamperStatus   [512]BYTE //防区防拆状态，0xff-无效，0-对应防区无报警，1-对应防区有报警
	ST_byEnableSubSystem       [32]BYTE  //子系统启用状态，0-无效，1-对应子系统未启用，2-对应子系统启用
	ST_bySubSystemGuardType    [32]BYTE  //子系统布防类型，0-无效，1-外出布防，2-即时布防，3-在家布防
	ST_bySubSystemAlarm        [32]BYTE  //子系统报警状态，0-无效，1-正常，2-报警
	ST_byAlarmOutCharge        [512]BYTE //触发器电量状态，(最大支持512个触发器查询)，0-无效，1-正常，2-电量低
	ST_byAlarmOutTamperStatus  [512]BYTE //触发器防拆状态，(最大支持512个触发器查询)，0-无效，1-防拆，2-无防拆
	ST_byAlarmInShieldedStatus [512]BYTE //防区屏蔽状态，0-无效，1-屏蔽，2-非屏蔽
	ST_byAlarmOutLinkage       [512]BYTE //触发器联动事件类型，(最大支持512个触发器查询)，0-无效，1-报警，2-布防，3-撤防，4-手动控制
	ST_byRes                   [512]BYTE //保留字节
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
