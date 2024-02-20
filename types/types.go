package types

type Experience struct {
	Designation string   `json:"designation"`
	Duration    string   `json:"duration"`
	Company     string   `json:"company"`
	Location    string   `json:"location"`
	Description []string `json:"description"`
}

type Education struct {
	Instituition string `json:"instituition"`
	Location     string `json:"location"`
	Degree       string `json:"degree"`
	Duration     string `json:"duration"`
}

type Project struct {
	Name        string   `json:"name"`
	Href        string   `json:"href"`
	TechStack   string   `json:"techstack"`
	SourceCode  string   `json:"sourcehref"`
	Description []string `json:"description"`
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
	Name        string            `json:"name"`
	Location    string            `json:"location"`
	Email       string            `json:"email"`
	Mobile      string            `json:"mobile"`
	Summary     string            `json:"summary"`
	Designation string            `json:"designation"`
	Contact     map[string]string `json:"contact"`
	Skills      map[string]string `json:"skills"`
	Experience  []Experience      `json:"experience"`
	Education   []Education       `json:"education"`
	Projects    []Project         `json:"projects"`
}
