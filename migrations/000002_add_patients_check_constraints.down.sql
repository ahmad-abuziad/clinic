ALTER TABLE patients DROP CONSTRAINT IF EXISTS patients_first_name_check;
ALTER TABLE patients DROP CONSTRAINT IF EXISTS patients_last_name_check;
ALTER TABLE patients DROP CONSTRAINT IF EXISTS patients_gender_check;
ALTER TABLE patients DROP CONSTRAINT IF EXISTS patients_date_of_birth_check;