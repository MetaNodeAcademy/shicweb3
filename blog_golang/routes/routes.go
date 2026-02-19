package routes

import (
	// "golang_learning_blog/controllers"
	// "golang_learning_blog/utils"

	"time"

	"github.com/gin-gonic/gin"
)

// health
// 临时的健康检查处理函数，实际项目中应该替换为真正的处理函数
func Health(c *gin.Context) {
	c.JSON(200,
		gin.H{
			"status":  "OK",
			"message": time.Now().Format(time.RFC3339),
		})
}

func SetupRoutes() *gin.Engine {
	r := gin.New()

	// 创建控制器实例
	//authController := &controllers.AuthController{}
	//postController := &controllers.PostController{}
	//commentController := &controllers.CommentController{}

	// API路由组
	api := r.Group("/api/v1")
	{
		// 1、不需要认证路由：注册、登录认证
		auth := api.Group("/auth")
		{
			// /api/v1/auth/register
			auth.POST("/register", Health)
			// /api/v1/auth/login
			auth.POST("/login", Health)
		}

		// 2、认证路由：用户信息、文章、评论
		authed := api.Group("")
		// +认证

		{
			// 用户信息
			// /api/v1/profile
			authed.GET("/profile", Health)

			// 文章
			posts := authed.Group("/posts")
			{
				// /api/v1/posts
				posts.POST("", Health)
				// /api/v1/posts/:id
				posts.PUT("/:id", Health)
				// /api/v1/posts/:id
				posts.DELETE("/:id", Health)

			}

			// 评论
			// /api/v1/posts/:post_id/comments
			comments := authed.Group("/posts/:post_id/comments")
			{
				comments.POST("", Health)
			}
		}

		// 3、不需要认证路由：文章
		public := api.Group("")
		{
			// /api/v1/posts
			public.GET("/posts", Health)
			// /api/v1/posts/:id
			public.GET("/posts/:id", Health)
		}

		// 4、不需要认证路由：公开评论
		// /api/v1/comments
		comments := api.Group("/comments")
		{
			// /api/v1/comments/post/:post_id
			comments.GET("/post/:post_id", Health)
		}
	}

	// 健康检查
	r.GET("/health", Health)

	return r
}
