package model

const (
	//索引名
	PurchaseItemIndexName = "purchase"
	IdleItemIndexName     = "idle"
)

type mi = map[string]interface{}

var ItemIndexMapping = mi{
	"settings": mi{
		"number_of_shards":   3,
		"number_of_replicas": 1,
	},
	"mappings": mi{
		"dynamic": "false",
		"properties": mi{
			"ID": mi{ //整形字段, 允许精确匹配
				"type": "long",
			},
			"name": mi{
				"type":            "text",        //字符串类型且进行分词, 允许模糊匹配
				"analyzer":        "ik_max_word", //设置分词工具
				"search_analyzer": "ik_max_word",
				"fields": mi{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
					"keyword": mi{
						"type":         "keyword",
						"ignore_above": 256,
					},
				},
			},
			"description": mi{
				"type":            "text",        //字符串类型且进行分词, 允许模糊匹配
				"analyzer":        "ik_max_word", //设置分词工具
				"search_analyzer": "ik_max_word",
				"fields": mi{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
					"keyword": mi{
						"type":         "keyword",
						"ignore_above": 256,
					},
				},
			},
			"categoryId": mi{ //整形字段, 允许精确匹配
				"type": "long",
			},
			"uid": mi{ //整形字段, 允许精确匹配
				"type": "long",
			},
			"price": mi{ //整形字段, 允许精确匹配
				"type": "float",
			},
			"minPrice": mi{
				"type": "float",
			},
			"maxPrice": mi{
				"type": "float",
			},
			"tradeWay": mi{
				"type": "float",
			},
			"sort": mi{ //整形字段, 允许精确匹配
				"type": "integer",
			},
			"createdTime": mi{ //整形字段, 允许精确匹配
				"type": "long",
			},
			"publishStatus": mi{ //整形字段, 允许精确匹配
				"type": "integer",
			},
			"saleStatus": mi{ //整形字段, 允许精确匹配
				"type": "integer",
			},
			"address": mi{ //地址
				"type": "integer",
			},
			/*"date_field": mi{  //时间类型, 允许精确匹配
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
			},*/
		},
	},
}
