-- Delete Prizes
DELETE FROM prize_placement WHERE id IN (
  'c9e1074f-6e77-41d9-8e3d-d28b6edbc5c9',
  '14bfa6bb-04c8-4f68-8a2a-f684a4f04c3d',
  '45c48cce-2e2d-4ad6-9fa1-8db36f3e6a4a'
);

-- Delete Tournaments
DELETE FROM tournament WHERE id IN (
  'a1f27b7a-7d2b-4e6b-8a92-1f29b5b09c15',
  '3b6e7e94-34f3-49d0-8ca2-c0b88e0a4eab'
);

-- Delete Players
DELETE FROM player WHERE id IN (
  'd9b1d7db-5b8b-4b4c-9533-c1e1772f83e9',
  '2f8e9a6a-9f24-4e79-b545-f64d2a3d52cd',
  '5a7fbb2e-6044-4e0f-9e5a-d839cce8d124'
);
