package number_cast

import (
	"errors"
)

/*
	@author sunyuw <leentingOne@gmail.com>
	@date 2013-9-24

	将byte字节数组转换为32位整型.
*/
func Bytes2Int32(IBytes []byte) (castedInt int32, err error) {
	castedInt = 0x00
	if len(IBytes) > 4 {
		err = errors.New("length of parameter array is larger than 4,and result overflowed!")
	}
	for _, ibyte := range IBytes {
		castedInt = (castedInt << 8) | int32(ibyte)
	}
	return
}
