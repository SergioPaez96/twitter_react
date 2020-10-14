package models

/*Tweet captura del Body, el nensaje que nos llega*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
