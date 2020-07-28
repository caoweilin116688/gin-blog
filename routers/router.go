package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/export"
	"gin-blog/pkg/qrcode"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"gin-blog/pkg/setting"
	"gin-blog/pkg/upload"

	_ "gin-blog/docs" //导入包，自动执行包内的init函数

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	//当访问host/upload/images/时，将会读取到GOPATH/src/gin-blog/runtime/upload/images 下的文件
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//下载xlsx文件 访问 host/export,将会访问到 GOPATH/src/gin-blog/runtime/export
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//访问二维码
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	r.GET("/auth", api.GetAuth)
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//图片上传
	r.POST("/upload", api.UploadImage)
	//Go 的语法糖,代表了它的作用域
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//导出标签,生成xlsx文件
		r.POST("/tags/export", v1.ExportTag)

		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//二维码
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
