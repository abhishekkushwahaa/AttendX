package models

type User struct {
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	Fingerprint string `json:"fingerprint" bson:"fingerprint"`
	FaceImage   []byte `json:"face_image" bson:"face_image"`
}
