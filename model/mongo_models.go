package model

type CandidateStore struct {
	Q1  string   `json:"q1" bson:"q1"`   //Желаемая должность
	Q2  bool     `json:"q2" bson:"q2"`   // У вас есть опыт работы кладовщиком?
	Q2d string   `json:"q2d" bson:"q2d"` // У вас есть опыт работы кладовщиком?
	Q3  int64    `json:"q3" bson:"q3"`   // Укажите Ваш стаж работы кладовщиком
	Q4  []string `json:"q4" bson:"q4"`   // Отрасль
	Q5  []string `json:"q5" bson:"q5"`
	Q6  []string `json:"q6" bson:"q6"`
	Q7  []string `json:"q7" bson:"q7"`
	Q8  []string `json:"q8" bson:"q8"`
	Q9  []string `json:"q9" bson:"q9"`
	Q10 int      `json:"q10" bson:"q10"` // Сколько времени вы потратили на поиск работы кладовщика?
	Q11 int      `json:"q11" bson:"q11"` // Уровень ЗП (желаемый)
}
