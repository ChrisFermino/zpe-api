package request_handler

type Error struct {
    Message string `json:"message" binding:"required"`
    Status  int    `json:"status" binding:"required"`
    Error   string `json:"error" binding:"required"`
}

type Request struct {
    Data  any   `json:"data"`
    Error Error `json:"error"`
}

type RequestSuccess struct {
    Data any `json:"data"`
}
