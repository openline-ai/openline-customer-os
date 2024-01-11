// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: organization.proto

package organization_grpc_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrganizationGrpcServiceClient is the client API for OrganizationGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationGrpcServiceClient interface {
	UpsertOrganization(ctx context.Context, in *UpsertOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	WebScrapeOrganization(ctx context.Context, in *WebScrapeOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	LinkPhoneNumberToOrganization(ctx context.Context, in *LinkPhoneNumberToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	LinkEmailToOrganization(ctx context.Context, in *LinkEmailToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	LinkLocationToOrganization(ctx context.Context, in *LinkLocationToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	LinkDomainToOrganization(ctx context.Context, in *LinkDomainToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	UpsertCustomFieldToOrganization(ctx context.Context, in *CustomFieldForOrganizationGrpcRequest, opts ...grpc.CallOption) (*CustomFieldIdGrpcResponse, error)
	HideOrganization(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	ShowOrganization(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	RefreshLastTouchpoint(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	RefreshRenewalSummary(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	RefreshArr(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	AddParentOrganization(ctx context.Context, in *AddParentOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	RemoveParentOrganization(ctx context.Context, in *RemoveParentOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	UpdateOnboardingStatus(ctx context.Context, in *UpdateOnboardingStatusGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	UpdateOrganization(ctx context.Context, in *UpdateOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	AddSocial(ctx context.Context, in *AddSocialGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	UpdateOrganizationOwner(ctx context.Context, in *UpdateOrganizationOwnerGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	CreateBillingProfile(ctx context.Context, in *CreateBillingProfileGrpcRequest, opts ...grpc.CallOption) (*BillingProfileIdGrpcResponse, error)
	LinkEmailToBillingProfile(ctx context.Context, in *LinkEmailToBillingProfileGrpcRequest, opts ...grpc.CallOption) (*BillingProfileIdGrpcResponse, error)
	UnlinkEmailFromBillingProfile(ctx context.Context, in *UnlinkEmailFromBillingProfileGrpcRequest, opts ...grpc.CallOption) (*BillingProfileIdGrpcResponse, error)
}

type organizationGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationGrpcServiceClient(cc grpc.ClientConnInterface) OrganizationGrpcServiceClient {
	return &organizationGrpcServiceClient{cc}
}

func (c *organizationGrpcServiceClient) UpsertOrganization(ctx context.Context, in *UpsertOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/UpsertOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) WebScrapeOrganization(ctx context.Context, in *WebScrapeOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/WebScrapeOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) LinkPhoneNumberToOrganization(ctx context.Context, in *LinkPhoneNumberToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/LinkPhoneNumberToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) LinkEmailToOrganization(ctx context.Context, in *LinkEmailToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/LinkEmailToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) LinkLocationToOrganization(ctx context.Context, in *LinkLocationToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/LinkLocationToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) LinkDomainToOrganization(ctx context.Context, in *LinkDomainToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/LinkDomainToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) UpsertCustomFieldToOrganization(ctx context.Context, in *CustomFieldForOrganizationGrpcRequest, opts ...grpc.CallOption) (*CustomFieldIdGrpcResponse, error) {
	out := new(CustomFieldIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/UpsertCustomFieldToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) HideOrganization(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/HideOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) ShowOrganization(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/ShowOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) RefreshLastTouchpoint(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/RefreshLastTouchpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) RefreshRenewalSummary(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/RefreshRenewalSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) RefreshArr(ctx context.Context, in *OrganizationIdGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/RefreshArr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) AddParentOrganization(ctx context.Context, in *AddParentOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/AddParentOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) RemoveParentOrganization(ctx context.Context, in *RemoveParentOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/RemoveParentOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) UpdateOnboardingStatus(ctx context.Context, in *UpdateOnboardingStatusGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/UpdateOnboardingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) UpdateOrganization(ctx context.Context, in *UpdateOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/UpdateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) AddSocial(ctx context.Context, in *AddSocialGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/AddSocial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) UpdateOrganizationOwner(ctx context.Context, in *UpdateOrganizationOwnerGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/UpdateOrganizationOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) CreateBillingProfile(ctx context.Context, in *CreateBillingProfileGrpcRequest, opts ...grpc.CallOption) (*BillingProfileIdGrpcResponse, error) {
	out := new(BillingProfileIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/CreateBillingProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) LinkEmailToBillingProfile(ctx context.Context, in *LinkEmailToBillingProfileGrpcRequest, opts ...grpc.CallOption) (*BillingProfileIdGrpcResponse, error) {
	out := new(BillingProfileIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/LinkEmailToBillingProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) UnlinkEmailFromBillingProfile(ctx context.Context, in *UnlinkEmailFromBillingProfileGrpcRequest, opts ...grpc.CallOption) (*BillingProfileIdGrpcResponse, error) {
	out := new(BillingProfileIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/UnlinkEmailFromBillingProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationGrpcServiceServer is the server API for OrganizationGrpcService service.
// All implementations should embed UnimplementedOrganizationGrpcServiceServer
// for forward compatibility
type OrganizationGrpcServiceServer interface {
	UpsertOrganization(context.Context, *UpsertOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	WebScrapeOrganization(context.Context, *WebScrapeOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	LinkPhoneNumberToOrganization(context.Context, *LinkPhoneNumberToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	LinkEmailToOrganization(context.Context, *LinkEmailToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	LinkLocationToOrganization(context.Context, *LinkLocationToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	LinkDomainToOrganization(context.Context, *LinkDomainToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	UpsertCustomFieldToOrganization(context.Context, *CustomFieldForOrganizationGrpcRequest) (*CustomFieldIdGrpcResponse, error)
	HideOrganization(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error)
	ShowOrganization(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error)
	RefreshLastTouchpoint(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error)
	RefreshRenewalSummary(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error)
	RefreshArr(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error)
	AddParentOrganization(context.Context, *AddParentOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	RemoveParentOrganization(context.Context, *RemoveParentOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	UpdateOnboardingStatus(context.Context, *UpdateOnboardingStatusGrpcRequest) (*OrganizationIdGrpcResponse, error)
	UpdateOrganization(context.Context, *UpdateOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	AddSocial(context.Context, *AddSocialGrpcRequest) (*OrganizationIdGrpcResponse, error)
	UpdateOrganizationOwner(context.Context, *UpdateOrganizationOwnerGrpcRequest) (*OrganizationIdGrpcResponse, error)
	CreateBillingProfile(context.Context, *CreateBillingProfileGrpcRequest) (*BillingProfileIdGrpcResponse, error)
	LinkEmailToBillingProfile(context.Context, *LinkEmailToBillingProfileGrpcRequest) (*BillingProfileIdGrpcResponse, error)
	UnlinkEmailFromBillingProfile(context.Context, *UnlinkEmailFromBillingProfileGrpcRequest) (*BillingProfileIdGrpcResponse, error)
}

// UnimplementedOrganizationGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrganizationGrpcServiceServer struct {
}

func (UnimplementedOrganizationGrpcServiceServer) UpsertOrganization(context.Context, *UpsertOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) WebScrapeOrganization(context.Context, *WebScrapeOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebScrapeOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) LinkPhoneNumberToOrganization(context.Context, *LinkPhoneNumberToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkPhoneNumberToOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) LinkEmailToOrganization(context.Context, *LinkEmailToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkEmailToOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) LinkLocationToOrganization(context.Context, *LinkLocationToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkLocationToOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) LinkDomainToOrganization(context.Context, *LinkDomainToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkDomainToOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) UpsertCustomFieldToOrganization(context.Context, *CustomFieldForOrganizationGrpcRequest) (*CustomFieldIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCustomFieldToOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) HideOrganization(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HideOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) ShowOrganization(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) RefreshLastTouchpoint(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshLastTouchpoint not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) RefreshRenewalSummary(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshRenewalSummary not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) RefreshArr(context.Context, *OrganizationIdGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshArr not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) AddParentOrganization(context.Context, *AddParentOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParentOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) RemoveParentOrganization(context.Context, *RemoveParentOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveParentOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) UpdateOnboardingStatus(context.Context, *UpdateOnboardingStatusGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOnboardingStatus not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) UpdateOrganization(context.Context, *UpdateOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) AddSocial(context.Context, *AddSocialGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSocial not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) UpdateOrganizationOwner(context.Context, *UpdateOrganizationOwnerGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganizationOwner not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) CreateBillingProfile(context.Context, *CreateBillingProfileGrpcRequest) (*BillingProfileIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBillingProfile not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) LinkEmailToBillingProfile(context.Context, *LinkEmailToBillingProfileGrpcRequest) (*BillingProfileIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkEmailToBillingProfile not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) UnlinkEmailFromBillingProfile(context.Context, *UnlinkEmailFromBillingProfileGrpcRequest) (*BillingProfileIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkEmailFromBillingProfile not implemented")
}

// UnsafeOrganizationGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationGrpcServiceServer will
// result in compilation errors.
type UnsafeOrganizationGrpcServiceServer interface {
	mustEmbedUnimplementedOrganizationGrpcServiceServer()
}

func RegisterOrganizationGrpcServiceServer(s grpc.ServiceRegistrar, srv OrganizationGrpcServiceServer) {
	s.RegisterService(&OrganizationGrpcService_ServiceDesc, srv)
}

func _OrganizationGrpcService_UpsertOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).UpsertOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/UpsertOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).UpsertOrganization(ctx, req.(*UpsertOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_WebScrapeOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebScrapeOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).WebScrapeOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/WebScrapeOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).WebScrapeOrganization(ctx, req.(*WebScrapeOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_LinkPhoneNumberToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkPhoneNumberToOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).LinkPhoneNumberToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/LinkPhoneNumberToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).LinkPhoneNumberToOrganization(ctx, req.(*LinkPhoneNumberToOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_LinkEmailToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkEmailToOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).LinkEmailToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/LinkEmailToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).LinkEmailToOrganization(ctx, req.(*LinkEmailToOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_LinkLocationToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkLocationToOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).LinkLocationToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/LinkLocationToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).LinkLocationToOrganization(ctx, req.(*LinkLocationToOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_LinkDomainToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkDomainToOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).LinkDomainToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/LinkDomainToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).LinkDomainToOrganization(ctx, req.(*LinkDomainToOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_UpsertCustomFieldToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomFieldForOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).UpsertCustomFieldToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/UpsertCustomFieldToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).UpsertCustomFieldToOrganization(ctx, req.(*CustomFieldForOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_HideOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).HideOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/HideOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).HideOrganization(ctx, req.(*OrganizationIdGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_ShowOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).ShowOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/ShowOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).ShowOrganization(ctx, req.(*OrganizationIdGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_RefreshLastTouchpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).RefreshLastTouchpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/RefreshLastTouchpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).RefreshLastTouchpoint(ctx, req.(*OrganizationIdGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_RefreshRenewalSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).RefreshRenewalSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/RefreshRenewalSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).RefreshRenewalSummary(ctx, req.(*OrganizationIdGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_RefreshArr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).RefreshArr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/RefreshArr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).RefreshArr(ctx, req.(*OrganizationIdGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_AddParentOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParentOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).AddParentOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/AddParentOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).AddParentOrganization(ctx, req.(*AddParentOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_RemoveParentOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveParentOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).RemoveParentOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/RemoveParentOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).RemoveParentOrganization(ctx, req.(*RemoveParentOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_UpdateOnboardingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOnboardingStatusGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).UpdateOnboardingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/UpdateOnboardingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).UpdateOnboardingStatus(ctx, req.(*UpdateOnboardingStatusGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/UpdateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).UpdateOrganization(ctx, req.(*UpdateOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_AddSocial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSocialGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).AddSocial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/AddSocial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).AddSocial(ctx, req.(*AddSocialGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_UpdateOrganizationOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationOwnerGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).UpdateOrganizationOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/UpdateOrganizationOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).UpdateOrganizationOwner(ctx, req.(*UpdateOrganizationOwnerGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_CreateBillingProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillingProfileGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).CreateBillingProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/CreateBillingProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).CreateBillingProfile(ctx, req.(*CreateBillingProfileGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_LinkEmailToBillingProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkEmailToBillingProfileGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).LinkEmailToBillingProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/LinkEmailToBillingProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).LinkEmailToBillingProfile(ctx, req.(*LinkEmailToBillingProfileGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_UnlinkEmailFromBillingProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkEmailFromBillingProfileGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).UnlinkEmailFromBillingProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/UnlinkEmailFromBillingProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).UnlinkEmailFromBillingProfile(ctx, req.(*UnlinkEmailFromBillingProfileGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationGrpcService_ServiceDesc is the grpc.ServiceDesc for OrganizationGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organizationGrpcService",
	HandlerType: (*OrganizationGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertOrganization",
			Handler:    _OrganizationGrpcService_UpsertOrganization_Handler,
		},
		{
			MethodName: "WebScrapeOrganization",
			Handler:    _OrganizationGrpcService_WebScrapeOrganization_Handler,
		},
		{
			MethodName: "LinkPhoneNumberToOrganization",
			Handler:    _OrganizationGrpcService_LinkPhoneNumberToOrganization_Handler,
		},
		{
			MethodName: "LinkEmailToOrganization",
			Handler:    _OrganizationGrpcService_LinkEmailToOrganization_Handler,
		},
		{
			MethodName: "LinkLocationToOrganization",
			Handler:    _OrganizationGrpcService_LinkLocationToOrganization_Handler,
		},
		{
			MethodName: "LinkDomainToOrganization",
			Handler:    _OrganizationGrpcService_LinkDomainToOrganization_Handler,
		},
		{
			MethodName: "UpsertCustomFieldToOrganization",
			Handler:    _OrganizationGrpcService_UpsertCustomFieldToOrganization_Handler,
		},
		{
			MethodName: "HideOrganization",
			Handler:    _OrganizationGrpcService_HideOrganization_Handler,
		},
		{
			MethodName: "ShowOrganization",
			Handler:    _OrganizationGrpcService_ShowOrganization_Handler,
		},
		{
			MethodName: "RefreshLastTouchpoint",
			Handler:    _OrganizationGrpcService_RefreshLastTouchpoint_Handler,
		},
		{
			MethodName: "RefreshRenewalSummary",
			Handler:    _OrganizationGrpcService_RefreshRenewalSummary_Handler,
		},
		{
			MethodName: "RefreshArr",
			Handler:    _OrganizationGrpcService_RefreshArr_Handler,
		},
		{
			MethodName: "AddParentOrganization",
			Handler:    _OrganizationGrpcService_AddParentOrganization_Handler,
		},
		{
			MethodName: "RemoveParentOrganization",
			Handler:    _OrganizationGrpcService_RemoveParentOrganization_Handler,
		},
		{
			MethodName: "UpdateOnboardingStatus",
			Handler:    _OrganizationGrpcService_UpdateOnboardingStatus_Handler,
		},
		{
			MethodName: "UpdateOrganization",
			Handler:    _OrganizationGrpcService_UpdateOrganization_Handler,
		},
		{
			MethodName: "AddSocial",
			Handler:    _OrganizationGrpcService_AddSocial_Handler,
		},
		{
			MethodName: "UpdateOrganizationOwner",
			Handler:    _OrganizationGrpcService_UpdateOrganizationOwner_Handler,
		},
		{
			MethodName: "CreateBillingProfile",
			Handler:    _OrganizationGrpcService_CreateBillingProfile_Handler,
		},
		{
			MethodName: "LinkEmailToBillingProfile",
			Handler:    _OrganizationGrpcService_LinkEmailToBillingProfile_Handler,
		},
		{
			MethodName: "UnlinkEmailFromBillingProfile",
			Handler:    _OrganizationGrpcService_UnlinkEmailFromBillingProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization.proto",
}
