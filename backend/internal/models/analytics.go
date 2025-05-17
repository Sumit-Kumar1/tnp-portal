package models

type PlacementStatsResponse struct {
	TotalStudents       int     `json:"totalStudents"`
	PlacedStudents      int     `json:"placedStudents"`
	PlacementPercentage float64 `json:"placementPercentage"`
	TopCompanies        []struct {
		CompanyID string `json:"companyId"`
		Name      string `json:"name"`
		Count     int    `json:"count"`
	} `json:"topCompanies"`
	DepartmentWiseStats map[string]struct {
		Total      int     `json:"total"`
		Placed     int     `json:"placed"`
		Percentage float64 `json:"percentage"`
	} `json:"departmentWiseStats"`
}

type StudentPerformanceResponse struct {
	TotalApplications  int            `json:"totalApplications"`
	InterviewsAttended int            `json:"interviewsAttended"`
	OffersReceived     int            `json:"offersReceived"`
	SuccessRate        float64        `json:"successRate"`
	SkillDistribution  map[string]int `json:"skillDistribution"`
}

type CompanyTrendsResponse struct {
	TotalJobsPosted       int            `json:"totalJobsPosted"`
	ApplicationsReceived  int            `json:"applicationsReceived"`
	HiringRate            float64        `json:"hiringRate"`
	DepartmentPreferences map[string]int `json:"departmentPreferences"`
	SkillDemand           map[string]int `json:"skillDemand"`
}
