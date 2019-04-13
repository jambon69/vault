// Code generated by protoc-gen-go. DO NOT EDIT.
// source: builtin/logical/database/dbplugin/database.proto

package dbplugin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Deprecated: Do not use.
type InitializeRequest struct {
	Config               []byte   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	VerifyConnection     bool     `protobuf:"varint,2,opt,name=verify_connection,json=verifyConnection,proto3" json:"verify_connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeRequest) Reset()         { *m = InitializeRequest{} }
func (m *InitializeRequest) String() string { return proto.CompactTextString(m) }
func (*InitializeRequest) ProtoMessage()    {}
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{0}
}

func (m *InitializeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeRequest.Unmarshal(m, b)
}
func (m *InitializeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeRequest.Marshal(b, m, deterministic)
}
func (m *InitializeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeRequest.Merge(m, src)
}
func (m *InitializeRequest) XXX_Size() int {
	return xxx_messageInfo_InitializeRequest.Size(m)
}
func (m *InitializeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeRequest proto.InternalMessageInfo

func (m *InitializeRequest) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *InitializeRequest) GetVerifyConnection() bool {
	if m != nil {
		return m.VerifyConnection
	}
	return false
}

type InitRequest struct {
	Config               []byte   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	VerifyConnection     bool     `protobuf:"varint,2,opt,name=verify_connection,json=verifyConnection,proto3" json:"verify_connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{1}
}

