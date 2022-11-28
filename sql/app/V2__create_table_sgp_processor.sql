
create table ${flyway:database}.his_historical
(
    his_id_patient int not null,
    his_id_file_patient int not null,
    his_first_name varchar(45) not null,
    his_second_name varchar(45) null,
    his_first_last_name varchar(45) not null,
    his_second_last_name varchar(45) null,
    his_admission_date date not null,
    his_high_date date null,
    his_low_date date null,
    his_creation_date timestamp null,
    his_created_by varchar(45) null,
    his_modification_date timestamp null,
    his_modified_by varchar(45) null,
    primary key(his_id_patient,his_id_file_patient, his_admission_date)
);


create trigger ${flyway:database}.his_historical_insert_aud
    BEFORE INSERT ON ${flyway:database}.his_historical
    FOR EACH ROW
    set NEW.his_created_by=USER(),
		NEW.his_creation_date=now();

create trigger ${flyway:database}.his_historical_update_aud
    BEFORE UPDATE ON ${flyway:database}.his_historical
    FOR EACH ROW
    set NEW.his_modification_date=now(),
		NEW.his_modified_by=USER();