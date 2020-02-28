package model

type Repos struct {
	Total int    `json:total`
	Repos []Repo `json:repos`
}

type Repo struct {
	SearchScore    string      `json:searchScore`
	Contact        Contact     `json:contact`
	Date           Date        `json:date`
	DownloadURL    string      `json:downladURL`
	HomepageURL    string      `json:homepageURL`
	DisclaimerURL  string      `json:disclaimerURL`
	DisclaimerText string      `json:disclaimerText`
	LaborHours     int         `json:laborHours`
	Languages      []string    `json:languages`
	Name           string      `json:name`
	Version        string      `json:version`
	Description    string      `json:description`
	Organization   string      `json:organization`
	Permissions    Permissions `json:permissions`
	RepositoryURL  string      `json:repositoryURL`
	Status         string      `json:status`
	Tags           []string    `json:tags`
	VCS            string      `json:vcs`
	Agency         Agency      `json:agency`
	RepoID         string      `json:repoID`
	Score          float32     `json:score`
}

type Contact struct {
	URL   string `json:URL`
	Email string `json:email`
	Name  string `json:name`
	Phone string `json:phone`
}

type Date struct {
	Created      string `json:created`
	LastModified string `json:lastModified`
}

type Permissions struct {
	Licenses  Licenses `json:licenses`
	UsageType string   `json:usageType`
}

type Licenses struct {
	Name string `json:name`
	URL  string `json:URL`
}

type Agency struct {
	Name                string       `json:name`
	Acronym             string       `json:acronym`
	Website             string       `json:website`
	CodeURL             string       `json:codeURL`
	FallbackFile        string       `json:fallback_file`
	Requirements        Requirements `json:requirements`
	ComplianceDashboard bool         `json:complianceDashboard`
}

type Requirements struct {
	AgencyWidePolicy      int     `json:agencyWidePolicy`
	OpenSourceRequirement int     `json:openSourceRequirement`
	InventoryRequirement  int     `json:inventoryRequirement`
	SchemaFormat          float32 `json:schemaFormat`
	OverallCompliance     float32 `json:overallCompliance`
}
