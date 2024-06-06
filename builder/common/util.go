package common

import (
	"context"
	"fmt"

	"github.com/wangweihong/packer-plugin-hci/hci/example"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/wangweihong/gotoolbox/pkg/netutil"
)

func FindAvailableIP(instanceCIDR string) (string, error) {
	ipNet, err := netutil.ValidateCIDR(instanceCIDR)
	if err != nil {
		return "", err
	}

	for _, ip := range netutil.GenerateIPs(ipNet) {
		if !netutil.Ping(ip, 5) {
			return ip, nil
		}
	}
	return "", fmt.Errorf("no valid ip")
}

func GetImageRepository(client *example.Client, cluster string, tenant string, name string) (string, error) {
	response, err := example.NewRepository(client).RepositoryList(context.Background(), &iexample.RepositoryListRequest{
		Cluster:     cluster,
		Tenant:      tenant,
		FilterUsage: "image",
		FilterName:  name,
	})
	if err != nil {
		return "", fmt.Errorf("error getting repository, err=%s", err)
	}
	for _, v := range response.Data.List {
		if v.Name == name {
			return v.UUID, nil
		}
	}

	return "", fmt.Errorf("repository %v not found", name)
}

func InspectServer(c *example.Client, cluster, tenant, server string) (*iexample.EcsInspectResponse, error) {
	resp, err := example.NewCompute(c).EcsInspect(context.Background(), &iexample.EcsInspectRequest{
		Cluster: cluster,
		Tenant:  tenant,
		ECS:     server,
		// get vnc password
		ShowPlaintext: true,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ShutoffServer(c *example.Client, cluster, tenant, server string) error {
	_, err := example.NewCompute(c).EcsShutoff(context.Background(), &iexample.EcsShutoffRequest{
		Cluster: cluster,
		Tenant:  tenant,
		ECS:     server,
	})
	return err
}

func GetImage(client *example.Client, cluster string, tenant string, repository, name, imageType string) (string, error) {
	response, err := example.NewRepository(client).RepositoryImageGet(context.Background(), &iexample.RepositoryImageGetRequest{
		Cluster:    cluster,
		Tenant:     tenant,
		Name:       name + "." + imageType,
		Repository: repository,
	})
	if err != nil {
		return "", fmt.Errorf("error getting image %v in %v , err=%s", name, repository, err)
	}

	return response.Data.Image.UUID, nil
}

func DeleteImage(client *example.Client, cluster string, tenant string, repository, image string) error {
	_, err := example.NewRepository(client).RepositoryImageDelete(context.Background(), &iexample.RepositoryImageDeleteRequest{
		Cluster:    cluster,
		Tenant:     tenant,
		UUID:       image,
		Repository: repository,
	})
	if err != nil {
		return err
	}

	return nil
}

func SyncImage(client *example.Client, cluster string, tenant string, repository string) error {
	_, err := example.NewRepository(client).RepositoryImageSync(context.Background(), &iexample.RepositoryImageSyncRequest{
		Cluster: cluster,
		Tenant:  tenant,
		Image:   struct{ Repository string }{Repository: repository},
	})
	if err != nil {
		return err
	}

	return nil
}

func IsExporting(client *example.Client, cluster string, tenant string, server string) (bool, error) {
	rawResp, err := example.NewCompute(client).EcsDiskExportProgress(context.Background(), &iexample.EcsDiskExportProgressRequest{
		Cluster: cluster,
		Tenant:  tenant,
		ECS:     server,
	})
	if err != nil {
		if rawResp != nil && rawResp.ErrorCode == 22151 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
