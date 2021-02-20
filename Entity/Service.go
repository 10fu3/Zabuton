package Entity

type service struct {
	uuid string
	name string
	subDomain string
	adminID string
}

type ServiceBuilder struct {
	s service
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{}
}

func EditService(s Service) *ServiceBuilder {
	b := ServiceBuilder{}
	b.s.uuid = s.GetUUID()
	b.s.adminID = s.GetAdminID()
	b.s.subDomain = s.GetSubDomain()
	b.s.name = s.GetName()
	return &b
}

func (r *ServiceBuilder) SetName(name string) *ServiceBuilder {
	r.s.name = name
	return r
}

func (r *ServiceBuilder) SetUUID(id string) *ServiceBuilder {
	r.s.uuid = id
	return r
}

func (r *ServiceBuilder) SetSubDomain(subDomain string) *ServiceBuilder {
	r.s.subDomain = subDomain
	return r
}

func (r *ServiceBuilder) SetAdminID (adminID string) *ServiceBuilder {
	r.s.adminID = adminID
	return r
}

func (r *ServiceBuilder) Build() Service {
	return &r.s
}

func (s *service) GetName() string {
	return s.name
}

func (s *service) GetSubDomain() string {
	return s.subDomain
}

func (s *service) GetAdminID() string {
	return s.adminID
}

func (s *service) GetUUID() string {
	return s.uuid
}

type Service interface {
	GetName() string
	GetSubDomain() string
	GetAdminID() string
	GetUUID() string
}


