package ads

import "sort"

type (
	User struct {
		Country string
		Browser string
	}

	Campaing struct {
		ClickUrl  string
		Price     float64
		Targeting Targeting
	}

	Targeting struct {
		Browser string
		Country string
	}

	filterFunc func(in []*Campaing, u *User) (out []*Campaing)
)

var (
	filters = []filterFunc{
		filterByBrowser,
		filterByCountry,
	}
)

func MakeAuction(in []*Campaing, u *User) (winner *Campaing) {
	campaings := make([]*Campaing, len(in))
	copy(campaings, in) // копируем компании

	for _, f := range filters {
		campaings = f(campaings, u)
	}
	if len(campaings) == 0 { // если кампаний не осталось, значит у нас нет победителя, нечего сортировать
		return nil
	}

	sort.Slice(campaings, func(i, j int) bool {
		return campaings[i].Price < campaings[j].Price
	})
	return campaings[0]
}

func filterByBrowser(in []*Campaing, u *User) (out []*Campaing) {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Browser) == 0 {
			continue // если браузер не указан,значит компания готова быть на любых браузерах
		}

		if in[i].Targeting.Browser == u.Browser {
			continue // если браузер пользователя подходит, тогда идем дальше
		}
		in[i] = in[0]
		in = in[1:]
	}
	return in
}

func filterByCountry(in []*Campaing, u *User) (out []*Campaing) {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Country) == 0 {
			continue // если страна не указана,значит компания готова быть на любых странах
		}

		if in[i].Targeting.Country == u.Country {
			continue // если страна пользователя подходит, тогда идем дальше
		}
		in[i] = in[0]
		in = in[1:]
	}
	return in
}

func GetCampaings() []*Campaing {
	return []*Campaing{
		{
			Price: 1,
			Targeting: Targeting{
				Country: "RU",
				Browser: "Chrome",
			},
			ClickUrl: "https://google.ua",
		},
		{
			Price: 1,
			Targeting: Targeting{
				Country: "DE",
				Browser: "Chrome",
			},
			ClickUrl: "https://google.de",
		},
		{
			Price: 1,
			Targeting: Targeting{
				Browser: "Firefox",
			},
			ClickUrl: "https://duckduckgo.com",
		},
	}
}
