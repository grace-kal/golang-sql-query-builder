package querybuilder

import "testing"

func TestCreateTable(t *testing.T) {
	//Arrange
	qb := NewQueryBuilder("users")
	tableStructure := map[string]string{
		"id":       "INT PRIMARY KEY",
		"username": "VARCHAR(255)",
		"email":    "VARCHAR(255)",
	}
	expectedQuery := "CREATE TABLE users (id INT PRIMARY KEY, username VARCHAR(255), email VARCHAR(255))"
	//Act
	actualQuery := qb.Create(tableStructure)
	//Assert
	if actualQuery != expectedQuery {
		t.Errorf("Expected query: %s, but got: %s", expectedQuery, actualQuery)
	}
}

func TestSelect(t *testing.T) {
	//Arrange
	qb := NewQueryBuilder("users")
	fields := []string{"id", "username"}
	expectedQuery := "SELECT id, username FROM users"
	//Act
	actualQuery := qb.Select(fields)
	//Assert
	if actualQuery != expectedQuery {
		t.Errorf("Expected query: %s, but got: %s", expectedQuery, actualQuery)
	}
}

func TestInsert(t *testing.T) {
	qb := NewQueryBuilder("users")
	data := map[string]interface{}{
		"username": "gg",
		"email":    "kalinina.grace@gmail.com",
	}
	expectedQuery := "INSERT INTO users (username, email) VALUES ('gg', 'kalinina.grace@gmail.com')"
	actualQuery := qb.Insert(data)
	if actualQuery != expectedQuery {
		t.Errorf("Expected query: %s, but got: %s", expectedQuery, actualQuery)
	}
}

func TestUpdate(t *testing.T) {
	qb := NewQueryBuilder("users")
	data := map[string]interface{}{
		"username": "gg",
		"email":    "kalinina.grace@gmail.com",
	}
	expectedQuery := "UPDATE users SET username = 'gg', email = 'kalinina.grace@gmail.com'"
	actualQuery := qb.Update(data)
	if actualQuery != expectedQuery {
		t.Errorf("Expected query: %s, but got: %s", expectedQuery, actualQuery)
	}
}
