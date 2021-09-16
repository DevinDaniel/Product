package main

import (
	"context"
	"fmt"
	consul2 "github.com/asim/go-micro/plugins/registry/consul/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/opentracing/opentracing-go"
	"log"
	"product/common"
	go_micro_service_product "product/proto/product"
)

func main(){
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t,io,err := common.NewTracer("go.micro.service.product.client","localhost:6831")
	if err!=nil{
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//添加注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		)

	productService := go_micro_service_product.NewProductService("go.micro.service.product",service.Client())

	productAdd := &go_micro_service_product.ProductInfo{
		ProductName:          "product",
		ProductSku:           "devin",
		ProductPrice:         12.1,
		ProductDescription:   "product-devin",
		ProductCategoryId:    1,
		ProductImage:         []*go_micro_service_product.ProductImage{
			{
				ImageName:"devin-image",
				ImageCode:"devinimage1",
				ImageUrl:"devinimage1",
			},
			{
				ImageName:"devin-image2",
				ImageCode:"devinimage2",
				ImageUrl:"devinimage2",
			},
		},
		ProductSize:          []*go_micro_service_product.ProductSize{
			{
				SizeName:"devin-size",
				SizeCode:"devin-size-code",
			},
		},
		ProductSeo:           &go_micro_service_product.ProductSeo{
			SeoTitle:             "devin_seo",
			SeoKeywords:          "devin_seo",
			SeoDescription:       "devin_seo",
			SeoCode:              "devin_seo",
		},
	}
	res,err := productService.AddProduct(context.TODO(),productAdd)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}
