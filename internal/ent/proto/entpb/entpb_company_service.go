// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	ent "refernet/internal/ent"
	company "refernet/internal/ent/company"
	workexperience "refernet/internal/ent/workexperience"
	strings "strings"
)

// CompanyService implements CompanyServiceServer
type CompanyService struct {
	client *ent.Client
	UnimplementedCompanyServiceServer
}

// NewCompanyService returns a new CompanyService
func NewCompanyService(client *ent.Client) *CompanyService {
	return &CompanyService{
		client: client,
	}
}

func toProtoCompany_Size(e company.Size) Company_Size {
	if v, ok := Company_Size_value[strings.ToUpper(string(e))]; ok {
		return Company_Size(v)
	}
	return Company_Size(0)
}

func toEntCompany_Size(e Company_Size) company.Size {
	if v, ok := Company_Size_name[int32(e)]; ok {
		return company.Size(strings.ToLower(v))
	}
	return ""
}

// toProtoCompany transforms the ent type to the pb type
func toProtoCompany(e *ent.Company) (*Company, error) {
	v := &Company{}
	createdat := timestamppb.New(e.CreatedAt)
	v.CreatedAt = createdat
	foundedat := int32(e.FoundedAt)
	v.FoundedAt = foundedat
	id := int32(e.ID)
	v.Id = id
	logourl := e.LogoURL
	v.LogoUrl = logourl
	name := e.Name
	v.Name = name
	overview := e.Overview
	v.Overview = overview
	size := toProtoCompany_Size(e.Size)
	v.Size = size
	updatedat := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updatedat
	website := e.Website
	v.Website = website
	for _, edg := range e.Edges.Staffs {
		id := int32(edg.ID)
		v.Staffs = append(v.Staffs, &WorkExperience{
			Id: id,
		})
	}
	return v, nil
}

// Create implements CompanyServiceServer.Create
func (svc *CompanyService) Create(ctx context.Context, req *CreateCompanyRequest) (*Company, error) {
	company := req.GetCompany()
	m := svc.client.Company.Create()
	companyCreatedAt := runtime.ExtractTime(company.GetCreatedAt())
	m.SetCreatedAt(companyCreatedAt)
	companyFoundedAt := int(company.GetFoundedAt())
	m.SetFoundedAt(companyFoundedAt)
	companyLogoURL := company.GetLogoUrl()
	m.SetLogoURL(companyLogoURL)
	companyName := company.GetName()
	m.SetName(companyName)
	companyOverview := company.GetOverview()
	m.SetOverview(companyOverview)
	companySize := toEntCompany_Size(company.GetSize())
	m.SetSize(companySize)
	companyUpdatedAt := runtime.ExtractTime(company.GetUpdatedAt())
	m.SetUpdatedAt(companyUpdatedAt)
	companyWebsite := company.GetWebsite()
	m.SetWebsite(companyWebsite)
	for _, item := range company.GetStaffs() {
		staffs := int(item.GetId())
		m.AddStaffIDs(staffs)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoCompany(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements CompanyServiceServer.Get
func (svc *CompanyService) Get(ctx context.Context, req *GetCompanyRequest) (*Company, error) {
	var (
		err error
		get *ent.Company
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetCompanyRequest_VIEW_UNSPECIFIED, GetCompanyRequest_BASIC:
		get, err = svc.client.Company.Get(ctx, id)
	case GetCompanyRequest_WITH_EDGE_IDS:
		get, err = svc.client.Company.Query().
			Where(company.ID(id)).
			WithStaffs(func(query *ent.WorkExperienceQuery) {
				query.Select(workexperience.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoCompany(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements CompanyServiceServer.Update
func (svc *CompanyService) Update(ctx context.Context, req *UpdateCompanyRequest) (*Company, error) {
	company := req.GetCompany()
	companyID := int(company.GetId())
	m := svc.client.Company.UpdateOneID(companyID)
	companyFoundedAt := int(company.GetFoundedAt())
	m.SetFoundedAt(companyFoundedAt)
	companyLogoURL := company.GetLogoUrl()
	m.SetLogoURL(companyLogoURL)
	companyName := company.GetName()
	m.SetName(companyName)
	companyOverview := company.GetOverview()
	m.SetOverview(companyOverview)
	companySize := toEntCompany_Size(company.GetSize())
	m.SetSize(companySize)
	companyUpdatedAt := runtime.ExtractTime(company.GetUpdatedAt())
	m.SetUpdatedAt(companyUpdatedAt)
	companyWebsite := company.GetWebsite()
	m.SetWebsite(companyWebsite)
	for _, item := range company.GetStaffs() {
		staffs := int(item.GetId())
		m.AddStaffIDs(staffs)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoCompany(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements CompanyServiceServer.Delete
func (svc *CompanyService) Delete(ctx context.Context, req *DeleteCompanyRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Company.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}