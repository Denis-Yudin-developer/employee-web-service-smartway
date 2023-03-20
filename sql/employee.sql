-- name: create-user
INSERT INTO employees
    (name, surname, phone, company_id)
    values ($1, $2, $3, $4)
RETURNING id;

--name: create-department
INSERT into department
    (employee_id, name, phone)
    values ($1, $2, $3)
RETURNING id;

--name: create-passport
INSERT into passport
(employee_id, passport_type, passport_number)
values ($1, $2, $3)
RETURNING id;

--name: add-department-to-user
UPDATE employees
    SET department_id = $1
WHERE id = $2;

--name: add-passport-to-user
UPDATE employees
    SET passport_id = $1
WHERE id = $2;

-- name: find-employees-by-company-id
SELECT e.id, e.name, e.surname, e.phone, e.company_id,
       d.id, d.name, d.phone,
       p.id, p.passport_type, p.passport_number
FROM
    employees as e
    INNER JOIN department d on d.id = e.department_id
    INNER JOIN passport p on e.id = p.employee_id
WHERE e.company_id = $1;

-- name: find-employees-by-department
SELECT e.id, e.name, e.surname, e.phone, e.company_id,
       d.id, d.name, d.phone,
       p.id, p.passport_type, p.passport_number
FROM
    employees as e
        INNER JOIN department d on d.id = e.department_id
        INNER JOIN passport p on e.id = p.employee_id
WHERE d.name = $1;

--name: is-employee-present
SELECT 1
    FROM employees
WHERE id = $1;

-- name: update-employee
UPDATE employees SET
    name = COALESCE($1, name),
    surname = COALESCE($2, surname),
    phone = COALESCE($3, phone),
    company_id = COALESCE($4, company_id)
WHERE id = $5;

--name: update-department-by-user-id
UPDATE department SET
    name = COALESCE($1, name),
    phone = COALESCE($2, phone)
WHERE employee_id = $3;

--name: update-passport-by-user-id
UPDATE passport SET
    passport_type = COALESCE($1, passport_type),
    passport_number = COALESCE($2, passport_number)
WHERE employee_id = $3;

--name: delete-user
DELETE FROM employees
    WHERE id = $1