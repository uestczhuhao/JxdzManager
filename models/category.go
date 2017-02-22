package models

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id       int
	Name     string `orm:"size(45)"`
	ParentID int
	ChildID  string `orm:"size(100)"`
	Depth    int
	Type     string
	Content  string `orm:"size(1000)"`
}

type CategorySorted struct {
	Id   int
	Name string `orm:"size(45)"`
}

type CategoryListOne struct {
	Id            int
	DepthNumber   int
	CateOne       *Category
	CateTwo       []*Category
	CateOnesChild []*Category
}

type CategoryLeftMenuTwo struct {
	Id        int
	CateTwo   *Category
	CateThree []*Category
}

//作为测试用的初始化函数
func InitAndClear() {
	AddCategory("学院概况", 0, "目录", "")
	AddCategory("学院简介", 1, "单页面", "学院的前身是建于1964年的电子机械系，2001年成立学院。 学院设有机械工程系、工业工程系、电力电子工程系，工程训练中心、“机电与控制工程”国家级实验教学示范中心，先进装备及其控制技术研究所、设备监控与健康管理研究所、机电测控一体化技术研究所、精密机电智能测试与控制研究所、智能光机电系统研究所、微系统与机器人研究所、可靠性工程研究所、数字化设计与仿真研究所、电力电子与智能感知研究所，建有四川智能服务机器人2011协同创新中心、四川省智能机电系统与工业信息化重点实验室、四川省安全生产智能服务机器人工程实验室、信息与产业部智能机械及其系统工程重点实验室等一批教学和科研机构。")
	AddCategory("学院机构", 1, "目录", "")
	AddCategory("行政机构", 3, "单页面", "科室职责：负责对外接待和宣传、财务的日常管理和结算；负责网络与信息化建设、监印管理、消防安全；负责党校、组织、统战、工会等党务全团日常管理工作；负责国资、机房、文书收发等工作；协调学院各科室工作。")
	AddCategory("教学机构", 3, "单页面", "机械工程系长期致力于机电一体化领域的人才培养、学科建设和科研工作，经过40余年的积淀，在教学和科研等方面项果累累。现有教职工30余人，其中博士生导师4人，教授4人，副教授11名，其中具有博士、硕士学位的教师比例达到80%。全系共有3个专业实验室，拥有仪器设备近200台（套），占地约800平方米。 ")

	AddCategory("师资队伍", 0, "目录", "")
	AddCategory("师资介绍", 6, "教师列表", "")
	AddCategory("杰出人才", 6, "单页面", "“千人计划”入选者左明健、许亮峰　“长江学者奖励计划”讲座教授左明健、Horacio Dante Espinosa、邢留冬")
	AddCategory("人事政策", 6, "文章列表", "")
	AddCategory("办事指南", 6, "文章列表", "")

	AddCategory("科学研究", 0, "目录", "")
	AddCategory("学科建设", 11, "单页面", "机械电子工程学院目前有机械工程四川省一级重点学科，下设“机械电子工程”博士学位授权点和“机械制造及自动化”、“机械电子工程”、“机械设计及理论”、“精密仪器及机械”、“电力电子与电力传动”5个硕士学位授权点，有2个省部级重点实验室和2个校级研究所。")
	AddCategory("科研成果", 11, "单页面", "  随着机械电子工程学院的不断壮大，我院的科研能力也得到了极大的提高！ 自2002年以来，我院先后承担了国家“863”项目、国家自然科学基金项目、国家自然科学基金国际（地区）合作交流资助项目、军事预研项目、教育部高等学校博士学科点专项科研基金项目、四川省应用基础研究项目、四川省科技支撑计划项目等各类科研项目100余项，经费3000余万元；在包括IEEE Transactions on Reliability、Concurrent Engineering: Research and Applications、《机械工程学报》、《中国机械工程》等国内外重要刊物、学术会议上共发表论文700余篇，其中被SCI、EI收录200余篇；出版专著和教材20余部；获得教育部自然科学二等奖1项、军内科技进步二等奖1项、省部级科技进步三等奖5项；申请国家专利29项，目前获得发明专利授权9项，实用新型专利授权4项。")
	AddCategory("资料下载", 11, "下载列表", "")
	AddCategory("办事指南", 11, "文章列表", "")
	AddCategory("合作与交流", 11, "文章列表", "")

	AddCategory("快速链接", 0, "链接列表", "")
	AddCategory("机械电子工程学院微信官方网站", 17, "链接列表", "http://www.weiweixiao.net/index.php?g=Wap&m=Index&a=index&token=4f7731d259beac5")
	AddCategory("电子科技大学", 17, "链接列表", "http://www.uestc.edu.cn/")
	AddCategory("电子科技大学信息门户登录", 17, "链接列表", "http://idas.uestc.edu.cn/authserver/login?service=http%3A%2F%2Fportal.uestc.edu.cn%2F")
	AddCategory("在线图书馆", 17, "链接列表", "http://www.lib.uestc.edu.cn/")
	AddCategory("后台入口", 17, "链接列表", "http://www.jxdz.uestc.edu.cn/index.php/admin")
	return
}

