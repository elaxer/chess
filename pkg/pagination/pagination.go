// Description: Пакет предоставляет код для работы с пагинацией объектов.
package pagination

import "sync"

// Pagination представляет пагинацию объектов.

type Pagination struct {
	// perPage - количество объектов на одной странице.
	perPage int
	// totalCount - общее количество всех объектов.
	totalCount int
	// page - номер текущей страницы.
	page int
	// pageButtonsCount - количество кнопок пагинации на одной странице.
	pageButtonsCount int

	pagesCount int
	pageStart  int
	pageEnd    int

	onceCount, onceStart, oncePage sync.Once
}

func New(perPage, totalCount, page, pageButtonsCount int) *Pagination {
	return &Pagination{
		perPage:          perPage,
		totalCount:       totalCount,
		page:             page,
		pageButtonsCount: pageButtonsCount,
	}
}

// Limit возвращает число максимально возможного количества объектов на одной странице.
func (p *Pagination) Limit() int {
	return p.perPage
}

// Offset возвращает смещение для запроса к репозиторию.
func (p *Pagination) Offset() int {
	return p.Page() - 1
}

// Page возвращает номер текущей страницы.
// Страница не может быть меньше 1 и больше общего количества страниц.
func (p *Pagination) Page() int {
	p.oncePage.Do(func() {
		p.page = min(max(p.page, 1), p.PagesCount())
	})

	return p.page
}

// PagesCount возвращает общее количество страниц.
func (p *Pagination) PagesCount() int {
	p.onceCount.Do(func() {
		if p.totalCount < p.perPage {
			p.pagesCount = 1
			return
		}

		p.pagesCount = p.totalCount / p.perPage
		if p.totalCount%p.perPage != 0 {
			p.pagesCount++
		}
	})

	return p.pagesCount
}

// PagesRange возвращает диапазон страниц для пагинации.
// Возвращает начальную и конечную страницу.
func (p *Pagination) PagesRange() (int, int) {
	p.onceStart.Do(func() {
		if p.totalCount < p.perPage {
			p.pageStart = 1
			p.pageEnd = 1
			return
		}

		p.pageStart = max(1, p.page-(p.pageButtonsCount/2))
		p.pageEnd = min(p.pageStart+p.pageButtonsCount-1, p.PagesCount())

		if p.pageEnd-p.pageStart+1 < p.pageButtonsCount {
			p.pageStart = max(1, p.pageEnd-p.pageButtonsCount+1)
		}
	})

	return p.pageStart, p.pageEnd
}
