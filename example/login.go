package main

import t "github.com/zmj159809/hikvision_CGO/include"

func login() {
	t.NewNET_DVR_ALARMER() ///
	loginfo := t.NewNET_DVR_USER_LOGIN_INFO()
	deviceinfo := t.NewNET_DVR_DEVICEINFO_V40()
	loginfo.SetSUserName(*username)
	loginfo.SetSPassword(*password)
	loginfo.SetSDeviceAddress(*ip)
	port := 8000

	loginfo.SetWPort(t.SwigcptrWORD(&port))
	uid := t.NET_DVR_Login_V40(loginfo, deviceinfo)
}
