package ip

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const rawDir = "./data/raw/"
const dataFile = "./data/db/ipv4.txt"

// 拉取原始数据
func Pull() {
	dir := getRawDir()
	urlArr := map[string]string{
		"delegated-apnic-latest":         "https://ftp.apnic.net/stats/apnic/delegated-apnic-latest",
		"delegated-arin-extended-latest": "https://ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest",
		"delegated-afrinic-latest":       "https://ftp.afrinic.net/pub/stats/afrinic/delegated-afrinic-latest",
		"delegated-lacnic-latest":        "https://ftp.lacnic.net/pub/stats/lacnic/delegated-lacnic-latest",
		"delegated-ripencc-latest":       "https://ftp.ripe.net/ripe/stats/delegated-ripencc-latest",
	}
	for filename, url := range urlArr {
		log.Println("[Pull]:", url)
		log.Println("[File]:", filename)

		saveFilename := dir + "/" + filename + ".txt"
		removeFile(saveFilename)
		readToRemoteCtx(saveFilename, url)
	}
	log.Println("[End]:拉取结束")
}

// 生成ipv4数据并保存在本地文件
func Create() {
	removeFile(dataFile)

	rawFileArr := []string{
		rawDir + "latest/delegated-afrinic-latest.txt",
		rawDir + "latest/delegated-apnic-latest.txt",
		rawDir + "latest/delegated-arin-extended-latest.txt",
		rawDir + "latest/delegated-lacnic-latest.txt",
		rawDir + "latest/delegated-ripencc-latest.txt",
	}
	for _, file := range rawFileArr {
		fmt.Println(file)
		f, err := os.Open(file)
		if err != nil {
			panic("读取文件异常...")
		}
		buf := bufio.NewReader(f)
		for {
			line, errR := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			if errR != nil {
				if errR == io.EOF {
					break
				}
			}

			// 过滤注释和非ipv4类型的数据
			if strings.Index(line, "#") == 0 || strings.Index(line, "ipv4") == -1 {
				continue
			}

			var organization string
			var country string
			var ipType string
			var ip string
			var length int
			var date string
			var status string
			infoArr := strings.Split(line, "|")
			if len(infoArr) == 0 {
				continue
			}
			if len(infoArr) >= 1 {
				organization = infoArr[0]
			}
			if len(infoArr) >= 2 {
				country = infoArr[1]
			}
			if len(infoArr) >= 3 {
				ipType = infoArr[2]
			}
			if len(infoArr) >= 4 {
				ip = infoArr[3]
			}
			if len(infoArr) >= 5 {
				length, _ = strconv.Atoi(infoArr[4])
			}
			if len(infoArr) >= 6 {
				date = infoArr[5]
			}
			if len(infoArr) >= 7 {
				status = infoArr[6]
			}
			if ipType != "ipv4" || (status != "assigned" && status != "allocated") {
				continue
			}
			fmt.Println(organization, country, ipType, ip, length, date, status)

			ipStart := Ip2long(ip)
			ipEnd := int(ipStart) + length - 1
			row := fmt.Sprintf("%s,%d,%d,%s,%s", country, ipStart, ipEnd, ip, Long2ip(ipStart))
			row = row + "\n"
			fmt.Println(row)
		}
	}
}

// ip转long
func Ip2long(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// long转ip
func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

// 获取原始数据保存路径
func getRawDir() string {
	folderPath := filepath.Join(rawDir, time.Now().Format("20060102"))
	mkdir(folderPath)
	return folderPath
}

// 创建文件目录
func mkdir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0777)
		_ = os.Chmod(path, 0777)
	}
}

// 移除文件
func removeFile(path string) {
	_, err := os.Lstat(path)
	if !os.IsNotExist(err) {
		errR := os.Remove(path)
		if errR != nil {
			panic("删除文件错误:" + err.Error())
		}
	}
}

// 读取远程地址内容
func readToRemoteCtx(saveFilename string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic("读取远程地址内容错误:" + err.Error())
	}
	defer resp.Body.Close()
	reader := bufio.NewReaderSize(resp.Body, 1024*32)
	for {
		b, errR := reader.ReadBytes('\n') // 按照行读取，遇到\n结束读取
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		lineData := string(b)
		//lineData := strings.TrimSuffix(strings.TrimSuffix(string(b), "\n"), "\r")
		fmt.Println(" -[Line]", string(b))
		_ = writeToFile(saveFilename, lineData)
	}
}

// 写入文件
func writeToFile(filename string, content string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic("文件创建错误:" + err.Error())
	} else {
		n, _ := f.Seek(0, 2)
		_, err = f.WriteAt([]byte(content), n)
		fmt.Println("Write succeeded")
		defer f.Close()
	}
	return err
}
