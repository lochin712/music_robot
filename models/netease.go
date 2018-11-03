package models

// NetEaseResponse NetEaseResponse
type NetEaseResponse struct {
	Result struct {
		Songs []struct {
			Name        string        `json:"name"`
			ID          int           `json:"id"`
			Position    int           `json:"position"`
			Alias       []interface{} `json:"alias"`
			Status      int           `json:"status"`
			Fee         int           `json:"fee"`
			CopyrightID int           `json:"copyrightId"`
			Disc        string        `json:"disc"`
			No          int           `json:"no"`
			Artists     []struct {
				Name      string        `json:"name"`
				ID        int           `json:"id"`
				PicID     int           `json:"picId"`
				Img1V1ID  int           `json:"img1v1Id"`
				BriefDesc string        `json:"briefDesc"`
				PicURL    string        `json:"picUrl"`
				Img1V1URL string        `json:"img1v1Url"`
				AlbumSize int           `json:"albumSize"`
				Alias     []interface{} `json:"alias"`
				Trans     string        `json:"trans"`
				MusicSize int           `json:"musicSize"`
			} `json:"artists"`
			Album struct {
				Name        string `json:"name"`
				ID          int    `json:"id"`
				Type        string `json:"type"`
				Size        int    `json:"size"`
				PicID       int64  `json:"picId"`
				BlurPicURL  string `json:"blurPicUrl"`
				CompanyID   int    `json:"companyId"`
				Pic         int64  `json:"pic"`
				PicURL      string `json:"picUrl"`
				PublishTime int64  `json:"publishTime"`
				Description string `json:"description"`
				Tags        string `json:"tags"`
				Company     string `json:"company"`
				BriefDesc   string `json:"briefDesc"`
				Artist      struct {
					Name      string        `json:"name"`
					ID        int           `json:"id"`
					PicID     int           `json:"picId"`
					Img1V1ID  int           `json:"img1v1Id"`
					BriefDesc string        `json:"briefDesc"`
					PicURL    string        `json:"picUrl"`
					Img1V1URL string        `json:"img1v1Url"`
					AlbumSize int           `json:"albumSize"`
					Alias     []interface{} `json:"alias"`
					Trans     string        `json:"trans"`
					MusicSize int           `json:"musicSize"`
				} `json:"artist"`
				Songs           []interface{} `json:"songs"`
				Alias           []interface{} `json:"alias"`
				Status          int           `json:"status"`
				CopyrightID     int           `json:"copyrightId"`
				CommentThreadID string        `json:"commentThreadId"`
				Artists         []struct {
					Name      string        `json:"name"`
					ID        int           `json:"id"`
					PicID     int           `json:"picId"`
					Img1V1ID  int           `json:"img1v1Id"`
					BriefDesc string        `json:"briefDesc"`
					PicURL    string        `json:"picUrl"`
					Img1V1URL string        `json:"img1v1Url"`
					AlbumSize int           `json:"albumSize"`
					Alias     []interface{} `json:"alias"`
					Trans     string        `json:"trans"`
					MusicSize int           `json:"musicSize"`
				} `json:"artists"`
			} `json:"album"`
			Starred         bool          `json:"starred"`
			Popularity      float64       `json:"popularity"`
			Score           int           `json:"score"`
			StarredNum      int           `json:"starredNum"`
			Duration        int           `json:"duration"`
			PlayedNum       int           `json:"playedNum"`
			DayPlays        int           `json:"dayPlays"`
			HearTime        int           `json:"hearTime"`
			Ringtone        interface{}   `json:"ringtone"`
			Crbt            interface{}   `json:"crbt"`
			Audition        interface{}   `json:"audition"`
			CopyFrom        string        `json:"copyFrom"`
			CommentThreadID string        `json:"commentThreadId"`
			RtURL           interface{}   `json:"rtUrl"`
			Ftype           int           `json:"ftype"`
			RtUrls          []interface{} `json:"rtUrls"`
			Copyright       int           `json:"copyright"`
			HMusic          struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int64       `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"hMusic"`
			MMusic struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int64       `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"mMusic"`
			LMusic struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int64       `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"lMusic"`
			BMusic struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int64       `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"bMusic"`
			Mp3URL string      `json:"mp3Url"`
			Mvid   int         `json:"mvid"`
			Rurl   interface{} `json:"rurl"`
			Rtype  int         `json:"rtype"`
		} `json:"songs"`
		SongCount int `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}

// Song Song
type Song struct {
	Author string `json:"author"`
	Lrc    string `json:"lrc"`
	Pic    string `json:"pic"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}

var TopListAll = [][]string{
	{"0", "云音乐新歌榜", "/api/playlist/detail?id=3779629"},
	{"1", "云音乐热歌榜", "/api/playlist/detail?id=3778678"},
	{"2", "网易原创歌曲榜", "/api/playlist/detail?id=2884035"},
	{"3", "云音乐飙升榜", "/api/playlist/detail?id=19723756"},
	{"4", "云音乐电音榜", "/api/playlist/detail?id=10520166"},
	{"5", "UK排行榜周榜", "/api/playlist/detail?id=180106"},
	{"6", "美国Billboard周榜", "/api/playlist/detail?id=60198"},
	{"7", "KTV嗨榜", "/api/playlist/detail?id=21845217"},
	{"8", "iTunes榜", "/api/playlist/detail?id=11641012"},
	{"9", "Hit FM Top榜", "/api/playlist/detail?id=120001"},
	{"10", "日本Oricon周榜", "/api/playlist/detail?id=60131"},
	{"11", "韩国Melon排行榜周榜", "/api/playlist/detail?id=3733003"},
	{"12", "韩国Mnet排行榜周榜", "/api/playlist/detail?id=60255"},
	{"13", "韩国Melon原声周榜", "/api/playlist/detail?id=46772709"},
	{"14", "中国TOP排行榜(港台榜)", "/api/playlist/detail?id=112504"},
	{"15", "中国TOP排行榜(内地榜)", "/api/playlist/detail?id=64016"},
	{"16", "香港电台中文歌曲龙虎榜", "/api/playlist/detail?id=10169002"},
	{"17", "华语金曲榜", "/api/playlist/detail?id=4395559"},
	{"18", "中国嘻哈榜", "/api/playlist/detail?id=1899724"},
	{"19", "法国 NRJ EuroHot 30周榜", "/api/playlist/detail?id=27135204"},
	{"20", "台湾Hito排行榜", "/api/playlist/detail?id=112463"},
	{"21", "Beatport全球电子舞曲榜", "/api/playlist/detail?id=3812895"},
}
