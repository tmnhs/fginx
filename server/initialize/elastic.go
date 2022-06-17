package initialize

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

//初始化
func Elastic() {

	host := fmt.Sprintf("http://%s:%d/", global.GV_CONFIG.Elastic.Path, global.GV_CONFIG.Elastic.Port)

	options := []elastic.ClientOptionFunc{
		elastic.SetURL(host),
		elastic.SetSniff(false),                                            //是否开启集群嗅探
		elastic.SetHealthcheckInterval(30 * time.Second),                   //设置两次运行状况检查之间的间隔, 默认60s
		elastic.SetGzip(false),                                             //启用或禁用gzip压缩
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)), //ERROR日志输出配置
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),          //INFO级别日志输出配置
	}
	options = append(options, elastic.SetBasicAuth(
		global.GV_CONFIG.Elastic.UserName, //账号
		global.GV_CONFIG.Elastic.Password, //密码
	))

	var err error

	global.GV_ES, err = elastic.NewClient(options...)

	if err != nil {
		global.GV_LOG.Error("Elastic err:", zap.Error(err))
		return
	}
	global.GV_LOG.Debug("Elasticsearch running on :", zap.String("elastic", host))

	info, code, err := global.GV_ES.Ping(host).Do(context.Background())

	if err != nil {
		global.GV_LOG.Error("Elastic err:", zap.Error(err))
		return
	}
	global.GV_LOG.Debug("Elasticsearch returned with code and version \n", zap.Int("elastic", code), zap.String("elastic", info.Version.Number))

	esversion, err := global.GV_ES.ElasticsearchVersion(host)
	if err != nil {
		global.GV_LOG.Error("Elastic err:", zap.Error(err))
		return
	}
	global.GV_LOG.Debug("Elasticsearch version \n", zap.String("elastic", esversion))
	//创建索引
	createIndex(model.PurchaseItemIndexName)
	createIndex(model.IdleItemIndexName)
}

func createIndex(indexName string) {
	exists, err := global.GV_ES.IndexExists(indexName).Do(context.Background())
	if err != nil {
		global.GV_LOG.Error("find index err:", zap.Error(err))
		return
	}
	if exists == true {
		global.GV_LOG.Debug(" index  exists")
		return
	}
	_, err = global.GV_ES.CreateIndex(indexName).BodyJson(model.ItemIndexMapping).Do(context.Background())
	if err != nil {
		global.GV_LOG.Error("createIndex err:", zap.Error(err))
		return
	}
	global.GV_LOG.Debug("create index success ,index:", zap.String("elastic", indexName))
}
