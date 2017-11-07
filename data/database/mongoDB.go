package database

import (
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"net"
	"gopkg.in/mgo.v2/bson"
	"github.com/adgluzdov/IMobyServer/model"
)

type MongoDB struct {
	Session *mgo.Session
}

var database *MongoDB

func GetMongoDB() *MongoDB {
	if database == nil {
		database = &MongoDB{}
		database.Init()
	}
	return database
}

func (this *MongoDB)Init()(err error){
	mongoURI := "mongodb://adgluzdoff:DmgKon490YedAg@cluster-shard-00-00-jz1qh.mongodb.net:27017,cluster-shard-00-01-jz1qh.mongodb.net:27017,cluster-shard-00-02-jz1qh.mongodb.net:27017/vkbot?replicaSet=Cluster-shard-0&authSource=admin"
	dialInfo, err := mgo.ParseURL(mongoURI)
	if err != nil {return err }
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	this.Session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {return err }
	this.Session.SetMode(mgo.Monotonic, true)
	return nil
}

func (this *MongoDB) FindAccount(uid string,account *model.Account)(err error)  {
	query := bson.M{"uid":uid}
	err = this.Session.DB("IMoby").C("Accounts").Find(query).One(account)
	return
}

func (this *MongoDB) FindAccountById(Id bson.ObjectId,account *model.Account)(err error)  {
	query := bson.M{"_id":Id}
	err = this.Session.DB("IMoby").C("Accounts").Find(query).One(account)
	return
}

func (this *MongoDB) FindMarketById(Id bson.ObjectId,market *model.Market)(err error)  {
	query := bson.M{"_id":Id}
	err = this.Session.DB("IMoby").C("Markets").Find(query).One(market)
	return
}

func (this *MongoDB) FindMarketAll(markets interface{})(err error)  {
	err = this.Session.DB("IMoby").C("Markets").Find(nil).All(markets)
	return
}

func (this *MongoDB) FindToken(accsses_token string,result interface{})(err error)  {
	err = this.Session.DB("IMoby").C("Tokens").Find(bson.M{"info.accsses_token":accsses_token}).One(result)
	return
}

func (this *MongoDB) DeleteTokenById(Id string)(err error)  {
	query := bson.M{"_id":Id}
	err = this.Session.DB("IMoby").C("Tokens").Remove(query)
	return
}

func (this *MongoDB) InsertAccount(user model.Account)(err error)  {
	err = this.Session.DB("IMoby").C("Accounts").Insert(user)
	return
}

func (this *MongoDB) CreateMarket(market model.Market) (bson.ObjectId,error) {
	market.Id = bson.NewObjectId()
	err := this.Session.DB("IMoby").C("Markets").Insert(market)
	return market.Id,err
}


func (this *MongoDB) InsertToken(token model.Token)(err error)  {
	err = this.Session.DB("IMoby").C("Tokens").Insert(token)
	return
}


func (db *MongoDB)Close(){
	db.Session.Close()
}

func (db *MongoDB)FindAll(collation string,query interface{},result interface{})(err error){
	c := db.Session.DB("IMoby").C(collation)
	err = c.Find(query).All(result)
	return
}

func (db *MongoDB)GetCount(collation string,query interface{})(n int,err error){
	c := db.Session.DB("IMoby").C(collation)
	n,err = c.Find(query).Count()
	return
}

func (db *MongoDB)Update(collation string,selector interface{},update interface{})(err error){
	c := db.Session.DB("IMoby").C(collation)
	err = c.Update(selector,update)
	return
}
