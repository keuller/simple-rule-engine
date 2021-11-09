package entity

type Rule struct {
	ID         string `bson:"id"`
	Title      string `bson:"title"`
	When       string `bson:"when"`
	Then       string `bson:"then"`
	Conditions string `bson:"-"`
}
