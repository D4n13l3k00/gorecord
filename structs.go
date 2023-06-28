package main

type StationsAPIResult struct {
	Result struct {
		Tags []struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			DetailPicture string `json:"detail_picture"`
			Picture       string `json:"picture"`
			Svg           string `json:"svg"`
			Pdf           string `json:"pdf"`
		} `json:"tags"`
		Genre []struct {
			ID            int         `json:"id"`
			Name          string      `json:"name"`
			DetailPicture interface{} `json:"detail_picture"`
			Picture       interface{} `json:"picture"`
			Svg           interface{} `json:"svg"`
			Pdf           interface{} `json:"pdf"`
		} `json:"genre"`
		Stations []struct {
			ID              int         `json:"id"`
			Prefix          string      `json:"prefix"`
			Title           string      `json:"title"`
			Tooltip         string      `json:"tooltip"`
			Sort            int         `json:"sort"`
			BgColor         interface{} `json:"bg_color"`
			BgImage         string      `json:"bg_image"`
			BgImageMobile   string      `json:"bg_image_mobile"`
			SvgOutline      string      `json:"svg_outline"`
			SvgFill         string      `json:"svg_fill"`
			PdfOutline      string      `json:"pdf_outline"`
			PdfFill         string      `json:"pdf_fill"`
			ShortTitle      string      `json:"short_title"`
			IconGray        string      `json:"icon_gray"`
			IconFillColored string      `json:"icon_fill_colored"`
			IconFillWhite   string      `json:"icon_fill_white"`
			New             bool        `json:"new"`
			NewDate         interface{} `json:"new_date,omitempty"`
			Stream64        string      `json:"stream_64"`
			Stream128       string      `json:"stream_128"`
			Stream320       string      `json:"stream_320"`
			StreamHls       string      `json:"stream_hls"`
			Genre           []struct {
				ID            int         `json:"id"`
				Name          string      `json:"name"`
				DetailPicture interface{} `json:"detail_picture"`
				Picture       interface{} `json:"picture"`
				Svg           interface{} `json:"svg"`
				Pdf           interface{} `json:"pdf"`
			} `json:"genre"`
			DetailPageURL string      `json:"detail_page_url"`
			ShareURL      string      `json:"shareUrl"`
			Mark          interface{} `json:"mark"`
			Updated       string      `json:"updated"`
		} `json:"stations"`
	} `json:"result"`
}
