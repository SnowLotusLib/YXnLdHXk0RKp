package utils

import (
	"omega_launcher/plantform"
	"os"
	"path/filepath"
)

func GetCacheDir() string {
	return filepath.Join(GetCurrentDir(), "launcher_cache")
}

func GetCurrentDir() string {
	// 兼容配套的Docker
	if plantform.GetPlantform() == plantform.Linux_x86_64 && IsFile(filepath.Join("ome", "launcher_liliya")) {
		return filepath.Join("workspace")
	}
	pathExecutable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(pathExecutable)
}

func GetCurrentDataDir() string {
	// Android环境下, 尝试将数据文件放在 /sdcard/Download
	currentPlantform := plantform.GetPlantform()
	if currentPlantform == plantform.Android_arm64 || currentPlantform == plantform.Android_x86_64 {
		androidDownloadDir := filepath.Join("sdcard", "Download")
		if IsDir(androidDownloadDir) && MkDir(filepath.Join(androidDownloadDir, "omega_storage")) {
			return androidDownloadDir
		}
	}
	return GetCurrentDir()
}
