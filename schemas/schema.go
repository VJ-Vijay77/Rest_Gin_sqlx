package schemas


 var Place = `CREATE TABLE details(
	Name text,
	Job text,
	Salary integer);`

var AlterAge = `ALTER TABLE details ADD COLUMN Age varchar;`	