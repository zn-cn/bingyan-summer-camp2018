
package models

import (
    "github.com/astaxie/beego/orm"

 //  "fmt"
   "strconv"
  _ "github.com/go-sql-driver/mysql"
) 

func RegisterDB(){
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:123456@/management?charset=utf8")
    orm.RegisterModel(new(User) )
}


type User struct {
    Id     int64    `form:"-"`
    Name   string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
    Password string `form:"password,text,password:" valid:"MinSize(6);MaxSize(20)"`
    Email   string `form:"client,text,client:"`
    Nick    string `form:"nick,text,nick:"`
}


func AddUser(username string,pass string,em string ,ni string) error {
	o := orm.NewOrm()
    //创建数据库表
	user := &User{    

      Name     : username,
      Password : pass,
      Email    : em,
      Nick     : ni,

    } 

	// 查询数据
	qs := o.QueryTable("user")
	err := qs.Filter("name", username).One(user)
	if err == nil { 
		return err
    } 


	// 插入数据
	_, err = o.Insert(user)
	if err != nil {
        
		return err
	} 

	return nil
}

//查询成员 
func FindUserByUserName(username string,password string) (bool,bool, User) {
	o := orm.NewOrm()
	var user User
    err1 := o.QueryTable(user).Filter("name", username).One(&user)
    err2 := o.QueryTable(user).Filter("password", password).One(&user)
	return err1 != orm.ErrNoRows, err2 != orm.ErrNoRows,user
}
/*
func ReadUser(username string) （error） {
    o :=orm.NewOrm()
    us := &User{ Name : username} 

    //查询数据
    qs := o.QueryTable("user")
    err := qs.Filter("name", username).One(us)
    if err == nil { 
    return err
   } 
   return nil
}
*/


func DeleteUser(id string) error {
    
	cid, err := strconv.ParseInt(id, 10,64)
	if err != nil {
		return err
    } 
    

	o := orm.NewOrm()

	us := &User{Id: cid}
	_, err = o.Delete(us)
	return err
}

func GetAllUser() ([]*User, error) {
	o := orm.NewOrm()

	users := make([]*User, 0)

	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}




    

































/*


package models

import (
	//"fmt"
//	_ "management/routers"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
//	_ "github.com/go-sql-driver/mysql"

)





type Article struct {
    Id     int    `form:"-"`
    Name   string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
    Client string `form:"client,text,client:"`
    Url    string `form:"url,text,url:"`
}

func (a *Article) TableName() string {
    return "articles"
}



/*

type User struct{
	Id int	
	Nick string
	Password string
     Email string
	Tel int
	
   }





	func init(){
		// set default database
		//orm.RegisterDataBase("default", "mysql", "root:123456@management?charset=utf8", 30)
	
		// register model
		orm.RegisterModel(new(Article))
	
		// create table
		orm.RunSyncdb("default", false, true)
	}





	*/
	
	



   // Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
   // Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系

/*
type Profile struct {
    Id          int
  //  Age         int16
    User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct { 
    Id    int
    Title string
    User  *User  `orm:"rel(fk)"`    //设置一对多关系
    Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
    Id    int
    Name  string
    Posts []*Post `orm:"reverse(many)"`
}
*/