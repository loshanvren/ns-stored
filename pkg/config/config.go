package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"os"
)

type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

// DefaultConfig 只在测试直接访问
var DefaultConfig *viper.Viper

func init() {
	DefaultConfig = readViperConfig("env")

	// global defaults
	DefaultConfig.SetDefault("json_logs", false)
	DefaultConfig.SetDefault("loglevel", "debug")
}

// 读取默认配置
func Config() Provider {
	return DefaultConfig
}

// 读取配置中的字符串值
func GetStringValue(key string, defaultValue string) string {
	value := DefaultConfig.GetString(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

// 读取配置中的字符串值
func GetBoolValue(key string, defaultValue bool) bool {
	if DefaultConfig.InConfig(key) {
		return DefaultConfig.GetBool(key)
	}
	return defaultValue
}

// 读取字符串
func GetString(key string) string {
	return DefaultConfig.GetString(key)
}

// GetInt 读取int值
func GetInt(key string) int {
	return DefaultConfig.GetInt(key)
}

// 初始化
func readViperConfig(configName string) *viper.Viper {
	v := viper.New()

	v.SetConfigName(configName) // name of config file (without extension)

	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		v.AddConfigPath(envPath)
	}
	v.AddConfigPath(".")                // optionally look for config in the working directory
	v.AddConfigPath("$HOME/.ns-stored") // call multiple times to add many search paths
	v.AddConfigPath("/etc/ns-stored")   // path to look for the config file in
	v.AddConfigPath("/opt/ns-stored")   // container path

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		log.Panicf("fatal error config file: %s", err)
	}

	v.AutomaticEnv() // 自动读取env
	return v
}
