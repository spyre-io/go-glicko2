package glicko

const (
	RATING_SCALE_PARAMETER = 173.7178
	RATING_BASE_R          = 1500
	RATING_BASE_RD         = 350
	RATING_BASE_SIGMA      = 0.06
)

type Rating struct {
	R     float64 `json:"r"`
	Mu    float64 `json:"mu"`
	Rd    float64 `json:"rd"`
	Phi   float64 `json:"phi"`
	Sigma float64 `json:"sigma"`
}

func (rating Rating) ConfidenceInterval() (float64, float64) {
	return rating.R - 2*rating.Rd, rating.R + 2*rating.Rd
}

func (rating *Rating) Update(mu float64, phi float64, sigma float64) {
	setMu(rating, mu)
	setPhi(rating, phi)
	setSigma(rating, sigma)
}

func (rating *Rating) Touch() {
	setPhi(rating, phiA(rating.Phi, rating.Sigma))
}

func setR(rating *Rating, r float64) {
	rating.R = r
	rating.Mu = (rating.R - RATING_BASE_R) / RATING_SCALE_PARAMETER
}

func setMu(rating *Rating, mu float64) {
	rating.Mu = mu
	rating.R = rating.Mu*RATING_SCALE_PARAMETER + RATING_BASE_R
}

func setRd(rating *Rating, rd float64) {
	rating.Rd = rd
	rating.Phi = rating.Rd / RATING_SCALE_PARAMETER
}

func setPhi(rating *Rating, phi float64) {
	rating.Phi = phi
	rating.Rd = rating.Phi * RATING_SCALE_PARAMETER
}

func setSigma(rating *Rating, sigma float64) {
	rating.Sigma = sigma
}

func NewRating(r float64, rd float64, sigma float64) *Rating {
	rating := &Rating{}

	setR(rating, r)
	setRd(rating, rd)
	setSigma(rating, sigma)

	return rating
}

func NewDefaultRating() *Rating {
	return NewRating(RATING_BASE_R, RATING_BASE_RD, RATING_BASE_SIGMA)
}
