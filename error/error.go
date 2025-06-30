package customerror

// ประกาศ type ErrorCode เป็นตัวแปรที่จะใช้เก็บค่าของ ErrorCode
type ErrorCode int

// การกำหนดความหมายของ ErrorCode ต่างๆ ตามที่เราต้องการและยังเป็นการกำหนดให้ใช้ไปในทางเดียวกัน
// การใช้ iota จะทำให้ค่า ErrorCode ตัวถัดไปเพิ่มขึ้นจากตัวก่อนหน้า 1 ค่า ใน code เริ่มต้นจาก 1
const (
	ErrNotFound ErrorCode = iota + 1
	ErrInvalidInput
	ErrInternal
	ErrUnauthorized
	ErrForbidden
	ErrBadRequest
	ErrConflict
)

// ประกาศ type CustomError เป็นตัวแปรที่จะใช้เก็บค่าของ CustomError
type CustomError struct {
	Message string
	Code    ErrorCode
}

// สร้าง method Error() เพื่อให้สามารถ implement interface error ได้
func (e *CustomError) Error() string {
	return e.Message
}

// เรียกใช้ function NewCustomError() เพื่อสร้างตัวแปร CustomError ใหม่
func NewCustomError(message string, code ErrorCode) *CustomError {
	return &CustomError{Message: message, Code: code}
}
