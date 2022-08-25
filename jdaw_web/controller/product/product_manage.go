package product

import (
	"fmt"
	"gin-vue-seckill/jdaw_product_srv/proto/product"
	"gin-vue-seckill/jdaw_user_srv/common/utils"
	util "gin-vue-seckill/jdaw_web/common/utils"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetProductList(ctx *gin.Context) {
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
	client := product.NewProductsService("jdaw.product.srv", s.Client())
	rep, err := client.GetProductList(ctx, &product.ProductListRequest{
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
			"products":     rep.Productlist,
			"total":        int(rep.Total),
			"current_page": int(rep.Current),
			"page_size":    int(rep.PageSize),
		})
	}
}

func AddProduct(ctx *gin.Context) {
	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	num := ctx.PostForm("num")
	unit := ctx.PostForm("unit")
	desc := ctx.PostForm("desc")
	file, err := ctx.FormFile("pic")
	if err != nil {
		ctx.JSON(http.StatusHTTPVersionNotSupported, gin.H{
			"code": 501,
			"msg":  "请上传商品图片",
		})
	}
	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "jdaw_web/upload/" + unix_str + file.Filename
	err = ctx.SaveUploadedFile(file, file_path)
	fmt.Println(err)
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := product.NewProductsService("jdaw.product.srv", s.Client())
	rep, _ := client.ProductAdd(ctx, &product.ProductAddRequest{
		Name:  name,
		Price: utils.StrToFloat32(price),
		Num:   utils.StrToInt(num),
		Unit:  unit,
		Pic:   file_path,
		Desc:  desc,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})

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
	client := product.NewProductsService("jdaw.product.srv", s.Client())
	rep, _ := client.ProductDelete(ctx, &product.ProductDeleteRequest{Id: int32(id)})
	ctx.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func ToProductEdit(ctx *gin.Context) {
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
	client := product.NewProductsService("jdaw.product.srv", s.Client())
	rep, err := client.GetProductInfo(ctx, &product.ProductDeleteRequest{Id: int32(id)})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "无法查到该商品数据",
		})
	} else {
		pro := product.ProductInfoResponse{
			Id:    rep.Id,
			Name:  rep.Name,
			Price: rep.Price,
			Num:   rep.Num,
			Unit:  rep.Unit,
			Pic:   rep.Pic,
			Desc:  rep.Desc,
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":       200,
			"msg":        "成功查询到商品数据",
			"product":    pro,
			"img_base64": util.Img2Base64(rep.Pic),
		})
	}
}

func UpdataProduct(ctx *gin.Context) {
	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	num := ctx.PostForm("num")
	unit := ctx.PostForm("unit")
	desc := ctx.PostForm("desc")
	file, err := ctx.FormFile("pic")
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := product.NewProductsService("jdaw.product.srv", s.Client())
	if err != nil {
		rep, _ := client.ProductUpdate(ctx, &product.ProductUpdateRequest{
			Id:    utils.StrToInt(id),
			Name:  name,
			Price: utils.StrToFloat32(price),
			Num:   utils.StrToInt(num),
			Unit:  unit,
			Desc:  desc,
		})
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	} else {
		unix_int64 := time.Now().Unix()
		unix_str := strconv.FormatInt(unix_int64, 10)
		file_path := "jdaw_web/upload/" + unix_str + file.Filename
		ctx.SaveUploadedFile(file, file_path)
		rep, _ := client.ProductUpdate(ctx, &product.ProductUpdateRequest{
			Id:    utils.StrToInt(id),
			Name:  name,
			Price: utils.StrToFloat32(price),
			Num:   utils.StrToInt(num),
			Unit:  unit,
			Pic:   file_path,
			Desc:  desc,
		})
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	}

}
