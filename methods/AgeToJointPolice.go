// TÍNH TUỔI ĐỂ TRỞ THÀNH ÔNG CAN
package methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Person struct {
	name      string
	birthYear int
	hometown  string
}

type Student struct {
	Person
	school string
}

type Police struct {
	Person
	station string
}

func CheckAgeToJointPolice() {
	color.Yellow("CHƯƠNG TRÌNH TÍNH SỐ NĂM ĐỂ TRỞ THÀNH ÔNG CAN")
	// Nhập dữ liệu từ bàn phím
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập họ và tên: ")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)

	// Nhập năm sinh
	var birthYearInput int
	for {
		fmt.Print("Nhập năm sinh: ")
		yearStr, _ := reader.ReadString('\n')
		yearStr = strings.TrimSpace(yearStr)

		year, err := strconv.Atoi(yearStr)
		if err != nil || year <= 0 {
			fmt.Println("Năm sinh phải là số nguyên dương, vui lòng nhập lại!")
		} else {
			birthYearInput = year
			break
		}
	}

	// Nhập quê quán
	var hometownInput string
	for {
		fmt.Print("Nhập quê quán: ")
		hometownInput, _ = reader.ReadString('\n')
		hometownInput = strings.TrimSpace(hometownInput)

		if len(hometownInput) > 0 {
			break
		}
		fmt.Println("Quê quán không được để trống, vui lòng nhập lại!")
	}

	// Nhập tên trường
	var schoolInput string
	for {
		fmt.Print("Nhập tên trường: ")
		schoolInput, _ = reader.ReadString('\n')
		schoolInput = strings.TrimSpace(schoolInput)

		if len(schoolInput) > 0 {
			break
		}
		fmt.Println("Tên trường không được để trống, vui lòng nhập lại!")
	}

	// Nhập đơn vị công tác
	var stationInput string
	for {
		fmt.Print("Nhập đơn vị công tác: ")
		stationInput, _ = reader.ReadString('\n')
		stationInput = strings.TrimSpace(stationInput)

		if len(stationInput) > 0 {
			break
		}
		fmt.Println("Đơn vị công tác không được để trống, vui lòng nhập lại!")
	}

	// Tính tuổi của sinh viên
	now := time.Now().Year()
	age := now - birthYearInput

	// Kiểm tra tuổi của sinh viên
	if age < 25 {
		color.Yellow("Tuổi yêu cầu: 25 tuổi trở lên")
		color.Yellow("Tuổi của %v: %v", nameInput, age)
		color.Yellow("Bạn cần thêm: %v năm nữa để có thể làm ông can", 25-age)
		
	} else {
		// Kiểm tra nếu là công an
		if age >= 25 {
			p := Police{
				Person: Person{
					name:      nameInput,
					birthYear: birthYearInput,
					hometown:  hometownInput,
				},
				station: stationInput,
			}
			color.Green("%v đang là công an, công tác tại %v\n", p.name, p.station)
			color.Green("Thâm niên công tác: %v năm", age-25)
		} else {
			// In ra các thông tin đã nhập
			s := Student{
				Person: Person{
					name:      nameInput,
					birthYear: birthYearInput,
					hometown:  hometownInput,
				},
				school: schoolInput,
			}
			fmt.Println("Tên:", s.name)
			fmt.Println("Năm sinh:", s.birthYear)
			fmt.Println("Quê quán:", s.hometown)
			fmt.Println("Tên trường:", s.school)
		}
	}
}
