CREATE TABLE jobseekers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100) UNIQUE,
    password VARCHAR(255),
    gender VARCHAR(10) CHECK (gender IN ('male', 'female')), 
    phone VARCHAR(20),
    date_of_birth DATE, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    jobseeker_id INT REFERENCES jobseekers(id) ON DELETE CASCADE,
    summary TEXT,
    city VARCHAR(100),
    country VARCHAR(100),
    education TEXT,
    experience TEXT
);

CREATE TABLE follows (
    jobseeker_id INT REFERENCES jobseekers(id) ON DELETE CASCADE,
    employer_id INT REFERENCES employers(id) ON DELETE CASCADE
);

ALTER TABLE applications
ADD CONSTRAINT fk_jobseeker
FOREIGN KEY (jobseeker_id) REFERENCES jobseekers(id) ON DELETE CASCADE;

INSERT INTO jobseekers (first_name, last_name, email, password, gender, phone, date_of_birth)
VALUES 
('John', 'Doe', 'john.doe@example.com', 'securepassword123', 'male', '+1234567890', '1990-05-15'),
('Jane', 'Smith', 'jane.smith@example.com', 'securepassword456', 'female', '+0987654321', '1988-08-22'),
('Alice', 'Johnson', 'alice.johnson@example.com', 'password789', 'female', '+1122334455', '1992-02-10'),
('Bob', 'Williams', 'bob.williams@example.com', 'passwordabc', 'male', '+6677889900', '1985-12-05'),
('Emily', 'Brown', 'emily.brown@example.com', 'passworddef', 'female', '+4433221100', '1995-09-25');

INSERT INTO profiles (jobseeker_id, summary, city, country, education, experience)
VALUES 
(1, 'Experienced software engineer with 5+ years in full-stack development.', 'New York', 'USA', 'B.Sc. in Computer Science', 'Worked at TechCorp for 3 years as a backend developer'),
(2, 'Marketing expert with a focus on digital marketing and social media strategies.', 'Los Angeles', 'USA', 'M.Sc. in Marketing', 'Digital marketing manager at MarketCorp for 4 years'),
(3, 'Human resources specialist with experience in employee relations and recruitment.', 'Chicago', 'USA', 'B.A. in Human Resources', 'HR manager at RecruitCorp for 5 years'),
(4, 'Data scientist with expertise in machine learning and data analytics.', 'San Francisco', 'USA', 'Ph.D. in Data Science', 'Data scientist at DataSolutions for 6 years'),
(5, 'Graphic designer skilled in Adobe Creative Suite and web design.', 'Miami', 'USA', 'B.F.A. in Graphic Design', 'Freelance graphic designer for 4 years');

INSERT INTO follows (jobseeker_id, employer_id) VALUES
(1, 1),
(1, 3),
(2, 2),
(3, 1),
(4, 4),
(5, 5);