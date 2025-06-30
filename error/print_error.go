package customerror

import "fmt"

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
