package base

import (
	"errors"
	"strings"
)

type BaseGet interface {
	GetInfo(resource string) (data interface{}, productName, productId, router string, err error)
}

func GetTrace(resource string) (data interface{}, productName, productId, router string, err error) {
	if len(resource) == 0 {
		return nil, "", "", "", errors.New("no resource")
	}

	if strings.HasPrefix(resource, "http://") || strings.HasPrefix(resource, "https://") {
		//URL
		if strings.HasPrefix(resource, "http://yu.jywykjgs.com") {
			//多宝鱼
			dby := BaseGetDby{}
			return dby.GetInfo(resource)
		}
	} else {
		return nil, "", "", "", errors.New("not yet")
	}
	return
}
