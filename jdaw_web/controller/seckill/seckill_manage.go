package seckill

import (
	"gin-vue-seckill/jdaw_product_srv/proto/seckill"
	"gin-vue-seckill/jdaw_user_srv/common/utils"
	util "gin-vue-seckill/jdaw_web/common/utils"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSeckillList(ctx *gin.Context) {
	cp := ctx.Query("currentPage")
	ps := ctx.Query("pageSize")
	currentPage, _ := strconv.ParseInt(cp, 10, 0)
	pageSize, _ := strconv.ParseInt(ps, 10, 0)

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, err := client.GetSeckillList(ctx, &seckill.SeckillListRequest{
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "没有获取到数据",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         rep.Code,
			"msg":          rep.Msg,
			"seckills":     rep.Seckillls,
			"total":        int(rep.Total),
			"current_page": int(rep.Current),
			"page_size":    int(rep.PageSize),
		})
	}
}

func GetProduct(ctx *gin.Context) {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, err := client.GetProducts(ctx, &seckill.ProductRequest{})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "无法查到任何商品数据",
		})
	} else {

		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"msg":      "成功查询到商品数据",
			"products": rep.Products,
		})
	}
}

func DeleteProduct(ctx *gin.Context) {
	id32 := ctx.PostForm("id")
	id, _ := strconv.ParseInt(id32, 10, 32)
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, _ := client.SeckillDelete(ctx, &seckill.SeckillDeleteRequest{Id: int32(id)})
	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func AddSeckill(ctx *gin.Context) {
	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	num := ctx.PostForm("num")
	pid := ctx.PostForm("pid")
	starttime := ctx.PostForm("start_time")
	endtime := ctx.PostForm("end_time")
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, _ := client.SecKillAdd(ctx, &seckill.SeckillDetail{
		Name:      name,
		Price:     utils.StrToFloat32(price),
		Num:       utils.StrToInt(num),
		Pid:       utils.StrToInt(pid),
		StartTime: starttime,
		EndTime:   endtime,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func ToEditSeckill(ctx *gin.Context) {
	id32 := ctx.Query("id")
	id, _ := strconv.ParseInt(id32, 10, 32)
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, err := client.SecKillToEdit(ctx, &seckill.SeckillDeleteRequest{Id: int32(id)})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "没有查询到数据",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":        rep.Code,
		"msg":         rep.Msg,
		"seckill":     rep.Seckillde,
		"products_no": rep.ProductsNo,
	})
}

func UpdateSeckill(ctx *gin.Context) {
	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	num := ctx.PostForm("num")
	pid := ctx.PostForm("pid")
	starttime := ctx.PostForm("start_time")
	endtime := ctx.PostForm("end_time")
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, _ := client.SecKillDoEdit(ctx, &seckill.SeckillDetail{
		Id:        utils.StrToInt(id),
		Name:      name,
		Price:     utils.StrToFloat32(price),
		Num:       utils.StrToInt(num),
		Pid:       utils.StrToInt(pid),
		StartTime: starttime,
		EndTime:   endtime,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func GetFrontSeckillList(ctx *gin.Context) {
	cp := ctx.Query("currentPage")
	ps := ctx.Query("pageSize")
	currentPage, _ := strconv.ParseInt(cp, 10, 0)
	pageSize, _ := strconv.ParseInt(ps, 10, 0)

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, err := client.FrontSecKillList(ctx, &seckill.SeckillListRequest{
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "没有获取到数据",
		})
	}

	for _, seckills := range rep.Seckillls {
		seckills.Pic = util.Img2Base64(seckills.Pic)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":         rep.Code,
		"msg":          rep.Msg,
		"current":      rep.Current,
		"page_size":    rep.PageSize,
		"total_page":   rep.Total,
		"seckill_list": rep.Seckillls,
	})
}

func GetSeckillInfo(ctx *gin.Context) {
	id32 := ctx.Query("id")
	id, _ := strconv.ParseInt(id32, 10, 32)
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := seckill.NewSeckillsService("jdaw.product.srv", s.Client())
	rep, err := client.FrontSecKillDetail(ctx, &seckill.SeckillDeleteRequest{Id: int32(id)})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "没有查询到数据",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    rep.Code,
		"msg":     rep.Msg,
		"seckill": rep.Seckillde,
	})
}
