package config

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ServerProperties 定义全局配置属性
type ServerProperties struct {
	Bind           string `cfg:"bind"`           // 绑定地址
	Port           int    `cfg:"port"`           // 端口号
	AppendOnly     bool   `cfg:"appendOnly"`     // 是否只追加模式
	AppendFilename string `cfg:"appendFilename"` // 追加存档文件名
	MaxClients     int    `cfg:"maxclients"`     // 最大客户端连接数
	RequirePass    string `cfg:"requirepass"`    // 认证密码
	Databases      int    `cfg:"databases"`      // 数据库数量

	Peers []string // 其他节点地址（非映射属性，没有 cfg 标签）
}

// Properties 全局配置属性的指针
var Properties *ServerProperties

func init() {
	Properties = &ServerProperties{
		Bind:       "127.0.0.1",
		Port:       6379,
		AppendOnly: false,
	}
}

// parse 解析配置文件并返回 ServerProperties 结构体的实例
func parse(src io.Reader) *ServerProperties {
	config := &ServerProperties{} // 创建一个新的 ServerProperties 实例

	// 读取配置文件
	rawMap := make(map[string]string) // 存储原始配置键值对
	scanner := bufio.NewScanner(src)  // 创建扫描器以逐行读取
	for scanner.Scan() {              // 循环读取每一行
		line := scanner.Text() // 获取当前行
		// 跳过以 # 开头的行（注释行）
		if len(line) > 0 && line[0] == '#' {
			continue
		}
		// 查找第一个空格的位置
		pivot := strings.IndexAny(line, " ")
		// 确保空格在合适的位置
		if pivot > 0 && pivot < len(line)-1 {
			key := line[0:pivot]                       // 获取键（空格前的部分）
			value := strings.Trim(line[pivot+1:], " ") // 获取值（空格后的部分，去掉首尾空格）
			rawMap[strings.ToLower(key)] = value       // 将键值对存入原始映射，键转换为小写
		}
	}
	// 检查扫描器是否出现错误
	if err := scanner.Err(); err != nil {
		// logger.Fatal(err)
	}
	// 解析格式
	t := reflect.TypeOf(config)  // 获取 config 的类型信息
	v := reflect.ValueOf(config) // 获取 config 的值信息
	n := t.Elem().NumField()     // 获取结构体字段的数量

	for i := 0; i < n; i++ { // 遍历每一个字段
		field := t.Elem().Field(i)         // 获取字段类型
		fieldVal := v.Elem().Field(i)      // 获取字段值
		key, ok := field.Tag.Lookup("cfg") // 从标签中查找 cfg
		if !ok {
			key = field.Name
		}
		value, ok := rawMap[strings.ToLower(key)]
		if ok {
			// 根据字段类型填充配置
			switch field.Type.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Int:
				intValue, err := strconv.ParseInt(value, 10, 64)
				if err == nil {
					fieldVal.SetInt(intValue)
				}
			case reflect.Bool:
				boolValue := "yes" == value
				fieldVal.SetBool(boolValue)
			case reflect.Slice:
				if field.Type.Elem().Kind() == reflect.String {
					slice := strings.Split(value, ",")
					fieldVal.Set(reflect.ValueOf(slice))
				}
			}
		}
	}
	return config // 返回解析后的配置
}

// SetupConfig 读取配置文件并将属性存储到 Properties
func SetupConfig(configFilename string) {
	file, err := os.Open(configFilename) // 打开配置文件
	if err != nil {
		panic(err) // 如果打开失败，抛出错误
	}
	defer file.Close()
	Properties = parse(file) // 解析文件，并将结果存储在 Properties 中
}
