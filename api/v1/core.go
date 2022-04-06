package v1

import (
	"catvod/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CoreHandler(ctx *gin.Context) {
	tvProvider := ctx.Param("tv")
	// 获取 id 参数，通过 GetQuery 获取参数，返回string类型，第二个参数，可判断是否存参数
	t, ok1 := ctx.GetQuery("t")       // 有t参数时走分类程序
	play, ok2 := ctx.GetQuery("play") // 有play参数时走播放程序
	ids, ok3 := ctx.GetQuery("ids")   // 有ids参数时走详情程序
	wd, ok4 := ctx.GetQuery("wd")     // 有ids参数时走详情程序
	fmt.Printf("%s,%v,%v,%v,%v\n", tvProvider, t, play, ids, wd)
	tv := service.NewTVProvider(tvProvider)
	if ok1 {
		fmt.Println("走分类程序", t)
		pg := ctx.DefaultQuery("pg", "1")
		ext := ctx.DefaultQuery("ext", "")
		fmt.Println(pg, ext)
		pgInt, err := strconv.Atoi(pg)
		if err != nil {
			pgInt = 1
		}
		if pgInt < 1 {
			pgInt = 1
		}
		data := tv.GetCategory(t, pgInt)
		fmt.Println(data)
		ctx.JSON(200, data)
	} else if ok2 {
		fmt.Println("走播放程序")
		fmt.Println(play)
		ctx.JSON(
			http.StatusOK,
			gin.H{"code": 0, "playUrl": "https://jx.blbo.cc:4433/?url=", "url": play, "parse": 1, "header": ""},
		)
	} else if ok3 {
		fmt.Println("走详情程序")
		fmt.Println(ids) // 传过来的ids为字符串数组
		idsArr := make([]string, 0)
		if strings.Contains(ids, ",") {
			idsArr = strings.Split(ids, ",")
		} else {
			idsArr = append(idsArr, ids)
		}

		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "走详情程序",
			"data":    "tv provider: " + tvProvider,
		})
	} else if ok4 {
		// https://xxx.xX.xx?ac=detail&quick=是否快搜&wd=搜索词
		fmt.Println("走搜索程序")
		fmt.Println(wd) // 传过来的ids为字符串数组
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "走搜索程序",
			"data":    "tv provider: " + tvProvider,
		})
	} else {
		fmt.Println("走首页程序")
		data := tv.GetHome()
		ctx.JSON(200, data)
	}

}
