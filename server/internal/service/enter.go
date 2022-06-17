package service

import (
	"github.com/tmnhs/fginx/server/internal/service/file"
	"github.com/tmnhs/fginx/server/internal/service/kafka"
	"github.com/tmnhs/fginx/server/internal/service/redis"
	"github.com/tmnhs/fginx/server/internal/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	FileServiceGroup   file.FileServiceGroup
	RedisServiceGroup  redis.ServiceGroup
	KafkaServiceGroup  kafka.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
