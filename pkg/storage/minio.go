package storage

import (
	"net/http"
	"sync"
	"warrant-api/pkg/config"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/utils"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var pool sync.Pool

func Init() {
	pool = sync.Pool{
		New: func() any {
			client, err := minio.New(config.Storage.Endpoint, &minio.Options{
				Creds:  credentials.NewStaticV4(config.Storage.AccessKeyID, config.Storage.SecretAccessKey, ""),
				Secure: config.Storage.UseSsl,
			})
			if err != nil {
				utils.Handle(messages.Errorf(http.StatusServiceUnavailable, "Storage system failure."))
			}
			return client
		},
	}
}

/*
func PutTransportCost(g *gin.Context, transportCost *model.TransportCost) error {
	if transportCost.FileName.Valid {
		url, err := presignedPut(g, transportCost.WarrantID+"/costs/"+transportCost.ID+"."+transportCost.FileName.String)
		if err == nil && url != nil {
			transportCost.Url = url.String()
		}
	}
	return nil
}
func GetTransportCost(g *gin.Context, transportCost *model.TransportCost) error {
	if transportCost.FileName.Valid {
		url, err := presignedGet(g, transportCost.WarrantID+"/costs/"+transportCost.ID+"."+transportCost.FileName.String)
		if err == nil && url != nil {
			transportCost.Url = url.String()
		}
	}
	return nil
}
func PutRepair(g *gin.Context, repair *model.Repair) error {
	if repair.FileName.Valid {
		url, err := presignedPut(g, repair.WarrantID+"/repair/"+repair.ID+"."+repair.FileName.String)
		if err == nil && url != nil {
			repair.Url = url.String()
		}
	}
	return nil
}
func GetRepair(g *gin.Context, repair *model.Repair) error {
	if repair.FileName.Valid {
		url, err := presignedGet(g, repair.WarrantID+"/repair/"+repair.ID+"."+repair.FileName.String)
		if err == nil && url != nil {
			repair.Url = url.String()
		}
	}
	return nil
}

func UpdateRepair(g *gin.Context, repair *model.Repair, oldName string) error {
	mc := pool.Get().(*minio.Client)
	defer pool.Put(mc)
	name := repair.WarrantID + "/repair/" + repair.ID + "." + oldName
	// Source object
	src := minio.CopySrcOptions{
		Bucket: config.Storage.Bucket,
		Object: name,
	}
	// Destination object
	dst := minio.CopyDestOptions{
		Bucket:      config.Storage.Bucket,
		Object:      name + "." + repo.GenUUID.Next().String(),
		ReplaceTags: true,
		UserTags: map[string]string{
			"status": "deleted",
		},
	}
	_, err := mc.CopyObject(context.Background(), dst, src)

	if err != nil {
		return err
	}
	url, err := mc.PresignedPutObject(context.Background(), config.Storage.Bucket, repair.FileName.String, 10*time.Second)
	if err != nil {
		return err
	}
	repair.Url = url.String()
	return nil

}

func UpdateTransportConst(g *gin.Context, costs *model.TransportCost, oldName string) error {
	mc := pool.Get().(*minio.Client)
	defer pool.Put(mc)
	name := costs.WarrantID + "/costs/" + costs.ID + "." + oldName
	// Source object
	src := minio.CopySrcOptions{
		Bucket: config.Storage.Bucket,
		Object: name,
	}
	// Destination object
	dst := minio.CopyDestOptions{
		Bucket:      config.Storage.Bucket,
		Object:      name + "." + repo.GenUUID.Next().String(),
		ReplaceTags: true,
		UserTags: map[string]string{
			"status": "deleted",
		},
	}
	_, err := mc.CopyObject(context.Background(), dst, src)
	if err != nil {
		return err
	}
	url, err := mc.PresignedPutObject(context.Background(), config.Storage.Bucket, costs.FileName.String, 10*time.Second)
	if err != nil {
		return err
	}
	costs.Url = url.String()
	return nil
}

func presignedPut(g *gin.Context, name string) (*url.URL, error) {
	mc := pool.Get().(*minio.Client)
	defer pool.Put(mc)

	return mc.PresignedPutObject(context.Background(), config.Storage.Bucket, name, 10*time.Second)
}
func presignedGet(g *gin.Context, name string) (*url.URL, error) {
	mc := pool.Get().(*minio.Client)
	defer pool.Put(mc)
	params := url.Values{}
	return mc.PresignedGetObject(context.Background(), config.Storage.Bucket, name, 10*time.Second, params)
}

// func Update(g *gin.Context, name string) (*url.URL, error) {
// 	mc := pool.Get().(*minio.Client)
// 	defer pool.Put(mc)
// 	// Source object
// 	src := minio.CopySrcOptions{
// 		Bucket: config.Storage.Bucket,
// 		Object: name,
// 	}
// 	// Destination object
// 	dst := minio.CopyDestOptions{
// 		Bucket:      config.Storage.Bucket,
// 		Object:      name + "." + repo.GenUUID.Next().String(),
// 		ReplaceTags: true,
// 		UserTags: map[string]string{
// 			"status": "deleted",
// 		},
// 	}
// 	_, err := mc.CopyObject(context.Background(), dst, src)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return mc.PresignedPutObject(context.Background(), config.Storage.Bucket, name, 10*time.Second)

// }

*/
