-- Loans Table
INSERT INTO loans (borrower_name, personal_id, loan_amount, term, status, created_at) VALUES
  ('John Doe', '123456789', 1000.0, 12, 'approved', '2023-03-14 10:00:00'),
  ('Jane Smith', '987654321', 500.0, 6, 'pending', '2023-03-13 14:30:00'),
  ('Bob Johnson', '456789123', 2000.0, 24, 'rejected', '2023-03-12 09:15:00'),
  ('John Doe', '123456789', 600.0, 8, 'approved', '2023-01-14 10:00:00'),
  ('John Doe', '123456789', 3500.0, 24, 'approved', '2022-03-14 10:00:00');


-- Blacklist Table
INSERT INTO blacklist (personal_id, reason, created_at) VALUES
  ('999999999', 'Fraud', '2023-03-10 16:45:00'),
  ('888888888', 'Default', '2023-03-09 11:20:00');

-- Loan Applications Table
INSERT INTO loan_applications (personal_id, created_at) VALUES
  ('123456789', '2023-03-14 09:30:00'),
  ('123456789', '2023-03-14 10:30:00'),
  ('987654321', '2023-03-13 12:45:00');