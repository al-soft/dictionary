create schema lang;

CREATE TABLE lang.tubalar
(
    id serial primary key,
    word varchar(120) NOT NULL default('')
);


CREATE TABLE lang.russian
(
    id serial primary key,
    word varchar(120) NOT NULL default('')
);

CREATE TABLE lang.translation
(
    id serial primary key,
    tubalar_id int not null,
	russian_id int not null,
    CONSTRAINT fk_tubalar
      FOREIGN KEY(tubalar_id) 
	  REFERENCES lang.tubalar(id),
    CONSTRAINT fk_russian
      FOREIGN KEY(russian_id) 
	  REFERENCES lang.russian(id)
);

select 
	t.word,
	r.word
from lang.tubalar t, 
	 lang.russian r,
	 lang.translation tr
where
	tr.tubalar_id=t.id
	and tr.russian_id=r.id
;	