package controller

import (
	"context"
	"gin-vue-seckill/jdaw_product_srv/dao"
	"gin-vue-seckill/jdaw_product_srv/models"
	"gin-vue-seckill/jdaw_product_srv/proto/product"
	"time"
)

type Products struct{}

var timeLayoutStr = "2006-01-02 15:04:05"

func (pro *Products) GetProductList(ctx context.Context, in *product.ProductListRequest, out *product.ProductListResponse) error {
	cp := in.CurrentPage
	ps := in.PageSize
	currentPage := int(cp)
	pageSize := int(ps)
	productlist, err, count := dao.GetProductList(currentPage, pageSize)
	if err != nil {
		out.Code = 200
		out.Msg = "当前无商品"
		return err
	} else {
		out.Code = 200
		out.Msg = "已查到商品"
		productreList := make([]*product.ProductDetail, 0, len(productlist))
		for _, products := range productlist {
			productre := &product.ProductDetail{
				Id:         int32(products.Id),
				Name:       products.Name,
				Price:      products.Price,
				Num:        int32(products.Num),
				Unit:       products.Unit,
				Pic:        products.Pic,
				Desc:       products.Desc,
				CreateTime: products.CreateTime.Format(timeLayoutStr),
			}
			productreList = append(productreList, productre)
		}
		out.Productlist = productreList
		out.PageSize = int32(pageSize)
		out.Current = int32(currentPage)
		out.Total = int32(count)
		return nil
	}
}

func (pro *Products) ProductAdd(ctx context.Context, in *product.ProductAddRequest, out *product.ProductAddResponse) error {
	name := in.Name
	price := in.Price
	num := in.Num
	uint := in.Unit
	desc := in.Desc
	pic := in.Pic

	product := models.Products{
		Name:       name,
		Price:      price,
		Num:        int(num),
		Unit:       uint,
		Pic:        pic,
		Desc:       desc,
		CreateTime: time.Now(),
	}
	err := dao.CreateProduct(&product)
	out.Code = 200
	if err != nil {
		out.Msg = "服务器错误无法创建商品"
	}
	out.Msg = "创建商品成功"
	return nil
}

func (pro *Products) ProductDelete(ctx context.Context, in *product.ProductDeleteRequest, out *product.ProductAddResponse) error {
	id32 := in.Id
	id := int(id32)
	err := dao.DeleteProduct(id)
	out.Code = 200
	if err != nil {
		out.Msg = "删除失败"
	}
	out.Msg = "删除成功"
	return nil
}

func (pro *Products) GetProductInfo(ctx context.Context, in *product.ProductDeleteRequest, out *product.ProductInfoResponse) error {
	id32 := in.Id
	id := int(id32)
	product, err := dao.GetProduct(id)
	if err == nil {
		out.Id = id32
		out.Name = product.Name
		out.Num = int32(product.Num)
		out.Unit = product.Unit
		out.Price = product.Price
		out.Pic = product.Pic
		out.Desc = product.Desc
		return nil
	}
	return err
}

func (pro *Products) ProductUpdate(ctx context.Context, in *product.ProductUpdateRequest, out *product.ProductAddResponse) error {
	id := in.Id
	name := in.Name
	price := in.Price
	num := in.Num
	uint := in.Unit
	desc := in.Desc
	pic := in.Pic
	product := models.Products{
		Id:         int(id),
		Name:       name,
		Price:      price,
		Num:        int(num),
		Unit:       uint,
		Pic:        pic,
		Desc:       desc,
		CreateTime: time.Now(),
	}
	err := dao.UpdateProduct(&product)
	out.Code = 200
	if err != nil {
		out.Msg = "服务器错误无法修改商品"
	}
	out.Msg = "修改商品成功"
	return nil
}
