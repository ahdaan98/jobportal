CREATE TABLE employers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255),
    phone VARCHAR(15),
    address VARCHAR(255),
    country VARCHAR(100),
    website VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    employer_id INT REFERENCES employers(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    employment_type VARCHAR(50),
    description TEXT NOT NULL,
    location VARCHAR(255),
    salary NUMERIC(10, 2),
    experience_level VARCHAR(50),
    posted_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    job_id INT REFERENCES jobs(id) ON DELETE CASCADE,
    jobseeker_id INT,
    resume TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO employers (name, email, password, phone, address, country, website) VALUES
('Tech Solutions', 'contact@techsolutions.com', 'techpass123', '+1234567890', '123 Tech Lane', 'USA', 'www.techsolutions.com'),
('HealthPlus Corp', 'info@healthplus.com', 'healthpass123', '+1987654321', '456 Health St', 'USA', 'www.healthplus.com'),
('EcoFriendly Inc', 'support@ecofriendly.com', 'ecopass123', '+1122334455', '789 Green Blvd', 'Canada', 'www.ecofriendly.com'),
('FinanceGuru', 'hello@financeguru.com', 'financepass123', '+1223344556', '321 Money St', 'UK', 'www.financeguru.com'),
('TravelNest', 'booking@travelnest.com', 'travelpass123', '+1334455667', '654 Journey Ave', 'Australia', 'www.travelnest.com');

INSERT INTO jobs (employer_id, title, employment_type, description, location, salary, experience_level) VALUES
(1, 'Software Engineer', 'Full-time', 'Develop and maintain software applications.', 'New York, NY', 95000.00, 'Mid-level'),
(1, 'Project Manager', 'Full-time', 'Oversee project timelines and deliverables.', 'New York, NY', 110000.00, 'Senior'),
(2, 'Registered Nurse', 'Part-time', 'Provide patient care and support.', 'Los Angeles, CA', 65000.00, 'Entry-level'),
(3, 'Environmental Scientist', 'Contract', 'Conduct research and analyze environmental data.', 'Toronto, ON', 75000.00, 'Mid-level'),
(4, 'Financial Analyst', 'Full-time', 'Analyze financial data and prepare reports.', 'London, UK', 80000.00, 'Mid-level'),
(5, 'Travel Consultant', 'Remote', 'Assist clients in planning their travel itineraries.', 'Sydney, Australia', 60000.00, 'Entry-level'),
(1, 'UX/UI Designer', 'Full-time', 'Design user-friendly interfaces for web and mobile applications.', 'San Francisco, CA', 90000.00, 'Mid-level'),
(2, 'Health Coach', 'Part-time', 'Guide clients in making healthier lifestyle choices.', 'Miami, FL', 50000.00, 'Entry-level'),
(3, 'Sustainability Consultant', 'Contract', 'Advise businesses on sustainability practices and compliance.', 'Vancouver, BC', 80000.00, 'Senior'),
(4, 'Accountant', 'Full-time', 'Prepare and analyze financial statements and reports.', 'Manchester, UK', 70000.00, 'Mid-level'),
(5, 'Digital Marketing Specialist', 'Remote', 'Develop and execute digital marketing strategies.', 'Remote', 65000.00, 'Mid-level'),
(1, 'DevOps Engineer', 'Full-time', 'Manage and automate infrastructure and deployment.', 'Austin, TX', 105000.00, 'Senior'),
(2, 'Pharmacist', 'Full-time', 'Dispense medications and provide health consultations.', 'Chicago, IL', 120000.00, 'Mid-level'),
(3, 'Wildlife Biologist', 'Contract', 'Research and study wildlife species and habitats.', 'Calgary, AB', 70000.00, 'Mid-level'),
(4, 'Marketing Manager', 'Full-time', 'Develop and implement marketing strategies.', 'Birmingham, UK', 90000.00, 'Senior'),
(5, 'Web Developer', 'Part-time', 'Create and maintain websites for various clients.', 'Brisbane, Australia', 50000.00, 'Entry-level'),
(1, 'Data Analyst', 'Remote', 'Analyze data trends and generate reports.', 'Remote', 80000.00, 'Mid-level'),
(3, 'Product Manager', 'Full-time', 'Lead product development and coordinate with teams.', 'Toronto, ON', 95000.00, 'Senior'),
(2, 'Nutritional Scientist', 'Full-time', 'Conduct research on nutrition and health outcomes.', 'Los Angeles, CA', 85000.00, 'Mid-level'),
(4, 'Customer Support Specialist', 'Part-time', 'Provide support to customers via chat and email.', 'London, UK', 40000.00, 'Entry-level'),
(5, 'Content Writer', 'Freelance', 'Create engaging content for blogs and websites.', 'Remote', 30000.00, 'Entry-level');