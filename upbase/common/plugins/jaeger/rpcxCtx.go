/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 16:18
@Description:

*********************************************/
package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/smallnest/rpcx/share"
	"log"
	"context"
	"net/http"
)

// RPCX 框架完事胜grpc
//创建符合RPCX格式的且带SPAN信息的content
func RpcxSpanWithContext(ctx context.Context, operationName string, r *http.Request) (opentracing.Span, context.Context, error) {
	// 请求元数据 metaData
	md := ctx.Value(share.ReqMetaDataKey)
	var span opentracing.Span
	var parentSpan opentracing.Span

	tracer := opentracing.GlobalTracer()
	if md != nil {
		carrier := opentracing.TextMapCarrier(md.(map[string]string))
		spanContext, err := tracer.Extract(opentracing.TextMap, carrier)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Printf("metadata error %s\n", err)

		} else {
			parentSpan = tracer.StartSpan(operationName, ext.RPCServerOption(spanContext))
		}
	}

	if parentSpan != nil {
		span = parentSpan
	} else {
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil && spanCtx != nil {
			span = tracer.StartSpan( operationName, ext.RPCServerOption(spanCtx) )
		}else{
			span = opentracing.StartSpan(operationName)
		}

	}

	metadata := opentracing.TextMapCarrier(make(map[string]string))
	err := tracer.Inject(span.Context(), opentracing.TextMap, metadata)
	if err != nil {
		return nil, nil, err
	}
	ctx = context.WithValue(context.Background(), share.ReqMetaDataKey, (map[string]string)(metadata))
	return span, ctx, nil

}
