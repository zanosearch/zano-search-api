package search

import (
	"github.com/bytedance/sonic"
	"github.com/zanosearch/zano-search-api/internal/base64"
	"github.com/zanosearch/zano-search-api/internal/zano"
	"strings"
)

type Rank struct {
	Score int   `json:"score"`
	Offer Offer `json:"offer"`
}

type Offer struct {
	Credentials struct {
		Alias        string `json:"alias"`
		Argon2IdHash string `json:"argon2id_hash"`
	} `json:"credentials"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Category    string `json:"category"`
	Url         struct {
		IsDarkweb   bool   `json:"is_darkweb"`
		Title       string `json:"title"`
		Url         string `json:"url"`
		Description string `json:"description"`
	} `json:"url"`
	PaymentMethods []struct {
		Crypto   bool   `json:"crypto"`
		Currency string `json:"currency"`
		MoreInfo string `json:"more_info"`
	} `json:"payment_methods"`
	Socials []struct {
		Platform string `json:"platform"`
		Url      string `json:"url"`
	} `json:"socials"`
	Contacts []struct {
		Type    string `json:"type"`
		Contact string `json:"contact"`
	} `json:"contacts"`
	Wares []struct {
		Title       string   `json:"title"`
		Id          string   `json:"id"`
		Description string   `json:"description"`
		Price       string   `json:"price"`
		Info        string   `json:"info"`
		Images      []string `json:"images"`
	} `json:"wares,omitempty"`
	Fulfillment struct {
		IsDigital             bool   `json:"is_digital"`
		ShipsFrom             string `json:"ships_from"`
		ShipsTo               string `json:"ships_to"`
		FulfillmentPledge     bool   `json:"fulfillment_pledge"`
		FulfillmentPledgeDays int    `json:"fulfillment_pledge_days"`
	} `json:"fulfillment"`
	EscrowAccepted struct {
		BazaarEscrow bool `json:"bazaar_escrow"`
		Contract     bool `json:"contract"`
	} `json:"escrow_accepted"`
	Type              string `json:"type"`
	FrontpageBoosting struct {
		CostPerClick int64 `json:"cost_per_click"`
	} `json:"frontpage_boosting"`
	BazaarInstanceId string `json:"bazaar_instance_id"`
	BazaarUuid       string `json:"bazaar_uuid"`
}

func OfferSearch(instanceId string, tokens []string, offers []zano.MarketplaceOffers) []Rank {

	var scoredOffers []Rank

	for _, offer := range offers {
		// 1. decode base64
		ds, err := base64.DecodeBase64(offer.Com)
		if err == nil {
			data := Offer{}
			if err = sonic.Unmarshal([]byte(ds), &data); err == nil {
				// 2. TODO: check title, description and wares for token matches and rank
				if data.BazaarInstanceId == instanceId {
					var score int
					for _, token := range tokens {
						if strings.Contains(strings.ToLower(data.Description), token) {
							score++
						}

						if strings.Contains(strings.ToLower(data.Title), token) {
							score++
						}

						for _, ware := range data.Wares {
							if strings.Contains(strings.ToLower(ware.Title), token) {
								score++
							}

							if strings.Contains(strings.ToLower(ware.Description), token) {
								score++
							}

							if strings.Contains(strings.ToLower(ware.Info), token) {
								score++
							}
						}
					}
					if score > 0 {
						scoredOffers = append(scoredOffers, Rank{
							Score: score,
							Offer: data,
						})
					}
				}
			}
		}
		//fmt.Println(offer)
	}
	return scoredOffers
}
