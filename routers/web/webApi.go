package web

import (
	"applyUpLoadFile/config"
	"applyUpLoadFile/middleware/log"
	"applyUpLoadFile/model"
	"applyUpLoadFile/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/onethefour/common/xutils"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// LoadWebRouter
// 加载Web路由
func LoadWebRouter(group *gin.RouterGroup) {
	group.POST("/apply/addApply", AddApply)
	group.GET("/apply/coins", GetCoins)
	group.POST("/upload", UpLoadApply)
	group.POST("/uploads", MultUploadFile)
}

func NewError(ctx *gin.Context, err string) {
	log.Info(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    500,
		"message": err,
		"data":    "",
	})
}

// AddApply
// @ 增加申请人信息
// @Description 增加申请人信息
// @Accept application/json
// @Produce application/json
// @Param   name        body string     true        "用户名"
// @Param   phone       body  string     true        "手机"
// @Param   email       body  string     true        "邮箱"
// @Param   coin_name   body   string     true        "币种名称"
// @Param   introduce   body   string     true        "介绍"
// @Param   id_card_picture body string     true  "身份证文件路径"
// @Param   business_picture body string     true "营业执照文件路径"
// @Success 200 object  Result    "ok"
// @Router /web/apply/addApply [post]
func AddApply(ctx *gin.Context) {
	err := xutils.LockMax(ctx.ClientIP(), 2)
	if err != nil {
		pass := AddRequest(ctx.ClientIP())
		if !pass {
			NewError(ctx, ctx.ClientIP()+",申请过于频繁,30分钟后重试")
			return
		}
		NewError(ctx, "申请过于频繁")
		return
	}
	defer xutils.UnlockDelay(ctx.ClientIP(), time.Second*10)
	var params = new(model.ApplyCompany)
	if err := ctx.ShouldBindJSON(params); err != nil {
		NewError(ctx, err.Error())
		return
	}
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, params.Phone)
	if !ok {
		NewError(ctx, "手机格式不对")
		return
	}

	if params.IdCardPicture == "" || params.BusinessPicture == "" {
		NewError(ctx, "身份证和执照不能为空")
		return
	}
	cName := ""
	coins := ""
	s := strings.Split(params.CoinName, ",")
	for i := 0; i < len(s); i++ {
		cId, err := strconv.Atoi(s[i])
		if err != nil {
			NewError(ctx, err.Error())
			return
		}
		getCInfo, e := model.GetCoinById(int64(cId))
		if e != nil {
			NewError(ctx, e.Error())
			return
		}
		cName += getCInfo.Name + "(" + getCInfo.FullName + ")"
		coins += getCInfo.Name
		if i+1 < len(s) {
			cName += ","
			coins += ","
		}
	}

	apply := &model.ApplyPending{
		Name:            params.Name,
		Phone:           params.Phone,
		Email:           params.Email,
		CoinName:        coins,
		Introduce:       params.Introduce,
		IdCardPicture:   params.IdCardPicture,
		BusinessPicture: params.BusinessPicture,
	}

	existApply, err := model.ExistApply(apply)
	if err != nil {
		NewError(ctx, err.Error())
		return
	}
	if existApply {
		ctx.JSON(http.StatusOK, &Result{
			Code:    200,
			Message: fmt.Sprintf("该信息，%s 已申请过", params.Name),
			Data:    "success",
		})
		return
	}
	pending, err := model.InsertApplyPending(apply)
	if err != nil {
		NewError(ctx, err.Error())
		return
	}
	if pending > 0 {
		bodyText := fmt.Sprintf("商户试用申请: \n商户: %s \n手机号: %s \n邮箱: %s \n币种名称: %s \n公司简介: %s",
			params.Name, params.Phone, params.Email, cName, params.Introduce)
		var wg sync.WaitGroup
		for i := 0; i < len(config.Cfg.Email.Recipient); i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				em := &utils.EmailConfig{
					IamUserName:  config.Cfg.Email.IamUserName,
					Recipient:    config.Cfg.Email.Recipient[i],
					SmtpUsername: config.Cfg.Email.SmtpUsername,
					SmtpPassword: config.Cfg.Email.SmtpPassword,
					Host:         config.Cfg.Email.Host,
				}
				em.SendEmail(bodyText)
			}(i)
		}
		wg.Wait()
		ctx.JSON(http.StatusOK, &Result{
			Code:    200,
			Message: fmt.Sprintf("%s 新增完成", params.Name),
			Data:    "success",
		})
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Code:    500,
		Message: "该商户已有同样申请",
		Data:    "failure",
	})
	return
}

