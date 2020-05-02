package grpc

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrInvalidRequestType = errors.New("RequestType has only four type: Add,Subtract,Multiply,Divide")
)

// Service Define a service interface
type Service interface {
	Calculate(ctx context.Context, reqType string, a, b int) (int, error)

	// HealthCheck check service health status
	//HealthCheck() bool
}

//ArithmeticService implement Service interface
type CalculateService struct {
}

// Calculate 实现Service接口
func (s CalculateService) Calculate(_ context.Context, reqType string, a, b int) (res int, err error) {

	if strings.EqualFold(reqType, "Add") {
		res = a + b
		return
	} else if strings.EqualFold(reqType, "Subtract") {
		res = a - b
		return
	} else if strings.EqualFold(reqType, "Multiply") {
		res = a * b
		return
	} else if strings.EqualFold(reqType, "Divide") {
		if b == 0 {
			res, err = 0, errors.New("the dividend can not be zero!")
			return
		}
		res, err = a/b, nil
	} else {
		res, err = 0, ErrInvalidRequestType
	}
	return
}

// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true。
//func (s CalculateService) HealthCheck() bool {
//	return true
//}

//type HealthImpl struct{}
//// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
//// https://github.com/hashicorp/consul/blob/master/agent/checks/grpc.go
//// consul 检查服务器的健康状态，consul 用 google.golang.org/grpc/health/grpc_health_v1.HealthServer 接口，实现了对 grpc健康检查的支持，所以我们只需要实现先这个接口，consul 就能利用这个接口作健康检查了
//func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
//	return &grpc_health_v1.HealthCheckResponse{
//		Status: grpc_health_v1.HealthCheckResponse_SERVING,
//	}, nil
//}
//// Watch HealthServer interface 有两个方法
//// Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
//// Watch(*HealthCheckRequest, Health_WatchServer) error
//// 所以在 HealthImpl 结构提不仅要实现 Check 方法，还要实现 Watch 方法
//func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
//	return nil
//}
