package core

import (
	"Nietzsche/internal/constant"
	"Nietzsche/internal/data_structure"
	"errors"
	"fmt"
	"strconv"
)

func cmdBFRESERVE(args []string) []byte {
	if !(len(args) == 3 || len(args) == 5) {
		return Encode(errors.New("(error) ERR wrong number of arguments for 'BF.RESERVE' command"), false)
	}
	key := args[0]
	errRate, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return Encode(errors.New(fmt.Sprintf("error rate must be a floating point number %s", args[1])), false)
	}
	capacity, err := strconv.ParseUint(args[2], 10, 64)
	if err != nil {
		return Encode(errors.New(fmt.Sprintf("capacity must be an integer number %s", args[2])), false)
	}
	_, exist := bloomStore[key]
	if exist {
		return Encode(errors.New(fmt.Sprintf("Bloom filter with key '%s' already exist", key)), false)
	}
	bloomStore[key] = data_structure.CreateBloomFilter(capacity, errRate)
	return constant.RespOk
}

func cmdBFMADD(args []string) []byte {
	if len(args) < 2 {
		return Encode(errors.New("(error) ERR wrong number of arguments for 'BF.MADD' command"), false)
	}
	key := args[0]
	bloom, exist := bloomStore[key]
	if !exist {
		bloom = data_structure.CreateBloomFilter(constant.BfDefaultInitCapacity,
			constant.BfDefaultErrRate)
		bloomStore[key] = bloom
	}
	var res []string
	for i := 1; i < len(args); i++ {
		item := args[i]
		bloom.Add(item)
		res = append(res, "1")
	}
	return Encode(res, false)
}

func cmdBFEXISTS(args []string) []byte {
	if len(args) != 2 {
		return Encode(errors.New("(error) ERR wrong number of arguments for 'BF.EXISTS' command"), false)
	}
	key, item := args[0], args[1]
	bloom, exist := bloomStore[key]
	if !exist {
		return constant.RespZero
	}
	if !bloom.Exist(item) {
		return constant.RespZero
	}
	return constant.RespOne
}
