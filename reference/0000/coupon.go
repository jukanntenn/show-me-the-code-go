/*
顺序循环方式生成 200 个兑换码，输出格式如下：
-----------------------------------
兑换码	lpsMkZ6F
生效时间	2020-02-27 12:55:44
过期时间	2020-03-28 12:55:44
-----------------------------------

主要参考资料：
[1] 生成随机字符串：https://blog.csdn.net/impressionw/article/details/72765756
[2] 时间处理：http://docs.studygolang.com/pkg/time/
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Coupon struct {
	code      string
	createdAt time.Time
	expireAt  time.Time
}

func main() {
	coupons := make([]Coupon, 200)

	for i := range coupons {
		createAt := time.Now()
		coupon := Coupon{
			code:      RandCode(8),
			createdAt: createAt,
			expireAt:  createAt.Add(time.Hour * 24 * 30),
		}
		coupons[i] = coupon
	}
	for i := range coupons {
		fmt.Println("-----------------------------------")
		fmt.Println("兑换码\t" + coupons[i].code)
		fmt.Println("生效时间\t" + coupons[i].createdAt.Format("2006-01-02 15:04:05"))
		fmt.Println("过期时间\t" + coupons[i].expireAt.Format("2006-01-02 15:04:05"))
		fmt.Println("-----------------------------------")
	}
}

func RandCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
