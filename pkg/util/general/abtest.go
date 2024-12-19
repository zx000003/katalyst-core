package general

import (
	"hash/crc32"
	"os"
)

const (
	ABEnabledKey = "AB_SAMPLING_ENABLED"
	ABEnabledVal = "true"
)

func EnableBorwein() bool {
	mod := ABTestMod()
	Infof("Node mod is %v", mod)
	if mod == -1 {
		return true
	}
	if mod > 40 {
		return true
	}
	return false
}

func EnableDynamicThreshold() bool {
	mod := ABTestMod()
	Infof("Node mod is %v", mod)
	if mod == -1 {
		return true
	}
	if mod > 10 {
		return true
	}
	return false
}

func ABTestMod() int {
	str := os.Getenv(ABEnabledKey)
	if str != ABEnabledVal {
		return -1
	}
	hostname, _ := os.Hostname()
	hash := crc32.ChecksumIEEE([]byte(hostname))
	return int(hash) % 100
}
