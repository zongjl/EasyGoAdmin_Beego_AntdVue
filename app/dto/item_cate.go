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

package dto

import "github.com/gookit/validate"

// 栏目查询条件
type ItemCateQueryReq struct {
	Name string `form:"name"` // 栏目名称
}

// 添加栏目
type ItemCateAddReq struct {
	Name    string `form:"name" validate:"required"`   // 栏目名称
	Pid     int    `form:"pid" validate:"int"`         // 父级ID
	ItemId  int    `form:"itemId" validate:"int"`      // 栏目ID
	Pinyin  string `form:"pinyin" validate:"required"` // 拼音(全)
	Code    string `form:"code" validate:"required"`   // 拼音(简)
	IsCover int    `form:"isCover" validate:"int"`     // 是否有封面：1是 2否
	Cover   string `form:"cover"`                      // 封面
	Status  int    `form:"status" validate:"int"`      // 状态：1启用 2停用
	Note    string `form:"note"`                       // 备注
	Sort    int    `form:"sort" validate:"int"`        // 排序
}

// 添加栏目表单验证
func (v ItemCateAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required":   "栏目名称不能为空.",
		"Pid.int":         "请选择上级栏目.",
		"ItemId.int":      "请选择栏目ID.",
		"Pinyin.required": "拼音全拼不能为空.",
		"Code.required":   "拼音简拼不能为空.",
		"IsCover.int":     "请选择是否有封面.",
		"Status.int":      "请选择栏目状态.",
		"Sort.int":        "排序不能为空.",
	}
}

// 修改栏目
type ItemCateUpdateReq struct {
	Id      int    `form:"id" validate:"int"`
	Name    string `form:"name" validate:"required"`   // 栏目名称
	Pid     int    `form:"pid" validate:"int"`         // 父级ID
	ItemId  int    `form:"itemId" validate:"int"`      // 栏目ID
	Pinyin  string `form:"pinyin" validate:"required"` // 拼音(全)
	Code    string `form:"code" validate:"required"`   // 拼音(简)
	IsCover int    `form:"isCover" validate:"int"`     // 是否有封面：1是 2否
	Cover   string `form:"cover"`                      // 封面
	Status  int    `form:"status" validate:"int"`      // 状态：1启用 2停用
	Note    string `form:"note"`                       // 备注
	Sort    int    `form:"sort" validate:"int"`        // 排序
}

// 更新栏目表单验证
func (v ItemCateUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":          "栏目ID不能为空.",
		"Name.required":   "栏目名称不能为空.",
		"Pid.int":         "请选择上级栏目.",
		"ItemId.int":      "请选择栏目ID.",
		"Pinyin.required": "拼音全拼不能为空.",
		"Code.required":   "拼音简拼不能为空.",
		"IsCover.int":     "请选择是否有封面.",
		"Status.int":      "请选择栏目状态.",
		"Sort.int":        "排序不能为空.",
	}
}
