package grpc_config

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/domain_product"
	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/product_proto"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ProductServer struct {
	product_proto.UnimplementedProductServiceServer
	database *gorm.DB
}

func (p *ProductServer) SendProduct(c context.Context, req *product_proto.ProductRequest) (*product_proto.ProductResponse, error) {

	log.Println(req.ProductId)

	productID, _ := strconv.Atoi(req.ProductId)
	var product domain_product.Product
	err := p.database.Find(&product, "product_id = ?", uint(productID)).Error

	if err != nil {
		res := &product_proto.ProductResponse{}
		return res, err
	}

	res := &product_proto.ProductResponse{

		Productresp: &product_proto.Product{
			ProductID:       uint32(product.ProductID),
			Price:           product.Price,
			ProductImageSrc: product.ProductImageSrc,
			Name:            product.Name,
			Description:     product.Description,
			Unit:            int32(product.Unit),
			AddedBy:         uint32(product.AddedBy),
			Available:       product.Available,
			Category:        product.Category,
		},
	}
	return res, nil
}

func GRPCListen(db *gorm.DB) {
	gRPCPort := 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()
	product_proto.RegisterProductServiceServer(s, &ProductServer{database: db})
	log.Printf("grpc server started on port %d", gRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC:%v", err)
	}
}
