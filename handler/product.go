package handler

import (
	"context"
	"fmt"
	"product/common"
	"product/domain/model"
	"product/domain/service"
	. "product/proto/product"
)

type Product struct{
	ProductDataService service.IProductDataService
}
//添加商品

func(h *Product)AddProduct(ctx context.Context,req *ProductInfo,res *ResponseProduct) error{
	productAdd := &model.Product{}
	fmt.Println(req)
	if err := common.SwapTo(req,productAdd);err!=nil{
		return err
	}
	fmt.Println(productAdd)
	productID,err := h.ProductDataService.AddProduct(productAdd)
	if err!=nil{
		return err
	}
	res.ProductId=productID
	return nil
}

//根据ID查找商品
func(h *Product)FindProductByID(ctx context.Context,req *RequestID, res *ProductInfo) error{
	productData,err := h.ProductDataService.FindProductByID(req.ProductId)
	if err!=nil{
		return err
	}
	if err := common.SwapTo(productData,res);err!=nil{
		return err
	}
	return nil
}

//商品更新
func(h *Product)UpdateProduct(ctx context.Context,req *ProductInfo, res *Response) error{
	productAdd := &model.Product{}
	if err := common.SwapTo(req,productAdd);err!=nil{
		return err
	}
	err := h.ProductDataService.UpdateProduct(productAdd)
	if err!=nil{
		return err
	}
	res.Msg="更新成功"
	return nil
}
//根据ID删除对应商品
func(h *Product)DeleteProductByID(ctx context.Context,req *RequestID,res *Response) error{
	if err := h.ProductDataService.DeleteProduct(req.ProductId);err!=nil{
		return err
	}
	res.Msg="删除成功"
	return nil
}

//查找所有商品
func(h *Product)FindAllProduct(ctx context.Context,req *RequestAll,res *AllProduct) error{
	productAll,err := h.ProductDataService.FindAllProduct()
	if err!=nil{
		return err
	}
	for _,v := range productAll{
		productInfo := &ProductInfo{}
		err := common.SwapTo(v,productInfo)
		if err!=nil{
			return err
		}
		res.ProductInfo=append(res.ProductInfo,productInfo)
	}
	return nil
}