func AddCategory(name string, parentid int, typ string, cont string) error {
	var err error
	// var desid string
	var resultid int
	var curid int64
	o := orm.NewOrm()
	if parentid == 0 {
		cur := &Category{Name: name, ParentID: 0, ChildID: "0", Depth: 1, Type: typ, Content: cont}
		_, err = o.Insert(cur)
		return err
	}

	resultid = parentid

	parentcate := Category{Id: resultid}

	if o.Read(&parentcate, "Id") == nil {
		depth := parentcate.Depth + 1

		cur := &Category{Name: name, ParentID: parentid, ChildID: "0", Depth: depth, Type: typ, Content: cont}

		curid, err = o.Insert(cur)
		if err != nil {
			return err
		}
	} else {
		beego.Debug("输入父节点错误！")
	}

	if o.Read(&parentcate, "Id") == nil {
		var newchildid string
		var curidstr string
		curidstr = strconv.FormatInt(curid, 10)

		if parentcate.ChildID != "0" {
			newchildid = parentcate.ChildID + "," + curidstr
			parentcate.ChildID = newchildid
			o.Update(&parentcate, "ChildID")
		} else {
			parentcate.ChildID = curidstr
			o.Update(&parentcate, "ChildID")
		}
	}
	return err
}

func DelCategory(id int, name string, parentid int) int {
	o := orm.NewOrm()
	q := o.QueryTable("Category")
	var category Category
	// var idstr string

	err := q.Filter("Id", id).Filter("Name", name).Filter("ParentID", parentid).One(&category)

	if err != nil {
		beego.Debug(err)
		return 1

	} else {
		//err==nil  即搜索到了目标栏目category
		if category.ChildID == "0" {
			//当要删除的这个栏目没有子栏目时才可以删除,否则拒绝删除命令
			var resultidDel int
			var newchildidOfDel string
			var DelIdStr string
			resultidDel = category.ParentID
			parentcateOfDel := Category{Id: resultidDel}
			if o.Read(&parentcateOfDel, "Id") == nil {
				if strings.Contains(parentcateOfDel.ChildID, ",") == true {
					newchildidOfDel = parentcateOfDel.ChildID
					DelIdStr = strconv.Itoa(category.Id)
					newchildidOfDel = strings.Replace(newchildidOfDel, DelIdStr, "", -1)
					newchildidOfDel = strings.Replace(newchildidOfDel, ",,", ",", -1)
					newchildidOfDel = strings.Trim(newchildidOfDel, ",")
					parentcateOfDel.ChildID = newchildidOfDel
					o.Update(&parentcateOfDel, "ChildID")
				} else {
					parentcateOfDel.ChildID = "0"
					o.Update(&parentcateOfDel, "ChildID")
				}
			}

			o.Delete(&category)
			return 0
		} else {
			return 1
		}
	}

}

