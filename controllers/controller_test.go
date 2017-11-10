package controllers

import (
	"testing"
	"surl/storage"
)

const (
	pass = "\u2713"
	fail = "\u2717"
)

func TestStoreUrl(t *testing.T) {

	str := storage.RedisPool.StoreLongUrl2Redis("https://www.surltesturl.com/s?ie=utf-8&f=8&rsv_bp=1&tn=heh&wd=%E5%8E%9F%E5%AD%" +
		"90%E6%93%8D%E4%BD%9C%20golang&oq=%25E5%258E%259F%25E5%25AD%2590%25E6%2593%258D%25E4%25BD%259C&rsv_pq=9567b597" +
		"0000079e&rsv_t=9a64OwwtbGj3KbOk7ThPlYkGmivy075mRyAXBOtazaeNimpLBlo14SxefbM&rqlang=cn&rsv_enter=1&rsv_sug3=" +
		"10&rsv_sug2=0&inputT=2383&rsv_sug4=3068")

	if str == "" {
		t.Fatal("store long url error", "\033[;31m", fail, "\033[;0m")
	} else {
		t.Log("store long url success", "\033[;32m", pass, "\033[;0m")
	}
}
