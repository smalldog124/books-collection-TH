package find_test

import (
	"crawler/find"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTableData_Input_HTML_Shold_Be_HTML_Teble(t *testing.T) {
	expected := `	<tr>
	<td>1</td>
	<td>978-616-213-808-9</td>
	<td>ศึกยุทธหัตถี (ปกแข็ง)</td>
	<td> สุภฤกษ์ บุญกอง</td>
	<td>บริษัท สกายบุ๊กส์ จำกัด</td>
	<td class="text-center"><a class="btn btn-default" href="Detail/205495"
			role="button"><span class="glyphicon glyphicon-info-sign"></span>ดูข้อมูล</a>
	</td>
</tr>
<tr>
	<td>2</td>
	<td>978-616-8269-16-9</td>
	<td>ยิหวายาใจ</td>
	<td> แอลลี่</td>
	<td>บริษัท มันดี กรุ๊ป จำกัด</td>
	<td class="text-center"><a class="btn btn-default" href="Detail/205493"
			role="button"><span class="glyphicon glyphicon-info-sign"></span>ดูข้อมูล</a>
	</td>
</tr>`
	body := []byte(`<tbody>
	<tr>
	<td>1</td>
	<td>978-616-213-808-9</td>
	<td>ศึกยุทธหัตถี (ปกแข็ง)</td>
	<td> สุภฤกษ์ บุญกอง</td>
	<td>บริษัท สกายบุ๊กส์ จำกัด</td>
	<td class="text-center"><a class="btn btn-default" href="Detail/205495"
			role="button"><span class="glyphicon glyphicon-info-sign"></span>ดูข้อมูล</a>
	</td>
</tr>
<tr>
	<td>2</td>
	<td>978-616-8269-16-9</td>
	<td>ยิหวายาใจ</td>
	<td> แอลลี่</td>
	<td>บริษัท มันดี กรุ๊ป จำกัด</td>
	<td class="text-center"><a class="btn btn-default" href="Detail/205493"
			role="button"><span class="glyphicon glyphicon-info-sign"></span>ดูข้อมูล</a>
	</td>
</tr>
</tbody>`)

	actual := find.GetTableData(body)

	assert.Equal(t, expected, actual)
}

func Test_BooksDetail_Input_HTML_Table_Shold_Be_Books_Detail(t *testing.T) {
	expected := [][]string{{"ศึกยุทธหัตถี (ปกแข็ง)", " สุภฤกษ์ บุญกอง", "บริษัท สกายบุ๊กส์ จำกัด"}, {"ยิหวายาใจ", " แอลลี่", "บริษัท มันดี กรุ๊ป จำกัด"}}

	table := `	<tr>
	<td>1</td>
	<td>978-616-213-808-9</td>
	<td>ศึกยุทธหัตถี (ปกแข็ง)</td>
	<td> สุภฤกษ์ บุญกอง</td>
	<td>บริษัท สกายบุ๊กส์ จำกัด</td>
	<td class="text-center"><a class="btn btn-default" href="Detail/205495"
			role="button"><span class="glyphicon glyphicon-info-sign"></span>ดูข้อมูล</a>
	</td>
</tr>
<tr>
	<td>2</td>
	<td>978-616-8269-16-9</td>
	<td>ยิหวายาใจ</td>
	<td> แอลลี่</td>
	<td>บริษัท มันดี กรุ๊ป จำกัด</td>
	<td class="text-center"><a class="btn btn-default" href="Detail/205493"
			role="button"><span class="glyphicon glyphicon-info-sign"></span>ดูข้อมูล</a>
	</td>
</tr>`

	actual := find.BooksDetail(table)

	assert.Equal(t, expected, actual)
}
