CREATE TABLE tasks (
    id SHORTKEY PRIMARY KEY,
    creator_id VARCHAR (50) NOT NULL, 
    created_at TIMESTAMP NOT NULL,
    title VARCHAR (150) NOT NULL,
    description TEXT,
    status INT NOT NULL
);

CREATE TRIGGER trigger_tasks_genid BEFORE INSERT ON tasks FOR EACH ROW EXECUTE PROCEDURE shortkey_generate();
