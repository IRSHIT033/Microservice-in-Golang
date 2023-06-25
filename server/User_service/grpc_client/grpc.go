package grpc_client

import (
	"context"
	"log"
	"strconv"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/product_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

func GetProductViaGRPC(productId uint) domain_user.Product {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return domain_user.Product{}
	}
	defer conn.Close()
	client := product_proto.NewProductServiceClient(conn)
	// ctx, cancel := context.WithTimeout(, time.SecDial()
	// defer cancel()

	log.Println(strconv.FormatUint(uint64(productId), 10))

	response, err := client.SendProduct(context.Background(), &product_proto.ProductRequest{ProductId: strconv.FormatUint(uint64(productId), 10)})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return domain_user.Product{}
	}

	return domain_user.Product{
		Model:           gorm.Model{ID: uint(response.Productresp.ProductID)},
		ProductImageSrc: response.Productresp.ProductImageSrc,
		Name:            response.Productresp.Name,
		Category:        response.Productresp.Category,
		Description:     response.Productresp.Description,
		Price:           response.Productresp.Price,
		Unit:            int(response.Productresp.Unit),
		Available:       response.Productresp.Available,
	}
}
