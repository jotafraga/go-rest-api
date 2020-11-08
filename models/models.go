package models

//Product é a estrutura de Produtos responsável por agrupar dados referente ao produto.
type Product struct {
	ProductID   int     `json:"productId,omitempty" db:"product_id"`
	Name        string  `json:"name,omitempty" db:"name"`
	Description string  `json:"description,omitempty" db:"description"`
	Price       float64 `json:"price,omitempty" db:"price"`
	Amount      int     `json:"amount,omitempty" db:"amount"`
}

//Response é a estrutura responsável por agrupar dados referente a resposta do servidor.
type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// Product response payload
// swagger:response productRes
type swaggProductRes struct {
	// in:body
	Body Product
}

// Success Response
// swagger:response okResp
type swaggRespOk struct {
	// in:body
	Body struct {
		// HTTP status code 200 - OK
		Code int `json:"code"`
	}
}

// Error Bad Request
// swagger:response badReq
type swaggReqBadRequest struct {
	// in:body
	Body struct {
		// HTTP status code 400 -  Bad Request
		Code int `json:"code"`
	}
}

// Error Not Found
// swagger:response notFoundReq
type swaggReqNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 -  Not Found
		Code int `json:"code"`
	}
}
