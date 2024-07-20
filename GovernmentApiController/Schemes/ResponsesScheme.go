package Schemes

import "time"

// GetIndexTreeDataResponse схема ответа данных показателя

type GetIndexTreeDataResponse struct {
	id          uint32 //Идентификатор элемента. Служит для заполнения поля p_parent_id для получения дочерних элементов;
	rownum      uint32
	text        string //Текстовое описание элемента
	leaf        bool   //True если узел не содержит дочерних элементов, false иначе
	expanded    bool   //True указывает что узел следует развернуть для leaf = true;
	measureName string
	y           map[int]int //yXXXXX - это значение для определенной даты (XXXXX соответствует одному элементу из значений dateList запроса GetIndexPeriods).
}

//В данном примере возвращает данные по РК (id=741880). Для того чтобы получить данные дочерних элементов нужно выполнить запрос следующего вида:
//http://taldau.stat.gov.kz/ru/Api/GetIndexTreeData?p_measure_id=1&p_index_id=2709379&p_period_id=7&p_terms=741880&p_term_id=741880&p_dicIds=67&idx=0&p_parent_id=741880
//где p_parent_id указывает на id из первого запроса GetIndexTreeData.

// GetSegmentListResponse Информация о разрезностях.
type GetSegmentListResponse struct {
	dicId      []int //":"67 + 915",
	dicClassId []int //:"213 + 4855",
	ids        []int
	names      []string //:"Регионы + Виды экономической деятельности ",
	fullNames  []string //:"<u><b>1.Регионы</b></u><br><b>Полное наименование:</b> Классификатор административно-территориальных объектов ГК РК 11-2009<br><b>Аббревиатура:</b> КАТО(по каталогу)<br><b>Категория:</b> Национальный",
	termIds    []int    //:" 741880,741885",
	termNames  []string //:" РЕСПУБЛИКА КАЗАХСТАН + Всего ",
	dicCount   int16    //:2,
	termPath   string   //:"/",
	idx        int8     //:0,
	order      int8     //:0,
	decFormat  int8     //:1
}

// GetPeriodListResponse Информация о доступных типах периодов.
type GetPeriodListResponse struct {
	id   int8   //Тип периода. В дальнейшем используется в качестве параметра periodId для других запросов;
	name string //Название типа периода.
}

// GetIndexPeriodsResponse Информация о доступных датах для отображения
type GetIndexPeriodsResponse struct {
	dateList       []time.Time //Это внутренние имена полей;
	periodNameList []time.Time //Наименования дат для отображения;
	datesToDraw    []time.Time //Даты в формате dd.mm.yyyy, если потребуется конвертировать в/из типа дата;
	datesIds       []int16     //Внутренние идентификаторы дат.
}

// GetIndexAttributesResponse метод получения информации о показателе
type GetIndexAttributesResponse struct {
	id                   uint32   //Указывается идентификатор показателя;
	path                 []uint32 //:"/0/700894/700895/700911" Иерархия id по схеме родительский/дочерный элемент
	decFormat            int8     // Указывает на точность количества знаков после запятой. -1 означает что точность не определена. Пример: 0;
	namePath             []string // Это иерархия родительских показателей до текущего показателя. Пример:"Национальные счета/Текущие счета/Счет производства";
	name                 string   // Наименование показателя. Пример: "Валовой внутренний продукт методом производства";
	cgParams             []int8   //Пример:"1,1,1,1",
	measureId            int16    //Measure - это префикс для единиц измерения; Пример:1,
	measureKfc           int8     //Kfc - множитель, коэфициент относительного корневой единицы измерения; Пример:1;
	measureSign          string
	measureName          string //Наименование единицы измерения. Пример: "тенге";
	preferredMeasureId   int16  //Пример: 1;
	preferredMeasureName string //Пример:"тенге";
	preferredMeasureKfc  int8   //Пример:1,
	preferredMeasureSign string
}

// IndexFeedBranchResponse Список показателей для селектора экономики/раздела. Ответ в XML.
type IndexFeedBranchResponse struct {
	title         string    // Наименование показателя. Пример: <title>Показатели - Статистика предприятий</title>
	description   string    // Описание показателя. Пример: <description> Раздел статистики. Периодичность: Месяц с накоплением </description>
	lastBuildDate time.Time // Пример: <lastBuildDate>30.04.2016</lastBuildDate>
	language      string    // Пример: <language>ru</language>
	pubDate       time.Time // Пример: <pubDate>30.04.2016</pubDate>
	growth        float64   // Пример: <growth>-0,0833838024769378</growth>
	trend         string    // Наименование тренда. Пример: <trend>Тренд не определен</trend>
	diff          int64     // item/diff - это прирост(+)/снижение(-). Пример:<diff>-190</diff>
	measure       string    //<measure>единица</measure>
}

// IndexFeedRegionResponse Список показателей для селектора регионов. Ответ в XML.
type IndexFeedRegionResponse struct {
	title         string    //<title>Основные показатели - Акмолинская область</title>
	description   string    //<description>	Статистические показатели - Акмолинская область. Периодичность: Год</description>
	lastBuildDate time.Time //<lastBuildDate>30.04.2016</lastBuildDate>
	language      string    // Пример: <language>ru</language>
	pubDate       time.Time // Пример: <pubDate>30.04.2016</pubDate>
	growth        float64   // Пример: <growth>-0,0833838024769378</growth>
	trend         string    // Наименование тренда. Пример: <trend>Стабильный рост</trend>
	diff          int64     // item/diff - это прирост(+)/снижение(-). Пример:<diff>0,1198</diff>
	measure       string    // <measure>человек</measure>
}

// GetRegionPeriodsResponse Информация о доступных периодах для регионов
type GetRegionPeriodsResponse struct {
	periodId   uint8  //7
	periodName string // "Квартал"
}

// GetIndustryPeriodsResponse Информация о доступных периодах для секторов экономики/разделов
type GetIndustryPeriodsResponse struct {
	periodId   uint8  //7
	periodName string // "Квартал"
}

// SearchResponse Поиск показателей по наименованию или коду
type SearchResponse struct {
	total   uint16      //"1",
	results map[any]any //[{ "id":3768962, "Name":"штатное количество должностей по оказанию социально-правовых услуг", "Code":"66210808"}]
}
