package domain

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/midnight-trigger/todo/api/definition"
	"github.com/midnight-trigger/todo/api/error_handling"

	"github.com/midnight-trigger/todo/infra/mysql"
	"github.com/midnight-trigger/todo/infra/mysql/mock_mysql"

	"github.com/stretchr/testify/assert"
)

func TestDeleteTodo_正常系(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	param := new(definition.DeleteTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	todo := new(mysql.Todos)
	todo.Id = param.TodoId
	todo.UserId = userId
	todo.Title = "Test Title"
	todo.Body = "Test Body"
	todo.Status = "todo"

	response := new(definition.DeleteTodoResponse)
	response.Id = param.TodoId

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	gomock.InOrder(
		mockedTodos.EXPECT().FindById(param.TodoId).Return(*todo, nil),
		mockedTodos.EXPECT().Delete(todo).Return(nil),
	)

	domain := new(Todo)
	domain.MTodos = mockedTodos

	result := domain.DeleteTodo(param, userId)
	assert.Equal(t, 200, result.Code)
	assert.Equal(t, response, result.Data.(*definition.DeleteTodoResponse))
}

func TestDeleteTodo_パスパラメータのTodoIdに紐づくレコードがDB上に存在しない場合エラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	param := new(definition.DeleteTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	mockedTodos.EXPECT().FindById(param.TodoId).Return(mysql.Todos{}, gorm.ErrRecordNotFound)

	expect := new(error_handling.ErrorHandling)
	expect.Code = 404
	expect.ErrMessage = "対象Todoが見つかりません"
	expect.ErrStack = errors.New("")

	domain := new(Todo)
	domain.MTodos = mockedTodos

	result := domain.DeleteTodo(param, userId)
	assert.Equal(t, *expect, result.ErrorHandling)
}

func TestDeleteTodo_ログインユーザにDB更新権限が無い場合エラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	param := new(definition.DeleteTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	todo := new(mysql.Todos)
	todo.Id = param.TodoId
	todo.UserId = "1802f638-53f2-4848-9859-a54a2bdf5160"
	todo.Title = "Title"
	todo.Body = "Body"
	todo.Status = "todo"

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	mockedTodos.EXPECT().FindById(param.TodoId).Return(*todo, nil)

	expect := new(error_handling.ErrorHandling)
	expect.Code = 400
	expect.ErrMessage = "必要な権限がありません"
	expect.ErrStack = errors.New("")

	domain := new(Todo)
	domain.MTodos = mockedTodos

	result := domain.DeleteTodo(param, userId)
	assert.Equal(t, *expect, result.ErrorHandling)
}

func TestDeleteTodo_サーバで問題が起きた場合エラーを返すか検証(t *testing.T) {
	s := GetNewTodoService()
	ctrl := s.TestInit(t)

	// リクエスト定義
	param := new(definition.DeleteTodoParam)
	param.TodoId = 1
	userId := "1802f638-53f2-4848-9859-a54a2bdf5163"

	todo := new(mysql.Todos)
	todo.Id = param.TodoId
	todo.UserId = userId
	todo.Title = "Test Title"
	todo.Body = "Test Body"
	todo.Status = "todo"

	expect := new(error_handling.ErrorHandling)
	expect.Code = 500
	expect.ErrMessage = "サーバーエラー"
	expect.ErrStack = errors.New("")
	expect.RawErrMessage = "not implemented"

	mockedTodos := mock_mysql.NewMockITodos(ctrl)
	for i := 0; i < 2; i++ {
		switch i {
		case 0:
			mockedTodos.EXPECT().FindById(param.TodoId).Return(*todo, errors.New("not implemented")).Times(2)
		case 1:
			gomock.InOrder(
				mockedTodos.EXPECT().FindById(param.TodoId).Return(*todo, nil),
				mockedTodos.EXPECT().Delete(todo).Return(errors.New("not implemented")).Times(1),
			)
		}
		domain := new(Todo)
		domain.MTodos = mockedTodos

		result := domain.DeleteTodo(param, userId)
		assert.Equal(t, *expect, result.ErrorHandling)
	}
}
