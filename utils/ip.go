package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// GetCidrIpRange 计算得到CIDR地址范围中最小和最大IP
func GetCidrIpRange(cidr string) (string, string) {
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	seg1MinIp, seg1MaxIp := getIpSeg1Range(ipSegs, maskLen)
	seg2MinIp, seg2MaxIp := getIpSeg2Range(ipSegs, maskLen)
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs, maskLen)

	return strconv.Itoa(seg1MinIp) + "." + strconv.Itoa(seg2MinIp) + "." + strconv.Itoa(seg3MinIp) + "." + strconv.Itoa(seg4MinIp),
		strconv.Itoa(seg1MaxIp) + "." + strconv.Itoa(seg2MaxIp) + "." + strconv.Itoa(seg3MaxIp) + "." + strconv.Itoa(seg4MaxIp)
}

// GetCidrHostNum 计算得到CIDR地址范围内可拥有的主机数量
func GetCidrHostNum(maskLen int) uint {
	cidrIpNum := uint(0)
	var i uint = uint(32 - maskLen - 1)
	for ; i >= 1; i-- {
		cidrIpNum += 1 << i
	}
	return cidrIpNum
}

// GetCidrIpMask 获取Cidr的掩码
func GetCidrIpMask(maskLen int) string {
	// ^uint32(0)二进制为32个比特1，通过向左位移，得到CIDR掩码的二进制
	cidrMask := ^uint32(0) << uint(32-maskLen)
	fmt.Println(fmt.Sprintf("%b \n", cidrMask))
	//计算CIDR掩码的四个片段，将想要得到的片段移动到内存最低8位后，将其强转为8位整型，从而得到
	cidrMaskSeg1 := uint8(cidrMask >> 24)
	cidrMaskSeg2 := uint8(cidrMask >> 16)
	cidrMaskSeg3 := uint8(cidrMask >> 8)
	cidrMaskSeg4 := uint8(cidrMask & uint32(255))

	return fmt.Sprint(cidrMaskSeg1) + "." + fmt.Sprint(cidrMaskSeg2) + "." + fmt.Sprint(cidrMaskSeg3) + "." + fmt.Sprint(cidrMaskSeg4)
}

//得到第一段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg1Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 8 {
		segIp, _ := strconv.Atoi(ipSegs[0])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[0])
	return getIpSegRange(uint8(ipSeg), uint8(8-maskLen))
}

//得到第二段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg2Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 16 {
		segIp, _ := strconv.Atoi(ipSegs[1])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[1])
	return getIpSegRange(uint8(ipSeg), uint8(16-maskLen))
}

//得到第三段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg3Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 24 {
		segIp, _ := strconv.Atoi(ipSegs[2])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[2])
	return getIpSegRange(uint8(ipSeg), uint8(24-maskLen))
}

//得到第四段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg4Range(ipSegs []string, maskLen int) (int, int) {
	ipSeg, _ := strconv.Atoi(ipSegs[3])
	segMinIp, segMaxIp := getIpSegRange(uint8(ipSeg), uint8(32-maskLen))
	return segMinIp, segMaxIp
}

//根据用户输入的基础IP地址和CIDR掩码计算一个IP片段的区间
func getIpSegRange(userSegIp, offset uint8) (int, int) {
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & userSegIp
	segMaxIp := userSegIp&(255<<offset) | ^(255 << offset)
	return int(segMinIp), int(segMaxIp)
}

func IsNetWorkOk(network string) bool {
	_, _, err := net.ParseCIDR(network)
	if err != nil {
		return false
	}
	return true
}

func IsIpOk(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	}
	return true
}

func IpToNum(ip string) uint32 {
	defer func() {
		if err := recover(); err != nil {
			return
		}

	}()
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return 0
	}
	var num uint32
	for i, octet := range octets {
		oct, err := strconv.Atoi(octet)
		if err != nil {
			return 0
		}
		if oct > 255 {
			return 0
		}
		num |= uint32(oct) << ((3 - i) * 8)
	}
	return num
}

func NumToIp(num uint32) string {
	return strconv.Itoa(int(num>>24)) + "." + strconv.Itoa(int(num>>16&0xFF)) + "." +
		strconv.Itoa(int(num>>8&0xFF)) + "." + strconv.Itoa(int(num&0xFF))
}
