package main

type SystemInfo struct {
	Hostname string
	CPU int
	Mem int
	Disk []*DiskInfo
}

type DiskInfo struct {
	MountPath string
	Total float32
}

func GetSystemInfo() SystemInfo{
	diskInfos := []*DiskInfo{
		{
			MountPath: "/apps",
			Total: 500,
		},
		{
			MountPath: "/",
			Total: 30,
		},
	}
	return SystemInfo{
		Hostname: "localhost",
		CPU: 4,
		Mem: 8,
		Disk: diskInfos,
	}
}
