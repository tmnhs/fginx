package initialize_test

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/tmnhs/fginx/server/global"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

var client *elastic.Client

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func TestElastic(t *testing.T) {

	global.GV_CONFIG.Elastic.Path = "114.55.178.217"
	global.GV_CONFIG.Elastic.Port = 9200
	global.GV_CONFIG.Elastic.UserName = "elastic"
	global.GV_CONFIG.Elastic.Password = "020821mnh"
	fmt.Println(global.GV_CONFIG.Elastic)
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
	client, err = elastic.NewClient(options...)
	if err != nil {
		t.Error("new client err", err)
	}
	//t.Log("client is ",client)

	//index()
	//create()
	//delete()
	//find()
	deleteIndex()
}

func index() {
	type mi = map[string]interface{}
	mapping := mi{
		"settings": mi{
			"number_of_shards":   3,
			"number_of_replicas": 2,
		},
		"mappings": mi{
			"properties": mi{
				"id": mi{ //整形字段, 允许精确匹配
					"type": "integer",
				},
				"name": mi{
					"type":            "text",     //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        "ik_smart", //设置分词工具
					"search_analyzer": "ik_smart",
					"fields": mi{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": mi{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"date_field": mi{ //时间类型, 允许精确匹配
					"type": "date",
				},
				"keyword_field": mi{ //字符串类型, 允许精确匹配
					"type": "keyword",
				},
				"nested_field": mi{ //嵌套类型
					"type": "nested",
					"properties": mi{
						"id": mi{
							"type": "integer",
						},
						"start_time": mi{ //长整型, 允许精确匹配
							"type": "long",
						},
						"end_time": mi{
							"type": "long",
						},
					},
				},
			},
		},
	}
	indexName := "test" //要创建的索引名
	_, err := client.CreateIndex(indexName).BodyJson(mapping).Do(context.Background())
	if err != nil {
		fmt.Println("client ", err)
		return
	}
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		fmt.Println("client ", err)
	}
	fmt.Println("the index exists:", exists)
}
func deleteIndex() {
	res, err := client.DeleteIndex("idle").Do(context.Background())
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("res:", res)
	res, err = client.DeleteIndex("purchase").Do(context.Background())
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("res:", res)
}
func create() {
	e1 := Employee{"Jane", "Jack", 30, "I like to collect rock albums", []string{"music"}}
	_, err := client.Index().Index("test").Id("1").BodyJson(e1).Do(context.Background())
	if err != nil {
		fmt.Println("create err:", err)
	}
}

func delete() {
	res, err := client.Delete().Index("test").Id("2").Do(context.Background())
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("res:", res)
}

func find() {
	//res,err:=client.Get().Index("test").Id("3").Do(context.Background())
	//if err != nil {
	//	fmt.Println("err:",err)
	//}

	//if res.Found {
	//	fmt.Printf("Got document %s in version %d from index %s, type %s\n", res.Id, res.Version, res.Index, res.Type)
	//	fmt.Println( res.Source)
	//}
	//res, err := client.Search("test").Do(context.Background())
	//printEmployee(res,err)
	//字段相等
	//q := elastic.NewQueryStringQuery("last_name:Jack")
	//res, err := client.Search("test").Query(q).Do(context.Background())
	//printEmployee(res,err)

	//条件查询
	//年龄大于30岁的
	//boolQ := elastic.NewBoolQuery()
	////boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	//boolQ.Filter(elastic.NewRangeQuery("age").Gt(25))
	//res, err := client.Search("test").Query(boolQ).Do(context.Background())
	//printEmployee(res, err)
	//短语搜索 搜索about字段中有 rock climbing
	//matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock")
	//res, err := client.Search("test").Query(matchPhraseQuery).Do(context.Background())
	//printEmployee(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err := client.Search("test").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}
