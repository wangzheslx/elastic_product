package data

import (
	"product/internal/conf"
	"testing"
)

func TestNewElasticsearch(t *testing.T) {
	// 模拟配置数据
	c := &conf.Data{
		Elastic: &conf.Data_Elastic{
			Addr: "http://192.168.223.128:9200",
		},
	}

	client := NewElasticsearch(c)

	if client == nil {
		t.Errorf("Expected client to be non-nil, but got nil")
	}

	// 断言客户端不为空

	// 可以进一步添加其他断言和测试逻辑
}
