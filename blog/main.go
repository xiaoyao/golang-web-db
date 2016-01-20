package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func initData(coll *mgo.Collection) {
	blogs := []interface{}{
		Blog{
			Id:        bson.NewObjectId(),
			Title:     "我的第一篇博客",
			Content:   "这是我的第一篇博客",
			CreatedAt: time.Now(),
		},
		Blog{
			Id:        bson.NewObjectId(),
			Title:     "今天天气不错",
			Content:   "今天天气虽然冷，但是很晴朗",
			CreatedAt: time.Now(),
		},
		Blog{
			Id:        bson.NewObjectId(),
			Title:     "周末参加创业松鼠聚会",
			Content:   "周末要去参加创业松鼠的聚会，可以认识很多朋友",
			CreatedAt: time.Now(),
		},
	}

	err := coll.Insert(blogs...)
	if err != nil {
		log.Printf("insert blogs with error: %s\n", err)
	}

	err = coll.EnsureIndexKey("title")
	if err != nil {
		log.Printf("ensure index with error: %s\n", err)
	}

	fmt.Println("初始化博客数据成功！")
}

func updateBlog(coll *mgo.Collection, title, content string) {
	err := coll.Update(bson.M{
		"title": title,
	}, bson.M{
		"content": content,
	})
	if err != nil {
		log.Printf("update blog(title: %s) with error: %s\n", title, err)
		return
	}

	fmt.Println("更新博客数据成功！")
}

func queryBlogs(coll *mgo.Collection) {
	blogs := make([]Blog, 0)
	err := coll.Find(bson.M{}).All(&blogs)
	if err != nil {
		log.Printf("query blogs with error: %s\n", err)
		return
	}
	fmt.Printf("blogs list: %#v\n", blogs)
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Printf("mongo dial with error: %s\n", err)
		return
	}

	defer session.Close()

	db := session.DB("qiniu_blog")

	coll := db.C("blog")

	fmt.Println("===== 欢迎来到七牛博客 =====")

	// 初始化博客数据
	fmt.Println("----- 初始化博客数据 -----")
	fmt.Scanln()
	initData(coll)
	fmt.Println("----- 初始化博客数据，完成！ -----")

	fmt.Println("================================================")

	// 获取博客数据
	fmt.Println("----- 获取博客数据 -----")
	fmt.Scanln()
	queryBlogs(coll)
	fmt.Println("----- 获取博客数据，完成！ -----")

	fmt.Println("================================================")

	// 更新博客数据
	fmt.Println("----- 更新博客数据 -----")
	fmt.Scanln()
	updateBlog(coll, "今天天气不错", "这是新的内容")
	fmt.Println("----- 更新博客数据，完成！ -----")

	fmt.Println("================================================")

	// 获取博客数据
	fmt.Println("----- 获取博客数据 -----")
	fmt.Scanln()
	queryBlogs(coll)
	fmt.Println("----- 获取博客数据，完成！ -----")

}
