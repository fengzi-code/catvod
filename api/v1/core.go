package v1

import (
	"catvod/service"
	"catvod/utils"
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
	wd, ok4 := ctx.GetQuery("wd")     // 有wd参数时走搜索程序
	fmt.Printf("%s,%v,%v,%v,%v\n", tvProvider, t, play, ids, wd)
	tv := service.NewTVProvider(tvProvider)

	switch {
	case ok1:

		pg := ctx.DefaultQuery("pg", "1")
		ext := ctx.DefaultQuery("ext", "")
		fmt.Printf("接口为: %s, 分类编号为: %s,页数为: %s,走分类程序\n", tvProvider, t, pg)
		pgInt, err := strconv.Atoi(pg)
		if err != nil {
			pgInt = 1
		}
		if pgInt < 1 {
			pgInt = 1
		}
		filters := tv.GetFilter(ext)
		fmt.Printf("接口为: %s, 分类筛选数据: %s \n", tvProvider, filters)
		data := tv.GetCategory(t, pgInt)
		ctx.JSON(200, data)
	case ok2:
		data := utils.GetPlayUrl(play)
		fmt.Printf("接口为: %s, 播放地址: %s,解析模式: %d,解析地址: %s \n", tvProvider, data.Url, data.Parse, data.PlayUrl)
		ctx.JSON(http.StatusOK, data)
	case ok3:
		idsArr := make([]string, 0)
		if strings.Contains(ids, ",") {
			idsArr = strings.Split(ids, ",")
		} else {
			idsArr = append(idsArr, ids)
		}
		fmt.Printf("接口为: %s, 走详情程序,视频ID数组: %s\n", tvProvider, idsArr)
		data := tv.GetDetails(idsArr)
		ctx.JSON(200, gin.H{"code": 0, "message": "视频详情", "list": data})
	case ok4:
		data := tv.Search(wd)
		fmt.Printf("接口为: %s, 走搜索程序,搜索关键字: %s\n", tvProvider, wd)
		ctx.JSON(200, gin.H{"code": 0, "total": len(data), "message": "视频搜索", "list": data})
	default:
		fmt.Printf("接口为: %s, 进入首页 \n", tvProvider)
		data := tv.GetHome()
		ctx.JSON(200, data)
	}
}
