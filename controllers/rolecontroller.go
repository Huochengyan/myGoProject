package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleC struct {
	Mgo *mongo.Database
	//RedisCli *redis.Client
}

///*   获得角色  */
//func (m RoleC) GetRoles(g *gin.Context) {
//	rsp := new(Rsp)
//	var roles [] Role
//	cur, err := m.Mgo.Collection(db.Role).Find(context.Background(), bson.D{}, nil)
//	if err == nil {
//		for cur.Next(context.Background()) {
//			elme := new(Role)
//			err := cur.Decode(elme)
//			if err == nil {
//				roles = append(roles, *elme)
//			}
//		}
//		rsp.Msg = "success"
//		rsp.Code = 200
//		rsp.Data = roles
//		g.JSON(http.StatusOK, rsp)
//		return
//	}
//	rsp.Msg = "falid"
//	rsp.Code = 201
//	g.JSON(http.StatusOK, rsp)
//	return
//}
