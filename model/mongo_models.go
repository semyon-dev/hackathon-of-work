package model

// Модель 1 - «Кладовищик»
type CandidateStore struct {
	Description string   `json:"description" bson:"description"`
	Id          int64    `json:"id" bson:"_id"`
	CandidateID string   `json:"candidate_id" bson:"candidate_id"`
	Q1          string   `json:"q1" bson:"q1"`   // Желаемая должность
	Q2          bool     `json:"q2" bson:"q2"`   // У вас есть опыт работы кладовщиком?
	Q2d         string   `json:"q2d" bson:"q2d"` // У вас есть опыт работы кладовщиком?
	Q3          int64    `json:"q3" bson:"q3"`   // Укажите Ваш стаж работы кладовщиком
	Q4          []string `json:"q4" bson:"q4"`   // Отрасль
	Q5          []string `json:"q5" bson:"q5"`   // Названия компаний
	Q6          []string `json:"q6" bson:"q6"`   // С какой системой хранения Вы имели опыт работы?
	Q7          []string `json:"q7" bson:"q7"`   // Опыт работы с программами
	Q8          []string `json:"q8" bson:"q8"`   // Опыт с инструментарием
	Q9          []string `json:"q9" bson:"q9"`   // Типы работ
	Q10         int      `json:"q10" bson:"q10"` // Сколько времени вы потратили на поиск работы кладовщика?
	Q11         int      `json:"q11" bson:"q11"` // Уровень ЗП (желаемый)
}

// Модель 2 - «Водитель погрузчика»
type CandidateDriver struct {
	Description string   `json:"description" bson:"description"`
	Id          int64    `json:"id" bson:"_id"`
	CandidateID string   `json:"candidate_id" bson:"candidate_id"`
	Q1          string   `json:"q1" bson:"q1"`   // Желаемая должность
	Q2          bool     `json:"q2" bson:"q2"`   // У вас есть водительские права?
	Q3          int64    `json:"q3" bson:"q3"`   // Водительское удостоверение
	Q4          string   `json:"q4" bson:"q4"`   // Удостоверение тракториста-машиниста нового образца
	Q5          string   `json:"q5" bson:"q5"`   // Удостоверение тракториста-машиниста старого образца
	Q6          bool     `json:"q6" bson:"q6"`   // У вас есть опыт работы водителем погрузчиком?
	Q7          int      `json:"q7" bson:"q7"`   // Ваш стаж работы водителем погрузчика
	Q8          []string `json:"q8" bson:"q8"`   // Сфера деятельности
	Q9          []string `json:"q9" bson:"q9"`   // Названия компаний
	Q10         []string `json:"q10" bson:"q10"` //  Какой техникой вы управляете?
	Q11         int      `json:"q11" bson:"q11"` // Максимальная высота поднятия груза
	Q12         []string `json:"q12" bson:"q12"` // С какой системой хранения Вы имели опыт работы?
	Q13         bool     `json:"q13" bson:"q13"` // удостоверение стропальщика?
	Q14         []string `json:"q14" bson:"q14"` // Типы работ
	Q15         bool     `json:"q15" bson:"q15"` // Возможность провести тест-драйв на собеседовании
	Q16         []string `json:"q16" bson:"q16"` // Ключевые качества
	Q17         []string `json:"q17" bson:"q17"` // Знание языков -	Уровень владения (4 - родной, 3 - разговорный, 2 - понимаю, 1 - умею читать)
	Q19         string   `json:"q19" bson:"q19"` // Хобби (футбол)
	Q20         int      `json:"q20" bson:"q20"` // Уровень ЗП (30000)
	Q21         int      `json:"q21" bson:"q21"` // График работы (1 - полный, 2 - неполный день, 3 - сменный, 4 - удалённая работа, 5 - вахтовый метод
	Q22         bool     `json:"q22" bson:"q22"` // Ночная смена
	Q23         bool     `json:"q23" bson:"q23"` // Командировки
	Q24         []string `json:"q24" bson:"q24"` // Станции метро
}

// Модель 3 - «Официант»
type CandidateRestaurant struct {
	Description string   `json:"desciption" bson:"desciption"`
	Id          int64    `json:"id" bson:"_id"`
	CandidateID string   `json:"candidate_id" bson:"candidate_id"`
	Q1          string   `json:"q1" bson:"q1"`   // Желаемая должность
	Q2          bool     `json:"q2" bson:"q2"`   // Есть ли у Вас профессиональное образование в сфере гостеприимства?
	Q3          string   `json:"q3" bson:"q3"`   // профессиональное образовательное учреждение, где вы учились, и специальность
	Q4          int      `json:"q4" bson:"q4"`   // Год окончания обучения
	Q5          bool     `json:"q5" bson:"q5"`   // Стаж работы официантом
	Q6          []string `json:"q6" bson:"q6"`   // Рестораны, где работал
	Q7          []string `json:"q7" bson:"q7"`   // дополнительные профессиональные знания и навыки
	Q8          []string `json:"q8" bson:"q8"`   // Знание языков - Уровень владения (5 - native speaker, 4 - fluent, 3 - advanced,  2- intermediate, 1- pre-intermediate)
	Q9          []string `json:"q9" bson:"q9"`   // C какой посудой работали
	Q10         []int    `json:"q10" bson:"q10"` // График работы (1 - полный, 2 - неполный день, 3 - сменный, 6 - ночная смена)
}

// Пример JSON
//
//{
//"q1": "Официант",
//"q2": "True",
//"q3": "МГУ, философский факультет",
//"q4": "2015",
//"q5": "True",
//"q6": ["Ромашка", "Цветочек"],
//"q7": ["Бариста", "Сомелье"],
//"q8": [{"lang":"Русский", "v":4}, {"lang":"Английский", "v":1}],
//"q9": ["Кость", "Snow white"],
//"q10": [2, 3]
//}
