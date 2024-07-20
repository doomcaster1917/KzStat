package Schemes

// GetIndexTreeData Документация АПИ: https://taldau.stat.gov.kz/ru/Api/Dev#_Toc500231954
type GetIndexTreeData struct {
	p_index_id  uint32  //Идентификатор показателя
	p_period_id uint8   //Идентификатор типа периода
	p_terms     uint32  //Список элементов разделенных через запятую для выборки (termIds из GetSegmentList);
	p_term_id   uint32  //Указывается один главный элемент по которому нужна детализация (один из p_terms);
	p_dicIds    []uint8 //Список справочников через запятую (dicId из GetSegmentList);
	idx         uint8   //индекс разрезности (idx из GetSegmentList);
	parent_id   uint32  //Идентификатор родительского элемента. Для получения корневого элемента следует оставлять пустым.
}

// GetSegmentList Информация о разрезностях
type GetSegmentList struct {
	indexId  uint32 //Идентификатор показателя
	periodId uint8  //Идентификатор типа периода (из запроса GetPeriodList).
}

// GetPeriodList информация о доступных типах периодов
type GetPeriodList struct {
	indexId uint32 //Указывается идентификатор показателя.
}

// GetIndexPeriods информация о доступных датах для отображения
type GetIndexPeriods struct {
	p_index_id  uint32  //Идентификатор показателя
	p_period_id uint8   //Идентификатор типа периода;
	p_terms     []uint8 //Список элементов разделенных через запятую для выборки (termIds из GetSegmentList);
	p_term_id   uint32  //Указывается один главный элемент по которому нужна детализация (один из p_terms)
	p_dicIds    []uint8 //Список справочников через запятую (dicId из GetSegmentList)
}

// GetIndexAttributes метод получения информации о показателе
type GetIndexAttributes struct {
	indexId  uint32 //Идентификатор показателя
	periodId uint8  //Тип периода
}

// IndexFeed разработчики АПИ перегрузили эндпойнт через аргумент pageType
// поэтому структура запроса будет иметь методы Branch и Region, удовлетворяющие 2 значениям pageType
type IndexFeed struct {
	pageType string //Значение задаётся константно методами
	periodID uint8  //Идентификатор показателя
	termID   uint32 //Идентификатор сектора экономики/раздела
}

// Branch Список показателей для селектора экономики/раздела
func (v IndexFeed) Branch(periodID uint8, termID uint32) IndexFeed {
	v.pageType = "Branch"
	v.periodID = periodID
	v.termID = termID
	return v
}

// Region Список показателей для селекторов регионов
func (v IndexFeed) Region(periodID uint8, termID uint32) IndexFeed {
	v.pageType = "Region"
	v.periodID = periodID
	v.termID = termID
	return v
}

type GetIndustryPeriods struct {
	termID uint32 //Идентификатор сектора экономики/раздела
}

type Search struct {
	keyword any //Часть наименования или кода показателя
}

// GetRegionPeriods Информация о доступных типах периодов для регионов
var GetRegionPeriods string = "http://taldau.stat.gov.kz/ru/Api/GetRegionPeriods"
