package types

type Experience struct {
	Designation string
	Duration    string
	Company     string
	Location    string
	Description []string
}

type Education struct {
	Instituition string
	Location     string
	Degree       string
	Duration     string
}

type Project struct {
	Name        string
	Href        string
	TechStack   string
	SourceCode  string
	Description []string
}

type Resume interface {
	GetName() string
	GetLocation() string
	GetEmail() string
	GetMobile() string
	GetDesignation() string
	GetSummary() string

	GetSkills() map[string]string
	GetContact() map[string]string

	GetExperience() []Experience
	GetEducation() []Education
	GetProjects() []Project
}

func (r MyResume) GetName() string {
	return r.Name
}

func (r MyResume) GetLocation() string {
	return r.Location
}

func (r MyResume) GetEmail() string {
	return r.Email
}

func (r MyResume) GetContact() map[string]string {
	return r.Contact
}

func (r MyResume) GetMobile() string {
	return r.Mobile
}

func (r MyResume) GetDesignation() string {
	return r.Designation
}

func (r MyResume) GetSummary() string {
	return r.Summary
}

func (r MyResume) GetSkills() map[string]string {
	return r.Skills
}

func (r MyResume) GetExperience() []Experience {
	return r.Experience
}

func (r MyResume) GetEducation() []Education {
	return r.Education
}

func (r MyResume) GetProjects() []Project {
	return r.Projects
}

type MyResume struct {
	Name        string
	Location    string
	Email       string
	Mobile      string
	Summary     string
	Designation string
	Contact     map[string]string
	Skills      map[string]string
	Experience  []Experience
	Education   []Education
	Projects    []Project
}
