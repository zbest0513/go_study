package model

import (
	"github.com/shopspring/decimal"
	"go_study/global"
	"sort"
	"strconv"
	"sync"
)

type Result struct {
	url string
	time float64
}

type Results []Result

func (r Results) String() string {
	res := "["
	for i,item := range r {
		res += strconv.FormatFloat(item.time,'f',2,64)
		if i < r.Len()-1 {
			res += ","
		}
	}
	res += "]"
	return res
}

func (r Results) Len() int {
	return len(r)
}

func (r Results) Less(i, j int) bool {
	return r[i].time < r[j].time
}

func (r Results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type Runner struct {
	Target      string
	ThreadCount int
	Results     Results
}

func (runner *Runner) Run() map[string]string {
	runner.Results = make(Results,runner.ThreadCount)
	var waitGroup sync.WaitGroup
	for i := 0; i < runner.ThreadCount; i++ {
		waitGroup.Add(1)
		var tmp = i
		go func() {
			task := UrlAccessTask{runner.Target}
			_, _, time := task.run()
			r := Result{runner.Target, time}
			runner.Results[tmp] = r
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	global.GVA_LOG.Println("results:",runner.Results)
	return runner.report()
}

func (runner *Runner) report() map[string]string {
	sort.Sort(runner.Results)
	intPart := decimal.NewFromInt(int64(runner.ThreadCount)).Mul(decimal.NewFromFloat(0.95)).IntPart() - 1

	var total = 0.0
	for _,item := range runner.Results {
		total += item.time
	}
	avg, _ := decimal.NewFromFloat(total).DivRound(decimal.NewFromInt(int64(runner.ThreadCount)),2).Float64()

	global.GVA_LOG.Println("intPart : ",intPart)
	global.GVA_LOG.Println("report : ")
	global.GVA_LOG.Println("	avg time : ",avg," s")
	global.GVA_LOG.Println("	95% time : ",runner.Results[intPart])
	res:= make(map[string]string, 3)
	res["avg_time"] = strconv.FormatFloat(avg,'f',2,64)
	res["target_time"] = strconv.FormatFloat(runner.Results[intPart].time,'f',2,64)
	res["results"] = runner.Results.String()

	global.GVA_LOG.Println("	results : ",runner.Results)
	return res
}
