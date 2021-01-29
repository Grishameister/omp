package domain

type Stat struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
	Rating  int    `json:"rating"`
}

type Result struct {
	MaxPrice         int
	MaxPriceProduct  string
	MaxRating        int
	MaxRatingProduct string
}

func (r *Result) Compare(st *Stat) {
	if r.MaxRating <= st.Rating {
		r.MaxRating = st.Rating
		r.MaxRatingProduct = st.Product
	}

	if r.MaxPrice <= st.Price {
		r.MaxPrice = st.Price
		r.MaxPriceProduct = st.Product
	}
}
