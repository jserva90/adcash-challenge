CREATE TABLE loans (
   id INTEGER PRIMARY KEY,
   borrower_name TEXT NOT NULL,
   personal_id TEXT NOT NULL,
   loan_amount REAL NOT NULL,
   monthly_interest REAL NOT NULL DEFAULT 0.05,
   term INTEGER NOT NULL,
   status TEXT NOT NULL DEFAULT 'pending',
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blacklist (
   id INTEGER PRIMARY KEY,
   personal_id TEXT NOT NULL UNIQUE,
   reason TEXT NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE loan_applications (
   id INTEGER PRIMARY KEY,
   personal_id TEXT NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);