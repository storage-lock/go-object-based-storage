package object_based_storage

import (
	"context"
	"time"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// ObjectBasedStorage 通用对象存储实现 (S3兼容)
type ObjectBasedStorage struct {
	*s3_storage.S3Storage
}

func NewObjectBasedStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*ObjectBasedStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &ObjectBasedStorage{S3Storage: s3Storage}, nil
}

const StorageName = "object-based-storage"

func (x *ObjectBasedStorage) GetName() string {
	return StorageName
}

func (x *ObjectBasedStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}

func (x *ObjectBasedStorage) Watch(ctx context.Context, lockId string) (<-chan *storage.LockInformation, error) {
	return storage.NewPollingLockWatcher(x, 500*time.Millisecond).Watch(ctx, lockId)
}
