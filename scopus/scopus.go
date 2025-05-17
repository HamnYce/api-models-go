/*
Date Created: 17/05/2025
*/

package scopus

// region Scopus JSON Response
type ScopusSearchResponse struct {
	SearchResults *SearchResults `json:"search-results"`
	ServiceError  *ServiceError  `json:"service-error"`
}

type SearchResults struct {
	TotalResults float64            `json:"opensearch:totalResults,string"`
	Entries      []PublicationEntry `json:"entry"`
}

type PublicationEntry struct {
	Url                string  `json:"prism:url"`
	Title              string  `json:"dc:title"`
	Creator            string  `json:"dc:creator"`
	Doi                string  `json:"prism:doi"`
	SourceID           float64 `json:"source-id,string"`
	PublicationName    string  `json:"prism:publicationName"`
	Issn               string  `json:"prism:issn"`
	Eissn              string  `json:"prism:eIssn"`
	Volume             string  `json:"prism:volume"`
	PageRange          string  `json:"prism:pageRange"`
	IsOpenAccess       bool    `json:"openaccessFlag"`
	CoverDate          string  `json:"prism:coverDate"`
	CoverDisplayDate   string  `json:"prism:coverDisplayDate"`
	CitedByCount       float64 `json:"citedby-count,string"`
	AggregationType    string  `json:"prism:aggregationType"`
	Subtype            string  `json:"subtype"`
	SubtypeDescription string  `json:"subtypeDescription"`

	Authors []PublicationAuthor `json:"author"`
}

type PublicationAuthor struct {
	AuthorID  float64 `json:"authid,string"`
	Name      string  `json:"authname"`
	SurName   string  `json:"surname"`
	GivenName string  `json:"given-name"`
	Initials  string  `json:"initials"`
	Url       string  `json:"author-url"`
}

type ScopusSerialTiteResponse struct {
	ServiceError           *ServiceError          `json:"service-error"`
	Links                  []Link                 `json:"links"`
	SerialMetadataResponse SerialMetadataResponse `json:"serial-metadata-response"`
}

type ServiceError struct {
	Status struct {
		Code string `json:"statusCode"`
		Text string `json:"statusText"`
	} `json:"status"`
}

type Link struct {
	Valid bool   `json:"@_fa,string"`
	Ref   string `json:"@ref"`
	Href  string `json:"@href"`
	Type  string `json:"@type"`
}

type SerialMetadataResponse struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	SourceID              float64               `json:"source-id,string"`
	Title                 string                `json:"dc:title"`
	Publisher             string                `json:"dc:publisher"`
	AggregationType       string                `json:"prism:aggregationType"`
	CoverageStartYear     float64               `json:"coverageStartYear,string"`
	CoverageEndYear       float64               `json:"coverageEndYear,string"`
	Issn                  string                `json:"prism:issn"`
	Eissn                 string                `json:"prism:eIssn"`
	Url                   string                `json:"prism:url"`
	IsOpenAccess          *float64              `json:"openaccess,string"`
	HasOpenAccessArticle  *bool                 `json:"openaccessArticle"`
	HasOpenArchiveArticle *bool                 `json:"openArchiveArticle"`
	OaAllowsAuthorPaid    *bool                 `json:"oaAllowsAuthorPaid"`
	OpenAccessType        *string               `json:"openAccessType"`
	OpenAccessStartDate   *string               `json:"openAccessStartDate"`
	SubjectAreas          []SubjectArea         `json:"subject-area"`
	SnipList              SnipList              `json:"snipList"`
	SjrList               SjrList               `json:"sjrList"`
	CiteScoreYearInfoList CiteScoreYearInfoList `json:"citeScoreYearInfoList"`
}

type SubjectArea struct {
	Code        float64 `json:"@code,string"`
	Abbrev      string  `json:"@abbrev"`
	SubjectArea string  `json:"$"`
	Valid       bool    `json:"@_fa,string"`
}

type SnipList struct {
	Snip []Snip `json:"snip"`
}
type Snip struct {
	Year  float64 `json:"@year,string"`
	Value float64 `json:"$,string"`
	Valid bool    `json:"@_fa,string"`
}

type SjrList struct {
	Sjr []Sjr `json:"sjr"`
}
type Sjr struct {
	Year  float64 `json:"@year,string"`
	Value float64 `json:"$,string"`
	Valid bool    `json:"@_fa,string"`
}

type CiteScoreYearInfoList struct {
	CiteScoreCurrentMetric     float64             `json:"citeScoreCurrentMetric,string"`
	CiteScoreCurrentMetricYear float64             `json:"citeScoreCurrentMetricYear,string"`
	CiteScoreTracker           float64             `json:"citeScoreTracker,string"`
	CiteScoreTrackerYear       float64             `json:"citeScoreTrackerYear,string"`
	CiteScoreYearInfo          []CiteScoreYearInfo `json:"citeScoreYearInfo"`
}

type CiteScoreYearInfo struct {
	Year                     float64                    `json:"@year,string"`
	Status                   string                     `json:"@status"`
	Valid                    bool                       `json:"@_fa,string"`
	CiteScoreInformationList []CiteScoreInformationList `json:"citeScoreInformationList"`
}

type CiteScoreInformationList struct {
	Valid         bool            `json:"@_fa,string"`
	CiteScoreInfo []CiteScoreInfo `json:"citeScoreInfo"`
}

type CiteScoreInfo struct {
	DocType              string                 `json:"docType"`
	ScholarlyOutput      float64                `json:"scholarlyOutput,string"`
	CitationCount        float64                `json:"citationCount,string"`
	CiteScore            float64                `json:"citeScore,string"`
	PercentCited         float64                `json:"percentCited,string"`
	Valid                bool                   `json:"@_fa,string"`
	CiteScoreSubjectRank []CiteScoreSubjectRank `json:"citeScoreSubjectRank"`
}

type CiteScoreSubjectRank struct {
	SubjectCode float64 `json:"subjectCode,string"`
	Rank        float64 `json:"rank,string"`
	Percentile  float64 `json:"percentile,string"`
	Valid       bool    `json:"@_fa,string"`
}

// endregion Scopus JSON Response
