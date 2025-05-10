# Video service - projecte final

Dockerfiles don't work

## Routes

```go
baseGroup := router.Group("api")

baseGroup.POST("login", handlers.Login)
baseGroup.POST("register", middlewares.Ultrapass(), handlers.Register)
{
	uploadGroup := router.Group("upload", middlewares.JWT())
	uploadGroup.POST("", handlers.UploadVideo)
}

baseGroup.GET("videos", handlers.GetRecentVideos)
{
	videoGroup := baseGroup.Group("video")
	videoGroup.GET("source/:id", handlers.GetVideo)
	videoGroup.GET("info/:id", handlers.GetVideoInfo)
	videoGroup.GET("thumb/:id", handlers.GetVideoThumbnail)
}
```

### /login && /register
- username
- password

### /upload
- 

