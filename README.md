# go-custom-error
create error with you struct and business logic error code

## สร้าง Error Code และ Custom Error Struct
ใน error/error.go มีการประกาศ const ของ ErrorCode เพื่อเป็นการกำหนด error code ของ error ต่างๆที่เราต้องการ

```go
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
```


สร้าง custom error ด้วย การประการ struct CustomError และเรียกใช้ ErrorCode ที่เราสร้างไว้ก่อนหน้า เมื่อต้องการสร้าง error ให้เรียกใช้ผ่าน Func NewCustomError

```go
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
```


## ใช้งาน Custom Error Struct
สร้าง func สำหรับ return error ใน main.go ซึ่ง func return type เป็น error แต่เราสามารถใช้ CustomError implements interface error ได้

```go
// function สำหรับสร้าง error ต่างๆ ตามค่า ErrorCode ที่ระบุ
func ReturnErrorNotFound() error {
	return customerror.NewCustomError("Not Found", customerror.ErrNotFound)
}
```

ใน func main() ทำการ print ค่า error message ที่ได้จาก func ReturnErrorNotFound() ตามคำที่ business ต้องการด้วย ErrorCode

```go
fmt.Println("Not Found: ", customerror.ReturnMessageFromErrorCode(err))
```
func ReturnMessageFromErrorCode(err) จาก error/print_error.go

```go
// สร้าง function PrintError() เพื่อพิมพ์ข้อความของ error ตามค่า ErrorCode ที่ระบุ
func ReturnMessageFromErrorCode(err error) string {
	if err == nil {
		fmt.Println("No error")
		return "No error"
	}

	// ตรวจสอบว่า err เป็นตัวแปรที่สร้างจาก type CustomError หรือไม่
	if customErr, ok := err.(*CustomError); ok {
		// ถ้า err เป็นตัวแปรที่สร้างจาก type CustomError ก็จะพิมพ์ข้อความของ error ตามค่า ErrorCode ที่เราต้องการได้
		switch customErr.Code {
		case ErrNotFound:
			return "ไม่พบข้อมูลในระบบ"
		case ErrInvalidInput:
			return "ข้อมูลที่ระบุไม่ถูกต้อง"
		case ErrInternal:
			return "เกิดข้อผิดพลาดภายในระบบ"
		case ErrUnauthorized:
			return "คุณไม่มีสิทธิ์เข้าถึงข้อมูลนี้ กรุณาตรวจสอบสิทธิ์การเข้าถึง"
		case ErrForbidden:
			return "คุณไม่มีสิทธิ์เข้าถึงข้อมูลนี้"
		case ErrBadRequest:
			return "คำขอของคุณไม่ถูกต้อง กรุณาตรวจสอบคำขอของคุณ"
		case ErrConflict:
			return "ข้อมูลที่ระบุมีปัญหา กรุณาตรวจสอบข้อมูลอีกครั้ง"
		default:
			return "เกิดข้อผิดพลาดภายในระบบ"
		}
	} else {
		// ถ้า err ไม่ใช่ตัวแปรที่สร้างจาก type CustomError ก็จะพิมพ์ข้อความของ error ตามที่ระบุ
		return err.Error()
	}
}
```
