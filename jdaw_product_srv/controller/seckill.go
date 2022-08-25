package controller

import (
	"context"
	"gin-vue-seckill/jdaw_product_srv/dao"
	"gin-vue-seckill/jdaw_product_srv/models"
	"gin-vue-seckill/jdaw_product_srv/proto/seckill"
	"gin-vue-seckill/jdaw_user_srv/common/utils"
	mysql "gin-vue-seckill/jdaw_user_srv/data_source"
	util "gin-vue-seckill/jdaw_web/common/utils"
	"strconv"
	"time"
)

type Seckills struct{}

func (se *Seckills) GetSeckillList(ctx context.Context, in *seckill.SeckillListRequest, out *seckill.SeckillListResponse) error {
	cp := in.CurrentPage
	ps := in.PageSize
	currentPage := int(cp)
	pageSize := int(ps)
	seckilllist, err, count := dao.GetSeckillList(currentPage, pageSize)
	if err != nil {
		out.Code = 200
		out.Msg = "当前无商品"
		return err
	} else {
		out.Code = 200
		out.Msg = "已查到商品"
		seckillreList := make([]*seckill.SeckillDetail, 0, len(seckilllist))
		for _, seckills := range seckilllist {
			product := models.Products{
				Id: int(utils.StrToInt(seckills.PId)),
			}
			mysql.Db.First(&product)
			seckillre := &seckill.SeckillDetail{
				Id:         int32(seckills.Id),
				Name:       seckills.Name,
				Price:      seckills.Price,
				Num:        int32(seckills.Num),
				Pid:        utils.StrToInt(seckills.PId),
				Pname:      product.Name,
				StartTime:  seckills.StartTime.Format(timeLayoutStr),
				EndTime:    seckills.EndTime.Format(timeLayoutStr),
				CreateTime: seckills.CreateTime.Format(timeLayoutStr),
			}
			seckillreList = append(seckillreList, seckillre)
		}
		out.Seckillls = seckillreList
		out.PageSize = int32(pageSize)
		out.Current = int32(currentPage)
		out.Total = int32(count)
		return nil
	}
}

func (se *Seckills) GetProducts(ctx context.Context, in *seckill.ProductRequest, out *seckill.ProductResponse) error {
	prol, err := dao.GetProductList2()
	if err == nil {
		productList := make([]*seckill.Product, 0, len(prol))
		for _, pro := range prol {
			products := &seckill.Product{
				Id:    int32(pro.Id),
				Pname: pro.Name,
			}
			productList = append(productList, products)
		}
		out.Code = 200
		out.Msg = "已查到商品"
		out.Products = productList
		return nil
	} else {
		out.Code = 200
		out.Msg = "当前无商品"
		return err
	}
}

func (se *Seckills) SeckillDelete(ctx context.Context, in *seckill.SeckillDeleteRequest, out *seckill.SeckillDeleteResponse) error {
	id32 := in.Id
	id := int(id32)
	err := dao.DeleteSeckill(id)
	out.Code = 200
	if err != nil {
		out.Msg = "删除失败"
	}
	out.Msg = "删除成功"
	return nil
}

func (se *Seckills) SecKillAdd(ctx context.Context, in *seckill.SeckillDetail, out *seckill.SeckillDeleteResponse) error {
	name := in.Name
	price := in.Price
	num := in.Num
	pid := in.Pid
	starttime := in.StartTime
	endtime := in.EndTime
	stime, _ := time.ParseInLocation(starttime, timeLayoutStr, time.Local)
	etime, _ := time.ParseInLocation(endtime, timeLayoutStr, time.Local)
	seckill := models.Seckills{
		Name:       name,
		Price:      price,
		Num:        int(num),
		PId:        string(pid),
		StartTime:  stime,
		EndTime:    etime,
		CreateTime: time.Now(),
	}
	err := dao.CreateSeckill(&seckill)
	out.Code = 200
	if err != nil {
		out.Msg = "服务器错误无法创建商品"
	}
	out.Msg = "创建套餐成功"
	return nil
}

