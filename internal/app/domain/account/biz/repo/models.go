package repo

type loginResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		Token struct {
			AccessToken      string `json:"Access_token"`
			RefrashToken     string `json:"Refrash_token"`
			RxpiresIn        int    `json:"Rxpires_in"`
			RefrashRxpiresIn int    `json:"Refrash_Rxpires_in"`
		} `json:"Token"`
		UserData struct {
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
		} `json:"UserData"`
	} `json:"Data"`
}

type getMemberStatusResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		StatusData struct {
			Memidno          string `json:"MEMIDNO"`
			Memname          string `json:"MEMNAME"`
			Login            string `json:"Login"`
			Register         int    `json:"Register"`
			Audit            int    `json:"Audit"`
			AuditID          int    `json:"Audit_ID"`
			AuditCar         int    `json:"Audit_Car"`
			AuditMotor       int    `json:"Audit_Motor"`
			AuditSelfie      int    `json:"Audit_Selfie"`
			AuditF01         int    `json:"Audit_F01"`
			AuditSignture    int    `json:"Audit_Signture"`
			BlackList        string `json:"BlackList"`
			MenuCTRL         int    `json:"MenuCTRL"`
			MenuStatusText   string `json:"MenuStatusText"`
			StatusTextCar    string `json:"StatusTextCar"`
			StatusTextMotor  string `json:"StatusTextMotor"`
			NormalRentCount  int    `json:"NormalRentCount"`
			AnyRentCount     int    `json:"AnyRentCount"`
			MotorRentCount   int    `json:"MotorRentCount"`
			TotalRentCount   int    `json:"TotalRentCount"`
			Score            int    `json:"Score"`
			BlockFlag        int    `json:"BlockFlag"`
			BlockEdate       string `json:"BLOCK_EDATE"`
			CMKStatus        string `json:"CMKStatus"`
			IsShowBuy        string `json:"IsShowBuy"`
			HasNoticeMsg     string `json:"HasNoticeMsg"`
			AuthStatus       string `json:"AuthStatus"`
			BindHotai        string `json:"BindHotai"`
			IsHIMS           string `json:"IsHIMS"`
			GameMode         int    `json:"GameMode"`
			GameMsg          string `json:"GameMsg"`
			GameCover        int    `json:"GameCover"`
			IsCompMemAudit   int    `json:"IsCompMemAudit"`
			InsuranceDefault int    `json:"InsuranceDefault"`
		} `json:"StatusData"`
	} `json:"Data"`
}
