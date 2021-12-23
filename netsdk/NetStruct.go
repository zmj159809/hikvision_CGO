package netsdk

import "unsafe"

type LONG = int32
type LLONG = int64
type DWORD = uint32
type WORD = uint16
type BYTE = byte
type BOOL = int32

const (
	COMM_ALARM_ACS = 0x5002 //门禁主机报警信息
)

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

//----------------------获取报警主机防区状态参数------------------------//
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
