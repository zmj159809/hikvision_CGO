# 说明

此代码为海康门禁CGO版本的接口开发

### 更新记录

#### v 0.0.1

初步实现了登录 查询门状态 控制门 查询设备卡信息   查询防区状态 布防  撤防 功能

#### v 0.0.2

1. 增加了事件类型的映射。

2. 修改了获取状态的传入参数。

3. 增加了recover，防止程序异常退出。

#### v 0.0.3

优化目录结构，方便包引用

#### v 0.0.5

tag测试

#### v 0.0.6

添加license 文件

#### v 1.0.0

1. 优化部分的函数输入，使输入更加直观

2. 调整代码结构顺序

3. 部分内存泄漏bug修复

4. 增加函数注释

### 接口说明

##### **1、初始化**

```
func NetInit(sdkLog string,ifSdkLog bool) error
输入
sdkLog   sdk日志路径 
ifSdkLog   是否开启sdk日志
输出
error 错误信息
```

##### 2、登录

```
func NetLoginV40(deviceIp, username, password string) (int32, error) 
输入
deviceIp  设备ip地址  （默认端口8000不用添加，只填写ip即可）
username  用户名
password  密码
输出
int32 登录成功返回用户uid，登录失败返回-1和error
error 错误信息
```

##### 3、登出

```
func NetLogout(uid int32) error
输入
int32  NetLoginV40返回的用户uid
输出
error 错误信息
```

##### 4、释放SDK资源

```
func NetCleanup() 无参数
```

##### 5、获取门状态

```
func GetDoorStatus(uid int32, pBuf unsafe.Pointer) (err error)
输入
uid  NetLoginV40返回的用户uid
pBuf  NET_DVR_ACS_WORK_STATUS 结构体指针转换成unsafe.Pointer
输出
error 错误信息

示例  
var  status NET_DVR_ACS_WORK_STATUS
 err := GetDoorStatus(userID, unsafe.Pointer(&status))
 if err != nil {
	fmt.Println(err)
	} else {
		for k,v :=range stuParam.ST_byMagneticStatus[0:4]{
			fmt.Printf("门磁%d状态:%s |",k+1,getDoorStatus(int(v)))
		}
		fmt.Print("\n")
		for k,v :=range stuParam.ST_byDoorLockStatus[0:4]{
			fmt.Printf("门锁%d状态:%s |",k+1,getDoorStatus(int(v)))
		}
		fmt.Print("\n")
		for k,v :=range stuParam.ST_byDoorStatus[0:4]{
			fmt.Printf("门%d状态:%s |",k+1,getDoorStatus1(int(v)))
		}
			
//getDoorStatus 门磁/门锁  状态
func getDoorStatus(status int) string {
	switch status {
	case 0:
		return "关"
	case 1:
		return "开"
	default:
		return "未知状态"
	}
}
```

##### 6、设置回调函数

```
func NewObjectId(obj interface{}) ObjectId
输入
obj 接口  传入参数 要求实现方法 
{    
     Invoke(lCommand int, ip string, pAlarmInfo unsafe.Pointer, dwBufLen int) bool
     输入
     lCommand 报警信息类型
     ip  报警设备ip
     pAlarmInfo 报警信息结构体，根据报警类型判断
     dwBufLen 用户结构体参数，布防 SetDVRMessCallBack 时传入的dwUser（函数NewObjectId返回值）
     输出
     bool  //这里不建议返回false 返回false会造成设备重新布防撤防
}
func SetDVRMessCallBack(dwUser ObjectId) error
输入
dwUser 函数NewObjectId返回值
输出
error 错误信息
```

##### 7、布防撤防

```
//布防 在设置回调函数之后才能使用
func DoDefence(uid int32) (int32, error)
输入
int32  NetLoginV40返回的用户uid
输出
error 错误信息

//撤防 
func CloseDefence(dId int32) 
输入
int32  DoDefence返回的用户dId
```

##### 8、获取设备支持的门禁卡信息

```
func GetCardInfo(uid int32) (cardInfo map[string]string, err error) 
输入
int32  NetLoginV40返回的用户uid
输出
cardInfo 返回的卡信息map   cardNO  to  Name
error 错误信息
```

##### 9、获取报警主机状态（暂无使用）

```
func GetAlarmHostMainStatus(uid int32, status unsafe.Pointer) (err error)
输入
int32  NetLoginV40返回的用户uid
status NET_DVR_ALARMHOST_MAIN_STATUS_V51 结构体指针
输出
error 错误信息
```

##### 10、获取报警事件类型

```
// GetMajorString 返回主类型字符串
func (d DWORD) GetMajorString() string
接收者  报警主类型
输出    事件中文名称

// GetMinorString 获取告警次类型
func (d DWORD) GetMinorString(major DWORD) string
接收者  报警次类型
输入 告警主类型
输出    事件中文名称
```

