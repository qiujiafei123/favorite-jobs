package utils

import (
	"bufio"
	"errors"
	"favorite-jobs/log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.ZapLog.Errorw(msg, "错误原因", err)
	}
}

func CalculateSalaryRange(salary string) (string, string, error){
	salarys := strings.Split(salary, "-")
	if len(salarys) != 2 {
		return "", "", errors.New("薪资区间计算失败")
	}
	return salarys[0], salarys[1], nil
}


// 解析职位详情页的url 找到职位的id
func GetLagouJobIdByPath(path string) (int64, error) {
	paths := strings.Split(path, "/")
	if len(paths) == 0 {
		return 0, errors.New("通过职位详情没有解析到html")
	}
	// 获取最后一个元素并用 . 拆分
	jobId := strings.Split(paths[len(paths)-1], ".")
	if len(jobId) == 0 {
		return 0, errors.New("通过html没有截取到id")
	}
	id, err := strconv.ParseInt(jobId[0], 10, 64)
	return id, err
}


func ReplaceString(str string) string {
	str  = strings.Trim(str, " ")
	str = strings.Trim(str, "\n")
	return str
}

func RandNum(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func GetConfigPath(path string) string {
	root, _:= os.Getwd()
	fileName, _ := filepath.Abs(filepath.Join(root, path))
	return fileName
}


func ReadProxyList(path string) ([]string, error) {
	var list []string
	fileName := GetConfigPath(path)
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fio := bufio.NewScanner(f)
	for fio.Scan() {
		list = append(list, fio.Text())
	}
	err = fio.Err()
	if err != nil {
		return nil ,err
	}
	return list, nil
}