func (se *Seckills) SecKillToEdit(ctx context.Context, in *seckill.SeckillDeleteRequest, out *seckill.SecKillToEditResponse) error {
	id32 := in.Id
	id := int(id32)
	seckills, err := dao.GetSeckill(id)
	if err != nil {
		out.Code = 200
		out.Msg = "无法获取当前活动"
		return err
	}
	seckillde := &seckill.SeckillDetail{
		Id:        int32(id),
		Name:      seckills.Name,
		Price:     seckills.Price,
		Num:       int32(seckills.Num),
		Pid:       utils.StrToInt(seckills.PId),
		StartTime: seckills.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:   seckills.EndTime.Format("2006-01-02 15:04:05"),
	}
	out.Seckillde = seckillde
	prol, err := dao.GetProductList2()
	if err == nil {
		productList := make([]*seckill.Product, 0, len(prol))
		for _, pro := range prol {
			products := &seckill.Product{
				Id:    int32(pro.Id),
				Pname: pro.Name,
			}
			productList = append(productList, products)
		}
		out.Code = 200
		out.Msg = "已查到商品"
		out.ProductsNo = productList
		return nil
	} else {
		out.Code = 200
		out.Msg = "当前无商品"
		return err
	}
}

func (se *Seckills) SecKillDoEdit(ctx context.Context, in *seckill.SeckillDetail, out *seckill.SeckillDeleteResponse) error {
	id := in.Id
	name := in.Name
	price := in.Price
	num := in.Num
	pid := in.Pid
	stime := in.StartTime
	etime := in.EndTime
	starttime, _ := time.Parse(stime, timeLayoutStr)
	endtime, _ := time.Parse(etime, timeLayoutStr)
	seckills := &models.Seckills{
		Id:        int(id),
		Name:      name,
		Price:     price,
		Num:       int(num),
		PId:       string(pid),
		StartTime: starttime,
		EndTime:   endtime,
	}
	err := dao.UpdateSeckill(seckills)
	out.Code = 200
	if err != nil {
		out.Msg = "服务器错误无法修改套餐"
	}
	out.Msg = "修改套餐成功"
	return nil
}

func (se *Seckills) FrontSecKillList(ctx context.Context, in *seckill.SeckillListRequest, out *seckill.SeckillListResponse) error {
	cp := in.CurrentPage
	ps := in.PageSize
	currentPage := int(cp)
	pageSize := int(ps)
	seckilllist, err, count := dao.GetSeckillList(currentPage, pageSize)
	if err != nil {
		out.Code = 200
		out.Msg = "当前无商品"
		return err
	} else {
		out.Code = 200
		out.Msg = "已查到商品"
		seckillreList := make([]*seckill.SeckillDetail, 0, len(seckilllist))
		for _, seckills := range seckilllist {
			product := models.Products{
				Id: int(utils.StrToInt(seckills.PId)),
			}
			mysql.Db.First(&product)
			seckillre := &seckill.SeckillDetail{
				Id:         int32(seckills.Id),
				Name:       seckills.Name,
				Price:      seckills.Price,
				Num:        int32(seckills.Num),
				Pid:        utils.StrToInt(seckills.PId),
				Pname:      product.Name,
				StartTime:  seckills.StartTime.Format(timeLayoutStr),
				EndTime:    seckills.EndTime.Format(timeLayoutStr),
				CreateTime: seckills.CreateTime.Format(timeLayoutStr),
				Pic:        product.Pic,
				PPrice:     product.Price,
				Pdesc:      product.Desc,
				Unit:       product.Unit,
			}
			seckillreList = append(seckillreList, seckillre)
		}
		out.Seckillls = seckillreList
		out.PageSize = int32(pageSize)
		out.Current = int32(currentPage)
		out.Total = int32(count)
		return nil
	}
}

func (se *Seckills) FrontSecKillDetail(ctx context.Context, in *seckill.SeckillDeleteRequest, out *seckill.FrontSecKillDetailResponse) error {
	id32 := in.Id
	id := int(id32)
	seckills, err := dao.GetSeckill(id)
	if err != nil {
		out.Code = 200
		out.Msg = "无法获取当前商品或活动"
		return err
	}
	pid, _ := strconv.ParseInt(seckills.PId, 0, 10)
	pro, err1 := dao.GetProduct(int(pid))
	if err1 != nil {
		out.Code = 200
		out.Msg = "无法获取当前商品或活动"
		return err
	}
	seckill_rep := &seckill.SeckillDetail{
		Id:         int32(id),
		Name:       seckills.Name,
		Price:      seckills.Price,
		Num:        int32(seckills.Num),
		Pid:        utils.StrToInt(seckills.PId),
		Pname:      pro.Name,
		StartTime:  seckills.StartTime.Format(timeLayoutStr),
		EndTime:    seckills.EndTime.Format(timeLayoutStr),
		CreateTime: seckills.CreateTime.Format(timeLayoutStr),
		Pic:        util.Img2Base64(pro.Pic),
		PPrice:     pro.Price,
		Pdesc:      pro.Desc,
		Unit:       pro.Unit,
	}
	out.Code = 200
	out.Msg = "查询成功"
	out.Seckillde = seckill_rep
	return nil
}