func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (m *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(m, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *InitRequest) GetVerifyConnection() bool {
	if m != nil {
		return m.VerifyConnection
	}
	return false
}

type CreateUserRequest struct {
	Statements           *Statements          `protobuf:"bytes,1,opt,name=statements,proto3" json:"statements,omitempty"`
	UsernameConfig       *UsernameConfig      `protobuf:"bytes,2,opt,name=username_config,json=usernameConfig,proto3" json:"username_config,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetStatements() *Statements {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *CreateUserRequest) GetUsernameConfig() *UsernameConfig {
	if m != nil {
		return m.UsernameConfig
	}
	return nil
}

func (m *CreateUserRequest) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type RenewUserRequest struct {
	Statements           *Statements          `protobuf:"bytes,1,opt,name=statements,proto3" json:"statements,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RenewUserRequest) Reset()         { *m = RenewUserRequest{} }
func (m *RenewUserRequest) String() string { return proto.CompactTextString(m) }
func (*RenewUserRequest) ProtoMessage()    {}
func (*RenewUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{3}
}

func (m *RenewUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewUserRequest.Unmarshal(m, b)
}
func (m *RenewUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewUserRequest.Marshal(b, m, deterministic)
}
func (m *RenewUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewUserRequest.Merge(m, src)
}
func (m *RenewUserRequest) XXX_Size() int {
	return xxx_messageInfo_RenewUserRequest.Size(m)
}
func (m *RenewUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenewUserRequest proto.InternalMessageInfo

func (m *RenewUserRequest) GetStatements() *Statements {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *RenewUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RenewUserRequest) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type RevokeUserRequest struct {
	Statements           *Statements `protobuf:"bytes,1,opt,name=statements,proto3" json:"statements,omitempty"`
	Username             string      `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RevokeUserRequest) Reset()         { *m = RevokeUserRequest{} }
func (m *RevokeUserRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeUserRequest) ProtoMessage()    {}
func (*RevokeUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{4}
}

func (m *RevokeUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeUserRequest.Unmarshal(m, b)
}
func (m *RevokeUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeUserRequest.Marshal(b, m, deterministic)
}
func (m *RevokeUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeUserRequest.Merge(m, src)
}
func (m *RevokeUserRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeUserRequest.Size(m)
}
func (m *RevokeUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeUserRequest proto.InternalMessageInfo

func (m *RevokeUserRequest) GetStatements() *Statements {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *RevokeUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RotateRootCredentialsRequest struct {
	Statements           []string `protobuf:"bytes,1,rep,name=statements,proto3" json:"statements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RotateRootCredentialsRequest) Reset()         { *m = RotateRootCredentialsRequest{} }
func (m *RotateRootCredentialsRequest) String() string { return proto.CompactTextString(m) }
func (*RotateRootCredentialsRequest) ProtoMessage()    {}
func (*RotateRootCredentialsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{5}
}

func (m *RotateRootCredentialsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RotateRootCredentialsRequest.Unmarshal(m, b)
}
func (m *RotateRootCredentialsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RotateRootCredentialsRequest.Marshal(b, m, deterministic)
}
func (m *RotateRootCredentialsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RotateRootCredentialsRequest.Merge(m, src)
}
func (m *RotateRootCredentialsRequest) XXX_Size() int {
	return xxx_messageInfo_RotateRootCredentialsRequest.Size(m)
}
func (m *RotateRootCredentialsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RotateRootCredentialsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RotateRootCredentialsRequest proto.InternalMessageInfo

func (m *RotateRootCredentialsRequest) GetStatements() []string {
	if m != nil {
		return m.Statements
	}
	return nil
}

type Statements struct {
	// DEPRECATED, will be removed in 0.12
	CreationStatements string `protobuf:"bytes,1,opt,name=creation_statements,json=creationStatements,proto3" json:"creation_statements,omitempty"` // Deprecated: Do not use.
	// DEPRECATED, will be removed in 0.12
	RevocationStatements string `protobuf:"bytes,2,opt,name=revocation_statements,json=revocationStatements,proto3" json:"revocation_statements,omitempty"` // Deprecated: Do not use.
	// DEPRECATED, will be removed in 0.12
	RollbackStatements string `protobuf:"bytes,3,opt,name=rollback_statements,json=rollbackStatements,proto3" json:"rollback_statements,omitempty"` // Deprecated: Do not use.
	// DEPRECATED, will be removed in 0.12
	RenewStatements      string   `protobuf:"bytes,4,opt,name=renew_statements,json=renewStatements,proto3" json:"renew_statements,omitempty"` // Deprecated: Do not use.
	Creation             []string `protobuf:"bytes,5,rep,name=creation,proto3" json:"creation,omitempty"`
	Revocation           []string `protobuf:"bytes,6,rep,name=revocation,proto3" json:"revocation,omitempty"`
	Rollback             []string `protobuf:"bytes,7,rep,name=rollback,proto3" json:"rollback,omitempty"`
	Renewal              []string `protobuf:"bytes,8,rep,name=renewal,proto3" json:"renewal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statements) Reset()         { *m = Statements{} }
func (m *Statements) String() string { return proto.CompactTextString(m) }
func (*Statements) ProtoMessage()    {}
func (*Statements) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{6}
}

func (m *Statements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statements.Unmarshal(m, b)
}
func (m *Statements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statements.Marshal(b, m, deterministic)
}
func (m *Statements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statements.Merge(m, src)
}
func (m *Statements) XXX_Size() int {
	return xxx_messageInfo_Statements.Size(m)
}
func (m *Statements) XXX_DiscardUnknown() {
	xxx_messageInfo_Statements.DiscardUnknown(m)
}

var xxx_messageInfo_Statements proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Statements) GetCreationStatements() string {
	if m != nil {
		return m.CreationStatements
	}
	return ""
}

// Deprecated: Do not use.
func (m *Statements) GetRevocationStatements() string {
	if m != nil {
		return m.RevocationStatements
	}
	return ""
}

// Deprecated: Do not use.
func (m *Statements) GetRollbackStatements() string {
	if m != nil {
		return m.RollbackStatements
	}
	return ""
}

// Deprecated: Do not use.
func (m *Statements) GetRenewStatements() string {
	if m != nil {
		return m.RenewStatements
	}
	return ""
}

func (m *Statements) GetCreation() []string {
	if m != nil {
		return m.Creation
	}
	return nil
}

func (m *Statements) GetRevocation() []string {
	if m != nil {
		return m.Revocation
	}
	return nil
}

func (m *Statements) GetRollback() []string {
	if m != nil {
		return m.Rollback
	}
	return nil
}

func (m *Statements) GetRenewal() []string {
	if m != nil {
		return m.Renewal
	}
	return nil
}

type UsernameConfig struct {
	DisplayName          string   `protobuf:"bytes,1,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	RoleName             string   `protobuf:"bytes,2,opt,name=RoleName,proto3" json:"RoleName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsernameConfig) Reset()         { *m = UsernameConfig{} }
func (m *UsernameConfig) String() string { return proto.CompactTextString(m) }
func (*UsernameConfig) ProtoMessage()    {}
func (*UsernameConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{7}
}

func (m *UsernameConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsernameConfig.Unmarshal(m, b)
}
func (m *UsernameConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsernameConfig.Marshal(b, m, deterministic)
}
func (m *UsernameConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsernameConfig.Merge(m, src)
}
func (m *UsernameConfig) XXX_Size() int {
	return xxx_messageInfo_UsernameConfig.Size(m)
}
func (m *UsernameConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UsernameConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UsernameConfig proto.InternalMessageInfo

func (m *UsernameConfig) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *UsernameConfig) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

type InitResponse struct {
	Config               []byte   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitResponse) Reset()         { *m = InitResponse{} }
func (m *InitResponse) String() string { return proto.CompactTextString(m) }
func (*InitResponse) ProtoMessage()    {}
func (*InitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{8}
}

func (m *InitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitResponse.Unmarshal(m, b)
}
func (m *InitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitResponse.Marshal(b, m, deterministic)
}
func (m *InitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitResponse.Merge(m, src)
}
func (m *InitResponse) XXX_Size() int {
	return xxx_messageInfo_InitResponse.Size(m)
}
func (m *InitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitResponse proto.InternalMessageInfo

func (m *InitResponse) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type CreateUserResponse struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{9}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type TypeResponse struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeResponse) Reset()         { *m = TypeResponse{} }
func (m *TypeResponse) String() string { return proto.CompactTextString(m) }
func (*TypeResponse) ProtoMessage()    {}
func (*TypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{10}
}

func (m *TypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeResponse.Unmarshal(m, b)
}
func (m *TypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeResponse.Marshal(b, m, deterministic)
}
func (m *TypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeResponse.Merge(m, src)
}
func (m *TypeResponse) XXX_Size() int {
	return xxx_messageInfo_TypeResponse.Size(m)
}
func (m *TypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TypeResponse proto.InternalMessageInfo

func (m *TypeResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type RotateRootCredentialsResponse struct {
	Config               []byte   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RotateRootCredentialsResponse) Reset()         { *m = RotateRootCredentialsResponse{} }
func (m *RotateRootCredentialsResponse) String() string { return proto.CompactTextString(m) }
func (*RotateRootCredentialsResponse) ProtoMessage()    {}
func (*RotateRootCredentialsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{11}
}

func (m *RotateRootCredentialsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RotateRootCredentialsResponse.Unmarshal(m, b)
}
func (m *RotateRootCredentialsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RotateRootCredentialsResponse.Marshal(b, m, deterministic)
}
func (m *RotateRootCredentialsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RotateRootCredentialsResponse.Merge(m, src)
}
func (m *RotateRootCredentialsResponse) XXX_Size() int {
	return xxx_messageInfo_RotateRootCredentialsResponse.Size(m)
}
func (m *RotateRootCredentialsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RotateRootCredentialsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RotateRootCredentialsResponse proto.InternalMessageInfo

func (m *RotateRootCredentialsResponse) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf7b4c7fef2f66e, []int{12}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InitializeRequest)(nil), "dbplugin.InitializeRequest")
	proto.RegisterType((*InitRequest)(nil), "dbplugin.InitRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "dbplugin.CreateUserRequest")
	proto.RegisterType((*RenewUserRequest)(nil), "dbplugin.RenewUserRequest")
	proto.RegisterType((*RevokeUserRequest)(nil), "dbplugin.RevokeUserRequest")
	proto.RegisterType((*RotateRootCredentialsRequest)(nil), "dbplugin.RotateRootCredentialsRequest")
	proto.RegisterType((*Statements)(nil), "dbplugin.Statements")
	proto.RegisterType((*UsernameConfig)(nil), "dbplugin.UsernameConfig")
	proto.RegisterType((*InitResponse)(nil), "dbplugin.InitResponse")
	proto.RegisterType((*CreateUserResponse)(nil), "dbplugin.CreateUserResponse")
	proto.RegisterType((*TypeResponse)(nil), "dbplugin.TypeResponse")
	proto.RegisterType((*RotateRootCredentialsResponse)(nil), "dbplugin.RotateRootCredentialsResponse")
	proto.RegisterType((*Empty)(nil), "dbplugin.Empty")
}

func init() {
	proto.RegisterFile("builtin/logical/database/dbplugin/database.proto", fileDescriptor_7bf7b4c7fef2f66e)
}

var fileDescriptor_7bf7b4c7fef2f66e = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0x95, 0x93, 0x00, 0xc9, 0x80, 0x80, 0xec, 0x05, 0x64, 0xf9, 0x72, 0x6f, 0x91, 0x1f, 0x28,
	0x55, 0xd5, 0xb8, 0x82, 0x56, 0x54, 0xa8, 0xa2, 0x2a, 0xa1, 0xaa, 0x2a, 0x55, 0x3c, 0x2c, 0xf0,
	0x52, 0x55, 0x42, 0x1b, 0x67, 0x49, 0x56, 0x38, 0x5e, 0xd7, 0xbb, 0x0e, 0x4d, 0x7f, 0xa0, 0xfd,
	0x8c, 0x7e, 0x4e, 0x1f, 0xfb, 0x49, 0x95, 0x37, 0x59, 0xef, 0x26, 0x86, 0xf2, 0x40, 0xfb, 0xe6,
	0xd9, 0x99, 0x33, 0x73, 0xe6, 0x78, 0x76, 0x16, 0x9e, 0x76, 0x32, 0x16, 0x49, 0x16, 0x07, 0x11,
	0xef, 0xb1, 0x90, 0x44, 0x41, 0x97, 0x48, 0xd2, 0x21, 0x82, 0x06, 0xdd, 0x4e, 0x12, 0x65, 0x3d,
	0x16, 0x17, 0x27, 0xad, 0x24, 0xe5, 0x92, 0xa3, 0xba, 0x76, 0x78, 0x0f, 0x7a, 0x9c, 0xf7, 0x22,
	0x1a, 0xa8, 0xf3, 0x4e, 0x76, 0x19, 0x48, 0x36, 0xa0, 0x42, 0x92, 0x41, 0x32, 0x0e, 0xf5, 0x3f,
	0x42, 0xf3, 0x5d, 0xcc, 0x24, 0x23, 0x11, 0xfb, 0x42, 0x31, 0xfd, 0x94, 0x51, 0x21, 0xd1, 0x06,
	0xcc, 0x87, 0x3c, 0xbe, 0x64, 0x3d, 0xd7, 0xd9, 0x72, 0x76, 0x96, 0xf0, 0xc4, 0x42, 0x8f, 0xa1,
	0x39, 0xa4, 0x29, 0xbb, 0x1c, 0x5d, 0x84, 0x3c, 0x8e, 0x69, 0x28, 0x19, 0x8f, 0xdd, 0xca, 0x96,
	0xb3, 0x53, 0xc7, 0xab, 0x63, 0x47, 0xbb, 0x38, 0x3f, 0xa8, 0xb8, 0x8e, 0x8f, 0x61, 0x31, 0xcf,
	0xfe, 0x27, 0xf3, 0xfa, 0x3f, 0x1c, 0x68, 0xb6, 0x53, 0x4a, 0x24, 0x3d, 0x17, 0x34, 0xd5, 0xa9,
	0x9f, 0x01, 0x08, 0x49, 0x24, 0x1d, 0xd0, 0x58, 0x0a, 0x95, 0x7e, 0x71, 0x77, 0xad, 0xa5, 0x75,
	0x68, 0x9d, 0x16, 0x3e, 0x6c, 0xc5, 0xa1, 0xd7, 0xb0, 0x92, 0x09, 0x9a, 0xc6, 0x64, 0x40, 0x2f,
	0x26, 0xcc, 0x2a, 0x0a, 0xea, 0x1a, 0xe8, 0xf9, 0x24, 0xa0, 0xad, 0xfc, 0x78, 0x39, 0x9b, 0xb2,
	0xd1, 0x01, 0x00, 0xfd, 0x9c, 0xb0, 0x94, 0x28, 0xd2, 0x55, 0x85, 0xf6, 0x5a, 0x63, 0xd9, 0x5b,
	0x5a, 0xf6, 0xd6, 0x99, 0x96, 0x1d, 0x5b, 0xd1, 0xfe, 0x77, 0x07, 0x56, 0x31, 0x8d, 0xe9, 0xf5,
	0xfd, 0x3b, 0xf1, 0xa0, 0xae, 0x89, 0xa9, 0x16, 0x1a, 0xb8, 0xb0, 0xef, 0x45, 0x91, 0x42, 0x13,
	0xd3, 0x21, 0xbf, 0xa2, 0x7f, 0x95, 0xa2, 0x7f, 0x08, 0x9b, 0x98, 0xe7, 0xa1, 0x98, 0x73, 0xd9,
	0x4e, 0x69, 0x97, 0xc6, 0xf9, 0x4c, 0x0a, 0x5d, 0xf1, 0xff, 0x99, 0x8a, 0xd5, 0x9d, 0x86, 0x9d,
	0xdb, 0xff, 0x59, 0x01, 0x30, 0x65, 0xd1, 0x1e, 0xfc, 0x13, 0xe6, 0x23, 0xc2, 0x78, 0x7c, 0x31,
	0xc3, 0xb4, 0x71, 0x54, 0x71, 0x1d, 0x8c, 0xb4, 0xdb, 0x02, 0xed, 0xc3, 0x7a, 0x4a, 0x87, 0x3c,
	0x2c, 0xc1, 0x2a, 0x05, 0x6c, 0xcd, 0x04, 0x4c, 0x57, 0x4b, 0x79, 0x14, 0x75, 0x48, 0x78, 0x65,
	0xc3, 0xaa, 0xa6, 0x9a, 0x76, 0x5b, 0xa0, 0x27, 0xb0, 0x9a, 0xe6, 0xbf, 0xde, 0x46, 0xd4, 0x0a,
	0xc4, 0x8a, 0xf2, 0x9d, 0x4e, 0x89, 0xa7, 0x29, 0xbb, 0x73, 0xaa, 0xfd, 0xc2, 0xce, 0xc5, 0x31,
	0xbc, 0xdc, 0xf9, 0xb1, 0x38, 0xe6, 0x24, 0xc7, 0x6a, 0x02, 0xee, 0xc2, 0x18, 0xab, 0x6d, 0xe4,
	0xc2, 0x82, 0x2a, 0x45, 0x22, 0xb7, 0xae, 0x5c, 0xda, 0xf4, 0x4f, 0x60, 0x79, 0x7a, 0xf4, 0xd1,
	0x16, 0x2c, 0x1e, 0x33, 0x91, 0x44, 0x64, 0x74, 0x92, 0xff, 0x43, 0xa5, 0x26, 0xb6, 0x8f, 0xf2,
	0x4a, 0x98, 0x47, 0xf4, 0xc4, 0xfa, 0xc5, 0xda, 0xf6, 0xb7, 0x61, 0x69, 0xbc, 0x0b, 0x44, 0xc2,
	0x63, 0x41, 0x6f, 0x5b, 0x06, 0xfe, 0x7b, 0x40, 0xf6, 0xf5, 0x9e, 0x44, 0xdb, 0xc3, 0xe3, 0xcc,
	0xcc, 0xb7, 0x07, 0xf5, 0x84, 0x08, 0x71, 0xcd, 0xd3, 0xae, 0xae, 0xaa, 0x6d, 0xdf, 0x87, 0xa5,
	0xb3, 0x51, 0x42, 0x8b, 0x3c, 0x08, 0x6a, 0x72, 0x94, 0xe8, 0x1c, 0xea, 0xdb, 0xdf, 0x87, 0xff,
	0x6e, 0x19, 0xbe, 0x3b, 0xa8, 0x2e, 0xc0, 0xdc, 0x9b, 0x41, 0x22, 0x47, 0xbb, 0x5f, 0x6b, 0x50,
	0x3f, 0x9e, 0xec, 0x60, 0x14, 0x40, 0x2d, 0x2f, 0x89, 0x56, 0xcc, 0x8d, 0x50, 0x51, 0xde, 0x86,
	0x39, 0x98, 0xe2, 0xf4, 0x16, 0xc0, 0x74, 0x8c, 0xfe, 0x35, 0x51, 0xa5, 0x35, 0xe7, 0x6d, 0xde,
	0xec, 0x9c, 0x24, 0x7a, 0x01, 0x8d, 0x62, 0x9d, 0x20, 0xcf, 0x84, 0xce, 0xee, 0x18, 0x6f, 0x96,
	0x5a, 0xbe, 0x22, 0xcc, 0x35, 0xb7, 0x29, 0x94, 0x2e, 0x7f, 0x19, 0xdb, 0x87, 0xf5, 0x1b, 0xe5,
	0x43, 0xdb, 0x56, 0x9a, 0xdf, 0x5c, 0x6e, 0xef, 0xe1, 0x9d, 0x71, 0x93, 0xfe, 0x9e, 0x43, 0x2d,
	0x1f, 0x21, 0xb4, 0x6e, 0x00, 0xd6, 0xf3, 0x62, 0xeb, 0x3b, 0x35, 0x69, 0x8f, 0x60, 0xae, 0x1d,
	0x71, 0x71, 0xc3, 0x1f, 0x29, 0xf5, 0xf2, 0x0a, 0xc0, 0x3c, 0x87, 0xb6, 0x0e, 0xa5, 0x47, 0xb2,
	0x84, 0xf5, 0xab, 0xdf, 0x2a, 0xce, 0xd1, 0xe1, 0x87, 0x97, 0x3d, 0x26, 0xfb, 0x59, 0xa7, 0x15,
	0xf2, 0x41, 0xd0, 0x27, 0xa2, 0xcf, 0x42, 0x9e, 0x26, 0xc1, 0x90, 0x64, 0x91, 0x0c, 0xee, 0x7c,
	0xc9, 0x3b, 0xf3, 0x6a, 0x1f, 0xef, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x49, 0x0b, 0x5b,
	0xf5, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseClient interface {
	Type(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TypeResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	RenewUser(ctx context.Context, in *RenewUserRequest, opts ...grpc.CallOption) (*Empty, error)
	RevokeUser(ctx context.Context, in *RevokeUserRequest, opts ...grpc.CallOption) (*Empty, error)
	RotateRootCredentials(ctx context.Context, in *RotateRootCredentialsRequest, opts ...grpc.CallOption) (*RotateRootCredentialsResponse, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*Empty, error)
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Type(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TypeResponse, error) {
	out := new(TypeResponse)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/Type", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) RenewUser(ctx context.Context, in *RenewUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/RenewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) RevokeUser(ctx context.Context, in *RevokeUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/RevokeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) RotateRootCredentials(ctx context.Context, in *RotateRootCredentialsRequest, opts ...grpc.CallOption) (*RotateRootCredentialsResponse, error) {
	out := new(RotateRootCredentialsResponse)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/RotateRootCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *databaseClient) Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dbplugin.Database/Initialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
type DatabaseServer interface {
	Type(context.Context, *Empty) (*TypeResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	RenewUser(context.Context, *RenewUserRequest) (*Empty, error)
	RevokeUser(context.Context, *RevokeUserRequest) (*Empty, error)
	RotateRootCredentials(context.Context, *RotateRootCredentialsRequest) (*RotateRootCredentialsResponse, error)
	Init(context.Context, *InitRequest) (*InitResponse, error)
	Close(context.Context, *Empty) (*Empty, error)
	Initialize(context.Context, *InitializeRequest) (*Empty, error)
}

// UnimplementedDatabaseServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (*UnimplementedDatabaseServer) Type(ctx context.Context, req *Empty) (*TypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Type not implemented")
}
func (*UnimplementedDatabaseServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedDatabaseServer) RenewUser(ctx context.Context, req *RenewUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewUser not implemented")
}
func (*UnimplementedDatabaseServer) RevokeUser(ctx context.Context, req *RevokeUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeUser not implemented")
}
func (*UnimplementedDatabaseServer) RotateRootCredentials(ctx context.Context, req *RotateRootCredentialsRequest) (*RotateRootCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateRootCredentials not implemented")
}
func (*UnimplementedDatabaseServer) Init(ctx context.Context, req *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedDatabaseServer) Close(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedDatabaseServer) Initialize(ctx context.Context, req *InitializeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initialize not implemented")
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_Type_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Type(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/Type",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Type(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_RenewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).RenewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/RenewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).RenewUser(ctx, req.(*RenewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_RevokeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).RevokeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/RevokeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).RevokeUser(ctx, req.(*RevokeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_RotateRootCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateRootCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).RotateRootCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/RotateRootCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).RotateRootCredentials(ctx, req.(*RotateRootCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Close(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Initialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Initialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbplugin.Database/Initialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Initialize(ctx, req.(*InitializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbplugin.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Type",
			Handler:    _Database_Type_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Database_CreateUser_Handler,
		},
		{
			MethodName: "RenewUser",
			Handler:    _Database_RenewUser_Handler,
		},
		{
			MethodName: "RevokeUser",
			Handler:    _Database_RevokeUser_Handler,
		},
		{
			MethodName: "RotateRootCredentials",
			Handler:    _Database_RotateRootCredentials_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Database_Init_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Database_Close_Handler,
		},
		{
			MethodName: "Initialize",
			Handler:    _Database_Initialize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "builtin/logical/database/dbplugin/database.proto",
}
