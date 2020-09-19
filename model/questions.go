package model

//type Question struct {
//	Name        string
//	Description string
//	Type        string
//	Example     string
//}

type CandidateStore struct {
	Q1  string   `json:"q1"`
	Q2  bool     `json:"q2" bson:"q2"`
	Q2d string   `json:"q2d" bson:"q2d"`
	Q3  int64    `json:"q3" bson:"q3"` // Укажите Ваш стаж работы кладовщиком
	Q4  []string `json:"q4"`
	Q5  []string
	Q6  []string
	Q7  []string `json:"q7"`
	Q8  []string
	Q9  []string `bson:"q9"`
	Q10 int      // Сколько времени вы потратили на поиск работы кладовщика?
	Q11 int      // Уровень ЗП (желаемый)
}
