// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 演示一管理-服务类
 * @author 半城风雨
 * @since 2022-05-14
 * @File : example
 */
package services

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

// 中间件管理服务
var Example = new(exampleService)

type exampleService struct{}

func (s *exampleService) GetList(req dto.ExamplePageReq) ([]vo.ExampleListVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Example)).Filter("mark", 1)

	// 测试名称
	
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	

	// 状态：1正常 2停用
	
	if req.Status > 0 {
		query = query.Filter("status", req.Status)
	}
	

	// 类型：1京东 2淘宝 3拼多多 4唯品会
	
	if req.Type > 0 {
		query = query.Filter("type", req.Type)
	}
	

	// 是否VIP：1是 2否
	
	if req.IsVip > 0 {
		query = query.Filter("is_vip", req.IsVip)
	}
	

	// 排序
	query = query.OrderBy("id")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.Example, 0)
	// 对象转换
	query.All(&lists)

	// 数据处理
	var result []vo.ExampleListVo
	for _, v := range lists {
		item := vo.ExampleListVo{}
		item.Example = v
		
		
		
		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}
		
		
		
		
		
		result = append(result, item)
	}

	// 返回结果
	return result, count, nil
}

func (s *exampleService) Add(req dto.ExampleAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity models.Example
	
	entity.Name = req.Name
	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "example")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.Content = req.Content
	
	entity.Status = req.Status
	
	
	entity.Type = req.Type
	
	
	entity.IsVip = req.IsVip
	
	
	entity.Sort = req.Sort
	
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *exampleService) Update(req dto.ExampleUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &models.Example{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	
	entity.Name = req.Name
	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "example")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.Content = req.Content
	
	entity.Status = req.Status
	
	
	entity.Type = req.Type
	
	
	entity.IsVip = req.IsVip
	
	
	entity.Sort = req.Sort
	
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

// 删除
func (s *exampleService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Example{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Example{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}









func (s *exampleService) Status(req dto.ExampleStatusReq, userId int) (int64, error) {
	// 查询记录是否存在
	entity := &models.Example{Id: req.Id}
	err := entity.Get()
	if err != nil {
	return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}





func (s *exampleService) IsVip(req dto.ExampleIsVipReq, userId int) (int64, error) {
	// 查询记录是否存在
	entity := &models.Example{Id: req.Id}
	err := entity.Get()
	if err != nil {
	return 0, errors.New("记录不存在")
	}
	entity.IsVip = req.IsVip
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}




