// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: utils/pb/jobseeker/jobseeker.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateJobseekerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname   string                 `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname    string                 `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email       string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password    string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Gender      string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Phone       string                 `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Dateofbirth *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty"`
}

func (x *CreateJobseekerReq) Reset() {
	*x = CreateJobseekerReq{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobseekerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobseekerReq) ProtoMessage() {}

func (x *CreateJobseekerReq) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobseekerReq.ProtoReflect.Descriptor instead.
func (*CreateJobseekerReq) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobseekerReq) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CreateJobseekerReq) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CreateJobseekerReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateJobseekerReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateJobseekerReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateJobseekerReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateJobseekerReq) GetDateofbirth() *timestamppb.Timestamp {
	if x != nil {
		return x.Dateofbirth
	}
	return nil
}

type CreateJobseekerRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname   string                 `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname    string                 `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email       string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Gender      string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Phone       string                 `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Dateofbirth *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty"`
	Createdat   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=createdat,proto3" json:"createdat,omitempty"`
	Updatedat   string                 `protobuf:"bytes,9,opt,name=updatedat,proto3" json:"updatedat,omitempty"`
}

func (x *CreateJobseekerRes) Reset() {
	*x = CreateJobseekerRes{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobseekerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobseekerRes) ProtoMessage() {}

func (x *CreateJobseekerRes) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobseekerRes.ProtoReflect.Descriptor instead.
func (*CreateJobseekerRes) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobseekerRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateJobseekerRes) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CreateJobseekerRes) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CreateJobseekerRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateJobseekerRes) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateJobseekerRes) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateJobseekerRes) GetDateofbirth() *timestamppb.Timestamp {
	if x != nil {
		return x.Dateofbirth
	}
	return nil
}

func (x *CreateJobseekerRes) GetCreatedat() *timestamppb.Timestamp {
	if x != nil {
		return x.Createdat
	}
	return nil
}

func (x *CreateJobseekerRes) GetUpdatedat() string {
	if x != nil {
		return x.Updatedat
	}
	return ""
}

type JobSeekerProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobseekerid int64  `protobuf:"varint,1,opt,name=jobseekerid,proto3" json:"jobseekerid,omitempty"`
	Summary     string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	City        string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country     string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Education   string `protobuf:"bytes,6,opt,name=education,proto3" json:"education,omitempty"`
	Experience  string `protobuf:"bytes,7,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *JobSeekerProfileReq) Reset() {
	*x = JobSeekerProfileReq{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobSeekerProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSeekerProfileReq) ProtoMessage() {}

func (x *JobSeekerProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSeekerProfileReq.ProtoReflect.Descriptor instead.
func (*JobSeekerProfileReq) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{2}
}

func (x *JobSeekerProfileReq) GetJobseekerid() int64 {
	if x != nil {
		return x.Jobseekerid
	}
	return 0
}

func (x *JobSeekerProfileReq) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *JobSeekerProfileReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *JobSeekerProfileReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *JobSeekerProfileReq) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *JobSeekerProfileReq) GetExperience() string {
	if x != nil {
		return x.Experience
	}
	return ""
}

type JobSeekerProfileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobseeker  *CreateJobseekerRes `protobuf:"bytes,1,opt,name=jobseeker,proto3" json:"jobseeker,omitempty"`
	Summary    string              `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	City       string              `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country    string              `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Education  string              `protobuf:"bytes,6,opt,name=education,proto3" json:"education,omitempty"`
	Experience string              `protobuf:"bytes,7,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *JobSeekerProfileRes) Reset() {
	*x = JobSeekerProfileRes{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobSeekerProfileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSeekerProfileRes) ProtoMessage() {}

func (x *JobSeekerProfileRes) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSeekerProfileRes.ProtoReflect.Descriptor instead.
func (*JobSeekerProfileRes) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{3}
}

func (x *JobSeekerProfileRes) GetJobseeker() *CreateJobseekerRes {
	if x != nil {
		return x.Jobseeker
	}
	return nil
}

func (x *JobSeekerProfileRes) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *JobSeekerProfileRes) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *JobSeekerProfileRes) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *JobSeekerProfileRes) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *JobSeekerProfileRes) GetExperience() string {
	if x != nil {
		return x.Experience
	}
	return ""
}

type JobSeekerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JobSeekerID) Reset() {
	*x = JobSeekerID{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobSeekerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSeekerID) ProtoMessage() {}

func (x *JobSeekerID) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSeekerID.ProtoReflect.Descriptor instead.
func (*JobSeekerID) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{4}
}

func (x *JobSeekerID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type JSLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *JSLoginReq) Reset() {
	*x = JSLoginReq{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JSLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSLoginReq) ProtoMessage() {}

func (x *JSLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSLoginReq.ProtoReflect.Descriptor instead.
func (*JSLoginReq) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{5}
}

func (x *JSLoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *JSLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type FollowEmployerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobseekerid int64 `protobuf:"varint,1,opt,name=jobseekerid,proto3" json:"jobseekerid,omitempty"`
	Employerid  int64 `protobuf:"varint,2,opt,name=employerid,proto3" json:"employerid,omitempty"`
}

func (x *FollowEmployerReq) Reset() {
	*x = FollowEmployerReq{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowEmployerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowEmployerReq) ProtoMessage() {}

func (x *FollowEmployerReq) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowEmployerReq.ProtoReflect.Descriptor instead.
func (*FollowEmployerReq) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{6}
}

func (x *FollowEmployerReq) GetJobseekerid() int64 {
	if x != nil {
		return x.Jobseekerid
	}
	return 0
}

func (x *FollowEmployerReq) GetEmployerid() int64 {
	if x != nil {
		return x.Employerid
	}
	return 0
}

type EmployerRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Country string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Website string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *EmployerRes) Reset() {
	*x = EmployerRes{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmployerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployerRes) ProtoMessage() {}

func (x *EmployerRes) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployerRes.ProtoReflect.Descriptor instead.
func (*EmployerRes) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{7}
}

func (x *EmployerRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EmployerRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmployerRes) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *EmployerRes) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EmployerRes) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *EmployerRes) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

type Employers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emp []*EmployerRes `protobuf:"bytes,1,rep,name=emp,proto3" json:"emp,omitempty"`
}

func (x *Employers) Reset() {
	*x = Employers{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Employers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employers) ProtoMessage() {}

func (x *Employers) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employers.ProtoReflect.Descriptor instead.
func (*Employers) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{8}
}

func (x *Employers) GetEmp() []*EmployerRes {
	if x != nil {
		return x.Emp
	}
	return nil
}

type Jobseekers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobseekers []*CreateJobseekerRes `protobuf:"bytes,1,rep,name=jobseekers,proto3" json:"jobseekers,omitempty"`
}

func (x *Jobseekers) Reset() {
	*x = Jobseekers{}
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Jobseekers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jobseekers) ProtoMessage() {}

func (x *Jobseekers) ProtoReflect() protoreflect.Message {
	mi := &file_utils_pb_jobseeker_jobseeker_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jobseekers.ProtoReflect.Descriptor instead.
func (*Jobseekers) Descriptor() ([]byte, []int) {
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP(), []int{9}
}

func (x *Jobseekers) GetJobseekers() []*CreateJobseekerRes {
	if x != nil {
		return x.Jobseekers
	}
	return nil
}

var File_utils_pb_jobseeker_jobseeker_proto protoreflect.FileDescriptor

var file_utils_pb_jobseeker_jobseeker_proto_rawDesc = []byte{
	0x0a, 0x22, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65,
	0x65, 0x6b, 0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f,
	0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x22, 0xb8, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x64,
	0x61, 0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x61, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x61,
	0x74, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62,
	0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0xd2, 0x01, 0x0a, 0x13, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x6a, 0x6f, 0x62,
	0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x65,
	0x6b, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0a, 0x4a, 0x53, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x55, 0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x6f,
	0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x69, 0x64, 0x22, 0x95, 0x01, 0x0a,
	0x0b, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x22, 0x2f, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x22, 0x0a, 0x03, 0x65, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x52, 0x03, 0x65, 0x6d, 0x70, 0x22, 0x45, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b,
	0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x73, 0x32, 0xdf, 0x04, 0x0a,
	0x09, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x65,
	0x65, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x65, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a,
	0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4a,
	0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a,
	0x53, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x10, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65,
	0x72, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x73,
	0x65, 0x65, 0x6b, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x65, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b,
	0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utils_pb_jobseeker_jobseeker_proto_rawDescOnce sync.Once
	file_utils_pb_jobseeker_jobseeker_proto_rawDescData = file_utils_pb_jobseeker_jobseeker_proto_rawDesc
)

func file_utils_pb_jobseeker_jobseeker_proto_rawDescGZIP() []byte {
	file_utils_pb_jobseeker_jobseeker_proto_rawDescOnce.Do(func() {
		file_utils_pb_jobseeker_jobseeker_proto_rawDescData = protoimpl.X.CompressGZIP(file_utils_pb_jobseeker_jobseeker_proto_rawDescData)
	})
	return file_utils_pb_jobseeker_jobseeker_proto_rawDescData
}

var file_utils_pb_jobseeker_jobseeker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_utils_pb_jobseeker_jobseeker_proto_goTypes = []any{
	(*CreateJobseekerReq)(nil),    // 0: job.CreateJobseekerReq
	(*CreateJobseekerRes)(nil),    // 1: job.CreateJobseekerRes
	(*JobSeekerProfileReq)(nil),   // 2: job.JobSeekerProfileReq
	(*JobSeekerProfileRes)(nil),   // 3: job.JobSeekerProfileRes
	(*JobSeekerID)(nil),           // 4: job.JobSeekerID
	(*JSLoginReq)(nil),            // 5: job.JSLoginReq
	(*FollowEmployerReq)(nil),     // 6: job.FollowEmployerReq
	(*EmployerRes)(nil),           // 7: job.EmployerRes
	(*Employers)(nil),             // 8: job.Employers
	(*Jobseekers)(nil),            // 9: job.Jobseekers
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_utils_pb_jobseeker_jobseeker_proto_depIdxs = []int32{
	10, // 0: job.CreateJobseekerReq.dateofbirth:type_name -> google.protobuf.Timestamp
	10, // 1: job.CreateJobseekerRes.dateofbirth:type_name -> google.protobuf.Timestamp
	10, // 2: job.CreateJobseekerRes.createdat:type_name -> google.protobuf.Timestamp
	1,  // 3: job.JobSeekerProfileRes.jobseeker:type_name -> job.CreateJobseekerRes
	7,  // 4: job.Employers.emp:type_name -> job.EmployerRes
	1,  // 5: job.Jobseekers.jobseekers:type_name -> job.CreateJobseekerRes
	0,  // 6: job.JobSeeker.CreateJobseeker:input_type -> job.CreateJobseekerReq
	2,  // 7: job.JobSeeker.CreateJobSeekerProfile:input_type -> job.JobSeekerProfileReq
	4,  // 8: job.JobSeeker.GetJobseekerProfile:input_type -> job.JobSeekerID
	5,  // 9: job.JobSeeker.LoginJobseeker:input_type -> job.JSLoginReq
	6,  // 10: job.JobSeeker.FollowEmployer:input_type -> job.FollowEmployerReq
	6,  // 11: job.JobSeeker.UnFollowEmployer:input_type -> job.FollowEmployerReq
	4,  // 12: job.JobSeeker.GetFollowingEmployers:input_type -> job.JobSeekerID
	4,  // 13: job.JobSeeker.GetJobseeker:input_type -> job.JobSeekerID
	11, // 14: job.JobSeeker.GetJobseekers:input_type -> google.protobuf.Empty
	1,  // 15: job.JobSeeker.CreateJobseeker:output_type -> job.CreateJobseekerRes
	3,  // 16: job.JobSeeker.CreateJobSeekerProfile:output_type -> job.JobSeekerProfileRes
	3,  // 17: job.JobSeeker.GetJobseekerProfile:output_type -> job.JobSeekerProfileRes
	1,  // 18: job.JobSeeker.LoginJobseeker:output_type -> job.CreateJobseekerRes
	7,  // 19: job.JobSeeker.FollowEmployer:output_type -> job.EmployerRes
	11, // 20: job.JobSeeker.UnFollowEmployer:output_type -> google.protobuf.Empty
	8,  // 21: job.JobSeeker.GetFollowingEmployers:output_type -> job.Employers
	1,  // 22: job.JobSeeker.GetJobseeker:output_type -> job.CreateJobseekerRes
	9,  // 23: job.JobSeeker.GetJobseekers:output_type -> job.Jobseekers
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_utils_pb_jobseeker_jobseeker_proto_init() }
func file_utils_pb_jobseeker_jobseeker_proto_init() {
	if File_utils_pb_jobseeker_jobseeker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_utils_pb_jobseeker_jobseeker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_utils_pb_jobseeker_jobseeker_proto_goTypes,
		DependencyIndexes: file_utils_pb_jobseeker_jobseeker_proto_depIdxs,
		MessageInfos:      file_utils_pb_jobseeker_jobseeker_proto_msgTypes,
	}.Build()
	File_utils_pb_jobseeker_jobseeker_proto = out.File
	file_utils_pb_jobseeker_jobseeker_proto_rawDesc = nil
	file_utils_pb_jobseeker_jobseeker_proto_goTypes = nil
	file_utils_pb_jobseeker_jobseeker_proto_depIdxs = nil
}
