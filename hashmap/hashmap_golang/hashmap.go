package hashmap_golang

import "fmt"

const defaultHashMapLen = 1000000

type HashMap struct {
	array   [][][]interface{}
	curSize int
}

func NewHashMap() HashMap {
	return NewHashMapWithLen(defaultHashMapLen)
}

func NewHashMapWithLen(len int) HashMap {
	a := make([][][]interface{}, len)
	return HashMap{
		array:   a,
		curSize: len,
	}
}

func (h *HashMap) Set(key string, value int) {
	hash := HashFuncStringToInt(key, h.curSize)
	if len(h.array[hash]) == 0 {
		// if new to this hash...
		newarrelm := [][]interface{}{{key, value}}
		h.array[hash] = newarrelm
		return
	}
	curval := h.array[hash]
	newarr := make([][]interface{}, len(curval)+1)
	copy(newarr, curval)
	newarrelm := []interface{}{key, value}
	newarr = append(newarr, newarrelm)
	h.array[hash] = newarr
}

func (h HashMap) Get(key string) int {
	hash := HashFuncStringToInt(key, h.curSize)
	sameHashArr := h.array[hash]
	for i := 0; i < len(sameHashArr); i++ {
		// 使われていないhash値
		if len(sameHashArr[i]) == 0 {
			continue
		}
		if sameHashArr[i][0] == key {
			return sameHashArr[i][1].(int)
		}
	}
	panic(fmt.Sprintf("Not found!!\nkey: %s\ninternal sameHashArr array: %+v", key, sameHashArr))
}

func HashFuncStringToInt(key string, arraySize int) int {
	var hash uint
	for i := 0; i < len(key); i++ {
		hash += (uint(key[i])>>32)*uint(key[i]) + (1<<62 - 1)
	}
	return int(hash>>1|1<<61) % arraySize
}