func CreateCateList() ([]*CategoryListOne, error) {
	cateOne := RemoveLink()
	var CateListOut []*CategoryListOne
	var err error
	for i := 0; i < len(cateOne); i++ {
		var temp CategoryListOne
		temp.Id = i + 1
		temp.CateOne = cateOne[i]

		temp.CateTwo, err = FindAllSon(cateOne[i].Id)
		temp.DepthNumber = 33*len(temp.CateTwo) + 1
		if err != nil {
			return CateListOut, err
		}
		temp.CateOnesChild, err = FindAllChild(cateOne[i].Id)
		if err != nil {
			return CateListOut, err
		}
		CateListOut = append(CateListOut, &temp)
	}
	return CateListOut, err

}

func FindDepthIsHigh() []*Category {
	var Cate []*Category

	cateall, _ := GetAllCategories()
	for i := 0; i < len(cateall); i++ {
		var temp *Category
		temp = cateall[i]
		if temp.Depth == 100 {
			Cate = append(Cate, temp)
		}
	}
	return Cate
}

func SearchCategory(id int) (cate Category, err error) {
	o := orm.NewOrm()

	cate = Category{Id: id}
	if o.Read(&cate, "Id") == nil {
		return cate, err
	}
	return cate, nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cate := make([]*Category, 0)
	_, err := o.QueryTable("category").All(&cate)
	return cate, err
}

func GetSortedCategories() ([]*Category, error) {
	CateNotSort, err := GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	CateSort := make([]*Category, 0)
	CateSort = MakeSort(0, CateNotSort)
	return CateSort, err
}

func FindAllFather(id int) []*Category {
	var CateFather []*Category

	cate, _ := SearchCategory(id)

	for cate.Id != 0 {
		var temp Category
		temp = cate

		CateFather = append(CateFather, &temp)
		cate, _ = SearchCategory(cate.ParentID)
	}
	var CateFatherOut []*Category
	var len1 int = len(CateFather) - 1
	for j := len1; j >= 0; j-- {
		var temp *Category
		temp = CateFather[j]
		CateFatherOut = append(CateFatherOut, temp)
	}
	return CateFatherOut
}

func FindLeftMenu(id int) []*Category {
	cate, _ := SearchCategory(id)
	catefather, _ := SearchCategory(cate.ParentID)
	LeftMenu, _ := FindAllChild(catefather.Id)
	if cate.Depth == 100 {
		LeftMenu = FindDepthIsHigh()
		return LeftMenu
	} else if cate.Depth == 200 {
		LeftMenu = FindDepthIsHigh()
		return LeftMenu
	}
	return LeftMenu
}

func FindLeftMenuDepthTwo(id int) []*Category {
	cate, _ := SearchCategory(id)
	catefather, _ := SearchCategory(cate.ParentID)
	LeftMenu, _ := FindAllSon(catefather.Id)
	if cate.Depth == 100 {
		LeftMenu = FindDepthIsHigh()
		return LeftMenu
	} else if cate.Depth == 200 {
		LeftMenu = FindDepthIsHigh()
		return LeftMenu
	}
	return LeftMenu
}

func FindLeftMenuDepthThree(id int) []*Category {
	cate, _ := SearchCategory(id)
	catefather, _ := SearchCategory(cate.ParentID)
	LeftMenu, _ := FindAllChild(catefather.Id)
	if cate.Depth == 100 {
		LeftMenu = FindDepthIsHigh()
		return LeftMenu
	} else if cate.Depth == 200 {
		LeftMenu = FindDepthIsHigh()
		return LeftMenu
	}
	LeftMenuOut := make([]*Category, 0)
	for i := 0; i < len(LeftMenu); i++ {
		if LeftMenu[i].Depth == 3 && LeftMenu[i].ParentID == 3 {
			var temp *Category
			temp = LeftMenu[i]
			LeftMenuOut = append(LeftMenuOut, temp)
		}
	}
	return LeftMenuOut
}

