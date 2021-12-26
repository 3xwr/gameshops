package model

import "time"

type SteamRequestModel struct {
	Name string `json:"name"`
}

type GamePriceResponse struct {
	StoreName    string    `json:"store_name"`
	StoreAppID   int       `json:"store_app_id,omitempty"`
	StoreAppName string    `json:"store_app_name"`
	StorePrice   string    `json:"store_price,omitempty"`
	StoreImage   string    `json:"store_image,omitempty"`
	StoreAppURL  string    `json:"store_app_url,omitempty"`
	Status       string    `json:"status,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

type SteamResponseModel struct {
	Num945360 struct {
		Success bool `json:"success"`
		Data    struct {
			Type                string `json:"type"`
			Name                string `json:"name"`
			SteamAppid          int    `json:"steam_appid"`
			RequiredAge         int    `json:"required_age"`
			IsFree              bool   `json:"is_free"`
			ControllerSupport   string `json:"controller_support"`
			DetailedDescription string `json:"detailed_description"`
			AboutTheGame        string `json:"about_the_game"`
			ShortDescription    string `json:"short_description"`
			SupportedLanguages  string `json:"supported_languages"`
			Reviews             string `json:"reviews"`
			HeaderImage         string `json:"header_image"`
			Website             string `json:"website"`
			PcRequirements      struct {
				Minimum string `json:"minimum"`
			} `json:"pc_requirements"`
			MacRequirements   []interface{} `json:"mac_requirements"`
			LinuxRequirements []interface{} `json:"linux_requirements"`
			Developers        []string      `json:"developers"`
			Publishers        []string      `json:"publishers"`
			PriceOverview     struct {
				Currency         string `json:"currency"`
				Initial          int    `json:"initial"`
				Final            int    `json:"final"`
				DiscountPercent  int    `json:"discount_percent"`
				InitialFormatted string `json:"initial_formatted"`
				FinalFormatted   string `json:"final_formatted"`
			} `json:"price_overview"`
			Packages      []int `json:"packages"`
			PackageGroups []struct {
				Name                    string `json:"name"`
				Title                   string `json:"title"`
				Description             string `json:"description"`
				SelectionText           string `json:"selection_text"`
				SaveText                string `json:"save_text"`
				DisplayType             int    `json:"display_type"`
				IsRecurringSubscription string `json:"is_recurring_subscription"`
				Subs                    []struct {
					Packageid                int    `json:"packageid"`
					PercentSavingsText       string `json:"percent_savings_text"`
					PercentSavings           int    `json:"percent_savings"`
					OptionText               string `json:"option_text"`
					OptionDescription        string `json:"option_description"`
					CanGetFreeLicense        string `json:"can_get_free_license"`
					IsFreeLicense            bool   `json:"is_free_license"`
					PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
				} `json:"subs"`
			} `json:"package_groups"`
			Platforms struct {
				Windows bool `json:"windows"`
				Mac     bool `json:"mac"`
				Linux   bool `json:"linux"`
			} `json:"platforms"`
			Metacritic struct {
				Score int    `json:"score"`
				URL   string `json:"url"`
			} `json:"metacritic"`
			Categories []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
			} `json:"categories"`
			Genres []struct {
				ID          string `json:"id"`
				Description string `json:"description"`
			} `json:"genres"`
			Screenshots []struct {
				ID            int    `json:"id"`
				PathThumbnail string `json:"path_thumbnail"`
				PathFull      string `json:"path_full"`
			} `json:"screenshots"`
			Movies []struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				Thumbnail string `json:"thumbnail"`
				Webm      struct {
					Num480 string `json:"480"`
					Max    string `json:"max"`
				} `json:"webm"`
				Mp4 struct {
					Num480 string `json:"480"`
					Max    string `json:"max"`
				} `json:"mp4"`
				Highlight bool `json:"highlight"`
			} `json:"movies"`
			Recommendations struct {
				Total int `json:"total"`
			} `json:"recommendations"`
			Achievements struct {
				Total       int `json:"total"`
				Highlighted []struct {
					Name string `json:"name"`
					Path string `json:"path"`
				} `json:"highlighted"`
			} `json:"achievements"`
			ReleaseDate struct {
				ComingSoon bool   `json:"coming_soon"`
				Date       string `json:"date"`
			} `json:"release_date"`
			SupportInfo struct {
				URL   string `json:"url"`
				Email string `json:"email"`
			} `json:"support_info"`
			Background         string `json:"background"`
			ContentDescriptors struct {
				Ids   []interface{} `json:"ids"`
				Notes interface{}   `json:"notes"`
			} `json:"content_descriptors"`
		} `json:"data"`
	} `json:"945360"`
}

type GOGResponseModel struct {
	Products []struct {
		CustomAttributes []interface{} `json:"customAttributes"`
		Developer        string        `json:"developer"`
		Publisher        string        `json:"publisher"`
		Gallery          []string      `json:"gallery"`
		Video            struct {
			ID       string `json:"id"`
			Provider string `json:"provider"`
		} `json:"video"`
		SupportedOperatingSystems []string    `json:"supportedOperatingSystems"`
		Genres                    []string    `json:"genres"`
		GlobalReleaseDate         interface{} `json:"globalReleaseDate"`
		IsTBA                     bool        `json:"isTBA"`
		Price                     struct {
			Amount                     string      `json:"amount"`
			BaseAmount                 string      `json:"baseAmount"`
			FinalAmount                string      `json:"finalAmount"`
			IsDiscounted               bool        `json:"isDiscounted"`
			DiscountPercentage         int         `json:"discountPercentage"`
			DiscountDifference         string      `json:"discountDifference"`
			Symbol                     string      `json:"symbol"`
			IsFree                     bool        `json:"isFree"`
			Discount                   int         `json:"discount"`
			IsBonusStoreCreditIncluded bool        `json:"isBonusStoreCreditIncluded"`
			BonusStoreCreditAmount     string      `json:"bonusStoreCreditAmount"`
			PromoID                    interface{} `json:"promoId"`
		} `json:"price"`
		IsDiscounted    bool        `json:"isDiscounted"`
		IsInDevelopment bool        `json:"isInDevelopment"`
		ID              int         `json:"id"`
		ReleaseDate     interface{} `json:"releaseDate"`
		Availability    struct {
			IsAvailable          bool `json:"isAvailable"`
			IsAvailableInAccount bool `json:"isAvailableInAccount"`
		} `json:"availability"`
		SalesVisibility struct {
			IsActive   bool `json:"isActive"`
			FromObject struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"fromObject"`
			From     int `json:"from"`
			ToObject struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"toObject"`
			To int `json:"to"`
		} `json:"salesVisibility"`
		Buyable    bool   `json:"buyable"`
		Title      string `json:"title"`
		Image      string `json:"image"`
		URL        string `json:"url"`
		SupportURL string `json:"supportUrl"`
		ForumURL   string `json:"forumUrl"`
		WorksOn    struct {
			Windows bool `json:"Windows"`
			Mac     bool `json:"Mac"`
			Linux   bool `json:"Linux"`
		} `json:"worksOn"`
		Category         string        `json:"category"`
		OriginalCategory string        `json:"originalCategory"`
		Rating           int           `json:"rating"`
		Type             int           `json:"type"`
		IsComingSoon     bool          `json:"isComingSoon"`
		IsPriceVisible   bool          `json:"isPriceVisible"`
		IsMovie          bool          `json:"isMovie"`
		IsGame           bool          `json:"isGame"`
		Slug             string        `json:"slug"`
		IsWishlistable   bool          `json:"isWishlistable"`
		ExtraInfo        []interface{} `json:"extraInfo"`
		AgeLimit         int           `json:"ageLimit"`
	} `json:"products"`
	Ts               interface{} `json:"ts"`
	Page             int         `json:"page"`
	TotalPages       int         `json:"totalPages"`
	TotalResults     string      `json:"totalResults"`
	TotalGamesFound  int         `json:"totalGamesFound"`
	TotalMoviesFound int         `json:"totalMoviesFound"`
}

type SteamPayResponseModel struct {
	Error    int `json:"error"`
	Products []struct {
		URL         string `json:"url"`
		Title       string `json:"title"`
		NumInStock  int    `json:"num_in_stock"`
		Activation  string `json:"activation"`
		IsAvailable bool   `json:"is_available"`
		Image       string `json:"image"`
		Prices      struct {
			Rub int     `json:"rub"`
			Usd float64 `json:"usd"`
			Eur float64 `json:"eur"`
			Grn int     `json:"grn"`
		} `json:"prices"`
	} `json:"products"`
}

type PlatiruResponseModel struct {
	Pagenum    int `json:"Pagenum"`
	Pagesize   int `json:"Pagesize"`
	Totalpages int `json:"Totalpages"`
	Items      []struct {
		ID                           int     `json:"id"`
		Name                         string  `json:"name"`
		NameEng                      string  `json:"name_eng"`
		NameTranslit                 string  `json:"name_translit"`
		NameTranslitEng              string  `json:"name_translit_eng"`
		PartnerCommiss               float64 `json:"partner_commiss"`
		PriceEur                     float64 `json:"price_eur"`
		PriceRur                     float64 `json:"price_rur"`
		PriceUah                     float64 `json:"price_uah"`
		PriceUsd                     float64 `json:"price_usd"`
		SectionID                    int     `json:"section_id"`
		URL                          string  `json:"url"`
		Description                  string  `json:"description"`
		DescriptionEng               string  `json:"description_eng"`
		Image                        string  `json:"image"`
		SellerID                     int     `json:"seller_id"`
		SellerName                   string  `json:"seller_name"`
		SellerRating                 float64 `json:"seller_rating"`
		Numsold                      int     `json:"numsold"`
		NumsoldHidden                int     `json:"numsold_hidden"`
		CountPositiveresponses       int     `json:"count_positiveresponses"`
		CountPositiveresponsesHidden int     `json:"count_positiveresponses_hidden"`
		CountNegativeresponses       int     `json:"count_negativeresponses"`
		CountNegativeresponsesHidden int     `json:"count_negativeresponses_hidden"`
		CountReturns                 int     `json:"count_returns"`
		CountReturnsHidden           int     `json:"count_returns_hidden"`
		SalesForm                    string  `json:"sales_form"`
		SaleInfo                     struct {
			CommonBasePrice string `json:"common_base_price"`
			CommonPriceUsd  string `json:"common_price_usd"`
			CommonPriceRur  string `json:"common_price_rur"`
			CommonPriceEur  string `json:"common_price_eur"`
			CommonPriceUah  string `json:"common_price_uah"`
			SalePercent     string `json:"sale_percent"`
		} `json:"sale_info"`
		HideSales interface{} `json:"hide_sales"`
		Payadv    float64     `json:"payadv"`
	} `json:"items"`
	Total int `json:"total"`
}
