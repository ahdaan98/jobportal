CREATE TABLE employers_newsletter (
    id SERIAL PRIMARY KEY,
    employer_id INT REFERENCES employers(id) ON DELETE CASCADE,
    content TEXT,
    isfree BOOLEAN DEFAULT true,
    amount DECIMAL(10, 2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    jobseeker_id INT REFERENCES jobseekers(id) ON DELETE CASCADE,
    newletter_id INT REFERENCES employers_newsletter(id) ON DELETE CASCADE,
    startdate TIMESTAMP,
    enddate TIMESTAMP,
    status VARCHAR(50) CHECK (status IN ('active', 'inactive', 'expired', 'canceled')) DEFAULT 'inactive'
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    subscription_id INT REFERENCES subscriptions(id) ON DELETE CASCADE,
    amount DECIMAL(10, 2),
    status VARCHAR(50) CHECK (status IN ('success', 'failure', 'pending', 'canceled')) DEFAULT 'pending',
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE razorpay_details (
    id SERIAL PRIMARY KEY,
    payment_id INT REFERENCES payments(id) ON DELETE CASCADE,
    pay_id varchar(255),
    order_id varchar(255),
    signature varchar(255)
);

INSERT INTO employers_newsletter (employer_id, content, isfree, amount) VALUES
(1, 'Weekly updates on the latest software engineering trends and job opportunities.', true, 0.00),
(1, 'Exclusive software development tips and job openings for experienced engineers.', false, 14.99),
(2, 'Health and wellness insights plus part-time job listings for registered nurses.', false, 9.99),
(2, 'HealthPlus monthly newsletter with healthcare news and job opportunities.', true, 0.00),
(3, 'Monthly newsletter on environmental conservation and scientist job opportunities.', true, 0.00),
(3, 'Exclusive reports on sustainability practices and new job openings in eco-friendly industries.', false, 12.99),
(4, 'Finance tips and job market analysis for financial professionals.', false, 19.99),
(4, 'Weekly insights on the global financial market with senior-level job openings.', true, 0.00),
(5, 'Travel tips and remote job opportunities in the travel industry.', true, 0.00),
(5, 'Exclusive travel deals and remote work opportunities for digital nomads.', false, 7.99);