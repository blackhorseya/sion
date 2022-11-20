package repo

type loginResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         *data  `json:"Data"`
}

type token struct {
	AccessToken      string `json:"Access_token"`
	RefrashToken     string `json:"Refrash_token"`
	RxpiresIn        int    `json:"Rxpires_in"`
	RefrashRxpiresIn int    `json:"Refrash_Rxpires_in"`
}

type userData struct {
	Memidno        string `json:"MEMIDNO"`
	Memcname       string `json:"MEMCNAME"`
	Memtel         string `json:"MEMTEL"`
	Memhtel        string `json:"MEMHTEL"`
	Membirth       string `json:"MEMBIRTH"`
	Memareaid      int    `json:"MEMAREAID"`
	Memaddr        string `json:"MEMADDR"`
	Mememail       string `json:"MEMEMAIL"`
	Memcomtel      string `json:"MEMCOMTEL"`
	Memcontract    string `json:"MEMCONTRACT"`
	Memcontel      string `json:"MEMCONTEL"`
	Memmsg         string `json:"MEMMSG"`
	Cardno         string `json:"CARDNO"`
	Unimno         string `json:"UNIMNO"`
	Memsendcd      int    `json:"MEMSENDCD"`
	Carrierid      string `json:"CARRIERID"`
	Npoban         string `json:"NPOBAN"`
	HasCheckMobile int    `json:"HasCheckMobile"`
	NeedChangePWD  int    `json:"NeedChangePWD"`
	HasBindSocial  int    `json:"HasBindSocial"`
	HasVaildEMail  int    `json:"HasVaildEMail"`
	Audit          int    `json:"Audit"`
	IrFlag         int    `json:"IrFlag"`
	PayMode        int    `json:"PayMode"`
	RentType       int    `json:"RentType"`
	IDPic          int    `json:"ID_pic"`
	DDPic          int    `json:"DD_pic"`
	MOTORPic       int    `json:"MOTOR_pic"`
	AAPic          int    `json:"AA_pic"`
	F01Pic         int    `json:"F01_pic"`
	SignturePic    int    `json:"Signture_pic"`
	SigntureCode   string `json:"SigntureCode"`
	Memrfnbr       string `json:"MEMRFNBR"`
	Signature      string `json:"SIGNATURE"`
}

type data struct {
	Token    *token    `json:"Token"`
	UserData *userData `json:"UserData"`
}
