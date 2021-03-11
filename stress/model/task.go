package model

import (
	"go_study/global"
	"go_study/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type Task interface {
	run() (string,string,float64)
}

type UrlAccessTask struct {
	url string
}

func (task *UrlAccessTask) run() (string,string,float64){
	t := time.Now()
	global.GVA_LOG.Println("url:",task.url)
	res, _ := http.Get(task.url)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	t2 := time.Since(t)
	return utils.Bytes2str(body),"err",t2.Seconds()
}
