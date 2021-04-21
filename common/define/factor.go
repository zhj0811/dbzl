package define

//保单
type Policy struct {
	Number     string `json:"number" binding:"required" gorm:"not null;primaryKey"` //保单号
	Type       string `json:"type" gorm:"not null"`                                 //保单类型
	Insured    string `json:"insured"`                                              //被保险人
	USCC       string `json:"uscc" gorm:"uscc"`                                     //统一社会信用代码
	StartAt    string `json:"start_at" gorm:"start_at"`                             //起保时间
	ExpireAt   string `json:"expire_at" gorm:"expire_at"`                           //终保时间
	Pooled     bool   `json:"pooled"`                                               //是否共保
	Insurer    string `json:"insurer"`                                              //承保人
	Amount     string `json:"amount"`                                               //保险金额（元）
	Premium    string `json:"premium"`                                              //保险费（元）
	Rate       string `json:"rate"`                                                 //保险费率
	Content    string `json:"content"`                                              //保障内容
	Extension1 string `json:"extension1"`                                           //预留字段1
	Extension2 string `json:"extension2"`                                           //预留字段2
	Extension3 string `json:"extension3"`                                           //预留字段3
}

type Service struct {
	//ID           string `json:"id" gorm:"column:id;primary_key;"`
	Number       string `json:"number" binding:"required" gorm:"not null;index"` //保单号
	Insured      string `json:"insured"`                                         //被保险人
	Date         string `json:"date"`                                            //服务时间
	Organization string `json:"organization"`                                    //服务机构
	Type         string `json:"type"`                                            //服务类型
	Extension1   string `json:"extension1"`                                      //预留字段1
	Extension2   string `json:"extension2"`                                      //预留字段2
	Extension3   string `json:"extension3"`                                      //预留字段3
}
