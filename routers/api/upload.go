//c.Request.FormFile：获取上传的图片（返回提供的表单键的第一个文件）
//CheckImageExt、CheckImageSize检查图片大小，检查图片后缀
//CheckImage：检查上传图片所需（权限、文件夹）
//SaveUploadedFile：保存图片

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/upload"
)

//curl -X POST http://127.0.0.1:8000/upload -H 'Content-Type:multipart/form-data' -F image=@/Users/caoweilin/Downloads/01.png
//{"code":200,"data":{"image_save_url":"upload/images/96a3be3cf272e017046d1b2674a52bd3.png","image_url":"http://127.0.0.1:8000/upload/images/96a3be3cf272e017046d1b2674a52bd3.png"},"msg":"ok"}
func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		//
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
