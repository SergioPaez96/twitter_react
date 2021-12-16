package models

import "time"

/*GraboTweet es la estructura que tendrá nuestro Tweet*/
type GraboTweet struct {
	UserID  string    `bson:"userId" json:"userId,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
