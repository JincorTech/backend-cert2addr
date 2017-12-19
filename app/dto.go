package app

type ServerResponse struct {
	Status int32       `json:"status"`
	Data   interface{} `json:"data"`
}

type CertificateAddressRequest struct {
	Pem string `json:"pem"`
}

type CertificateAddressResponse struct {
	Address string `json:"address"`
}
