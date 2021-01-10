package scanners

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/omekov/geminitesttask/internal/parsesite/model"
)

// Branch ...
func Branch() ([]model.Branch, error) {
	var err error
	c := colly.NewCollector()
	var branchs []model.Branch
	c.OnHTML("div.location-list-item", func(e *colly.HTMLElement) {
		about := e.ChildAttr("article.node--location-content-type", "about")
		linkStaff := e.Request.AbsoluteURL(about + "/about")
		staff, _ := Staff(linkStaff)
		branch := model.Branch{
			Name:      e.ChildText("h2.location-item--title"),
			Borough:   e.ChildText("div.field-borough"),
			Address:   e.ChildText("div.field-location-direction"),
			StaffLink: linkStaff,
			Staff:     staff,
		}
		branchs = append(branchs, branch)
	})
	if err != c.Visit("https://ymcanyc.org/locations?type&amenities") {
		return nil, err
	}
	return branchs, nil
}

// Staff ...
func Staff(link string) (staffs []string, err error) {
	c := colly.NewCollector()
	staffBlocks := []string{"div.field-prgf-2c-left", "div.field-prgf-description"}
	staffHeaders := []model.StaffHeader{
		{
			Element: "h3",
			Value:   "Leadership",
		},
		{
			Element: "h2",
			Value:   "Leadership Staff",
		},
		{
			Element: "h3",
			Value:   "Staff",
		},
		{
			Element: "div.field-sb-title",
			Value:   "Leadership Staff",
		},
	}
	for _, e := range staffBlocks {
		c.OnHTML(e, func(e *colly.HTMLElement) {
			for _, h := range staffHeaders {
				if e.ChildText(h.Element) == h.Value {
					e.ForEach("p", func(i int, el *colly.HTMLElement) {
						staffs = append(staffs, fmt.Sprintf("%s", el.Text))
					})
				}
			}
		})
	}
	if err != c.Visit(link) {
		return nil, err
	}
	return staffs, nil
}