// UpLoadApply
// @单个文件上传
// @Description 单个文件上传
// @Accept 	multipart/form-data
// @Produce multipart/form-data
// @Param   file     	 formData file    true        "文件"
// @Param   groupName    formData string     true     "文件储存的文件夹名"
// @Success 200 object Result    "ok"
// @Router /web/upload [post]
func UpLoadApply(c *gin.Context) {
	err := xutils.LockMax(c.ClientIP(), 2)
	if err != nil {
		NewError(c, err.Error())
		return
	}
	defer xutils.UnlockDelay(c.ClientIP(), time.Second*10)
	group, e := c.GetPostForm("groupName")
	groupName := ""
	if !e {
		groupName = fmt.Sprintf("%s-%d", strings.Split(uuid.New(), "-")[0], time.Now().Unix())
	} else {
		groupName = fmt.Sprintf("%s-%d", group, time.Now().Unix())
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	basePath := utils.CreateDateDir(config.GetConfig().UpLoad.Url + "/" + groupName)
	str := strings.Split(file.Filename, ".")
	pathName := fmt.Sprintf("%d.%s", time.Now().UnixNano(), str[len(str)-1])
	//filename := basePath + "/" + filepath.Base(file.Filename)
	filename := basePath + "/" + pathName
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	b := utils.VailDataFileMd5(basePath, file.Filename)
	if !b {
		_ = os.RemoveAll(basePath)
		NewError(c, fmt.Sprintf("同样文件上传频繁请30秒后重试"))
		return
	}
	data := make(map[string]interface{})
	data["path"] = filename
	//res := &
	c.JSON(http.StatusOK, Result{
		Code:    200,
		Message: fmt.Sprintf("文件 %s 上传成功 ", file.Filename),
		Data:    data,
	})
	return
}

// MultUploadFile
// @title 多个文件上传
// @Description 多个文件上传
// @Accept multipart/form-data
// @Produce multipart/form-data
// @Param   files    	 formData     file      true     "多个文件"
// @Param   groupName    formData     string    true     "文件储存的文件夹名"
// @Success 200 object  Result    "ok"
// @Router /web/upload [post]
func MultUploadFile(c *gin.Context) {
	err := xutils.LockMax(c.ClientIP(), 2)
	if err != nil {
		NewError(c, err.Error())
		return
	}
	defer xutils.UnlockDelay(c.ClientIP(), time.Second*10)
	group, e := c.GetPostForm("groupName")
	groupName := ""
	if !e {
		groupName = fmt.Sprintf("%s-%d", RandString(10), time.Now().Unix())
	} else {
		groupName = fmt.Sprintf("%s-%s", group, RandString(10))
	}
	//获取到所有的文件
	form, _ := c.MultipartForm()
	//获取到所有的文件数组
	//files := form.File["upload[]"]
	files := form.File["files"]
	path := utils.CreateDateDir(config.GetConfig().UpLoad.Url + "/" + groupName)
	ps := ""
	//遍历数组进行处理
	for i, file := range files {
		//进行文件保存
		str := strings.Split(file.Filename, ".")
		pathName := fmt.Sprintf("%d.%s", time.Now().Unix(), str[len(str)-1])
		err = c.SaveUploadedFile(file, path+"/"+pathName)
		if err != nil {
			NewError(c, err.Error())
			return
		}
		b := utils.VailDataFileMd5(path, pathName)
		if !b {
			_ = os.RemoveAll(path)
			NewError(c, fmt.Sprintf("同样文件上传频繁请30秒后重试"))
			return
		}
		if i == 0 {
			ps += "/" + path + "/" + pathName
		} else {
			ps += "," + "/" + path + "/" + pathName
		}

	}
	data := make(map[string]interface{})
	data["path"] = ps
	if len(files) > 0 {
		c.JSON(http.StatusOK, Result{
			Code:    200,
			Message: fmt.Sprintf("%d files uploaded!", len(files)),
			Data:    data,
		})
		return
	}
	c.JSON(http.StatusOK, Result{
		Code:    500,
		Message: fmt.Sprintf("%d files uploaded!", len(files)),
		Data:    "",
	})
	return
}

// GetCoins
// @title 获取币种
// @Description 获取币种
// @Accept application/json
// @Produce application/json
// @Success 200 object  Result  "ok"
// @Router /web/apply/coins [get]
func GetCoins(c *gin.Context) {

	var ls []map[string]interface{}

	all, err := model.FindCoinInfoAll()
	if err != nil {
		return
	}

	for i := 0; i < len(all); i++ {
		mp := map[string]interface{}{
			"id":   all[i].Id,
			"name": all[i].Name,
		}
		ls = append(ls, mp)
	}
	data := make(map[string]interface{})
	data["list"] = ls
	c.JSON(http.StatusOK, Result{
		Code:    200,
		Message: "",
		Data:    data,
	})
	return
}

func AddRequest(ip string) bool {
	utils.CacheConf.CacheUtil.DeleteExpired()
	reIp, ok := utils.CacheConf.CacheUtil.Get(ip)
	if ok {
		add := reIp.(int)
		if add > 3 {
			return false
		}
		add += 1
		utils.CacheConf.CacheUtil.Set(ip, add, cache.DefaultExpiration)
	} else {
		utils.CacheConf.CacheUtil.Set(ip, 1, cache.DefaultExpiration)
	}
	return true
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
