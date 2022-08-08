package api

import (
	"mime/multipart"
	"strings"
	"time"

	"{{ web_framework_import }}"
	"github.com/spf13/cast"

	"{{ main_module }}/internal/form"
)

func init() {
	AddRouteRegistrar(UploadFiles)
}

// UploadFiles
// @Summary      上传文件
// @Description  可上传一个或多个文件，大小不超过 12MB
// @Tags         程序设置
// @Accept       multipart/form-data
// @Security     ApiKeyAuth
// @Param        upload  formData  file  true  "大小不超过 12MB 的文件"
// @Produce      json
// @Success      200  {object}  query.Response  "返回以 ; 分割的字符串，前端获取时的相对路径为 /uploads，比如返回的是 x.jpg，则实际路径为 /uploads/x.jpg"
// @Router       /upload/files [post]
{%- if web_framework == ".iris" %}
func UploadFiles(router iris.Party) {
	router.HandleDir("/uploads", iris.Dir(conf.UploadPath()))

	router.Post("/upload/files", conf.JWTMiddleware(), func(c iris.Context) {
		c.SetMaxRequestBodySize(12 * iris.MB)

		var filePaths []string

		// https://github.com/kataras/iris/blob/master/_examples/file-server/upload-files/main.go
		_, _, err := c.UploadFormFiles(conf.UploadPath(),
			func(ctx iris.Context, header *multipart.FileHeader) bool {
				header.Filename = cast.ToString(time.Now().UnixNano()) + "-" + header.Filename
				filePaths = append(filePaths, header.Filename)
				return true
			})
		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendJSON(c, strings.Join(filePaths, ";"))
		return
	})
}
{%- endif %}
