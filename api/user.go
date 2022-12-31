package api

import (
	"exam/dao"
	"exam/model"
	"exam/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func User() {
	Router := gin.Default()
	Router.GET("/login", func(c *gin.Context) {
		name := c.Query("name")
		password := c.Query("password")
		if service.IsNameExist(name) {
			rows, err := dao.NameQuery(name)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"msg": "不存在账号",
				})
			}
			var user model.User
			rows.Scan(&user.Id, &user.Name, &user.Password, &user.Balance)
			if user.Password == password {
				c.SetCookie("123", "321", 3600, "/", "localhost", false, true)
				c.JSON(http.StatusOK, gin.H{
					"msg": "登录成功",
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "已存在该用户不存在",
			})
		}
	})
	Router.GET("/signUp", func(c *gin.Context) {
		name := c.Query("name")
		password := c.Query("password")

		if service.IsNameExist(name) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "已存在该用户已经存在",
			})
		}
		err := dao.ExecName(name)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "创建名字失败",
			})
		}
		err = dao.ExecPassword(password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "创建密码失败",
			})
		}
	})

	User := Router.Group("/user", func(c *gin.Context) {
		_, err := c.Cookie("123")
		if err != nil {
			c.Abort()
		}
		c.Next()
	})
	{
		User.POST("/recharge", func(c *gin.Context) {
			recharge := c.Query("recharge")
			name := c.Query("name")

			money, err := strconv.ParseInt(recharge, 10, 64)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"msg": "失败",
				})
			}
			dao.Recharge(int(money), name)
		})
		User.POST("/transfer", func(c *gin.Context) {
			transfer := c.Query("transfer")
			sender := c.Query("sender")
			reciever := c.Query("reciever")
			money, err := strconv.ParseInt(transfer, 10, 64)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"msg": "错误",
				})
			}
			dao.Recharge(int(money), reciever)
			dao.ReduceMoney(int(money), sender)
			err = dao.Checkmoney(sender)
			if err != nil {
				log.Println(err)
				dao.Recharge(int(money), sender)
				dao.ReduceMoney(int(money), reciever)
			}
		})
	}
	Router.Run()
}
