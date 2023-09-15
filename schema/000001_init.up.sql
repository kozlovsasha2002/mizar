CREATE TABLE events (
    event_id BIGSERIAL PRIMARY KEY ,
    link TEXT NOT NULL ,
    deadline_date DATE NOT NULL ,
    deadline_time TIME NOT NULL ,
    is_completed BOOLEAN DEFAULT FALSE ,
    description TEXT NOT NULL
);

INSERT INTO events (link, deadline_date, deadline_time, description) VALUES
('link1', '2023-08-11', '01:59:59', 'Это текстовое описание первого события.'),
('link2', '2021-09-11', '11:59:59', 'Это текстовое описание второго события.'),
('link3', '2020-10-11', '13:59:59', 'Это текстовое описание третьего события.'),
('link4', '2023-11-11', '23:59:59', 'Это текстовое описание четвертого события.'),
('link5', '2023-12-11', '23:59:59', 'Это текстовое описание пятого события.'),
('link6', '2023-12-11', '23:59:59', 'Это текстовое описание шестого события.');
