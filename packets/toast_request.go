package packets

const ToastRequestID uint32 = 186

type ToastRequest struct {
	Title   string
	Content string
}
