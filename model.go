package main

import "time"

type GetCraftResponse struct {
	CategoryProductsCount struct {
		Five9316F4F60613D7F4A005Bd0 int `json:"59316f4f60613d7f4a005bd0"`
		Five9316F7860613D7F4A005Bd4 int `json:"59316f7860613d7f4a005bd4"`
		Five9316F7F60613D7F4A005Bd6 int `json:"59316f7f60613d7f4a005bd6"`
		Five9316F8660613D7F4A005Bd8 int `json:"59316f8660613d7f4a005bd8"`
		Five9316F8E60613D7F4A005Bda int `json:"59316f8e60613d7f4a005bda"`
		Five9774C2082Dd76D441D7Cddd int `json:"59774c2082dd76d441d7cddd"`
		Five99722Dd743D2B70B2F1B048 int `json:"599722dd743d2b70b2f1b048"`
		Five99722Dd743D2B70B2F1B049 int `json:"599722dd743d2b70b2f1b049"`
		Five99722Dd743D2B70B2F1B04A int `json:"599722dd743d2b70b2f1b04a"`
		FiveA7997Dc10A769001A2D8C62 int `json:"5a7997dc10a769001a2d8c62"`
		FiveB712E90B6C952001C883B31 int `json:"5b712e90b6c952001c883b31"`
		FiveE1E85Ed37Bdc8001C002F65 int `json:"5e1e85ed37bdc8001c002f65"`
		FiveF33Bd3B39Bb830022F81Ff4 int `json:"5f33bd3b39bb830022f81ff4"`
	} `json:"categoryProductsCount"`
	FilterValues struct {
		Cpv struct {
			Max int `json:"max"`
			Min int `json:"min"`
		} `json:"CPV"`
		CPVUsd struct {
			Max float64 `json:"max"`
			Min int     `json:"min"`
		} `json:"CPVUsd"`
		BusinessType struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"businessType"`
		BusinessTypeWithSlug struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"businessTypeWithSlug"`
		Category struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"category"`
		CategoryIDWithSlug struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"categoryIdWithSlug"`
		CategoryWithSlug struct {
			Terms interface{} `json:"terms"`
		} `json:"categoryWithSlug"`
		City struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"city"`
		Country struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"country"`
		CountryWithSlug struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"countryWithSlug"`
		EngagementRate struct {
			Max int `json:"max"`
			Min int `json:"min"`
		} `json:"engagementRate"`
		IgFollowers struct {
			Max int `json:"max"`
			Min int `json:"min"`
		} `json:"igFollowers"`
		IndustryExpertise struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"industryExpertise"`
		IndustryExpertiseWithSlug struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"industryExpertiseWithSlug"`
		Occupation struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"occupation"`
		Price struct {
			Max int64 `json:"max"`
			Min int   `json:"min"`
		} `json:"price"`
		PriceUsd struct {
			Max int64 `json:"max"`
			Min int   `json:"min"`
		} `json:"priceUsd"`
		ProfileIDName struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"profileIdName"`
		ServiceLanguage struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"serviceLanguage"`
		ServiceLanguageWithSlug struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"serviceLanguageWithSlug"`
		SubCategory struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"subCategory"`
		SubCategoryIDWithSlug struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"subCategoryIdWithSlug"`
		SubCategoryWithSlug struct {
			Terms []struct {
				Key   string `json:"key"`
				Count int    `json:"count"`
			} `json:"terms"`
		} `json:"subCategoryWithSlug"`
		Viewers struct {
			Max int `json:"max"`
			Min int `json:"min"`
		} `json:"viewers"`
		YoutubeFollowers struct {
			Max int `json:"max"`
			Min int `json:"min"`
		} `json:"youtubeFollowers"`
	} `json:"filterValues"`
	Method   string `json:"method"`
	Products []struct {
		Cpv                       int       `json:"CPV,omitempty"`
		CPVUsd                    float64   `json:"CPVUsd,omitempty"`
		Availability              bool      `json:"availability"`
		Banner                    string    `json:"banner"`
		BusinessType              string    `json:"businessType"`
		BusinessTypeWithSlug      string    `json:"businessTypeWithSlug"`
		Category                  string    `json:"category"`
		CategoryDesc              string    `json:"categoryDesc"`
		CategoryID                string    `json:"categoryId"`
		CategoryIDWithSlug        string    `json:"categoryIdWithSlug"`
		CategoryIds               []string  `json:"categoryIds"`
		CategorySlug              string    `json:"categorySlug"`
		City                      string    `json:"city"`
		Country                   string    `json:"country"`
		CountryWithSlug           string    `json:"countryWithSlug"`
		CreatedAt                 time.Time `json:"createdAt"`
		Currency                  string    `json:"currency"`
		Description               string    `json:"description"`
		FacebookURL               string    `json:"facebookURL"`
		FollowerShow              []string  `json:"followerShow,omitempty"`
		ID                        string    `json:"id"`
		IgFollowers               int       `json:"igFollowers"`
		ImageLink                 string    `json:"image_link"`
		IndustryExpertise         []string  `json:"industryExpertise"`
		IndustryExpertiseWithSlug []string  `json:"industryExpertiseWithSlug"`
		InnerHits                 struct {
			Services struct {
				Hits struct {
					Total    int     `json:"total"`
					MaxScore float64 `json:"max_score"`
					Hits     []struct {
						Score     float64     `json:"_score"`
						Index     string      `json:"_index"`
						Type      string      `json:"_type"`
						ID        string      `json:"_id"`
						UID       string      `json:"_uid"`
						Routing   string      `json:"_routing"`
						Parent    string      `json:"_parent"`
						Version   interface{} `json:"_version"`
						Sort      interface{} `json:"sort"`
						Highlight interface{} `json:"highlight"`
						Source    struct {
							ServiceLanguage           string      `json:"serviceLanguage"`
							FollowerShow              []string    `json:"followerShow"`
							Country                   string      `json:"country"`
							SubCategory               string      `json:"subCategory"`
							Occupation                []string    `json:"occupation"`
							LinkedinURL               string      `json:"linkedinURL"`
							SubCategoryIDWithSlug     string      `json:"subCategoryIdWithSlug"`
							TwitterURL                string      `json:"twitterURL"`
							Cpv                       int         `json:"CPV"`
							ProfileImage              string      `json:"profileImage"`
							Type                      string      `json:"type"`
							CreatedAt                 time.Time   `json:"createdAt"`
							YoutubeFollowers          int         `json:"youtubeFollowers"`
							Price                     int         `json:"price"`
							YoutubeURL                string      `json:"youtubeURL"`
							TwitterFollowers          int         `json:"twitterFollowers"`
							ID                        string      `json:"id"`
							PriceTo                   int         `json:"priceTo"`
							Slug                      string      `json:"slug"`
							UpdatedAt                 time.Time   `json:"updatedAt"`
							All                       string      `json:"all"`
							ImageLink                 string      `json:"image_link"`
							PriceUsd                  float64     `json:"priceUsd"`
							TypeSlug                  string      `json:"typeSlug"`
							SubCategoryID             string      `json:"subCategoryId"`
							CategorySlug              string      `json:"categorySlug"`
							WeightBoosting            int         `json:"weightBoosting"`
							Viewers                   int         `json:"viewers"`
							DeletedAt                 interface{} `json:"deletedAt"`
							ProfileID                 string      `json:"profileId"`
							IndustryExpertise         []string    `json:"industryExpertise"`
							ProfileIDName             string      `json:"profileIdName"`
							Name                      string      `json:"name"`
							SubCategoryWithSlug       string      `json:"subCategoryWithSlug"`
							TypeID                    string      `json:"typeId"`
							IsPremium                 bool        `json:"isPremium"`
							ProfileName               string      `json:"profileName"`
							City                      string      `json:"city"`
							BusinessTypeWithSlug      string      `json:"businessTypeWithSlug"`
							Link                      string      `json:"link"`
							Description               string      `json:"description"`
							Availability              bool        `json:"availability"`
							Title                     string      `json:"title"`
							IsPriceHidden             bool        `json:"isPriceHidden"`
							CategoryDesc              string      `json:"categoryDesc"`
							CategoryIDWithSlug        string      `json:"categoryIdWithSlug"`
							ProfileLink               string      `json:"profileLink"`
							FacebookURL               string      `json:"facebookURL"`
							LastProcessedAt           time.Time   `json:"lastProcessedAt"`
							ProfileDesc               string      `json:"profileDesc"`
							PriceFrom                 int         `json:"priceFrom"`
							CPVUsd                    float64     `json:"CPVUsd"`
							Currency                  string      `json:"currency"`
							CategoryTree              []string    `json:"categoryTree"`
							SubCategorySlug           string      `json:"subCategorySlug"`
							IgFollowers               int         `json:"igFollowers"`
							Banner                    string      `json:"banner"`
							ServiceLanguageWithSlug   string      `json:"serviceLanguageWithSlug"`
							IndustryExpertiseWithSlug []string    `json:"industryExpertiseWithSlug"`
							CategoryIds               []string    `json:"categoryIds"`
							CountryWithSlug           string      `json:"countryWithSlug"`
							InstagramURL              string      `json:"instagramURL"`
							Category                  string      `json:"category"`
							BusinessType              string      `json:"businessType"`
							CategoryID                string      `json:"categoryId"`
						} `json:"_source"`
						Fields         interface{} `json:"fields"`
						Explanation    interface{} `json:"_explanation"`
						MatchedQueries interface{} `json:"matched_queries"`
						InnerHits      interface{} `json:"inner_hits"`
						Nested         interface{} `json:"_nested"`
					} `json:"hits"`
				} `json:"hits"`
			} `json:"services"`
		} `json:"inner_hits"`
		InstagramURL            string   `json:"instagramURL"`
		IsPremium               bool     `json:"isPremium"`
		IsPriceHidden           bool     `json:"isPriceHidden"`
		Link                    string   `json:"link"`
		LinkedinURL             string   `json:"linkedinURL"`
		Name                    string   `json:"name"`
		Occupation              []string `json:"occupation"`
		Price                   int      `json:"price"`
		PriceFrom               int      `json:"priceFrom"`
		PriceTo                 int      `json:"priceTo"`
		PriceUsd                float64  `json:"priceUsd"`
		ProfileDesc             string   `json:"profileDesc"`
		ProfileID               string   `json:"profileId"`
		ProfileIDName           string   `json:"profileIdName"`
		ProfileImage            string   `json:"profileImage"`
		ProfileLink             string   `json:"profileLink"`
		ProfileName             string   `json:"profileName"`
		ServiceLanguage         string   `json:"serviceLanguage"`
		ServiceLanguageWithSlug string   `json:"serviceLanguageWithSlug"`
		Slug                    string   `json:"slug"`
		SubCategory             string   `json:"subCategory"`
		SubCategoryID           string   `json:"subCategoryId"`
		SubCategoryIDWithSlug   string   `json:"subCategoryIdWithSlug"`
		SubCategorySlug         string   `json:"subCategorySlug"`
		SubCategoryWithSlug     string   `json:"subCategoryWithSlug"`
		Title                   string   `json:"title"`
		TwitterFollowers        int      `json:"twitterFollowers"`
		TwitterURL              string   `json:"twitterURL"`
		Type                    string   `json:"type"`
		TypeID                  string   `json:"typeId"`
		TypeSlug                string   `json:"typeSlug"`
		Viewers                 int      `json:"viewers,omitempty"`
		YoutubeFollowers        int      `json:"youtubeFollowers"`
		YoutubeURL              string   `json:"youtubeURL"`
		AvgCreatorRating        int      `json:"avgCreatorRating,omitempty"`
	} `json:"products"`
	Promoted      []interface{} `json:"promoted"`
	SearchID      string        `json:"searchId"`
	Total         int           `json:"total"`
	TotalCollapse int           `json:"totalCollapse"`
}

/**
influencer name, details2 yg di display disana ttg influencer profile
- nama
- country
- city
- occupation
instagram link mereka
linkedinurl
twitterurl
youtubeurl
youtubesubscriber
twitterfollower
igfollower
cpv
views
packages: [
	package offered,
	package rate
]
*/

type InfluencerCSV struct {
	Name               string
	Country            string
	City               string
	Occupation         string
	InstagramURL       string
	LinkedInURL        string
	TwitterURL         string
	YoutubeURL         string
	YoutubeSubscribers int
	TwitterFollowers   int
	InstagramFollowers int
	Cpv                int
	Views              int
	Packages           string
}
