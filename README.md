# Book Collection Project

# Function And Parameter Naming Conventions

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
book
```

## File Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
api.go
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
books
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_BookScanHandler_Input_ISBN_978_616_18_2996_4_Shold_Be_Book_ID_1
```

## Variable Name
- ชื่อตัวแปรเป็น **camelCase** เช่น
```
request, booksAPI, bookReview, bookShelf
```

- ชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
books
```

- ชื่อตัวแปร struct ให้ตั้งชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ **camelCase** เช่น
```
BookResponse, BookRequest
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวพิมพ์เล็กก่อน เว้นแต่เมื่อมีการใช้ข้าม package ถึงจะใช้ Capital Case เช่น
```
Hour, Minute, url
```

## Struct ที่มี slice อยู่ด้านใน 
- ให้ทำการสร้าง struct ที่มี slice เป็น attribute ให้สร้างเป็น slice เปล่า
```
type PlanDetail struct {
    PlanList []Plan
}

func NewPlanDetail() PlanDetail {
    return PlanDetail{
        PlanList: make([]Plan, 0)
    }
}

```

## รูปแบบข้อมูล json 

ใช้เป็น **snakeCase** เช่น
```
book_id, edition_note
```

# Error Message Pattern
- ใช้รูปแบบ verb + noun + "error" เช่น
```
Get subproducts error
```

## ข้อตกลง Commit Message ร่วมกัน
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน
