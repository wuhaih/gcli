package util

import "unsafe"

/**
* 
* @author willian
* @created 2017-01-26 14:24
* @email 18702515157@163.com  
**/

func ByteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}