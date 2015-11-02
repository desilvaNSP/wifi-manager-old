package dao
import (
	"wifi-manager/core/utils"
)

type NameValue struct{
	Name string  `db:"name"json:"name"`
	Value int  `db:"value"json:"value"`
}

type User struct {
	TenantId int `db:"tenantid"json:"tenantid"`
	Username string `db:"username"json:"username"`
	Password string `db:"password"json:"password"`
	Email string `db:"email"json:"email"`
	Status string `db:"status"json:"status"`
}

type PortalUser struct{
	Username string `db:"username"json:"username"`
	Password string `json:"password"`
	Acctstarttime utils.NullString `db:"acctstarttime"json:"acctstarttime"`
	Acctlastupdatedtime utils.NullString `db:"acctlastupdatedtime"json:"acctlastupdatedtime"`
	Acctactivationtime utils.NullString `db:"acctactivationtime"json:"acctactivationtime"`
	Acctstoptime utils.NullString `db:"acctstoptime"json:"acctstoptime"`
	Location utils.NullString  `db:"location"json:"location"`
	Visits int64     `db:"visits"json:"visits"`
}

type Role struct{
   BitMask int `json:"bitMask"`
   Title string `json:"title"`
}

type AuthUser struct {
	Username string `json:"username"`
	Role Role   `json:"role"`
}

type Constrains struct {
	From string `json:"from"`
	To string   `json:"to"`
	LocationId string    `json:"locationid"`
}

type Location struct {
	LocationId string    `json:"locationid"`
	LocationName string  `json:"locationname"`
	NasIP string         `json:"nasip"`
	IPFrom string        `json:"ipfrom"`
	IPTo string          `json:"ipto"`

}

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type DBConfigs struct{
	Username string
	Password string
	DBName string
	Host string
	Port int
}

type ServerConfigs struct{
	Sample string
	HttpPort int
	HttpsPort int
	ReadTimeOut int
	WriteTimeOut int
}