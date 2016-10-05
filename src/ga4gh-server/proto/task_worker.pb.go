// Code generated by protoc-gen-go.
// source: task_worker.proto
// DO NOT EDIT!

/*
Package ga4gh_task_ref is a generated protocol buffer package.

It is generated from these files:
	task_worker.proto

It has these top-level messages:
	WorkerInfo
	JobRequest
	JobResponse
	UpdateStatusRequest
	QueuedTaskInfoRequest
	QueuedTaskInfo
*/
package ga4gh_task_ref

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ga4gh_task_exec "ga4gh-tasks"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Worker Info
type WorkerInfo struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LastPing int64  `protobuf:"varint,2,opt,name=last_ping" json:"last_ping,omitempty"`
	Hostname string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *WorkerInfo) Reset()                    { *m = WorkerInfo{} }
func (m *WorkerInfo) String() string            { return proto.CompactTextString(m) }
func (*WorkerInfo) ProtoMessage()               {}
func (*WorkerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type JobRequest struct {
	Worker    *WorkerInfo                `protobuf:"bytes,1,opt,name=worker" json:"worker,omitempty"`
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,2,opt,name=resources" json:"resources,omitempty"`
}

func (m *JobRequest) Reset()                    { *m = JobRequest{} }
func (m *JobRequest) String() string            { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()               {}
func (*JobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobRequest) GetWorker() *WorkerInfo {
	if m != nil {
		return m.Worker
	}
	return nil
}

func (m *JobRequest) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

type JobResponse struct {
	Job *ga4gh_task_exec.Job `protobuf:"bytes,2,opt,name=job" json:"job,omitempty"`
}

func (m *JobResponse) Reset()                    { *m = JobResponse{} }
func (m *JobResponse) String() string            { return proto.CompactTextString(m) }
func (*JobResponse) ProtoMessage()               {}
func (*JobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobResponse) GetJob() *ga4gh_task_exec.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type UpdateStatusRequest struct {
	Id       string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Step     int64                   `protobuf:"varint,2,opt,name=step" json:"step,omitempty"`
	State    ga4gh_task_exec.State   `protobuf:"varint,3,opt,name=state,enum=ga4gh_task_exec.State" json:"state,omitempty"`
	Log      *ga4gh_task_exec.JobLog `protobuf:"bytes,4,opt,name=log" json:"log,omitempty"`
	WorkerId string                  `protobuf:"bytes,5,opt,name=worker_id" json:"worker_id,omitempty"`
}

func (m *UpdateStatusRequest) Reset()                    { *m = UpdateStatusRequest{} }
func (m *UpdateStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateStatusRequest) ProtoMessage()               {}
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateStatusRequest) GetLog() *ga4gh_task_exec.JobLog {
	if m != nil {
		return m.Log
	}
	return nil
}

type QueuedTaskInfoRequest struct {
	MaxTasks int32 `protobuf:"varint,1,opt,name=max_tasks" json:"max_tasks,omitempty"`
}

func (m *QueuedTaskInfoRequest) Reset()                    { *m = QueuedTaskInfoRequest{} }
func (m *QueuedTaskInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*QueuedTaskInfoRequest) ProtoMessage()               {}
func (*QueuedTaskInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type QueuedTaskInfo struct {
	Inputs    []string                   `protobuf:"bytes,1,rep,name=inputs" json:"inputs,omitempty"`
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,2,opt,name=resources" json:"resources,omitempty"`
}

func (m *QueuedTaskInfo) Reset()                    { *m = QueuedTaskInfo{} }
func (m *QueuedTaskInfo) String() string            { return proto.CompactTextString(m) }
func (*QueuedTaskInfo) ProtoMessage()               {}
func (*QueuedTaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueuedTaskInfo) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkerInfo)(nil), "ga4gh_task_ref.WorkerInfo")
	proto.RegisterType((*JobRequest)(nil), "ga4gh_task_ref.JobRequest")
	proto.RegisterType((*JobResponse)(nil), "ga4gh_task_ref.JobResponse")
	proto.RegisterType((*UpdateStatusRequest)(nil), "ga4gh_task_ref.UpdateStatusRequest")
	proto.RegisterType((*QueuedTaskInfoRequest)(nil), "ga4gh_task_ref.QueuedTaskInfoRequest")
	proto.RegisterType((*QueuedTaskInfo)(nil), "ga4gh_task_ref.QueuedTaskInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Scheduler service

type SchedulerClient interface {
	GetQueueInfo(ctx context.Context, in *QueuedTaskInfoRequest, opts ...grpc.CallOption) (Scheduler_GetQueueInfoClient, error)
	GetJobToRun(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error)
	UpdateJobStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*ga4gh_task_exec.JobId, error)
	WorkerPing(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*WorkerInfo, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetQueueInfo(ctx context.Context, in *QueuedTaskInfoRequest, opts ...grpc.CallOption) (Scheduler_GetQueueInfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Scheduler_serviceDesc.Streams[0], c.cc, "/ga4gh_task_ref.Scheduler/GetQueueInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerGetQueueInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_GetQueueInfoClient interface {
	Recv() (*QueuedTaskInfo, error)
	grpc.ClientStream
}

type schedulerGetQueueInfoClient struct {
	grpc.ClientStream
}

func (x *schedulerGetQueueInfoClient) Recv() (*QueuedTaskInfo, error) {
	m := new(QueuedTaskInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) GetJobToRun(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/GetJobToRun", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UpdateJobStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*ga4gh_task_exec.JobId, error) {
	out := new(ga4gh_task_exec.JobId)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/UpdateJobStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) WorkerPing(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*WorkerInfo, error) {
	out := new(WorkerInfo)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/WorkerPing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	GetQueueInfo(*QueuedTaskInfoRequest, Scheduler_GetQueueInfoServer) error
	GetJobToRun(context.Context, *JobRequest) (*JobResponse, error)
	UpdateJobStatus(context.Context, *UpdateStatusRequest) (*ga4gh_task_exec.JobId, error)
	WorkerPing(context.Context, *WorkerInfo) (*WorkerInfo, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_GetQueueInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueuedTaskInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).GetQueueInfo(m, &schedulerGetQueueInfoServer{stream})
}

type Scheduler_GetQueueInfoServer interface {
	Send(*QueuedTaskInfo) error
	grpc.ServerStream
}

type schedulerGetQueueInfoServer struct {
	grpc.ServerStream
}

func (x *schedulerGetQueueInfoServer) Send(m *QueuedTaskInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Scheduler_GetJobToRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetJobToRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/GetJobToRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetJobToRun(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UpdateJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UpdateJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/UpdateJobStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UpdateJobStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_WorkerPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).WorkerPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/WorkerPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).WorkerPing(ctx, req.(*WorkerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ga4gh_task_ref.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobToRun",
			Handler:    _Scheduler_GetJobToRun_Handler,
		},
		{
			MethodName: "UpdateJobStatus",
			Handler:    _Scheduler_UpdateJobStatus_Handler,
		},
		{
			MethodName: "WorkerPing",
			Handler:    _Scheduler_WorkerPing_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetQueueInfo",
			Handler:       _Scheduler_GetQueueInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("task_worker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5f, 0xef, 0xd2, 0x30,
	0x14, 0x65, 0x0c, 0x88, 0xbb, 0x90, 0x29, 0x15, 0x71, 0x99, 0x89, 0xd1, 0x2a, 0x89, 0x21, 0x71,
	0x21, 0xe8, 0xab, 0xaf, 0x2a, 0xc6, 0x44, 0x04, 0x8c, 0xf1, 0x69, 0x19, 0xec, 0x32, 0x26, 0xb0,
	0xce, 0xb5, 0x8d, 0x7c, 0x0b, 0xbf, 0x97, 0x9f, 0xca, 0xae, 0x63, 0x91, 0x3f, 0x0b, 0xc9, 0xef,
	0x6d, 0x69, 0xcf, 0x3d, 0xe7, 0xf4, 0x9c, 0x3b, 0xe8, 0x8a, 0x80, 0x6f, 0xfd, 0xdf, 0x2c, 0xdb,
	0x62, 0xe6, 0xa5, 0x19, 0x13, 0x8c, 0xd8, 0x51, 0xf0, 0x36, 0xda, 0xf8, 0xfa, 0x22, 0xc3, 0xb5,
	0xdb, 0xd3, 0x5f, 0x78, 0xc0, 0x95, 0x14, 0x31, 0x4b, 0x0a, 0x14, 0x7d, 0x07, 0xf0, 0x5d, 0x4f,
	0x4d, 0x92, 0x35, 0x23, 0x00, 0xf5, 0x38, 0x74, 0x8c, 0x67, 0xc6, 0x2b, 0x8b, 0x74, 0xc1, 0xda,
	0x05, 0x5c, 0xf8, 0x69, 0x9c, 0x44, 0x4e, 0x5d, 0x1d, 0x99, 0xe4, 0x01, 0xdc, 0xdb, 0x30, 0x2e,
	0x92, 0x60, 0x8f, 0x8e, 0x99, 0x83, 0x68, 0x04, 0xf0, 0x89, 0x2d, 0x67, 0xf8, 0x4b, 0x22, 0x17,
	0x64, 0x08, 0xad, 0xc2, 0x82, 0xa6, 0x68, 0x8f, 0x5d, 0xef, 0xdc, 0x83, 0x77, 0x22, 0xf5, 0x1a,
	0xac, 0x0c, 0x39, 0x93, 0xd9, 0x0a, 0xb9, 0xa6, 0xbf, 0x80, 0xe7, 0x46, 0xbd, 0x59, 0x89, 0xa0,
	0x23, 0x68, 0x6b, 0x21, 0x9e, 0xb2, 0x84, 0x23, 0x79, 0x0e, 0xe6, 0x4f, 0xb6, 0x3c, 0xce, 0xf5,
	0xae, 0xe6, 0x14, 0x94, 0xfe, 0x31, 0xe0, 0xe1, 0xb7, 0x34, 0x0c, 0x04, 0xce, 0x45, 0x20, 0x24,
	0x2f, 0x4d, 0x9e, 0xbe, 0xb1, 0x03, 0x0d, 0x2e, 0x30, 0x3d, 0x3e, 0x6f, 0x00, 0x4d, 0xae, 0xa0,
	0xc5, 0xdb, 0xec, 0x71, 0xff, 0x8a, 0x36, 0x27, 0x42, 0xf2, 0x12, 0xcc, 0x1d, 0x8b, 0x9c, 0x86,
	0xd6, 0x7e, 0x5c, 0xa5, 0xfd, 0x99, 0x45, 0x79, 0x7c, 0x45, 0x16, 0xbe, 0x52, 0x6b, 0xea, 0xb0,
	0x86, 0xf0, 0xe8, 0xab, 0x44, 0x89, 0xe1, 0x42, 0x81, 0xf3, 0x10, 0x4a, 0x4b, 0x0a, 0xbb, 0x0f,
	0x0e, 0x9a, 0x83, 0x6b, 0x67, 0x4d, 0xfa, 0x05, 0xec, 0x73, 0x2c, 0xb1, 0xa1, 0x15, 0x27, 0xa9,
	0x14, 0x39, 0xc2, 0x54, 0xde, 0xef, 0x16, 0xe0, 0xf8, 0x6f, 0x1d, 0xac, 0xf9, 0x6a, 0x83, 0xa1,
	0xdc, 0x61, 0x46, 0x7e, 0x40, 0xe7, 0x03, 0x0a, 0xad, 0xa0, 0xc9, 0x07, 0x97, 0x4d, 0x55, 0x1a,
	0x75, 0x9f, 0xde, 0x86, 0xd1, 0xda, 0xc8, 0x20, 0x1f, 0xa1, 0xad, 0xa8, 0x55, 0x0a, 0x0b, 0x36,
	0x93, 0x09, 0xb9, 0xda, 0x81, 0xff, 0xfb, 0xe2, 0x3e, 0xa9, 0xbc, 0x2b, 0x2a, 0xa6, 0x35, 0x32,
	0x85, 0xfb, 0x45, 0x81, 0xea, 0xb8, 0xe8, 0x90, 0xbc, 0xb8, 0x9c, 0xa8, 0x68, 0xd8, 0xed, 0x57,
	0x75, 0x32, 0x09, 0x15, 0xe3, 0xfb, 0x72, 0xdb, 0xa7, 0x6a, 0xa9, 0xc9, 0x8d, 0xf5, 0x74, 0x6f,
	0xdc, 0xd1, 0xda, 0xb2, 0xa5, 0x7f, 0x9e, 0x37, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x48, 0x8d,
	0x10, 0x54, 0x77, 0x03, 0x00, 0x00,
}
