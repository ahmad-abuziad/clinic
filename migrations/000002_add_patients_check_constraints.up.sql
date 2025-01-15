ALTER TABLE patients ADD CONSTRAINT patients_first_name_check CHECK (octet_length(first_name) > 0 AND octet_length(first_name) <= 50);
ALTER TABLE patients ADD CONSTRAINT patients_last_name_check CHECK (octet_length(last_name) > 0 AND octet_length(last_name) <= 50);
ALTER TABLE patients ADD CONSTRAINT patients_gender_check CHECK (gender IN ('M', 'F'));
ALTER TABLE patients ADD CONSTRAINT patients_date_of_birth_check CHECK (date_of_birth >= '1900-01-01' AND date_of_birth <= CURRENT_DATE);
