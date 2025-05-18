-- Insert Players
INSERT INTO player (id, name, email, account_balance) VALUES
  ('d9b1d7db-5b8b-4b4c-9533-c1e1772f83e9', 'Alice Johnson', 'alice@example.com', 1100.50),
  ('2f8e9a6a-9f24-4e79-b545-f64d2a3d52cd', 'Bob Smith', 'bob@example.com', 750.50),
  ('5a7fbb2e-6044-4e0f-9e5a-d839cce8d124', 'Charlie Lee', 'charlie@example.com', 1250.75);

-- Insert Tournaments
INSERT INTO tournament (id, name, prize_pool, start_date, end_date) VALUES
  ('a1f27b7a-7d2b-4e6b-8a92-1f29b5b09c15', 'Footbal Tournament', 5000.00, '2025-04-01 10:00:00', '2025-04-10 18:00:00'),
  ('3b6e7e94-34f3-49d0-8ca2-c0b88e0a4eab', 'Basketball Tournament', 7500.00, '2025-06-15 09:00:00', '2025-06-25 20:00:00');

-- Insert Prizes
INSERT INTO prize_placement (id, playerID, tournamentID, placement) VALUES
  ('c9e1074f-6e77-41d9-8e3d-d28b6edbc5c9', 'd9b1d7db-5b8b-4b4c-9533-c1e1772f83e9', 'a1f27b7a-7d2b-4e6b-8a92-1f29b5b09c15', 1),
  ('14bfa6bb-04c8-4f68-8a2a-f684a4f04c3d', '2f8e9a6a-9f24-4e79-b545-f64d2a3d52cd', 'a1f27b7a-7d2b-4e6b-8a92-1f29b5b09c15', 2),
  ('45c48cce-2e2d-4ad6-9fa1-8db36f3e6a4a', '5a7fbb2e-6044-4e0f-9e5a-d839cce8d124', 'a1f27b7a-7d2b-4e6b-8a92-1f29b5b09c15', 3);
