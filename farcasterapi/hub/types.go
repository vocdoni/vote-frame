package hub

type HubCastAddBody struct {
	Text              string   `json:"text"`
	ParentURL         string   `json:"parentUrl"`
	Mentions          []uint64 `json:"mentions"`
	MentionsPositions []uint64 `json:"mentionsPositions"`
}

type HubMessageData struct {
	Type        string          `json:"type"`
	From        uint64          `json:"fid"`
	Timestamp   uint64          `json:"timestamp"`
	CastAddBody *HubCastAddBody `json:"castAddBody,omitempty"`
}

type HubMessage struct {
	Data    *HubMessageData `json:"data"`
	HexHash string          `json:"hash"`
}

type HubMessageResponse struct {
	Messages []*HubMessage `json:"messages"`
}

type UsernameProofs struct {
	Username       string `json:"name"`
	CustodyAddress string `json:"owner"`
	FID            uint64 `json:"fid"`
	Type           string `json:"type"`
	Timestamp      uint64 `json:"timestamp"`
}

type UserdataResponse struct {
	Proofs []*UsernameProofs `json:"proofs"`
}

type Verification struct {
	Address string `json:"address"`
}

type VerificationData struct {
	Type         string        `json:"type"`
	Verification *Verification `json:"verificationAddEthAddressBody"`
	Signer       string        `json:"signer"`
}

type VerificationMessage struct {
	Data *VerificationData `json:"data"`
}

type VerificationsResponse struct {
	Messages []*VerificationMessage `json:"messages"`
}

type FidResponse struct {
	FID uint64 `json:"fid"`
}
