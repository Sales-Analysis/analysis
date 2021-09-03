# Sales-Analysis

<img align="right" width="159px" src="https://github.com/Sales-Analysis/analysis/blob/master/data/sa.svg">


Не требует узкоспециализированных знаний в области экономики, статистики, маркетинга, просто загрузите имеющиеся данные. 
Библиотека обработает их с помощью классических статистических методов. Поможет эффективно распорядиться ресурсами, в том числе и собственным временем. ☺️

# Контент
  - [Установка](#Установка)
  - [Типы анализов](#Типы-анализов)

# Установка

1. Сначала необходимо установить [Go](https://golang.org/) (требуется версия 1.15+), затем вы можете использовать приведенную ниже команду Go для установки Analysis.
``` golang
go get github.com/Sales-Analysis/analysis
```

2. Импорт
``` golang
import analysis "github.com/Sales-Analysis/analysis"
```

# Типы анализов
  - [ABC](#ABC)
    - [Что такое АВС анализ?](#Что-тако-АВС-анализ?)
    - [Схема расчёта](#Схема-расчёта)
    - [Получение отчета ABC](#Получение-отчета-ABC)
        - [Функция для расчета анализа abc](#Функция-для-расчета-анализа-abc)
        - [Формат возвращаемой структуры](#Формат-возвращаемой-структуры)
  - ~~XYZ~~
  - ~~ABC-XYZ~~
  - ~~FMR~~
  - ~~RFM~~
  - ~~VEN~~

## ABC
### Что такое АВС анализ?

ABC анализ продаж выявляет долю каждого из товаров (или категории) в общем обороте, ранжирует ассортиментные позиции (или категории) по степени значимости их вклада в общий оборот, присваивая условные классы: A, B, C

A – имеющие максимальную долю в обороте, суммарно составляющие 80% от общего оборота

B – следующие по степени важности вносимой доли, суммарно составляющие 15% оборота

C – наименее эффективные товары с минимально вносимой долей в оборот, суммарно 5%

Для расчёта требуется

Список товаров (и/или товарных категорий) с указанием продаж/прибыли по каждому

Пример исходных данных:

| Наименование  |  Продажи      |
| ------------- | ------------- |
| Коробка 20*20 | 5 850,00р.    |
| Коробка 20*12 | 4 500,00р.    |
| Коробка 10*10 | 2 000,00р.    |

#### Схема расчёта 

1. Сортируется список по убыванию значения продаж

2. Определяется общая сумму продаж

3. Определяется доля продаж по каждой анализируемой позиции (продажи по каждой делим на общую сумму, выводятся значение в %). Значения выносятся новым столбцом «доля»

4. Считается накопленная доля нарастающим итогом. Для первой позиции (первая – с максимальной долей, учитывая убывающую сортировку!) накопленная доля равнозначна доли, рассчитанной ранее. Начиная со второй, накопленная доля определяется суммой доли данной позиции и накопительной доли предыдущей позиции. Значения выносятся новым столбцом «накопленная доля»

5. Присваиваются категории A/B/C:

Группа A – позиции, наколенный итог, которых составляет до 80% (начиная с первой позиции в списке до той, значение, которой приближено или равно 80% присваивается «А»)

Группа B – позиции, накопленный итог которых от 80% до 95% (первая позиция, которой будет присвоена группа «B» - следующая, за нижней границей группы «А», последняя - значение, которой приближено или равно 95%)

Группа C – позиции, накопленный итог 95% - 100% (начиная с первой позиции после нижней границы группы B до последней в списке присваивается «С»)


| Наименование  |    Продажи    |  Доля продаж  |  Накопительная доля  |  Категории  |
| ------------- | ------------- | ------------- | -------------------- | ----------- |
| Коробка 1     |    100,00р    |      57%      |         57%          |      A      |
| Коробка 2     |     50,00р    |      28%      |         85%          |      B      |
| Коробка 3     |     20,00р    |      11%      |         97%          |      C      |
| Коробка 4     |      5,00р    |       2%      |        100%          |      C      |


#### Получение отчета ABC

##### Функция для расчета анализа abc

``` golang
// pluID - код идентификатор PLU
// measures - наименование анализируемых позиций
// dimensions - данные по анализируемому критерию (продажи/оборот/прибыль)
func ABC(pluID []int64, measures []string, dimensions []float64) (*abc.ABC, error)
```

##### Формат возвращаемой структуры

``` golang
type ABC struct {
	Measures []string // Наименование анализируемых позиций
	Dimensions []float64 // Данные по анализируемому критерию (продажи/оборот/прибыль)
	Deposit []float64 // Доля продаж
	CumulativeShare []float64 // Накопительная доля
	Group []string // Категории
	Duplicates []duplicate // Дубликаты анализируемых позиций, которые не попали в расчет
}
```
##### Коды ошибок

``` golang
measuresNotEqualZero       = "Measure size is 0"
dimensionsNotEqualZero     = "Dimension size is 0"
measuresNotEqualDimensions = "Size measure and dimension to equal"
negativeValueDimensions    = "In the dimensions is a negative value"
```