// func CreateLeftMenu(){}
//寻找所有的儿子节点节点
func FindAllSon(id int) ([]*Category, error) {
	var son []*Category
	father, err := SearchCategory(id)

	if father.ChildID != "0" {
		sonid := strings.Split(father.ChildID, ",")

		for i := 0; i < len(sonid); i++ {
			j, _ := strconv.Atoi(sonid[i])
			temp, _ := SearchCategory(j)
			son = append(son, &temp)
		}
		return son, err
	}
	return son, err
}

//寻找所有的子节点
func FindAllChild(id int) ([]*Category, error) {
	var descendant []*Category
	cate, err := FindAllSon(id)
	if err != nil {
		return cate, err
	}
	for i := 0; i < len(cate); i++ {
		temp, _ := SearchCategory(cate[i].Id)
		descendant = append(descendant, &temp)
		cate1, _ := FindAllSon(cate[i].Id)
		for j := 0; j < len(cate1); j++ {
			temp, _ := SearchCategory(cate1[j].Id)
			descendant = append(descendant, &temp)
			cate2, _ := FindAllSon(cate1[j].Id)

			for k := 0; k < len(cate2); k++ {
				temp, _ := SearchCategory(cate2[k].Id)
				descendant = append(descendant, &temp)
				cate3, _ := FindAllSon(cate2[k].Id)
				for ii := 0; ii < len(cate3); ii++ {
					temp, _ := SearchCategory(cate3[ii].Id)
					descendant = append(descendant, &temp)
				}
			}

		}
	}
	return descendant, err
}

//输出排序好的并已经处理好格式的name和分栏
func StandardOut() ([]string, []*Category) {
	StanOut := make([]string, 0)
	var i int
	var CateOutPut []*Category
	CateSort, err := GetSortedCategories()
	CateNotSort, err := GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	for i = 0; i < len(CateNotSort); i++ {
		var j int
		if CateSort[i].Depth > 20 {
			var depth int = CateSort[i].Depth
			depth = depth % 50
			j = depth
		} else {
			j = CateSort[i].Depth - 1
		}
		StanOut = append(StanOut, strings.Repeat("--", j)+CateSort[i].Name)
	}
	CateOutPut = CateSort[0:i]
	return StanOut, CateOutPut
}

func GetAllCategoriesDepthIsOne() []*Category {
	_, cate := StandardOut()

	CateDepthOne := make([]*Category, 0)
	for i := 0; i < len(cate); i++ {
		var temp *Category
		if cate[i].Depth == 1 {
			temp = cate[i]
			CateDepthOne = append(CateDepthOne, temp)
		}

	}
	return CateDepthOne
}

func RemoveLink() []*Category {
	cateDepthOne := GetAllCategoriesDepthIsOne()
	CateOut := make([]*Category, 0)
	for i := 0; i < len(cateDepthOne); i++ {
		if cateDepthOne[i].Type != "链接列表" {
			CateOut = append(CateOut, cateDepthOne[i])
		}
	}
	return CateOut
}

//新建的struct,绑定id和处理好的name
func SortCategory() []*CategorySorted {
	CateName, CateSorted := StandardOut()
	// var CateOut []*CategorySorted
	CateOut := make([]*CategorySorted, 0)
	var i int = 0
	for i = 0; i < len(CateName); i++ {
		var temp CategorySorted
		temp.Id = CateSorted[i].Id
		temp.Name = CateName[i]
		CateOut = append(CateOut, &temp)

	}
	return CateOut
}

var CateSort []*Category //全局变量，用以存储已经找到的栏目分类
func MakeSort(id int, CateNotSort []*Category) []*Category {
	var i int
	if id == 0 {
		CateSort = make([]*Category, 0)
	}
	for i = 0; i < len(CateNotSort); i++ {
		if id == CateNotSort[i].ParentID {
			CateSort = append(CateSort, CateNotSort[i])
			MakeSort(CateNotSort[i].Id, CateNotSort)
		}
	}
	return CateSort
}
