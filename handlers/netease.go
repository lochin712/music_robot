package handlers

import (
	"encoding/json"
	"fmt"
	"music_robot/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var (
		err           error
		limitStr      = "10"
		offsetStr     = "0"
		searchTypeStr = "1"
		req           models.BearyChatRequest
		resp          models.NetEaseResponse
		atts          []models.Attachments
		result        models.BearyChatResponse
	)

	defer func() {
		if err != nil {
			c.JSON(400, gin.H{"errcode": 400, "errmsg": err.Error()})
		}
	}()

	err = c.BindJSON(&req)
	if err != nil {
		return
	}

	keywords := strings.TrimSpace(strings.TrimPrefix(req.Text, req.TriggerWord))

	str, err := models.HttpPost("http://music.163.com/api/search/pc", "offset="+offsetStr+"&limit="+limitStr+"&type="+searchTypeStr+"&s="+keywords)
	if err != nil {
		return
	}
	err = json.Unmarshal(str, &resp)
	if err != nil {
		return
	}
	for i, v := range resp.Result.Songs {
		var artists []string
		for _, artist := range v.Artists {
			artists = append(artists, artist.Name)
		}
		musicURL := "http://music.163.com/song/media/outer/url?id=" + strconv.Itoa(v.ID) + ".mp3"
		artistsStr := strings.Join(artists, ",")
		att := models.Attachments{
			// Title: "",
			Text: fmt.Sprintf("%d | [%s](%s) | %s | [%s](%s)",
				i+1, v.Name, musicURL, artistsStr, v.Album.Name, v.Album.PicURL),
			Color: "#666666",
			// Images: []models.ImageUrl{
			// 	{URL: v.Album.PicURL},
			// },
		}

		atts = append(atts, att)
	}

	result.Text = fmt.Sprintf("已为您找到%d条,《%s》相关的歌曲", len(atts), keywords)
	result.Attachments = atts

	res, err := json.Marshal(result)
	if err != nil {
		return
	}
	c.Writer.Write(res)

}
