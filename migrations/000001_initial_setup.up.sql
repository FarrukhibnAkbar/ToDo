CREATE TABLE IF NOT EXISTS tasks (
	id uuid PRIMARY KEY,
	assignee VARCHAR(64),
	title VARCHAR(64),
	summary VARCHAR(512),
	deadline TIMESTAMP,
	status VARCHAR(50)
);
