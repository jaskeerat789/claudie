package kuber

import (
	"context"
	"fmt"

	"github.com/Berops/claudie/proto/pb"
)

func SetUpStorage(c pb.KuberServiceClient, req *pb.SetUpStorageRequest) (*pb.SetUpStorageResponse, error) {
	res, err := c.SetUpStorage(context.Background(), req) //sending request to the server and receiving response
	if err != nil {
		return nil, fmt.Errorf("error while calling SetUpStorage on Kuber: %w", err)
	}
	return res, nil
}

func StoreKubeconfig(c pb.KuberServiceClient, req *pb.StoreKubeconfigRequest) (*pb.StoreKubeconfigResponse, error) {
	res, err := c.StoreKubeconfig(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error while calling StoreKubeconfig on kuber: %w", err)
	}
	return res, nil
}

func DeleteKubeconfig(c pb.KuberServiceClient, req *pb.DeleteKubeconfigRequest) (*pb.DeleteKubeconfigResponse, error) {
	res, err := c.DeleteKubeconfig(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error while calling DeleteKubeconfig on kuber: %w", err)
	}
	return res, nil
}

func DeleteNodes(c pb.KuberServiceClient, req *pb.DeleteNodesRequest) (*pb.DeleteNodesResponse, error) {
	res, err := c.DeleteNodes(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error while calling DeleteNodes on kuber: %w", err)
	}
	return res, nil
}

func StoreClusterMetadata(c pb.KuberServiceClient, req *pb.StoreClusterMetadataRequest) (*pb.StoreClusterMetadataResponse, error) {
	res, err := c.StoreClusterMetadata(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error while calling StoreClusterMetadata: %w", err)
	}

	return res, nil
}

func DeleteClusterMetadata(c pb.KuberServiceClient, req *pb.DeleteClusterMetadataRequest) (*pb.DeleteClusterMetadataResponse, error) {
	res, err := c.DeleteClusterMetadata(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error while calling DeleteClusterMetadata: %w", err)
	}

	return res, nil
}
