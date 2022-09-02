package schemas


 var Place = `CREATE TABLE details(
	Name text,
	Job text,
	Salary integer);`

var AlterAge = `ALTER TABLE details ADD COLUMN Age varchar;`	

var Users = `CREATE TABLE users(
	ID serial not null primary key,
	Name text,
	Password text);`

var DropUserTable = `DROP TABLE users;`
