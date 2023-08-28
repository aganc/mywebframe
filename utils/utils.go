package utils

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"airport/zlog"
)

/*
@Time : 2021/5/13 9:03 下午
@Author : pangshaoliang
@File : base
@Software: GoLand
*/

// GetNowTime 13位时间戳 单位ms
func GetNowTime() int64 {
	return time.Now().Unix() * 1000
}

func GetLimitAndOffset(page, size int) (limit int, offset int) {
	if page == 0 || size == 0 {
		limit = -1
		offset = -1
	} else {
		limit = size
		offset = (page - 1) * size
	}
	return
}

func BytesToSizeString(bytes int64) string {
	unitArray := []string{"KB", "MB", "GB", "TB", "PB"}

	size := float64(bytes) / 1024
	for _, unit := range unitArray {
		if size < 1024 {
			return fmt.Sprintf("%.2f%s", size, unit)
		}

		size = size / 1024
	}

	return fmt.Sprintf("%.2fEB", size)
}

func BytesToSizeStringV2(bytes int64) string {
	if bytes <= 0 {
		return ""
	}

	return BytesToSizeString(bytes)
}

func ExternalAddr(ctx context.Context) string {
	return fmt.Sprintf("%v", zlog.GetExternalAddress(ctx))
}

// ConvertInt2Bool 将int8转换为bool
func ConvertInt2Bool(i int8) bool {
	if i == 1 {
		return true
	}
	return false
}

// Float64PercentString 计算百分比的float64转为字符串
func Float64PercentString(f float64) string {
	fs := fmt.Sprint(f)
	fss := strings.Split(fs, ".")
	fz := fss[0]
	fx := ""
	if len(fss) == 2 {
		fx = fss[1]
		if len(fx) > 2 {
			fx = fx[:2]
		}
	}
	if len(fx) > 0 {
		fx = "." + fx
	}
	return fz + fx + "%"
}

func IsPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func ConvertStr2Int(s string) int {
	res, _ := strconv.ParseInt(s, 10, 64)
	return int(res)
}

func ConvertStr2Int32(s string) int32 {
	res, _ := strconv.ParseInt(s, 10, 32)
	return int32(res)
}

func ConvertStr2Int64(s string) (res int64) {
	res, _ = strconv.ParseInt(s, 10, 64)
	return
}

func ConvertStr2Uint8(s string) uint8 {
	res, _ := strconv.ParseUint(s, 10, 8)
	return uint8(res)
}

func ConvertStr2Uint32(s string) uint32 {
	res, _ := strconv.ParseUint(s, 10, 32)
	return uint32(res)
}

// GetPkgVersion pkgFilename 格式为 avrir_V1.3.1.5.tar.gz ，返回1.3.1.5
func GetPkgVersion(pkgFilename string) string {
	if len(pkgFilename) == 0 {
		return ""
	}
	tmpVersion1 := strings.Split(pkgFilename, "_V")
	if len(tmpVersion1) < 2 {
		return ""
	}
	tmpVersion2 := strings.Split(tmpVersion1[1], ".tar.gz")
	if len(tmpVersion2) < 2 {
		return ""
	}
	return tmpVersion2[0]
}

// CheckVersion 版本比较 pkgVersion 1.3.5.2 nodeVersion V3.6.1
func CheckVersion(pkgVersion, nodeVersion string) bool {
	curVersion := strings.Trim(nodeVersion, "V")
	curVersion = "1." + curVersion

	if pkgVersion == curVersion {
		return false
	}
	pkgVersionList := strings.Split(pkgVersion, ".")
	if len(pkgVersionList) != 4 {
		return false
	}
	nodeVersionList := strings.Split(curVersion, ".")
	if len(nodeVersionList) != 4 {
		return false
	}
	i := 0
	for i < 4 {
		vv1, _ := strconv.Atoi(pkgVersionList[i])
		vv2, _ := strconv.Atoi(nodeVersionList[i])
		if vv1 < vv2 {
			return false
		}
		i++
	}
	return true
}

// ConvertPkgVersion pkgVersion 1.3.5.2 ---> V3.5.2
func ConvertPkgVersion(pkgVersion string) string {
	pkgVersionList := strings.Split(pkgVersion, ".")
	if len(pkgVersionList) != 4 {
		return ""
	}
	tmpVersion := pkgVersionList[1:]
	resVersion := fmt.Sprintf("%v%v", "V", strings.Join(tmpVersion, "."))
	return resVersion
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		*files = append(*files, path)
		return nil
	}
}

// GetFileList 通过文件目录获取文件列表
func GetFileList(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, visit(&files))
	if err != nil {
		return files, err
	}
	return files, nil
}
